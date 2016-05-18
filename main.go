package main

import (
	"fmt"
	"time"

	"github.com/yarbelk/cgol2/game"
)

func main() {
	startBoard := game.NewBoard()
	startBoard.SetAlive(0, 0)
	startBoard.SetAlive(1, 1)
	startBoard.SetAlive(2, 1)
	startBoard.SetAlive(0, 2)
	startBoard.SetAlive(1, 2)

	world := game.NewWorld(*startBoard)

	for {
		fmt.Println(world)

		world.Next()
		time.Sleep(1 * time.Second)
	}
}
