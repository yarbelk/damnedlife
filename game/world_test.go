package game_test

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/yarbelk/deadlife/game"
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
░░█`
	Expect(board.String()).To(Equal(expected))
}

/* Make sure that the following occures
a

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
	t.Log(fmt.Sprintf("genO\n\n%s\n\n", board.String()))

	var world = game.NewWorld(*board)
	world.Next()

	Expect(world.Generation()).To(Equal(1), "generation")
	genOne := world.CurrentGen()
	Expect(genOne.Get(1, 0)).To(BeTrue(), "%#v", game.Point{1, 0})
	Expect(genOne.Get(1, 1)).To(BeTrue(), "%#v", game.Point{1, 1})
	Expect(genOne.Get(1, 2)).To(BeTrue(), "%#v", game.Point{1, 2})
	t.Log(fmt.Sprintf("gen1\n\n%s\n\n", genOne.String()))

	world.Next()
	genTwo := world.CurrentGen()
	Expect(world.Generation()).To(Equal(2), "generation")
	Expect(genTwo.Get(0, 1)).To(BeTrue(), "%#v", game.Point{0, 1})
	Expect(genTwo.Get(1, 1)).To(BeTrue(), "%#v", game.Point{1, 1})
	Expect(genTwo.Get(2, 1)).To(BeTrue(), "%#v", game.Point{2, 1})
	t.Log(fmt.Sprintf("gen2\n\n%s\n\n", genTwo.String()))
}
