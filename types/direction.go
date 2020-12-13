package types

// Direction represents a direction the ship can move in
type Direction int

// Enums to represent all the directions
const (
	North Direction = iota
	South
	East
	West
	Left
	Right
	Forward
)
