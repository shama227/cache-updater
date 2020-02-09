package controllers;

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/go-pg/pg/v9"
	"github.com/ventuary-lab/cache-updater/src/entities"
)

type DbController struct {
	UcDelegate *UpdateController
	DbConnection *pg.DB
}

func (this *DbController) ConnectToDb () {
	dbuser, dbpass, dbdatabase := entities.GetDBCredentials()
  
	db := pg.Connect(&pg.Options{
		User:     dbuser,
		Password: dbpass,
		Database: dbdatabase,
	})

	this.DbConnection = db
}

func (this *DbController) HandleRecordsUpdate (byteValue []byte) {
	var records []entities.DAppStringRecord
	var numberRecords []entities.DAappNumberRecord

	json.Unmarshal([]byte(byteValue), &records)
	json.Unmarshal([]byte(byteValue), &numberRecords)

	nodeData := map[string]string{};

	for i := 0; i < len(records); i++ {
		record := records[i]

		if *record.Value == "" {
			numberRecord := numberRecords[i]

			*record.Value = strconv.Itoa(*numberRecord.Value)
		}

		nodeData[record.Key] = *record.Value;
	}

	var bondsorders []entities.BondsOrder
	var orderheights []uint64

	rawbo := entities.BondsOrder{}
	bondsorders = rawbo.UpdateAll(&nodeData)

	this.HandleBondsOrdersUpdate(&bondsorders)

	for _, order := range bondsorders {
		orderheights = append(orderheights, order.Height)
	}

	this.HandleBlocksMapUpdate(&orderheights)
}

func (this *DbController) HandleBondsOrdersUpdate (freshData *[]entities.BondsOrder) {
	var existingRecords []entities.BondsOrder

	_, getRecordsErr := this.DbConnection.
		Query(&existingRecords, fmt.Sprintf("SELECT * FROM %v;", entities.BONDS_ORDERS_NAME))

	if getRecordsErr != nil {
		return
	}

	isEmpty := len(existingRecords) == 0

	// Base case when table is empty, just upload and return
	if isEmpty {
		fmt.Printf("0 records exist \n")
		insertErr := this.DbConnection.Insert(freshData)

		if insertErr != nil {
			fmt.Printf("Error occured on Insert... %v \n", insertErr)
		} else {
			fmt.Printf("Successfully inserted %v rows \n", len(*freshData))
		}
	} else {
		var recordsToAdd []entities.BondsOrder
		updatedRecordsCount := 0
		
		for _, newRecord := range *freshData {
			exists := false
			for _, oldRecord := range existingRecords {
				if newRecord.Order_id == oldRecord.Order_id && (
					newRecord.Status != oldRecord.Status ||
					newRecord.Index != oldRecord.Index ||
					newRecord.Filledamount != oldRecord.Filledamount) {
					updateErr := this.DbConnection.Update(&newRecord)

					if updateErr != nil {
						fmt.Printf("Error occured on update... %v \n", updateErr)       
					} else {
						updatedRecordsCount++
					}

					exists = true
				} else if newRecord.Order_id == oldRecord.Order_id {
					exists = true
					break
				}
			}

			if !exists {
				recordsToAdd = append(recordsToAdd, newRecord)
			}
		}

		this.DbConnection.Insert(&recordsToAdd)

		fmt.Printf("Added %v, Updated %v rows... \n", len(recordsToAdd), updatedRecordsCount)
	}
}

func (this *DbController) TestUpdateBlocksMap () {
	// bm := entities.BlocksMap{}
	// bm.GetTimestampByHeight("77777")
}

func (this *DbController) HandleBlocksMapUpdate (heightarr *[]uint64) {
	var existingRecords []entities.BlocksMap
	var bondsOrders []entities.BondsOrder

	_, getRecordsErr := this.DbConnection.
		Query(&existingRecords, fmt.Sprintf("SELECT * FROM %v;", entities.BLOCKS_MAP_NAME))
	_, getBondsOrdersErr := this.DbConnection.
		Query(&bondsOrders, fmt.Sprintf("SELECT height FROM %v GROUP BY height ORDER BY height DESC;", entities.BONDS_ORDERS_NAME))

	if getRecordsErr != nil || getBondsOrdersErr != nil {
		fmt.Printf("Error occured on Query Select... %v; %v \n", getRecordsErr, getBondsOrdersErr)
		return
	}

	if len(bondsOrders) == 0 {
		fmt.Printf("%v table is empty... \n", entities.BONDS_ORDERS_NAME)
		return
	}

	bm := entities.BlocksMap{}

	if len(existingRecords) == 0 {
		var freshData []entities.BlocksMap

		minHeightBm := bondsOrders[0]
		maxHeightBm := bondsOrders[len(bondsOrders) - 1]
		maxRecordsCount := uint64(99)

		// GetBlocksMapSequenceByRange
		minHeight := minHeightBm.Height
		maxHeight := minHeightBm.Height + maxRecordsCount
		index := 1

		for {
			fetchedBlocksMap := bm.GetBlocksMapSequenceByRange(fmt.Sprintf("%v", minHeight), fmt.Sprintf("%v", maxHeight))

			freshData = append(freshData, *fetchedBlocksMap...)
			minHeight := maxHeight + 1
			maxHeight := minHeight + maxRecordsCount

			if maxHeight == maxHeightBm.Height {
				break
			}
			if maxHeight > maxHeightBm.Height {
				maxHeight = maxHeightBm.Height
			}

			index++

			if index == 3 {
				break
			}
		}

		fmt.Printf("blocks count: %v", len(freshData))
		insertErr := this.DbConnection.Insert(&freshData)

		if insertErr != nil {
			fmt.Printf("Error occured on Insert... %v \n", insertErr)
		} else {
			fmt.Printf("Successfully inserted %v rows \n", len(freshData))
		}
	}
}