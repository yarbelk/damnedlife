package main

import (
	"fmt"
	"time"

	"github.com/yarbelk/cgol2/game"
)

func main() {
	startBoard := game.NewBoard()
	game.Glider(startBoard, 0, 0)
	game.Glider(startBoard, 5, 0)
	game.Glider(startBoard, 10, 0)
	game.Glider(startBoard, 15, 0)

	game.LWSS(startBoard, 0, 5)

	world := game.NewWorld(*startBoard)

	for {
		fmt.Println(world)

		world.Next()
		time.Sleep(100 * time.Millisecond)
	}
}
