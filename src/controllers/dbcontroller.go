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

	_, getRecordsErr := this.DbConnection.
		Query(&existingRecords, fmt.Sprintf("SELECT * FROM %v;", entities.BLOCKS_MAP_NAME))

	if getRecordsErr != nil {
		fmt.Printf("Error occured on Query Select... %v \n", getRecordsErr)
		return
	}

	bm := entities.BlocksMap{}

	if len(existingRecords) == 0 {
		var freshData []entities.BlocksMap

		for index, heightValue := range *heightarr {
			strheight := strconv.Itoa(int(heightValue))
			blocksMap := bm.GetBlocksMapByHeight(strheight)
			freshData = append(freshData, *blocksMap)

			// TODO
			if index == 10 {
				break
			}
		}

		insertErr := this.DbConnection.Insert(freshData)

		if insertErr != nil {
			fmt.Printf("Error occured on Insert... %v \n", insertErr)
		} else {
			fmt.Printf("Successfully inserted %v rows \n", len(freshData))
		}
	}
}