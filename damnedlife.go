package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	gc "github.com/rthornton128/goncurses"
	"github.com/yarbelk/cgol2/game"
)

const (
	ALIVE = '█'
	DEAD  = '░'
)

const (
	TITLE_HEIGHT  = 5
	FOOTER_HEIGHT = 3
)

// ncurses version; thus the 'damned' part of the life

func setupTitle(win *gc.Window) {
	win.Erase()
	// func (w *Window) Border(ls, rs, ts, bs, tl, tr, bl, br Char) error
	win.Border(gc.ACS_VLINE, gc.ACS_VLINE, gc.ACS_HLINE, gc.ACS_HLINE, gc.ACS_ULCORNER, gc.ACS_URCORNER, gc.ACS_LTEE, gc.ACS_RTEE)
	win.Keypad(false)
	_, x := win.MaxYX()
	title := "Conways Game of Life"
	win.MovePrint(2, (x/2 - len(title)/2), title)
	win.MovePrint(3, (x/2 - len(title)/2), "(press Q to exit)")
}

func setupField(win *gc.Window) *gc.Window {
	win.Color(2)
	win.Erase()
	win.Timeout(200)
	// func (w *Window) Border(ls, rs, ts, bs, tl, tr, bl, br Char) error
	win.Border(gc.ACS_VLINE, gc.ACS_VLINE, ' ', ' ', gc.ACS_VLINE, gc.ACS_VLINE, gc.ACS_VLINE, gc.ACS_VLINE)
	y, x := win.MaxYX()
	gameBoard := win.Derived(y, x-2, 0, 1)
	gameBoard.SetBackground(gc.ColorPair(2) | gc.A_BOLD)
	gameBoard.Touch()
	gameBoard.Sync(gc.SYNC_DOWN)
	return gameBoard
}

func updateField(win *gc.Window, world *game.World) {
	board := world.CurrentGen()
	win.Color(2)
	win.Erase()
	win.SetBackground(gc.ColorPair(2) | gc.A_BOLD)
	y, x := win.MaxYX()
	y, x = y-2, x-2
	for i := 0; i <= x; i++ {
		for j := 0; j <= y; j++ {
			win.Move(j, i)
			if board.Get(i, j) {
				win.AddChar(gc.ACS_BOARD)
			}
		}
	}
	win.NoutRefresh()
}

func updateFooter(win *gc.Window, world *game.World, y, x int) {
	win.Erase()
	_, cols := win.MaxYX()

	// func (w *Window) Border(ls, rs, ts, bs, tl, tr, bl, br Char) error
	win.Border(gc.ACS_VLINE, gc.ACS_VLINE, gc.ACS_HLINE, gc.ACS_HLINE, gc.ACS_VLINE, gc.ACS_VLINE, gc.ACS_LLCORNER, gc.ACS_LRCORNER)
	win.MovePrint(1, 3, fmt.Sprintf("Generation: %d", world.Generation()))
	win.MovePrint(1, cols/2, fmt.Sprintf("Size (0, 0) -> (%d, %d)", x, y))
	win.NoutRefresh()
}

func setupFooter(win *gc.Window) {
	win.Erase()
	_, x := win.MaxYX()

	// func (w *Window) Border(ls, rs, ts, bs, tl, tr, bl, br Char) error
	win.Border(gc.ACS_VLINE, gc.ACS_VLINE, gc.ACS_HLINE, gc.ACS_HLINE, gc.ACS_VLINE, gc.ACS_VLINE, gc.ACS_LLCORNER, gc.ACS_LRCORNER)
	win.MovePrint(1, 3, "Generation: 0")
	win.MovePrint(1, x/2, "Size x,x -> y,y")
}

/* want the following

   ┌────────────────────┐
   │        TITLE       │
   ├────────────────────┤
   │                    │
   │                    │
   │                    │
   │                    │
   │                    │
   ├────────────────────┤
   │ G:1 (0,0)->(15,15) │
   └────────────────────┘
*/
func main() {
	f, err := os.Create("err.log")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	log.SetOutput(f)

	var stdscrn *gc.Window
	stdscrn, err = gc.Init()
	if err != nil {
		log.Println("Failed to init screen", err)
	}
	defer gc.End()

	rand.Seed(time.Now().Unix())
	gc.StartColor()

	// this has to be after the StartColor, or it breaks.
	var title, field, footer *gc.Window

	gc.InitPair(2, gc.C_YELLOW, gc.C_BLUE)

	// No echo or visiable stuff
	gc.Echo(false)
	gc.CBreak(true)
	gc.Cursor(0)

	stdscrn.Keypad(true)
	rows, cols := stdscrn.MaxYX()

	title, err = gc.NewWindow(TITLE_HEIGHT, cols, 0, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer title.Delete()

	field, err = gc.NewWindow(
		rows-(TITLE_HEIGHT+FOOTER_HEIGHT),
		cols,
		TITLE_HEIGHT, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer field.Delete()

	log.Printf("rows-(T+F): %d\n", rows-(TITLE_HEIGHT+FOOTER_HEIGHT))
	log.Printf("rows:       %d\n", rows)
	log.Printf("cols:       %d\n", cols)
	footer, err = gc.NewWindow(
		FOOTER_HEIGHT,
		cols,
		rows-FOOTER_HEIGHT,
		0)
	if err != nil {
		log.Fatal(err)
	}
	defer footer.Delete()

	setupTitle(title)
	gameBoard := setupField(field)
	setupFooter(footer)

	stdscrn.NoutRefresh()

	gc.Update()

	if err != nil {
		log.Fatal(err)
	}

	// setup world.
	startBoard := game.NewBoard()
	game.Glider(startBoard, 0, 0)
	game.Glider(startBoard, 5, 0)
	game.Glider(startBoard, 10, 0)
	game.Glider(startBoard, 15, 0)

	game.LWSS(startBoard, 0, 5)

	world := game.NewWorld(*startBoard)

	var y, x int
main:
	for {
		// Clear the section of screen where the box is currently located so
		// that it is blanked by calling Erase on the window and refreshing it
		// so that the chances are sent to the virtual screen but not actually
		// output to the terminal

		title.NoutRefresh()
		updateField(gameBoard, world)
		y, x = gameBoard.MaxYX()
		updateFooter(footer, world, y, x)
		world.Next()
		gc.Update()
		switch field.GetChar() {
		case 0:
			continue main
		case 'q':
			break main
		}
	}
}
