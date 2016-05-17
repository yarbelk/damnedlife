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

	board.Set(1, 1, true)
	Expect(board.Get(1, 1)).To(BeTrue())
}

func TestDiesWithLessThanTwoNeighbors(t *testing.T) {
	RegisterTestingT(t)
	var board = game.NewBoard()

	board.Set(1, 1, true)
	failOnAlive(board, t)

	board.Set(0, 0, true)
	failOnAlive(board, t)
}

func TestDiesWithOneNeighbor(t *testing.T) {
	RegisterTestingT(t)
	var board = game.NewBoard()
	board.Set(1, 1, true)

	board.Set(0, 0, true)

	failOnAlive(board, t)
}

func TestLivesWithTwoOrThreeNeighbor(t *testing.T) {
	RegisterTestingT(t)
	var board = game.NewBoard()
	board.Set(1, 1, true)

	board.Set(0, 0, true)
	board.Set(0, 1, true)

	failOnDead(board, t)

	board.Set(0, 2, true)

	failOnDead(board, t)
}

func TestDiesWhenMorThanThreeNeighbors(t *testing.T) {
	RegisterTestingT(t)
	var board = game.NewBoard()
	board.Set(1, 1, true)

	board.Set(0, 0, true)
	board.Set(0, 1, true)
	board.Set(0, 2, true)
	board.Set(1, 0, true)

	failOnAlive(board, t)

	board.Set(1, 2, true)
	failOnAlive(board, t)

	board.Set(2, 0, true)
	failOnAlive(board, t)

	board.Set(2, 1, true)
	failOnAlive(board, t)

	board.Set(2, 2, true)
	failOnAlive(board, t)
}

func TestNoNewCellWithTwoNeighbors(t *testing.T) {
	RegisterTestingT(t)
	var board = game.NewBoard()

	board.Set(0, 0, true)
	board.Set(0, 1, true)

	failOnAlive(board, t)
}

func TestNewCellWhenThreeNeighbors(t *testing.T) {
	RegisterTestingT(t)
	var board = game.NewBoard()

	board.Set(0, 0, true)
	board.Set(0, 1, true)
	board.Set(0, 2, true)

	failOnDead(board, t)
}

func TestGetBoardLimits(t *testing.T) {
	RegisterTestingT(t)
	var board = game.NewBoard()
	var topLeft, bottomRight game.Point

	board.Set(-10, -10, true)
	board.Set(10, 10, true)

	topLeft, bottomRight = board.GetLimits()

	Expect(topLeft).To(Equal(game.Point{-10, -10}))
	Expect(bottomRight).To(Equal(game.Point{10, 10}))
}
