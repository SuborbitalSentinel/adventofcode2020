package types

import (
	"fmt"
	"strings"
)

// Plane is the plane
type Plane struct {
	Layout [][]rune
}

// NewPlane creates a new seat layout
func NewPlane(input []string) *Plane {
	layout := Plane{}
	for _, row := range input {
		layout.Layout = append(layout.Layout, []rune(row))
	}
	return &layout
}

// Print prints out the planes layout
func (p *Plane) Print() {
	fmt.Println(strings.Repeat("~", len(p.Layout[0])))
	for _, row := range p.Layout {
		fmt.Println(string(row))
	}
	fmt.Println(strings.Repeat("~", len(p.Layout[0])))
}

// NextSeatingArrangement returns the number of changed seats
func (p *Plane) NextSeatingArrangement() (changedSeats int) {
	newLayout := [][]rune{}
	for y := range p.Layout {
		newRow := []rune{}
		for x := range p.Layout[y] {
			changed, symbol := p.modifyState(x, y)
			newRow = append(newRow, symbol)
			if changed {
				changedSeats++
			}
		}
		newLayout = append(newLayout, newRow)
	}
	p.Layout = newLayout
	fmt.Println("Changed Seats:", changedSeats)
	return
}

// CountOccupiedSeats returns the number of non-empty seats
func (p *Plane) CountOccupiedSeats() (count int) {
	for y := range p.Layout {
		for x := range p.Layout[y] {
			if p.isOccupiedSeat(x, y) {
				count++
			}
		}
	}
	return
}

func (p *Plane) modifyState(x, y int) (bool, rune) {
	if p.isFloor(x, y) {
		return false, '.'
	}
	if p.isEmptySeat(x, y) && p.shouldBecomeOccupiedLos(x, y) {
		return true, '#'
	}
	if p.isOccupiedSeat(x, y) && p.shouldBecomeEmptyLos(x, y) {
		return true, 'L'
	}
	return false, p.Layout[y][x]
	//if p.isEmptySeat(x, y) && p.shouldBecomeOccupied(x, y) {
	//	return true, '#'
	//}
	//if p.isOccupiedSeat(x, y) && p.shouldBecomeEmpty(x, y) {
	//	return true, 'L'
	//}
}

func (p *Plane) isEmptySeat(x, y int) bool {
	return p.Layout[y][x] == 'L'
}

func (p *Plane) isOccupiedSeat(x, y int) bool {
	return p.Layout[y][x] == '#'
}

func (p *Plane) isFloor(x, y int) bool {
	return p.Layout[y][x] == '.'
}

func (p *Plane) shouldBecomeOccupiedLos(x, y int) bool {
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			found, seat := p.findSeat(x, j, y, i)
			if found && seat == '#' {
				return false
			}
		}
	}
	return true
}

func (p *Plane) shouldBecomeEmptyLos(x, y int) bool {
	numberOfEmptySeats := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}

			found, seat := p.findSeat(x, j, y, i)
			if found && seat == 'L' {
				numberOfEmptySeats++
			}
		}
	}
	return numberOfEmptySeats >= 5
}

func (p *Plane) findSeat(x, xDir, y, yDir int) (bool, rune) {
	xPos, yPos := x+xDir, y+yDir
	for i := 1; 0 <= xPos && xPos < len(p.Layout[0]) && 0 <= yPos && yPos < len(p.Layout); i++ {
		if !p.isFloor(xPos, yPos) {
			return true, p.Layout[yPos][xPos]
		}
		xPos = x + (xDir * i)
		yPos = y + (yDir * i)
	}
	return false, rune(0)
}

func (p *Plane) shouldBecomeOccupied(x, y int) bool {
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if !p.isFloor(x+j, y+i) && !p.isEmptySeat(x+j, y+i) {
				return false
			}
		}
	}
	return true
}

func (p *Plane) shouldBecomeEmpty(x, y int) bool {
	numberOfEmptySeats := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if !p.isFloor(x+j, y+i) && !p.isEmptySeat(x+j, y+i) {
				numberOfEmptySeats++
			}
		}
	}
	return numberOfEmptySeats >= 4
}
