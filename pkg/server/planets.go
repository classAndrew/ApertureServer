package server

import (
	"math/rand"
)

// Planet planet struct
type Planet struct {
	Mass       float64
	PlanetType string
	Name       string
	Pos        Position
}

// Alnum alphanumeric
const Alnum string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ012345689"
const alphaIDMIN int = 3
const alphaIDMAX int = 5
const numIDLEN int = 4
const massMIN float64 = 20 // I guess these could be measured as metric Megatons
const massMAX float64 = 5000

// GeneratePlanet Generates a planet with random attributes.
func GeneratePlanet(star Star, ind int) Planet {
	name := generatePlanetName()
	mass := rand.Float64()*massMAX + massMIN
	star.pos.Index = ind
	return Planet{mass, "planet", star.Name + name, star.pos}
}

func generatePlanetName() string {
	name := make([]byte, 0)
	for i := 0; i < rand.Intn(alphaIDMAX-alphaIDMIN+1)+alphaIDMIN; i++ {
		name = append(name, Alnum[rand.Intn(26)])
	}
	for i := 0; i < rand.Intn(numIDLEN); i++ {
		name = append(name, Alnum[rand.Intn(10)+26])
	}
	return string(name)
}
