package server

import (
	"math/big"
)

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
	CelestialBodies []string   `json:"bodies"`
	Balance         string     `json:"balance"`
	Balance128      *big.Float `json:"-"`
}

// PlayerManager struct of default values to give to newly created players
type PlayerManager struct {
	DefaultOre     int     `json:"ore"`
	DefaultMetal   int     `json:"metal"`
	DefaultFood    int     `json:"food"`
	DefaultPlanets int     `json:"planets"`
	DefaultOil     float64 `json:"oil"`
}

// PM Singleton of the playermanager struct
var PM PlayerManager = PlayerManager{100, 100, 100, 1, 100}

// CreateUser creates a user with default starting resources
func (pm *PlayerManager) CreateUser(name string, starname string) *UserData {
	planets := make([]string, 1)
	planets[0] = starname
	user := &UserData{
		name,
		PM.DefaultOre,
		PM.DefaultMetal,
		PM.DefaultFood,
		0,
		PM.DefaultOil,
		100,
		PM.DefaultPlanets,
		planets,
		"",
		nil}
	return user
}
