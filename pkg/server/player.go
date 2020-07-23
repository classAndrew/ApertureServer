package server

import (
	"math/rand"

	"github.com/classAndrew/ApertureServer/pkg/data"
)

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
func (pm *PlayerManager) CreateUser(name string) *data.UserData {
	planets := make([][]int, 0)
	for i := 0; i < PM.DefaultPlanets; i++ {
		coordinates := []int{rand.Intn(MAXSIZE), rand.Intn(MAXSIZE)}
		planets = append(planets, coordinates)
	}

	user := &data.UserData{
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
