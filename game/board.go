package game

import (
	"bytes"
	"sort"
)

type Cell bool

func (c Cell) Rune() rune {
	if c {
		return '█'
	}
	return '░'
}

type Board struct {
	world map[int]map[int]bool
}

type Point [2]int

func (p Point) X() int { return p[0] }
func (p Point) Y() int { return p[1] }

type points []Point

func (p points) Len() int {
	return len(p)
}

func (p points) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p points) Less(i, j int) bool {
	return (p[i][1] < p[j][1]) || (p[i][1] == p[j][1]) && (p[i][0] < p[j][0])
}

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

// SetAlive makes a given x/y coord to be alive
func (b *Board) SetAlive(x, y int) {
	if b.world[x] == nil {
		b.world[x] = make(map[int]bool)
	}
	b.world[x][y] = true
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

func (b *Board) AllAlive() (points []Point) {
	for x := range b.world {
		for y := range b.world[x] {
			points = append(points, Point{x, y})
		}
	}
	return
}

func (b *Board) String() string {
	var buffer bytes.Buffer = bytes.Buffer{}

	tl, br := b.GetLimits()
	for x := tl.X(); x <= br.X(); x++ {
		for y := tl.Y(); y <= br.Y(); y++ {
			buffer.WriteRune(Cell(b.Get(x, y)).Rune())
		}
		buffer.WriteRune('\n')
	}
	return buffer.String()
}
