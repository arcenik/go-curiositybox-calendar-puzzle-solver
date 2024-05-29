package board

import (
	"errors"
	"fmt"

	"github.com/fatih/color"
)

type delta struct {
	y uint8
	x uint8
}

const (
	NumLines = 7
	NumColls = 7
)

var (
	back [NumLines][NumColls]string = [NumLines][NumColls]string{
		{"Jan", "Feb", "Mar", "Apr", "May", "Jun", ""},
		{"Jul", "Aug", "Sep", "Oct", "Nov", "Dec", ""},
		{" 1", " 2", " 3", " 4", " 5", " 6", " 7"},
		{" 8", " 9", "10", "11", "12", "13", "14"},
		{"15", "16", "17", "18", "19", "20", "21"},
		{"22", "23", "24", "25", "26", "27", "28"},
		{"29", "30", "31", "", "", "", ""},
	}
	monthDayNumber [NumLines][NumColls]int = [NumLines][NumColls]int{
		{1, 2, 3, 4, 5, 6, -1},
		{7, 8, 9, 10, 11, 12, -1},
		{1, 2, 3, 4, 5, 6, 7},
		{8, 9, 10, 11, 12, 13, 14},
		{15, 16, 17, 18, 19, 20, 21},
		{22, 23, 24, 25, 26, 27, 28},
		{29, 30, 31, -1, -1, -1, -1},
	}
	Pieces = map[uint8][]delta{
		11: {{0, 0}, {1, 2}, {0, 2}, {1, 1}, {1, 0}},
		12: {{0, 0}, {2, 1}, {2, 0}, {0, 1}, {1, 0}},
		13: {{0, 0}, {1, 2}, {0, 2}, {0, 1}, {1, 0}},
		14: {{0, 0}, {2, 1}, {2, 0}, {1, 1}, {0, 1}},
		21: {{0, 0}, {1, 2}, {0, 2}, {1, 1}, {0, 1}, {1, 0}},
		22: {{0, 0}, {2, 1}, {2, 0}, {1, 1}, {0, 1}, {1, 0}},
		31: {{0, 0}, {0, 2}, {2, 0}, {0, 1}, {1, 0}},
		32: {{0, 0}, {2, 2}, {1, 2}, {0, 2}, {0, 1}},
		33: {{2, 2}, {1, 2}, {0, 2}, {2, 1}, {2, 0}},
		34: {{0, 0}, {2, 2}, {2, 1}, {2, 0}, {1, 0}},
		41: {{3, 1}, {2, 1}, {2, 0}, {1, 1}, {0, 1}},
		42: {{1, 3}, {1, 2}, {1, 1}, {0, 1}, {1, 0}},
		43: {{3, 0}, {0, 0}, {2, 0}, {1, 1}, {1, 0}},
		44: {{0, 3}, {0, 0}, {1, 2}, {0, 2}, {0, 1}},
		45: {{3, 0}, {0, 0}, {2, 1}, {2, 0}, {1, 0}},
		46: {{0, 3}, {0, 0}, {0, 2}, {1, 1}, {0, 1}},
		47: {{3, 1}, {2, 1}, {1, 1}, {0, 1}, {1, 0}},
		48: {{1, 3}, {1, 2}, {0, 2}, {1, 1}, {1, 0}},
		51: {{2, 1}, {2, 0}, {1, 1}, {0, 1}, {1, 0}},
		52: {{0, 0}, {1, 2}, {1, 1}, {0, 1}, {1, 0}},
		53: {{0, 0}, {2, 0}, {1, 1}, {0, 1}, {1, 0}},
		54: {{0, 0}, {1, 2}, {0, 2}, {1, 1}, {0, 1}},
		55: {{0, 0}, {2, 1}, {2, 0}, {1, 1}, {1, 0}},
		56: {{0, 0}, {0, 2}, {1, 1}, {0, 1}, {1, 0}},
		57: {{0, 0}, {2, 1}, {1, 1}, {0, 1}, {1, 0}},
		58: {{1, 2}, {0, 2}, {1, 1}, {0, 1}, {1, 0}},
		61: {{1, 2}, {0, 2}, {2, 0}, {1, 1}, {1, 0}},
		62: {{0, 0}, {2, 2}, {2, 1}, {1, 1}, {0, 1}},
		65: {{0, 0}, {2, 2}, {1, 2}, {1, 1}, {1, 0}},
		66: {{0, 2}, {2, 1}, {2, 0}, {1, 1}, {0, 1}},
		71: {{0, 0}, {3, 1}, {2, 1}, {1, 1}, {1, 0}},
		72: {{0, 3}, {1, 2}, {0, 2}, {1, 1}, {1, 0}},
		73: {{0, 0}, {3, 1}, {2, 1}, {2, 0}, {1, 0}},
		74: {{0, 3}, {0, 2}, {1, 1}, {0, 1}, {1, 0}},
		75: {{3, 0}, {2, 0}, {1, 1}, {0, 1}, {1, 0}},
		76: {{0, 0}, {1, 3}, {1, 2}, {0, 2}, {0, 1}},
		77: {{3, 0}, {2, 1}, {2, 0}, {1, 1}, {0, 1}},
		78: {{0, 0}, {1, 3}, {1, 2}, {1, 1}, {0, 1}},
		81: {{3, 0}, {0, 0}, {3, 1}, {2, 0}, {1, 0}},
		82: {{0, 3}, {0, 0}, {0, 2}, {0, 1}, {1, 0}},
		83: {{0, 0}, {3, 1}, {2, 1}, {1, 1}, {0, 1}},
		84: {{0, 3}, {1, 3}, {1, 2}, {1, 1}, {1, 0}},
		85: {{3, 0}, {3, 1}, {2, 1}, {1, 1}, {0, 1}},
		86: {{0, 0}, {1, 3}, {1, 2}, {1, 1}, {1, 0}},
		87: {{3, 0}, {0, 0}, {2, 0}, {0, 1}, {1, 0}},
		88: {{0, 3}, {0, 0}, {1, 3}, {0, 2}, {0, 1}},
	}
	PiecesRotations = [][]uint8{
		{11, 12, 13, 14},
		{21, 22},
		{31, 32, 33, 34},
		{41, 42, 43, 44},
		{51, 52, 53, 54, 55, 56, 57, 58},
		{61, 62, 65, 66},
		{71, 72, 73, 74, 75, 76, 77, 78},
		{81, 82, 83, 84, 85, 86, 87, 88},
	}
	PiecesColors = map[uint8]color.Attribute{
		11: color.FgHiRed,
		12: color.FgHiRed,
		13: color.FgHiRed,
		14: color.FgHiRed,
		21: color.FgHiYellow,
		22: color.FgHiYellow,
		31: color.FgHiGreen,
		32: color.FgHiGreen,
		33: color.FgHiGreen,
		34: color.FgHiGreen,
		41: color.FgHiCyan,
		42: color.FgHiCyan,
		43: color.FgHiCyan,
		44: color.FgHiCyan,
		45: color.FgHiCyan,
		46: color.FgHiCyan,
		47: color.FgHiCyan,
		48: color.FgHiCyan,
		51: color.FgHiBlue,
		52: color.FgHiBlue,
		53: color.FgHiBlue,
		54: color.FgHiBlue,
		55: color.FgHiBlue,
		56: color.FgHiBlue,
		57: color.FgHiBlue,
		58: color.FgHiBlue,
		61: color.FgHiMagenta,
		62: color.FgHiMagenta,
		65: color.FgHiMagenta,
		66: color.FgHiMagenta,
		71: color.FgGreen,
		72: color.FgGreen,
		73: color.FgGreen,
		74: color.FgGreen,
		75: color.FgGreen,
		76: color.FgGreen,
		77: color.FgGreen,
		78: color.FgGreen,
		81: color.FgWhite,
		82: color.FgWhite,
		83: color.FgWhite,
		84: color.FgWhite,
		85: color.FgWhite,
		86: color.FgWhite,
		87: color.FgWhite,
		88: color.FgWhite,
	}
)

type Board struct {
	Board            [NumLines][NumColls]uint8 // [7 lines][7 columns]
	LastPutSucceeded bool
	Completed        bool
}

func NewBoard() (res Board) {
	res = Board{}
	res.Board = [NumLines][NumColls]uint8{
		{0, 0, 0, 0, 0, 0, 99},
		{0, 0, 0, 0, 0, 0, 99},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 99, 99, 99, 99},
	}
	res.LastPutSucceeded = false
	res.Completed = false
	return res
}

// Print the board to stdOut using color when possible
func (b Board) PrintBoard() {

	colorBorder := color.New(color.FgHiWhite)
	colorBack := color.New(color.FgBlue)

	colorBorder.Println("┌──────────────────────────────╖")

	for y := range b.Board {
		colorBorder.Print("│ ")
		for x := range b.Board[y] {
			if b.Board[y][x] == 0 {
				colorBack.Printf("%4s", back[y][x])
			} else if b.Board[y][x] == 99 {
				fmt.Print("    ")
			} else {
				id := b.Board[y][x]
				colorPiece := color.New(PiecesColors[id])
				if color.NoColor {
					colorPiece.Printf("[%02x]", id)
				} else {
					colorPiece.Print("▓▓▓▓")
				}
			}
		}
		colorBorder.Println(" ║")
	}
	colorBorder.Println("╘══════════════════════════════╝")
}

// CanPutPiece returns a boolean telling if the given piece can be put on given
// y,x position. The pieceId is the key from board.Pieces map.
func (b Board) CanPutPiece(pieceId, y, x uint8) (bool, error) {
	piece, ok := Pieces[pieceId]
	if !ok {
		return false, fmt.Errorf("pieceId %v does not exists", pieceId)
	}
	if y >= NumLines || x >= NumColls {
		return false, fmt.Errorf("position %v,%v is off limits", y, x)
	}
	for _, d := range piece {
		if (y+d.y) >= NumLines || (x+d.x) >= NumColls {
			return false, nil
		} else if b.Board[y+d.y][x+d.x] != 0 {
			return false, nil
		}
	}
	return true, nil
}

// PutPiece returns a boolean telling if the given piece has been put on given
// y,x position. The pieceId is the key from board.Pieces map.
func (b *Board) PutPiece(pieceId, y, x uint8) (bool, error) {
	ok, err := b.CanPutPiece(pieceId, y, x)
	if !ok {
		return false, err
	}
	for _, d := range Pieces[pieceId] {
		b.Board[y+d.y][x+d.x] = pieceId
	}
	return true, nil
}

// return the Month/Day from a completed board
func (b Board) GetDate() (month, day int, err error) {
	numMonths, numDays := 0, 0
	if !b.Completed {
		return -1, -1, errors.New("board.Completed is false")
	}
	for y, line := range b.Board {
		for x, item := range line {
			if item == 0 {
				if y > 1 { // day
					day = monthDayNumber[y][x]
					numDays++
				} else { // month
					month = monthDayNumber[y][x]
					numMonths++
				}
			}
		}
	}
	if numDays == 1 && numMonths == 1 {
		return month, day, nil
	} else {
		return -1, -1, errors.New("the solution is not valid")
	}
}
