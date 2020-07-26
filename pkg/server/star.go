package server

import (
	"math/rand"
)

// StarSystem The star system holds data of planets and the star(s) (perhaps I could add binary, ternary, quaternary stars and moons as well)
type StarSystem struct {
	Name        string
	Pos         Position
	Star        Star
	Planets     []Planet `json:"-"`
	PlanetNames []string
}

// Star star struct
type Star struct {
	Name        string
	pos         Position
	typeStar    string
	Mass        float64
	planetCount int
	Luminosity  float64
}

const starNameLen int = 8
const minMass float64 = 200   // 200 metic kilotons
const maxMass float64 = 10000 // Pretty large stars out there
const maxPlanets int = 10     // min of 1 planet

// NewStarSystem allocates memory for the star system and return a pointer to an empty one
func NewStarSystem() *StarSystem {
	return new(StarSystem)
}

// GenerateStarSystem generates a star system
func GenerateStarSystem() StarSystem {
	star := GenerateStar()
	planNames := []string{}
	planets := make([]Planet, star.planetCount)
	for i := 0; i < star.planetCount; i++ {
		planets[i] = GeneratePlanet(star, i+1, "")
		planNames = append(planNames, planets[i].Name)
	}
	return StarSystem{star.Name, star.pos, star, planets, planNames}
}

// GenerateStar generates a star and returns its pointer
func GenerateStar() Star {
	pos := GeneratePos()
	luminosity := rand.Float64()
	mass := (rand.Float64() + minMass/maxMass) * maxMass
	name := generateStarName()
	planetCount := rand.Intn(maxPlanets) + 1
	return Star{name, pos, "star", mass, planetCount, luminosity}
}

func generateStarName() string {
	name := make([]byte, 0)
	for i := 0; i < starNameLen; i++ {
		name = append(name, Alnum[rand.Intn(26)])
	}
	return string(name)
}
