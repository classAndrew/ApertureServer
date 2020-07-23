package server

import (
	"math/rand"
)

// Planet planet struct
type Planet struct {
	mass       float64
	planetType string
	name       string
	pos        Position
}

const alnum string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ012345689"
const alphaIDMIN int = 3
const alphaIDMAX int = 5
const numIDLEN int = 4
const massMIN int = 20 // I guess these could be measured as metric Megatons
const massMAX int = 5000

// GeneratePlanet Generates a planet with random attributes
func GeneratePlanet() *Planet {
	// name := generatePlanetName()
	// mass := rand.Intn(massMAX-massMIN+1) + massMIN
	return nil
}

func generatePlanetName() string {
	name := make([]byte, 0)
	for i := 0; i < rand.Intn(alphaIDMAX-alphaIDMIN+1)+alphaIDMIN; i++ {
		name = append(name, alnum[rand.Intn(26)])
	}
	for i := 0; i < rand.Intn(numIDLEN); i++ {
		name = append(name, alnum[rand.Intn(10)+26])
	}
	return string(name)
}
