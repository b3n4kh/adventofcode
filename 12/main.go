package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Ship position and facing direction
type Ship struct {
	facing    rune
	latitudo  int
	longitudo int
}

// Waypoint position
type Waypoint struct {
	latitudo  int
	longitudo int
}

func (ship *Ship) getShip() {
	fmt.Printf("\nShip Facing: %v\nLong: %v\tLat: %v\n", string(ship.facing), ship.longitudo, ship.latitudo)
	fmt.Printf("Manhatten Distance to start: %v\n", ship.getDistance())
}

func (ship *Ship) move(direction rune, steps int) {
	switch direction {
	case 'N':
		ship.latitudo += steps
	case 'S':
		ship.latitudo -= steps
	case 'W':
		ship.longitudo += steps
	case 'E':
		ship.longitudo -= steps
	}
}

func (ship *Ship) moveToWaypoint(waypoint Waypoint, steps int) {
	ship.latitudo += waypoint.latitudo * steps
	ship.longitudo += waypoint.longitudo * steps
}

func (waypoint *Waypoint) move(direction rune, steps int) {
	switch direction {
	case 'N':
		waypoint.latitudo += steps
	case 'S':
		waypoint.latitudo -= steps
	case 'W':
		waypoint.longitudo += steps
	case 'E':
		waypoint.longitudo -= steps
	}
}

func (ship *Ship) tacking(direction rune, degree int) {
	rotation := degree / 90
	if direction == 'R' {
		rotation = 4 - rotation
	}
	//fmt.Printf("facing: %v direction: %v rotation: %v degree: %v\n", string(ship.facing), string(direction), rotation, degree)

	for i := 0; i < rotation; i++ {
		switch ship.facing {
		case 'N':
			ship.facing = 'W'
		case 'S':
			ship.facing = 'E'
		case 'W':
			ship.facing = 'S'
		case 'E':
			ship.facing = 'N'
		}
	}
}

func (waypoint *Waypoint) rotate(direction rune, degree int) {
	rotation := degree / 90
	if direction == 'R' {
		rotation = 4 - rotation
	}
	//fmt.Printf("facing: %v direction: %v rotation: %v degree: %v\n", string(ship.facing), string(direction), rotation, degree)
	var newLat, newLon int
	for i := 0; i < rotation; i++ {
		if waypoint.longitudo > 0 { // West
			newLat = waypoint.longitudo * -1 // South
		} else { // East
			newLat = waypoint.longitudo * -1 // North
		}
		if waypoint.latitudo > 0 { // North
			newLon = waypoint.latitudo // West
		} else { // South
			newLon = waypoint.latitudo
		}
		waypoint.latitudo = newLat
		waypoint.longitudo = newLon
	}
}

func parseInstructionsTwo(ship Ship, waypoint Waypoint, instruction string) (Ship, Waypoint) {
	direction := []rune(instruction[:1])[0]
	steps, _ := strconv.Atoi(instruction[1:])
	//fmt.Printf("Ship before: %+v\n", ship)
	switch direction {
	case 'F':
		ship.moveToWaypoint(waypoint, steps)
	case 'L', 'R':
		waypoint.rotate(direction, steps)
	default:
		waypoint.move(direction, steps)
	}
	//fmt.Printf("lon: %v lat: %v facing: %v\t\tInstruction: %v\n", ship.longitudo, ship.latitudo, string(ship.facing), instruction)
	return ship, waypoint
}

func (ship *Ship) parseInstructionsOne(instruction string) {
	direction := []rune(instruction[:1])[0]
	steps, _ := strconv.Atoi(instruction[1:])
	//fmt.Printf("Ship before: %+v\n", ship)
	switch direction {
	case 'F':
		ship.move(ship.facing, steps)
	case 'L', 'R':
		ship.tacking(direction, steps)
	default:
		ship.move(direction, steps)
	}
	//fmt.Printf("lon: %v lat: %v facing: %v\t\tInstruction: %v\n", ship.longitudo, ship.latitudo, string(ship.facing), instruction)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func (ship *Ship) getDistance() (distance int) {
	return abs(ship.latitudo) + abs(ship.longitudo)
}

func main() {
	dat, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer dat.Close()

	scanner := bufio.NewScanner(dat)
	var line string
	shipOne := Ship{'E', 0, 0}
	shipTwo := Ship{'E', 0, 0}
	waypoint := Waypoint{1, -10}
	for scanner.Scan() {
		line = scanner.Text()
		shipOne.parseInstructionsOne(line)
		shipTwo, waypoint = parseInstructionsTwo(shipTwo, waypoint, line)
	}
	shipOne.getShip()
	shipTwo.getShip()
	//fmt.Printf("%+v\n", allSeats.getNeighbors(3, 3))

}
