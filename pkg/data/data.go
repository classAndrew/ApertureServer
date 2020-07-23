package data

import (
	"math/big"
)

// DataManager sets the configurations for big-floating points
type DataManager struct {
	prec      uint
	base      int
	roundMode big.RoundingMode
	monHandle *MongoHandler
}

// UserData contains User's information
type UserData struct {
	Name            string     `json:"name"`
	NetOre          int        `json:"ore"`
	NetMetal        int        `json:"metal"`
	NetFood         int        `json:"food"`
	NetWattage      float64    `json:"wattage"`
	NetOil          float64    `json:"oil"`
	NetPopulation   int        `json:"population"`
	NetBodies       int        `json:"numberBodies"`
	CelestialBodies [][]int    `json:"bodies"`
	Balance         string     `json:"balance"`
	Balance128      *big.Float `json:"-"`
}

var monHand *MongoHandler = NewMongoHandler("localhost", "27017")

// DataMngr data manager variable for other parts of the server to use
var DataMngr DataManager = DataManager{128, 10, big.ToZero, monHand}

// NewUserData creates a new UserData struct with empty values
func NewUserData() UserData {
	return UserData{"", 0, 0, 0, 0, 0, 0, 0, [][]int{}, "", nil}
}

// RegisterUser method that adds a newly created User
func (dm *DataManager) RegisterUser(user *UserData) string {
	usr := dm.monHandle.GetUserMon(user.Name)
	if usr.Name != "" {
		return "Name already taken"
	}
	dm.monHandle.InsertUserMon(user)
	return "Success"
}

// GetUser will get the user with necessary precision applied
func (dm *DataManager) GetUser(name string) UserData {
	usrDat := dm.monHandle.GetUserMon(name)
	usrDat.Balance128, _, _ = big.ParseFloat(usrDat.Balance, 10, 128, big.ToZero)
	return usrDat
}
