// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RewardStatus reward status
// swagger:model RewardStatus
type RewardStatus struct {

	// current reward
	// Required: true
	CurrentReward *int64 `json:"currentReward"`

	// height
	// Required: true
	Height *int32 `json:"height"`

	// min increment
	// Required: true
	MinIncrement *int64 `json:"minIncrement"`

	// next check
	// Required: true
	NextCheck *int32 `json:"nextCheck"`

	// term
	// Required: true
	Term *int32 `json:"term"`

	// total waves amount
	// Required: true
	TotalWavesAmount *int32 `json:"totalWavesAmount"`

	// votes
	// Required: true
	Votes *RewardStatusVotes `json:"votes"`

	// voting interval
	// Required: true
	VotingInterval *int32 `json:"votingInterval"`

	// voting interval start
	// Required: true
	VotingIntervalStart *int32 `json:"votingIntervalStart"`

	// voting threshold
	// Required: true
	VotingThreshold *int32 `json:"votingThreshold"`
}

// Validate validates this reward status
func (m *RewardStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCurrentReward(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHeight(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMinIncrement(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNextCheck(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTerm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalWavesAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVotes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVotingInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVotingIntervalStart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVotingThreshold(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RewardStatus) validateCurrentReward(formats strfmt.Registry) error {

	if err := validate.Required("currentReward", "body", m.CurrentReward); err != nil {
		return err
	}

	return nil
}

func (m *RewardStatus) validateHeight(formats strfmt.Registry) error {

	if err := validate.Required("height", "body", m.Height); err != nil {
		return err
	}

	return nil
}

func (m *RewardStatus) validateMinIncrement(formats strfmt.Registry) error {

	if err := validate.Required("minIncrement", "body", m.MinIncrement); err != nil {
		return err
	}

	return nil
}

func (m *RewardStatus) validateNextCheck(formats strfmt.Registry) error {

	if err := validate.Required("nextCheck", "body", m.NextCheck); err != nil {
		return err
	}

	return nil
}

func (m *RewardStatus) validateTerm(formats strfmt.Registry) error {

	if err := validate.Required("term", "body", m.Term); err != nil {
		return err
	}

	return nil
}

func (m *RewardStatus) validateTotalWavesAmount(formats strfmt.Registry) error {

	if err := validate.Required("totalWavesAmount", "body", m.TotalWavesAmount); err != nil {
		return err
	}

	return nil
}

func (m *RewardStatus) validateVotes(formats strfmt.Registry) error {

	if err := validate.Required("votes", "body", m.Votes); err != nil {
		return err
	}

	if m.Votes != nil {
		if err := m.Votes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("votes")
			}
			return err
		}
	}

	return nil
}

func (m *RewardStatus) validateVotingInterval(formats strfmt.Registry) error {

	if err := validate.Required("votingInterval", "body", m.VotingInterval); err != nil {
		return err
	}

	return nil
}

func (m *RewardStatus) validateVotingIntervalStart(formats strfmt.Registry) error {

	if err := validate.Required("votingIntervalStart", "body", m.VotingIntervalStart); err != nil {
		return err
	}

	return nil
}

func (m *RewardStatus) validateVotingThreshold(formats strfmt.Registry) error {

	if err := validate.Required("votingThreshold", "body", m.VotingThreshold); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RewardStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RewardStatus) UnmarshalBinary(b []byte) error {
	var res RewardStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RewardStatusVotes reward status votes
// swagger:model RewardStatusVotes
type RewardStatusVotes struct {

	// decrease
	// Required: true
	Decrease *int32 `json:"decrease"`

	// increase
	// Required: true
	Increase *int32 `json:"increase"`
}

// Validate validates this reward status votes
func (m *RewardStatusVotes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDecrease(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIncrease(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RewardStatusVotes) validateDecrease(formats strfmt.Registry) error {

	if err := validate.Required("votes"+"."+"decrease", "body", m.Decrease); err != nil {
		return err
	}

	return nil
}

func (m *RewardStatusVotes) validateIncrease(formats strfmt.Registry) error {

	if err := validate.Required("votes"+"."+"increase", "body", m.Increase); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RewardStatusVotes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RewardStatusVotes) UnmarshalBinary(b []byte) error {
	var res RewardStatusVotes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
