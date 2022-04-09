package grid

import "math/rand"

// 3 types of grid location
// EMPTY, BARRIER, INDIV

type locType int

const (
	EMPTY locType = iota
	BARRIER
	INDIV
)

type Grid struct {
	rows []locType
	cols []locType
}

type Loc struct {
	x int
	y int
}

func (grid *Grid) isEmptyAt(x int, y int) bool {
	return grid.rows[x] == EMPTY && grid.cols[y] == EMPTY
}

// creates a new grid from the given config
func CreateGrid(sizeX int, sizeY int) *Grid {
	grid := Grid{
		rows: make([]locType, sizeX),
		cols: make([]locType, sizeY),
	}
	return &grid
}

// finds a random location in the grid
func (grid *Grid) FindRandomEmptyLoc() Loc {
	x := 0
	y := 0
	for true {
		x = rand.Intn(len(grid.rows))
		y = rand.Intn(len(grid.cols))
		if grid.isEmptyAt(x, y) {
			break
		}
	}
	return Loc{x, y}
}
