package types

import (
	"adventOfCode2020/util"
	"math"
)

// WaypointNavigationSystem represents the boats navigation system
type WaypointNavigationSystem struct {
	ShipPosition      Position
	WaypointPosition  Position
	Move              map[Direction]func(int)
	waypointDirection RolloverCounter
}

// NewWaypointNavigationSystem creates a configured waypoint navigation system
func NewWaypointNavigationSystem() *WaypointNavigationSystem {
	ns := WaypointNavigationSystem{
		ShipPosition:      Position{X: 0, Y: 0},
		WaypointPosition:  Position{X: 10, Y: 1},
		waypointDirection: RolloverCounter{},
	}
	ns.Move = map[Direction]func(int){
		North:   func(distance int) { ns.moveWaypoint(North, distance) },
		South:   func(distance int) { ns.moveWaypoint(South, distance) },
		East:    func(distance int) { ns.moveWaypoint(East, distance) },
		West:    func(distance int) { ns.moveWaypoint(West, distance) },
		Forward: func(distance int) { ns.moveShip(distance) },
		Left:    func(degrees int) { ns.rotateWaypoint(Left, degrees) },
		Right:   func(degrees int) { ns.rotateWaypoint(Right, degrees) },
	}

	return &ns
}

// ManhattenDistance tells you how far the ship has traveled from it's starting position
func (ns *WaypointNavigationSystem) ManhattenDistance() int {
	return util.Abs(ns.ShipPosition.X) + util.Abs(ns.ShipPosition.Y)
}

func (ns *WaypointNavigationSystem) rotateWaypoint(d Direction, degrees int) {
	oldDirection := ns.waypointDirection
	switch d {
	case Left:
		ns.waypointDirection.Change(degrees)
	case Right:
		ns.waypointDirection.Change(-degrees)
	}

	toRad := func(input int) float64 { return float64(input) * (math.Pi / 180.0) }
	angle := toRad(ns.waypointDirection.Value() - oldDirection.Value())
	x, y := float64(ns.WaypointPosition.X), float64(ns.WaypointPosition.Y)
	ns.WaypointPosition.X = int(math.Round(x*math.Cos(angle)) - math.Round(y*math.Sin(angle)))
	ns.WaypointPosition.Y = int(math.Round(y*math.Cos(angle)) + math.Round(x*math.Sin(angle)))
}

func (ns *WaypointNavigationSystem) moveShip(count int) {
	ns.ShipPosition.X += count * ns.WaypointPosition.X
	ns.ShipPosition.Y += count * ns.WaypointPosition.Y
}

func (ns *WaypointNavigationSystem) moveWaypoint(d Direction, distance int) {
	switch d {
	case North:
		ns.WaypointPosition.Y += distance
	case South:
		ns.WaypointPosition.Y -= distance
	case East:
		ns.WaypointPosition.X += distance
	case West:
		ns.WaypointPosition.X -= distance
	}
}
