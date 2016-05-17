package game_test

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/yarbelk/cgol2/game"
)

func TestWorldGetsGeneration(t *testing.T) {
	RegisterTestingT(t)
	var board = game.NewBoard()
	board.SetAlive(1, 1)

	var world = game.NewWorld(*board)
	world.Next()

	Expect(world.Generation()).To(Equal(1))

	current := world.CurrentGen()
	Expect(current.Get(1, 1)).To(BeFalse())
}

func TestWorldPrints(t *testing.T) {
	RegisterTestingT(t)
	var board = game.NewBoard()
	board.SetAlive(0, 0)
	board.SetAlive(2, 2)
	expected := `█░░
░░░
░░█
`
	Expect(board.String()).To(Equal(expected))
}

/* Make sure that the following occures

░░░    ░█░    ░░░    ░█░
███ => ░█░ => ███ => ░█░
░░░    ░█░    ░░░    ░█░
*/
func TestWorldGetsGenerationBlinker(t *testing.T) {
	RegisterTestingT(t)
	var board = game.NewBoard()
	board.SetAlive(0, 1)
	board.SetAlive(1, 1)
	board.SetAlive(2, 1)

	var world = game.NewWorld(*board)
	world.Next()

	Expect(world.Generation()).To(Equal(1))
	genOne := world.CurrentGen()
	Expect(genOne.Get(1, 0)).To(BeTrue())
	Expect(genOne.Get(1, 1)).To(BeTrue())
	Expect(genOne.Get(1, 2)).To(BeTrue())
	t.Log(fmt.Sprintf("\n\n%s\n\n", genOne.String()))

	Expect(world.Generation()).To(Equal(2))
	genTwo := world.CurrentGen()
	Expect(genTwo.Get(0, 1)).To(BeTrue())
	Expect(genTwo.Get(1, 1)).To(BeTrue())
	Expect(genTwo.Get(2, 1)).To(BeTrue())
	t.Log(fmt.Sprintf("\n\n%s\n\n", genTwo.String()))
}
