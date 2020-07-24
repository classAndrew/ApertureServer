package data

import (
	"math/big"

	"github.com/classAndrew/ApertureServer/pkg/server"
)

// DataManager sets the configurations for big-floating points
type DataManager struct {
	prec      uint
	base      int
	roundMode big.RoundingMode
	monHandle *MongoHandler
}

var monHand *MongoHandler = NewMongoHandler("localhost", "27017")

// DataMngr data manager variable for other parts of the server to use
var DataMngr DataManager = DataManager{128, 10, big.ToZero, monHand}

// NewUserData creates a new UserData struct with empty values
func NewUserData() server.UserData {
	return server.UserData{"", 0, 0, 0, 0, 0, 0, 0, []string{}, "", nil}
}

// RegisterUser method that adds a newly created User
func (dm *DataManager) RegisterUser(user *server.UserData) (string, bool) {
	usr := dm.monHandle.GetUserMon(user.Name)
	if usr.Name != "" {
		return "Name already taken", false
	}
	dm.monHandle.InsertUserMon(user)
	return "Success", true
}

// RegisterStarSystem registers an entire star system
func (dm *DataManager) RegisterStarSystem(starsys *server.StarSystem) string {
	sys := dm.monHandle.GetStarSystemMon(starsys.Name)
	if sys.Name != "" {
		return dm.RegisterStarSystem(starsys) // This is probably not the best way to handle duplicate starsystems
	}
	dm.monHandle.InsertStarSystemMon(starsys)
	return "Success"
}

// GetUser will get the user with necessary precision applied
func (dm *DataManager) GetUser(name string) server.UserData {
	usrDat := dm.monHandle.GetUserMon(name)
	usrDat.Balance128, _, _ = big.ParseFloat(usrDat.Balance, 10, 128, big.ToZero)
	return usrDat
}
