package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/arcenik/go-curiositybox-calendar-puzzle-solver/board"
	"github.com/arcenik/go-curiositybox-calendar-puzzle-solver/dbsqlite3"
)

type PrintCmd struct {
	Month int `arg:"" help:"month number" default:"-1" type:"int"`
	Day   int `arg:"" help:"day number" default:"-1"`
}

func (c *PrintCmd) Run(globals *Globals) error {
	var (
		b      board.Board = board.NewBoard()
		s      board.Solution
		count  int
		err    error
		months []string = []string{
			"January", "February", "Mars",
			"April", "May", "June",
			"Jully", "August", "September",
			"October", "November", "December",
		}
		// debug bool = globals.Debug
	)

	db := dbsqlite3.InitDB("results.db") // FIXME moves to parameters and set optional
	defer db.Close()

	month := c.Month
	day := c.Day

	if month == -1 && day == -1 {
		_, nowMonth, nowDay := time.Now().Date()
		month = int(nowMonth)
		day = nowDay
	} else if month == -1 || day == -1 {
		return errors.New("you must pass boths month and day")
	} else if month <= 0 || month > 12 {
		return errors.New("month must be between 1 and 12")
	} else if day <= 0 || day > 31 {
		return errors.New("day must be between 1 and 31")
	}

	s, err = db.GetOneSolution(month, day)
	if err != nil {
		fmt.Println("GetOneSolution returned an error:")
		return err
	}

	count, err = db.GetSolutionsCount(month, day)
	if err != nil {
		fmt.Println("GetOneSolution returned an error:")
		return err
	}

	for i := range s {
		s[i].OnBoard = true
	}

	fmt.Printf("One solution for %v %v among %v\n", months[month-1], day, count)
	b.LoadSolution(s)
	b.PrintBoard()

	return nil
}
