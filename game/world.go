package game

type World struct {
	last, current *Board
	gen           int
}

// NewWorld creates a new world from a starting board
// takes the board as a value so you don't have random
// pointers left the the same base board.
func NewWorld(b Board) *World {
	return &World{last: NewBoard(), current: &b}
}

// Next increments the world state to the next generation
func (w *World) Next() {
	w.gen++

	tl, br := w.current.GetLimits()
	w.last, w.current = w.current, NewBoard()

	for x := tl[0]; x <= tl[len(br)-1]; x++ {
		for y := tl[0]; y <= tl[len(tl)-1]; y++ {
			if w.last.NextState(x, y) {
				w.current.SetAlive(x, y)
			}
		}
	}

}

// CurrentGen of the board, as a value
func (w *World) CurrentGen() Board {
	return *w.current
}

// Generation of the world state.  Starts at 0, increases
// each call to next
func (w *World) Generation() int {
	return w.gen
}
