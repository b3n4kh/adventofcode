package main

import (
	"bufio"
	"fmt"
	"os"
)

// Coord Row and Col of Seat
type Coord struct {
	Row, Col int
}

var partB = false
var neighbors = make(map[Coord][]Coord)

func (c Coord) adjacent(occupied map[Coord]bool) []Coord {
	if neighbors[c] != nil {
		return neighbors[c]
	}
	var ret []Coord
	for rOffset := -1; rOffset <= 1; rOffset++ {
		for cOffset := -1; cOffset <= 1; cOffset++ {
			if cOffset == 0 && rOffset == 0 {
				continue
			}
			if !partB {
				ret = append(ret, Coord{c.Row + rOffset, c.Col + cOffset})
			} else {
				for i := 1; ; i++ {
					pos := Coord{c.Row + rOffset*i, c.Col + cOffset*i}
					if pos.Col < 0 || pos.Row < 0 || pos.Col > 100 || pos.Row > 100 {
						break
					}
					if _, seat := occupied[pos]; seat {
						ret = append(ret, pos)
						break
					}
				}
			}
		}
	}
	neighbors[c] = ret
	return ret
}

func (c Coord) occupiedNeighbors(occupied map[Coord]bool) int {
	count := 0
	for _, n := range c.adjacent(occupied) {
		if occupied[n] {
			count++
		}
	}
	return count
}

func main() {
	dat, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer dat.Close()
	scanner := bufio.NewScanner(dat)
	var line string

	occupied := make(map[Coord]bool)
	var split []string
	for scanner.Scan() {
		line = scanner.Text()
		split = append(split, line)
	}

	for r, s := range split {
		for c, v := range s {
			switch v {
			case 'L':
				occupied[Coord{r, c}] = false
			case '#':
				occupied[Coord{r, c}] = true
			case '.':
				delete(occupied, Coord{r, c})
			}
		}
	}

	occupyThreshold := 4
	if partB {
		occupyThreshold = 5
	}

	for {
		changed := false
		prev := make(map[Coord]bool)
		for k, v := range occupied {
			prev[k] = v
		}
		for k, occ := range prev {
			if occ {
				if k.occupiedNeighbors(prev) >= occupyThreshold {
					occupied[k] = false
					changed = true
				}
			} else {
				if k.occupiedNeighbors(prev) == 0 {
					occupied[k] = true
					changed = true
				}
			}
		}
		if !changed {
			break
		}
	}

	result := 0
	for _, v := range occupied {
		if v {
			result++
		}
	}
	fmt.Println(result)
}
