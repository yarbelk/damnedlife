package game

import "sort"

type Board struct {
	world map[int]map[int]bool
}

type Point [2]int

func NewBoard() *Board {
	return &Board{make(map[int]map[int]bool)}
}

func neighbors(x, y int) (neighbors []Point) {
	neighbors = []Point{
		Point{x - 1, y - 1}, Point{x - 1, y + 0}, Point{x - 1, y + 1},
		Point{x, y - 1}, Point{x, y + 1},
		Point{x + 1, y - 1}, Point{x + 1, y + 0}, Point{x + 1, y + 1},
	}
	return
}

// Set a given x/y coord to the passed in state
func (b *Board) Set(x, y int, alive bool) {
	if b.world[x] == nil {
		b.world[x] = make(map[int]bool)
	}
	b.world[x][y] = alive
}

// Get the state of a point
func (b *Board) Get(x, y int) bool {
	if b.world[x] == nil {
		return false
	}
	return b.world[x][y]
}

// NextState of the cell at x, y
func (b *Board) NextState(x, y int) bool {
	var c int
	var currentState bool

	for _, p := range neighbors(x, y) {
		i, j := p[0], p[1]
		if b.world[i] == nil {
			continue
		} else if b.world[i][j] {
			c++
		}
	}
	if b.world[x] != nil {
		currentState = b.world[x][y]
	}
	return currentState && (c == 2 || c == 3) || (!currentState && c == 3)
}

// GetLimits returns the Top Left and Bottom right extents of
// of the board.  Not the most efficent implementation, but
// the easiest to read
func (b *Board) GetLimits() (Point, Point) {
	var xs, ys []int

	for x := range b.world {
		for y := range b.world[x] {
			xs = append(xs, x)
			ys = append(ys, y)
		}
	}
	sort.Ints(xs)
	sort.Ints(ys)
	return Point{xs[0], ys[0]}, Point{xs[len(xs)-1], ys[len(ys)-1]}
}

// func (b *Board) String() string {
// 	var board [][]byte
// 	tl, br := b.GetLimits()
// }
