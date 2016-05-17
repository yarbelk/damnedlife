package game_test

import (
	"testing"

	. "github.com/onsi/gomega"
	"github.com/yarbelk/cgol2/game"
)

func failOnAlive(b *game.Board, t *testing.T) {
	nextState := b.NextState(1, 1)
	Expect(nextState).To(BeFalse())
}

func failOnDead(b *game.Board, t *testing.T) {
	nextState := b.NextState(1, 1)
	Expect(nextState).To(BeTrue())
}

func TestGetState(t *testing.T) {
	RegisterTestingT(t)
	var board = game.NewBoard()
	Expect(board.Get(1, 1)).To(BeFalse())

	board.SetAlive(1, 1)
	Expect(board.Get(1, 1)).To(BeTrue())
}

func TestDiesWithLessThanTwoNeighbors(t *testing.T) {
	RegisterTestingT(t)
	var board = game.NewBoard()

	board.SetAlive(1, 1)
	failOnAlive(board, t)

	board.SetAlive(0, 0)
	failOnAlive(board, t)
}

func TestDiesWithOneNeighbor(t *testing.T) {
	RegisterTestingT(t)
	var board = game.NewBoard()
	board.SetAlive(1, 1)

	board.SetAlive(0, 0)

	failOnAlive(board, t)
}

func TestLivesWithTwoOrThreeNeighbor(t *testing.T) {
	RegisterTestingT(t)
	var board = game.NewBoard()
	board.SetAlive(1, 1)

	board.SetAlive(0, 0)
	board.SetAlive(0, 1)

	failOnDead(board, t)

	board.SetAlive(0, 2)

	failOnDead(board, t)
}

func TestDiesWhenMorThanThreeNeighbors(t *testing.T) {
	RegisterTestingT(t)
	var board = game.NewBoard()
	board.SetAlive(1, 1)

	board.SetAlive(0, 0)
	board.SetAlive(0, 1)
	board.SetAlive(0, 2)
	board.SetAlive(1, 0)

	failOnAlive(board, t)

	board.SetAlive(1, 2)
	failOnAlive(board, t)

	board.SetAlive(2, 0)
	failOnAlive(board, t)

	board.SetAlive(2, 1)
	failOnAlive(board, t)

	board.SetAlive(2, 2)
	failOnAlive(board, t)
}

func TestNoNewCellWithTwoNeighbors(t *testing.T) {
	RegisterTestingT(t)
	var board = game.NewBoard()

	board.SetAlive(0, 0)
	board.SetAlive(0, 1)

	failOnAlive(board, t)
}

func TestNewCellWhenThreeNeighbors(t *testing.T) {
	RegisterTestingT(t)
	var board = game.NewBoard()

	board.SetAlive(0, 0)
	board.SetAlive(0, 1)
	board.SetAlive(0, 2)

	failOnDead(board, t)
}

func TestGetBoardLimits(t *testing.T) {
	RegisterTestingT(t)
	var board = game.NewBoard()
	var topLeft, bottomRight game.Point

	board.SetAlive(-10, -10)
	board.SetAlive(10, 10)

	topLeft, bottomRight = board.GetLimits()

	Expect(topLeft).To(Equal(game.Point{-10, -10}))
	Expect(bottomRight).To(Equal(game.Point{10, 10}))
}
