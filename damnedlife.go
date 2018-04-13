package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	gc "github.com/rthornton128/goncurses"
	"github.com/yarbelk/damnedlife/game"
)

const (
	// ALIVE is the symbole for alive
	ALIVE = '#'
)

// GenerationAge is how long a generation lasts in terms
// of clock wall time.  It is optimistic, and as the complexity increases,
// it will be longer
const GenerationAge = time.Millisecond * 50

const (
	// TitleHeight sets the ncurses title height.  This shows
	// some information
	TitleHeight = 5
	// FooterHeight sets the ncurses footer height. this
	// shows some information
	FooterHeight = 3
)

// ncurses version; thus the 'damned' part of the life

// setupTitle for inital drawing etc.
func setupTitle(win *gc.Window) {
	win.Erase()
	// func (w *Window) Border(ls, rs, ts, bs, tl, tr, bl, br Char) error
	win.Border(gc.ACS_VLINE, gc.ACS_VLINE, gc.ACS_HLINE, gc.ACS_HLINE, gc.ACS_ULCORNER, gc.ACS_URCORNER, gc.ACS_LTEE, gc.ACS_RTEE)
	win.Keypad(false)
	_, x := win.MaxYX()
	title := "Conways Game of Life"
	win.MovePrint(2, (x/2 - len(title)/2), title)
	win.MovePrint(3, (x/2 - len(title)/2), "(press Q to exit; hjkl to move), 0 to go to origin")
}

func updateTitle(win *gc.Window) {
	win.NoutRefresh()
}

// setupField to handle the actual cellular atomita.  Returns a Derived
// window, which is used to update the actual cells.
func setupField(win *gc.Window) *gc.Window {
	win.Color(2)
	win.Erase()
	// func (w *Window) Border(ls, rs, ts, bs, tl, tr, bl, br Char) error
	win.Border(gc.ACS_VLINE, gc.ACS_VLINE, ' ', ' ', gc.ACS_VLINE, gc.ACS_VLINE, gc.ACS_VLINE, gc.ACS_VLINE)
	y, x := win.MaxYX()
	gameBoard := win.Derived(y, x-2, 0, 1)
	gameBoard.SetBackground(gc.ColorPair(2) | gc.A_BOLD)
	gameBoard.Touch()
	gameBoard.Sync(gc.SYNC_DOWN)
	return gameBoard
}

// updateField with the new game world state.
func updateField(win *gc.Window, world *game.World, originY, originX int) {
	board := world.CurrentGen()
	win.Erase()
	y, x := win.MaxYX()
	y, x = y-2, x-2
	for i := 0; i <= x; i++ {
		for j := 0; j <= y; j++ {
			win.Move(j, i)
			if board.Get(i+originX, j+originY) {
				win.AddChar(ALIVE)
			}
		}
	}
	win.NoutRefresh()
}

// updateFooter with new window extents and current generation
func updateFooter(win *gc.Window, world *game.World, originY, originX, y, x int) {
	win.Erase()
	_, cols := win.MaxYX()

	// func (w *Window) Border(ls, rs, ts, bs, tl, tr, bl, br Char) error
	win.Border(gc.ACS_VLINE, gc.ACS_VLINE, gc.ACS_HLINE, gc.ACS_HLINE, gc.ACS_VLINE, gc.ACS_VLINE, gc.ACS_LLCORNER, gc.ACS_LRCORNER)
	win.MovePrint(1, 3, fmt.Sprintf("Generation: %d", world.Generation()))
	win.MovePrint(1, cols/2, fmt.Sprintf("Size (%d, %d) -> (%d, %d)", originX, originY, originX+x, originY+y))
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

// generationTimer waits for each tick of 100ms, updates the world state, and sends out
// a update command to the redraw (async`ly)
func generationTimer(world *game.World, newGeneration chan<- bool, quit <-chan bool, worldLocker *sync.RWMutex) {
	for {
		select {
		case <-time.After(GenerationAge):
			worldLocker.Lock()
			world.Next()
			worldLocker.Unlock()
			newGeneration <- true
			continue
		case <-quit:
			close(newGeneration)
			log.Println("[GT] leaving timer")
			return
		}
	}
}

func keyPresses(field *gc.Window, originMoved chan<- game.Point, quit chan bool) {
	origin := game.Point{X: 0, Y: 0}
keys:
	for {
		// get a char, flush input when you do get one to prevent being blocked
		// by a huge pipe of chars waiting to be processed when you hold down
		// a key
		switch field.GetChar() {
		case 'h':
			origin.X--
			originMoved <- origin
			continue keys
		case 'j':
			origin.Y++
			originMoved <- origin
			continue keys
		case 'k':
			origin.Y--
			originMoved <- origin
			continue keys
		case 'l':
			origin.X++
			originMoved <- origin
			continue keys
		case '0':
			origin.X, origin.Y = 0, 0
			originMoved <- origin
			continue keys
		case 'q':
			close(originMoved)
			close(quit)
			log.Println("[KP] leaving keyPresses, closed the channels")
			return
		case 0:
			select {
			case <-time.After(time.Millisecond * 1):
				continue keys
			case <-quit:
				return
			}
		}
	}
}

func redrawConsumer(title, gameBoard, footer *gc.Window, world *game.World, originMoved <-chan game.Point, newGeneration <-chan bool, worldLocker *sync.RWMutex) {
	redrawScreen := func(origin game.Point) {
		worldLocker.RLock()
		defer worldLocker.RUnlock()
		updateTitle(title)
		updateField(gameBoard, world, origin.Y, origin.X)
		boardRows, boardCols := gameBoard.MaxYX()
		updateFooter(footer, world, origin.Y, origin.X, boardRows, boardCols)
		gc.Update()
	}

	origin := game.Point{X: 0, Y: 0}
	redrawScreen(origin)
redraw:
	for {
		if originMoved == nil || newGeneration == nil {
			log.Println("[RC] Cowardly not litening to nil channels")
			return
		}
		select {
		case newOrigin, ok := <-originMoved:
			if !ok {
				log.Println("[RC] leaving redraw (originMoved not ok)")
				return
			}
			origin = newOrigin
			redrawScreen(newOrigin)
			continue redraw
		case _, ok := <-newGeneration:
			if !ok {
				log.Println("[RC] leaving redraw (newGeneration not ok)")
				return
			}
			redrawScreen(origin)
			continue redraw
		}
	}

}

func catchSigInt(quit chan bool) {
	var c chan os.Signal
	c = make(chan os.Signal)
	gc.CBreak(true)
	signal.Notify(c, syscall.SIGINT)
	for {
		<-c
		log.Println("trying to quit")
		close(quit)
		return
	}
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

	title, err = gc.NewWindow(TitleHeight, cols, 0, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer title.Delete()

	field, err = gc.NewWindow(
		rows-(TitleHeight+FooterHeight),
		cols,
		TitleHeight, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer field.Delete()

	footer, err = gc.NewWindow(
		FooterHeight,
		cols,
		rows-FooterHeight,
		0)
	if err != nil {
		log.Fatal(err)
	}
	defer footer.Delete()

	setupTitle(title)
	gameBoard := setupField(field)
	defer gameBoard.Delete()
	setupFooter(footer)

	stdscrn.NoutRefresh()

	gc.Update()

	if err != nil {
		log.Fatal(err)
	}

	// setup world.
	startBoard := game.NewBoard()
	// game.Glider(startBoard, 0, 0)
	// game.Glider(startBoard, 5, 0)
	// game.Glider(startBoard, 10, 0)
	// game.Glider(startBoard, 15, 0)

	// game.LWSS(startBoard, 0, 5)

	game.GliderGun(startBoard, 2, 2)

	world := game.NewWorld(*startBoard)

	var originMoved = make(chan game.Point)
	var newGeneration = make(chan bool)
	var quit = make(chan bool)
	var worldLocker = &sync.RWMutex{}

	go catchSigInt(quit)
	go keyPresses(field, originMoved, quit)                                                  // producter of origini change commands, and quit.  closes originMoved when done
	go generationTimer(world, newGeneration, quit, worldLocker)                              // produces ticks on the newGeneration channel, waits for a quit command to end. closes newGeneration when done
	redrawConsumer(title, gameBoard, footer, world, originMoved, newGeneration, worldLocker) // listens to orignMoved and newGeneration, quits when both are done.
}
