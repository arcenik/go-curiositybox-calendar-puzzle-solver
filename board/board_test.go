package board_test

import (
	"testing"

	"github.com/arcenik/go-curiositybox-calendar-puzzle-solver/board"
	"github.com/stretchr/testify/assert"
)

func TestBoardCanPutPiece(t *testing.T) {

	b := board.NewBoard()

	// should work
	ok, err := b.CanPutPiece(66, 4, 0)
	assert.True(t, ok, "b.CanPutPiece(66,4,0) should return (true, nil)")
	assert.Nil(t, err, "b.CanPutPiece(66,4,0) should return (true, nil)")

	// false: out of bound
	ok, err = b.CanPutPiece(66, 6, 0)
	assert.False(t, ok, "b.CanPutPiece(66, 6, 0) should return (false, nil)")
	assert.Nil(t, err)

	// false: out of bound
	ok, err = b.CanPutPiece(66, 10, 0)
	assert.False(t, ok, "b.CanPutPiece(66, 10, 0) should return (false, error)")
	assert.EqualError(t, err, "position 10,0 is off limits")

	// false: out of bound
	ok, err = b.CanPutPiece(66, 2, 10)
	assert.False(t, ok, "b.CanPutPiece(66, 2, 10) should return (false, error)")
	assert.EqualError(t, err, "position 2,10 is off limits")

	// false: out of bound
	ok, err = b.CanPutPiece(95, 2, 2)
	assert.False(t, ok, "b.CanPutPiece(95, 2, 2) should return (false, error)")
	assert.EqualError(t, err, "pieceId 95 does not exists")
}

func TestBoardPutPiece(t *testing.T) {

	b := board.NewBoard()

	// false: out of bound
	ok, err := b.PutPiece(66, 10, 0)
	assert.False(t, ok, "b.CanPutPiece(66, 10, 0) should return (false, error)")
	assert.EqualError(t, err, "position 10,0 is off limits")

	// false: wrong piece
	ok, err = b.PutPiece(95, 2, 2)
	assert.False(t, ok, "b.CanPutPiece(66, 6, 0) should return (false, error)")
	assert.EqualError(t, err, "pieceId 95 does not exists")

	// should work
	ok, err = b.PutPiece(66, 3, 0)
	expected := [board.NumLines][board.NumColls]uint8{
		{00, 00, 00, 00, 00, 00, 99},
		{00, 00, 00, 00, 00, 00, 99},
		{00, 00, 00, 00, 00, 00, 00},
		{00, 66, 66, 00, 00, 00, 00},
		{00, 66, 00, 00, 00, 00, 00},
		{66, 66, 00, 00, 00, 00, 00},
		{00, 00, 00, 99, 99, 99, 99},
	}
	assert.True(t, ok, "b.PutPiece(66,3,0) should return (true, nil)")
	assert.Nil(t, err, "b.PutPiece(66,3,0) should return (true, nil)")
	assert.Equal(t, expected, b.Board)

	// piece overlap
	ok, err = b.PutPiece(42, 3, 1)
	assert.False(t, ok, "b.PutPiece(66,3,0) should return (true, nil)")
	assert.Nil(t, err, "b.PutPiece(66,3,0) should return (true, nil)")
}

func TestBoardPrint(t *testing.T) {

	b := board.NewBoard()
	b.PutPiece(66, 3, 0)
	b.PrintBoard()

	// FIXME: test print output
}

func TestGetDate(t *testing.T) {

	// Solution{
	// 0:  0, 0, 0, true
	// 1:  0, 4, 4, true
	// 2:  2, 0, 1, true
	// 3:  2, 2, 0, true
	// 4:  4, 1, 5, true
	// 5:  0, 3, 1, true
	// 6:  1, 5, 0, true
	// 7:  6, 0, 4, true
	// }
	// ┌──────────────────────────────╖
	// │ ▓▓▓▓ Feb▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓     ║
	// │ ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓     ║
	// │ ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓ ║
	// │ ▓▓▓▓▓▓▓▓  10▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓ ║
	// │ ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓ ║
	// │ ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓ ║
	// │ ▓▓▓▓▓▓▓▓▓▓▓▓                 ║
	// ╘══════════════════════════════╝
	// the solution is not valid

	// working solution
	b := board.NewBoard()
	s, err := board.NewSolution(
		board.NewSolutionElement(0, 0, 0, true), // 0
		board.NewSolutionElement(0, 4, 4, true), // 1
		board.NewSolutionElement(2, 0, 1, true), // 2
		board.NewSolutionElement(2, 2, 0, true), // 3
		board.NewSolutionElement(4, 1, 5, true), // 4
		board.NewSolutionElement(0, 3, 1, true), // 5
		board.NewSolutionElement(1, 5, 0, true), // 6
		board.NewSolutionElement(6, 0, 4, true), // 7
	)
	assert.Nil(t, err, "NewSolution should not return an error")
	err = b.LoadSolution(s)
	b.Completed = true
	assert.Nil(t, err, "LoadSolution should not return an error")
	month, day, err := b.GetDate()
	assert.Nil(t, err)
	assert.Equal(t, 2, month)
	assert.Equal(t, 10, day)

	// Solution{
	// 0:  0, 0, 0, true
	// 1:  0, 4, 4, true
	// 2:  2, 0, 1, true
	// 3:  1, 2, 3, true
	// 4:  1, 3, 1, true
	// 5:  1, 0, 4, true
	// 6:  1, 5, 0, true
	// 7:  0, 2, 0, true
	// }
	// ┌──────────────────────────────╖
	// │ ▓▓▓▓ Feb▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓     ║
	// │ ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓ Nov▓▓▓▓     ║
	// │ ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓ ║
	// │ ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓ ║
	// │ ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓ ║
	// │ ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓ ║
	// │ ▓▓▓▓▓▓▓▓▓▓▓▓                 ║
	// ╘══════════════════════════════╝
	// the solution is not valid

	// 2 months
	b = board.NewBoard()
	s, err = board.NewSolution(
		board.NewSolutionElement(0, 0, 0, true), // 0
		board.NewSolutionElement(0, 4, 4, true), // 1
		board.NewSolutionElement(2, 0, 1, true), // 2
		board.NewSolutionElement(1, 2, 3, true), // 3
		board.NewSolutionElement(1, 3, 1, true), // 4
		board.NewSolutionElement(1, 0, 4, true), // 5
		board.NewSolutionElement(1, 5, 0, true), // 6
		board.NewSolutionElement(0, 2, 0, true), // 7
	)
	assert.Nil(t, err, "NewSolution should not return an error")
	err = b.LoadSolution(s)
	b.Completed = true
	assert.Nil(t, err, "LoadSolution should not return an error")
	month, day, err = b.GetDate()
	assert.EqualError(t, err, "the solution is not valid")
	assert.Equal(t, -1, month)
	assert.Equal(t, -1, day)

	b.Completed = false
	_, _, err = b.GetDate()
	assert.EqualError(t, err, "board.Completed is false")
}
