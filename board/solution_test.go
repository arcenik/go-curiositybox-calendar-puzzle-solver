package board_test

import (
	"testing"

	"github.com/arcenik/go-curiositybox-calendar-puzzle-solver/board"
	"github.com/stretchr/testify/assert"
)

func TestNewSolution(t *testing.T) {
	s, err := board.NewSolution()
	expectedNew := board.Solution{
		{-1, 0, 0, false},
		{-1, 0, 0, false},
		{-1, 0, 0, false},
		{-1, 0, 0, false},
		{-1, 0, 0, false},
		{-1, 0, 0, false},
		{-1, 0, 0, false},
		{-1, 0, 0, false},
	}
	assert.Equal(t, expectedNew, s)
	assert.Nil(t, err, "NewSolution() should not return an error")

	s, err = board.NewSolution(
		board.NewSolutionElement(0, 0, 0, true),
		board.NewSolutionElement(0, 1, 0, true),
		board.NewSolutionElement(0, 2, 0, true),
		board.NewSolutionElement(0, 3, 0, true),
		board.NewSolutionElement(0, 4, 0, true),
		board.NewSolutionElement(0, 0, 1, true),
		board.NewSolutionElement(0, 1, 1, true),
		board.NewSolutionElement(0, 2, 1, true),
	)
	expectedOk := board.Solution{
		{0, 0, 0, true},
		{0, 1, 0, true},
		{0, 2, 0, true},
		{0, 3, 0, true},
		{0, 4, 0, true},
		{0, 0, 1, true},
		{0, 1, 1, true},
		{0, 2, 1, true},
	}
	assert.Equal(t, expectedOk, s)
	assert.Nil(t, err, "NewSolution() with 8 arguments should not return an error")

	_, err = board.NewSolution(
		board.NewSolutionElement(0, 0, 0, true),
		board.NewSolutionElement(0, 1, 0, true),
		board.NewSolutionElement(0, 2, 0, true),
		board.NewSolutionElement(0, 3, 0, true),
		board.NewSolutionElement(0, 4, 0, true),
		board.NewSolutionElement(0, 0, 1, true),
		board.NewSolutionElement(0, 1, 1, true),
		board.NewSolutionElement(0, 2, 1, true),
		board.NewSolutionElement(0, 3, 1, true),
	)
	assert.EqualError(t, err, "solution cannot contains more than 8 pieces")
}

func TestSolutionBootstrap(t *testing.T) {
	b := board.NewBoard()
	s, _ := board.NewSolution()
	b.PutPiece(66, 3, 0)
	b.SolutionBootstrap(&s)
	expected := board.Solution{
		{0, 0, 0, true},
		{-1, 0, 0, false},
		{-1, 0, 0, false},
		{-1, 0, 0, false},
		{-1, 0, 0, false},
		{-1, 0, 0, false},
		{-1, 0, 0, false},
		{-1, 0, 0, false},
	}
	assert.Equal(t, expected, s)
}

func TestLoadSolution(t *testing.T) {

	expected := [board.NumLines][board.NumColls]uint8{
		{22, 22, 34, 11, 0, 11, 99},
		{22, 22, 34, 11, 11, 11, 99},
		{22, 22, 34, 34, 34, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 99, 99, 99, 99},
	}

	// working solution
	b := board.NewBoard()
	s, err := board.NewSolution(
		board.NewSolutionElement(0, 0, 3, true),
		board.NewSolutionElement(1, 0, 0, true),
		board.NewSolutionElement(3, 0, 2, true),
		board.NewSolutionElement(0, 0, 0, false),
	)
	assert.Nil(t, err, "NewSolution should not return an error")
	err = b.LoadSolution(s)
	assert.Nil(t, err, "LoadSolution should not return an error")
	assert.False(t, b.LastPutSucceeded, "b.LastPutSucceeded should be false after LoadSolution")
	assert.Equal(t, expected, b.Board)

	// out of board limits
	b = board.NewBoard()
	s, err = board.NewSolution(
		board.NewSolutionElement(0, 0, 3, true),
		board.NewSolutionElement(1, 0, 10, true),
		board.NewSolutionElement(3, 0, 2, true),
	)
	assert.Nil(t, err, "NewSolution should not return an error")
	err = b.LoadSolution(s)
	assert.EqualError(t, err, "error on piece 1: position 0,10 is off limits")

	// piece overlap
	b = board.NewBoard()
	s, err = board.NewSolution(
		board.NewSolutionElement(0, 0, 3, true),
		board.NewSolutionElement(1, 0, 0, true),
		board.NewSolutionElement(3, 0, 2, true),
		board.NewSolutionElement(0, 0, 0, true),
	)
	assert.Nil(t, err, "NewSolution should not return an error")

	err = b.LoadSolution(s)
	assert.EqualError(t, err, "piece 3 was not put")
}

func TestPrintSolution(t *testing.T) {
	b := board.NewBoard()
	s, err := board.NewSolution(
		board.NewSolutionElement(0, 0, 3, true),
		board.NewSolutionElement(1, 0, 0, true),
		board.NewSolutionElement(3, 0, 2, true),
	)
	assert.Nil(t, err, "NewSolution should not return an error")
	b.PrintSolution(s)

	// FIXME test print output
}

func TestLastPieceOnBoard(t *testing.T) {
	s, err := board.NewSolution(
		board.NewSolutionElement(0, 0, 3, true),
		board.NewSolutionElement(1, 0, 0, true),
		board.NewSolutionElement(3, 0, 2, true),
	)
	assert.Nil(t, err, "NewSolution should not return an error")
	res := s.LastPieceOnBoard()
	assert.Equal(t, res, int8(2))

	s, err = board.NewSolution(
		board.NewSolutionElement(0, 0, 0, true),
		board.NewSolutionElement(0, 1, 0, true),
		board.NewSolutionElement(0, 2, 0, true),
		board.NewSolutionElement(0, 3, 0, true),
		board.NewSolutionElement(0, 4, 0, true),
		board.NewSolutionElement(0, 0, 1, true),
		board.NewSolutionElement(0, 1, 1, true),
		board.NewSolutionElement(0, 2, 1, true),
	)
	assert.Nil(t, err, "NewSolution should not return an error")
	res = s.LastPieceOnBoard()
	assert.Equal(t, res, int8(7))
}

func TestIncrement(t *testing.T) {
	// increment x
	se := board.NewSolutionElement(0, 0, 3, false)
	expected := board.NewSolutionElement(0, 0, 4, false)
	res := se.Increment(2, board.NumLines, board.NumColls)
	assert.False(t, res)
	assert.Equal(t, expected, se)

	// increment y
	se = board.NewSolutionElement(0, 0, 6, false)
	expected = board.NewSolutionElement(0, 1, 0, false)
	res = se.Increment(2, board.NumLines, board.NumColls)
	assert.False(t, res)
	assert.Equal(t, expected, se)

	// increment rotation
	se = board.NewSolutionElement(0, 6, 6, false)
	expected = board.NewSolutionElement(1, 0, 0, false)
	res = se.Increment(2, board.NumLines, board.NumColls)
	assert.False(t, res)
	assert.Equal(t, expected, se)

	// carry over
	se = board.NewSolutionElement(1, 6, 6, false)
	expected = board.NewSolutionElement(-1, 0, 0, false)
	res = se.Increment(2, board.NumLines, board.NumColls)
	assert.True(t, res)
	assert.Equal(t, expected, se)
}

func TestSolutionWalker(t *testing.T) {

	b := board.NewBoard()
	s, err := board.NewSolution()
	assert.Nil(t, err, "NewSolution should not return an error")
	b.SolutionBootstrap(&s)

	res := b.SolutionWalker(&s)
	assert.True(t, res, "SolutionWalker should return true")

	// 0:  2, 0, 0, true
	// 1:  0, 0, 3, true
	// 2:  2, 4, 0, true
	// 3:  2, 2, 3, true
	// 4:  0, 3, 4, true
	// 5:  3, 3, 0, true
	// 6:  4, 1, 0, true
	// 7:  2, 2, 5, true
	s, err = board.NewSolution(
		board.NewSolutionElement(2, 0, 0, true),
		board.NewSolutionElement(0, 0, 3, true),
		board.NewSolutionElement(2, 4, 0, true),
		board.NewSolutionElement(2, 2, 3, true),
		board.NewSolutionElement(0, 3, 4, true),
		board.NewSolutionElement(3, 3, 0, true),
		board.NewSolutionElement(4, 1, 0, true),
		board.NewSolutionElement(2, 2, 3, false), // 2 steps from solution
	)
	assert.Nil(t, err, "NewSolution should not return an error")

	expectedStep1, err := board.NewSolution(
		board.NewSolutionElement(2, 0, 0, true),
		board.NewSolutionElement(0, 0, 3, true),
		board.NewSolutionElement(2, 4, 0, true),
		board.NewSolutionElement(2, 2, 3, true),
		board.NewSolutionElement(0, 3, 4, true),
		board.NewSolutionElement(3, 3, 0, true),
		board.NewSolutionElement(4, 1, 0, true),
		board.NewSolutionElement(2, 2, 4, false), // 1 steps from solution
	)
	assert.Nil(t, err, "NewSolution should not return an error")

	expectedStep2, err := board.NewSolution(
		board.NewSolutionElement(2, 0, 0, true),
		board.NewSolutionElement(0, 0, 3, true),
		board.NewSolutionElement(2, 4, 0, true),
		board.NewSolutionElement(2, 2, 3, true),
		board.NewSolutionElement(0, 3, 4, true),
		board.NewSolutionElement(3, 3, 0, true),
		board.NewSolutionElement(4, 1, 0, true),
		board.NewSolutionElement(2, 2, 5, true), // actual solution
	)
	assert.Nil(t, err, "NewSolution should not return an error")

	err = b.LoadSolution(s)
	assert.Nil(t, err, "LoadSolution should not return an error")
	assert.False(t, b.LastPutSucceeded, "b.LastPutSucceeded should be false after LoadSolution")

	res = b.SolutionWalker(&s)
	assert.True(t, res, "SolutionWalker should return true")
	assert.Equal(t, expectedStep1, s)

	res = b.SolutionWalker(&s)
	assert.True(t, res, "SolutionWalker should return true")
	assert.Equal(t, expectedStep2, s)
	assert.True(t, b.Completed)

	// TODO: test when SolutionWalker return false because the path is over
	// TODO: test when SolutionWalker panic because PutPiece returned an error

}

func TestToJSON(t *testing.T) {
	expected := `[{"rotation":2,"y":0,"x":0},{"rotation":0,"y":0,"x":3},{"rotation":2,"y":4,"x":0},{"rotation":2,"y":2,"x":3},{"rotation":0,"y":3,"x":4},{"rotation":3,"y":3,"x":0},{"rotation":4,"y":1,"x":0},{"rotation":2,"y":2,"x":5}]`
	s, err := board.NewSolution(
		board.NewSolutionElement(2, 0, 0, true),
		board.NewSolutionElement(0, 0, 3, true),
		board.NewSolutionElement(2, 4, 0, true),
		board.NewSolutionElement(2, 2, 3, true),
		board.NewSolutionElement(0, 3, 4, true),
		board.NewSolutionElement(3, 3, 0, true),
		board.NewSolutionElement(4, 1, 0, true),
		board.NewSolutionElement(2, 2, 5, true),
	)
	assert.Nil(t, err, "NewSolution should not return an error")
	result, err := s.ToJSON()
	assert.Nil(t, err, "ToJSON should not return an error")
	assert.Equal(t, expected, result)
}
