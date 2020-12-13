package types

// NavigationSystem represents the boats navigation system
type NavigationSystem struct {
	ShipPosition Position
	Move         map[Direction]func(int)
	direction    RolloverCounter
}

// NewNavigationSystem creates a configured navigation system
func NewNavigationSystem() *NavigationSystem {
	ns := NavigationSystem{
		ShipPosition: Position{X: 0, Y: 0},
		direction:    RolloverCounter{},
	}
	ns.Move = map[Direction]func(int){
		North:   func(distance int) { ns.move(North, distance) },
		South:   func(distance int) { ns.move(South, distance) },
		East:    func(distance int) { ns.move(East, distance) },
		West:    func(distance int) { ns.move(West, distance) },
		Forward: func(distance int) { ns.move(ns.currentDirection(), distance) },
		Left:    func(degrees int) { ns.rotate(Left, degrees) },
		Right:   func(degrees int) { ns.rotate(Right, degrees) },
	}

	return &ns
}

// ManhattenDistance tells you how far the ship has traveled from it's starting position
func (ns *NavigationSystem) ManhattenDistance() int {
	return abs(ns.ShipPosition.X) + abs(ns.ShipPosition.Y)
}

func (ns *NavigationSystem) currentDirection() Direction {
	switch ns.direction.Value() {
	case 0:
		return East
	case 90:
		return North
	case 180:
		return West
	case 270:
		return South
	}
	panic("Direction not specified")
}

func (ns *NavigationSystem) rotate(d Direction, degrees int) {
	switch d {
	case Left:
		ns.direction.Change(degrees)
	case Right:
		ns.direction.Change(-degrees)
	}
}

func (ns *NavigationSystem) move(d Direction, distance int) {
	switch d {
	case North:
		ns.ShipPosition.Y += distance
	case South:
		ns.ShipPosition.Y -= distance
	case East:
		ns.ShipPosition.X += distance
	case West:
		ns.ShipPosition.X -= distance
	}
}

// Go REALLY doesn't have a built in abs function for integers....WHAII
func abs(input int) int {
	if input < 0 {
		return -input
	}
	return input
}
