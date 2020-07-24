package server

import (
	"math/rand"
)

// Position of a celestial body
type Position struct {
	X     int
	Y     int // Bodies orbiting a star will inherit its star's galactic position
	Index int // This is the index of some body orbiting a system. (Stars will be 0)
}

// Body interface for general celestial / built bodies in space
type Body interface {
	GetMass() float64
	GetType() string
	GetName() string
	GetPos() Position
}

// GeneratePos by default will set index to 0
func GeneratePos() Position {
	posX, posY := rand.Intn(MAXSIZE), rand.Intn(MAXSIZE)
	return Position{posX, posY, 0}
}
