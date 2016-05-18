package main

import (
	"log"
	"os"

	gc "github.com/rthornton128/goncurses"
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
}

func setupField(win *gc.Window) {
	win.Erase()

	win.Border(gc.ACS_VLINE, gc.ACS_VLINE, ' ', ' ', gc.ACS_VLINE, gc.ACS_VLINE, gc.ACS_VLINE, gc.ACS_VLINE)
	win.MovePrint(2, 2, "bob bob")
}

func setupFooter(win *gc.Window) {
	win.Erase()
	_, x := win.MaxYX()
	win.Border(gc.ACS_VLINE, gc.ACS_VLINE, gc.ACS_HLINE, gc.ACS_HLINE, gc.ACS_VLINE, gc.ACS_VLINE, gc.ACS_VLINE, gc.ACS_VLINE)
	win.MovePrint(2, 2, "bob bob")
	win.MovePrint(1, 3, "Generation: 0")
	win.MovePrint(1, x/2, "Size x,x -> y,y")
}

func main() {
	f, err := os.Create("err.log")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if err != nil {
		log.Fatal(err)
	}

	var title, field, footer *gc.Window

	log.SetOutput(f)

	var stdscrn *gc.Window
	stdscrn, err = gc.Init()

	if err != nil {
		log.Println("Failed to init screen", err)
	}
	defer gc.End()

	// No echo or visiable stuff
	gc.Echo(false)
	gc.CBreak(true)
	gc.Cursor(0)

	stdscrn.Keypad(true)
	cols, rows := stdscrn.MaxYX()

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

	title, err = gc.NewWindow(TITLE_HEIGHT, rows, 0, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer title.Delete()

	field, err = gc.NewWindow(
		cols-(TITLE_HEIGHT+FOOTER_HEIGHT),
		rows,
		TITLE_HEIGHT, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer field.Delete()

	footer, err = gc.NewWindow(
		3,
		rows,
		rows-3, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer footer.Delete()

	setupTitle(title)
	setupField(field)
	setupFooter(footer)

	stdscrn.NoutRefresh()

	gc.Update()

	if err != nil {
		log.Fatal(err)
	}
main:
	for {
		// Clear the section of screen where the box is currently located so
		// that it is blanked by calling Erase on the window and refreshing it
		// so that the chances are sent to the virtual screen but not actually
		// output to the terminal
		title.NoutRefresh()
		field.NoutRefresh()
		footer.NoutRefresh()
		gc.Update()
		switch stdscrn.GetChar() {
		case 'q':
			break main
		}
	}
}
