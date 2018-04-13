package game

import (
	"bytes"
	"sort"
)

// Cell is an alive or dead type
type Cell bool

// Rune return a Human readable rune for the cell
// like String, but for runes.
func (c Cell) Rune() rune {
	if c {
		return '█'
	}
	return '░'
}

// Board holds a particular itteration/state of the world
type Board struct {
	world map[Point]bool
}

// Point is a bog-standard cartesional point
type Point struct {
	X, Y int
}

type points []Point

func (p points) Len() int {
	return len(p)
}

func (p points) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p points) Less(i, j int) bool {
	return (p[i].Y < p[j].Y) || (p[i].Y == p[j].Y) && (p[i].X < p[j].X)
}

// NewBoard creates an empty board
func NewBoard() *Board {
	return &Board{make(map[Point]bool)}
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
	b.world[Point{x, y}] = true
}

// Get the state of a point
func (b *Board) Get(x, y int) bool {
	return b.world[Point{x, y}]
}

// NextState of the cell at x, y
func (b *Board) NextState(x, y int) bool {
	var c int
	var currentState bool

	for _, p := range neighbors(x, y) {
		if b.world[Point{p.X, p.Y}] {
			c++
		}
	}
	currentState = b.world[Point{x, y}]
	return currentState && (c == 2 || c == 3) || (!currentState && c == 3)
}

// GetLimits returns the Top Left and Bottom right extents of
// of the board.  Not the most efficent implementation, but
// the easiest to read
func (b *Board) GetLimits() (Point, Point) {
	var xs, ys []int

	for p := range b.world {
		xs = append(xs, p.X)
		ys = append(ys, p.Y)
	}
	sort.Ints(xs)
	sort.Ints(ys)
	if len(xs) == 0 {
		return Point{0, 0}, Point{0, 0}
	}
	return Point{xs[0], ys[0]}, Point{xs[len(xs)-1], ys[len(ys)-1]}
}

// GetOpenCells return all cells currently alives, and their neighbors
func (b *Board) GetOpenCells() map[Point]bool {
	var allCells = make(map[Point]bool)
	for _, cell := range b.AllAlive() {
		allCells[cell] = true
		for _, neighbor := range neighbors(cell.X, cell.Y) {
			allCells[neighbor] = true
		}
	}

	return allCells
}

// AllAlive cells on the board
func (b *Board) AllAlive() (points []Point) {

	for p := range b.world {
		if b.world[p] { // should always be true; but check anyway
			points = append(points, p)
		}
	}
	return
}

// sanePrintLimits (0,0) - (2, 2) are returned if the passed
// in ones are too small. EG, always start at 0,0 unless
// there is stuff further up.
func sanePrintLimits(tl, br Point) (Point, Point) {
	if tl.X > 0 {
		tl.X = 0
	}
	if tl.Y > 0 {
		tl.Y = 0
	}
	if br.X < 2 {
		br.X = 2
	}
	if br.Y < 2 {
		br.Y = 2
	}
	return tl, br
}

func (b *Board) String() string {
	var buffer = bytes.Buffer{}

	tl, br := sanePrintLimits(b.GetLimits())
	for x := tl.X; x <= br.X; x++ {
		for y := tl.Y; y <= br.Y; y++ {
			buffer.WriteRune(Cell(b.Get(x, y)).Rune())
		}
		buffer.WriteRune('\n')
	}
	buffer.Truncate(buffer.Len() - 1) // drop final newline
	return buffer.String()
}
