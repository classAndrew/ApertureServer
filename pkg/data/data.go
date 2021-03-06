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
	return server.UserData{"", 0, 0, 0, 0, 0, 0, 0, "", []string{}, "", nil}
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
		// FIX, move recursive outside to caller
		return dm.RegisterStarSystem(starsys) // This is probably not the best way to handle duplicate starsystems
	}
	dm.monHandle.InsertStarSystemMon(starsys)
	for _, planet := range starsys.Planets {
		dm.monHandle.InsertPlanetMon(&planet)
	}
	return "Success"
}

// GetStarSystem will get the star system
func (dm *DataManager) GetStarSystem(name string) *server.StarSystem {
	sysDat := dm.monHandle.GetStarSystemMon(name)
	return sysDat
}

// RandomNovelPlanet returns a novel planet that's unexplored / uncolonized
func (dm *DataManager) RandomNovelPlanet() *server.Planet {
	return nil
}

// GetPlanet will get a planet by name
func (dm *DataManager) GetPlanet(name string) *server.Planet {
	return dm.monHandle.GetPlanetMon(name)
}

// SetRandomPlanetAttribute returns planet name, sets random novel planet's attribute (this really is only needed for setting initial home planet)
func (dm *DataManager) SetRandomPlanetAttribute(attrib string, value string) string {
	planetName := dm.monHandle.SetRandomPlanetNovelMon(attrib, value)
	dm.monHandle.UpdatePlanetMon(planetName, attrib, value)
	return planetName
}

// SetPlanetAttribute will set a planet attribute
func (dm *DataManager) SetPlanetAttribute(name string, attrib string, value string) {
	dm.monHandle.UpdatePlanetMon(name, attrib, value)
}

// RandomStarSystem will get a random star system that's unexplored (DOESN'T WORK YET)
func (dm *DataManager) RandomStarSystem() *server.StarSystem {
	sysDat := dm.monHandle.GetRandomStarSystem()
	return sysDat
}

// GetUser will get the user with necessary precision applied
func (dm *DataManager) GetUser(name string) server.UserData {
	usrDat := dm.monHandle.GetUserMon(name)
	usrDat.Balance128, _, _ = big.ParseFloat(usrDat.Balance, 10, 128, big.ToZero)
	return usrDat
}
