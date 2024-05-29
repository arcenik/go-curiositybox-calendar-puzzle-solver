package board

import (
	"encoding/json"
	"errors"
	"fmt"
)

const (
	SolutionLen = 8
	SolutionMax = SolutionLen - 1
)

type SolutionElement struct {
	Rotation int8  `json:"rotation"`
	Y        uint8 `json:"y"`
	X        uint8 `json:"x"`
	OnBoard  bool  `json:"-"`
}

type Solution [SolutionLen]SolutionElement

func NewSolution(pieces ...SolutionElement) (res Solution, err error) {
	res = Solution{
		{-1, 0, 0, false},
		{-1, 0, 0, false},
		{-1, 0, 0, false},
		{-1, 0, 0, false},
		{-1, 0, 0, false},
		{-1, 0, 0, false},
		{-1, 0, 0, false},
		{-1, 0, 0, false},
	}
	if len(pieces) > 8 {
		return Solution{}, errors.New("solution cannot contains more than 8 pieces")
	}
	for i, e := range pieces {
		res[i] = e
	}
	return res, nil
}

func NewSolutionElement(r int8, y, z uint8, onBoard bool) (res SolutionElement) {
	res = SolutionElement{r, y, z, onBoard}
	return res
}

func (s Solution) LastPieceOnBoard() (res int8) {
	for i, e := range s {
		if !e.OnBoard {
			return int8(i - 1)
		}
	}
	return SolutionMax
}

func (e *SolutionElement) Increment(maxRotation int8, maxY, maxX uint8) (carryOver bool) {
	if e.X+1 < maxX {
		e.X++
		return false
	} else {
		e.X = 0
		if e.Y+1 < maxY {
			e.Y++
			return false
		} else {
			e.Y = 0
			if e.Rotation+1 < maxRotation {
				e.Rotation++
				return false
			} else {
				e.Rotation = -1
				return true
			}
		}
	}
}

func (b *Board) SolutionWalker(s *Solution) (canContinue bool) {
	index := s.LastPieceOnBoard() + 1
	if b.LastPutSucceeded { // lastPut succeeded -> add a new element
		s[index] = NewSolutionElement(0, 0, 0, false)
	} else { // lastPut failed -> increment
		for s[index].Increment(int8(len(PiecesRotations[index])), 7, 7) {
			s[index] = NewSolutionElement(-1, 0, 0, false)
			index--
			if index >= 0 {
				s[index].OnBoard = false
			}
			b.LoadSolution(*s)
			if index < 0 { // carry over for 1st piece -> end of game
				return false
			}
		}
	}
	// try to put the piece on the board
	pieceId := PiecesRotations[index][s[index].Rotation]
	putResult, err := b.PutPiece(pieceId, s[index].Y, s[index].X)
	if err != nil {
		panic(fmt.Sprintln("PutPiece returned an error", err))
	}
	b.LastPutSucceeded = putResult
	s[index].OnBoard = putResult
	if (putResult) && (index == 7) {
		b.Completed = true
	}
	return true
}

// Print the solution
func (b Board) PrintSolution(s Solution) {
	fmt.Println("// Solution{")
	for i, e := range s {
		if e.Rotation == -1 {
			break
		}
		fmt.Printf("// %d:  %v, %v, %v, %v\n", i, e.Rotation, e.Y, e.X, e.OnBoard)
	}
	fmt.Println("// }")
}

// LoadSolution load the solution on the board
func (b *Board) LoadSolution(sol Solution) error {
	newBoard := NewBoard()
	for i, e := range sol {
		newBoard.LastPutSucceeded = e.OnBoard
		if e.Rotation < 0 || !e.OnBoard {
			break
		}
		pieceId := PiecesRotations[i][e.Rotation]
		ok, err := newBoard.PutPiece(pieceId, e.Y, e.X)
		if err != nil {
			return fmt.Errorf("error on piece %v: %v", i, err)
		}
		if !ok {
			return fmt.Errorf("piece %v was not put", i)
		}
		newBoard.PutPiece(pieceId, e.Y, e.X)
	}
	b.Board = newBoard.Board
	b.LastPutSucceeded = newBoard.LastPutSucceeded
	return nil
}

// SolutionBootstrap put the first piece on an empty board
func (b *Board) SolutionBootstrap(s *Solution) {
	s[0] = NewSolutionElement(0, 0, 0, true)
	b.LastPutSucceeded = true
}

// Return the solution as JSON string
func (s Solution) ToJSON() (string, error) {
	bytes, err := json.Marshal(s)
	return string(bytes), err
}
