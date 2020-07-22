package main

import (
	"math/big"
)

// DataManager sets the configurations for big-floating points
type DataManager struct {
	prec      uint
	base      int
	roundMode big.RoundingMode
	monHandle MongoHandler
}

// UserData contains User's information
type UserData struct {
	Balance    string     `json:"balance"`
	Balance128 *big.Float `json:"-"`
}

// NewUserData creates a new UserData struct with default values
func NewUserData() UserData {
	return UserData{"", nil}
}

// GetUser will get the user with necessary precision applied
func (dm *DataManager) GetUser(name string) UserData {
	usrDat := dm.monHandle.GetUserMon(name)
	usrDat.Balance128, _, _ = big.ParseFloat(usrDat.Balance, 10, 128, big.ToZero)
	return usrDat
}
