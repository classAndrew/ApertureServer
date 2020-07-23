package server

// StarSystem The star system holds data of planets and the star(s) (perhaps I could add binary, ternary, quaternary stars and moons as well)
type StarSystem struct {
	mass    float64
	star    *Star
	planets []*Planet
}

// Star star struct
type Star struct {
	name       string
	pos        Position
	typeStar   string
	mass       float64
	maxPlanets int
	luminosity float64
}

// GenerateStarSystem generates a star system
func GenerateStarSystem() *StarSystem {
	return nil
}
