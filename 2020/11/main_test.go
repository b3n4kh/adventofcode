package main

import (
	"strings"
	"testing"
)

const testdata = `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`

func Test_main(t *testing.T) {
	occupied := make(map[Coord]bool)

	split := strings.Split(testdata, "\n")
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

	if result != 37 {
		t.Errorf("The result was: %v wanted %v ", result, 37)
	}

}
