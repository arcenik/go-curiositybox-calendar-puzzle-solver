package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/arcenik/go-curiositybox-calendar-puzzle-solver/board"
	"github.com/arcenik/go-curiositybox-calendar-puzzle-solver/dbsqlite3"
)

type SolveAllCmd struct {
}

func (c *SolveAllCmd) Run(globals *Globals) error {
	var (
		b     board.Board = board.NewBoard()
		s     board.Solution
		err   error
		debug bool = globals.Debug
	)

	db := dbsqlite3.InitDB("results.db") // FIXME moves to parameters and set optional
	defer db.Close()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	go func() {
		for range sigChan {
			fmt.Println("Ctrl-C catched, closing DB")
			db.Close()
			os.Exit(1)
		}
	}()

	s, err = board.NewSolution()
	if err != nil {
		fmt.Println("New returned an error:", err)
		return err
	}

	fmt.Println("Starting to solve all solutions")

	b.SolutionBootstrap(&s)

	for b.SolutionWalker(&s) {

		if b.Completed {

			if debug {
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				b.PrintSolution(s)
				b.PrintBoard()
				// fmt.Println("completed", b.Completed)
				// fmt.Println("check and save")
			}

			month, day, err := b.GetDate()
			if err != nil {
				if debug {
					fmt.Println("the solution is not valid")
				}
			} else {
				if debug {
					fmt.Printf("Solution for %v %v\n", month, day)
				}
				json, err := s.ToJSON()
				if err != nil {
					fmt.Println("Error extract JSON from solution:", err)
				} else {
					_ = db.InsertSolution(month, day, json)
					// TODO: ignore just the error for unique constraint on solutions.json
					//       UNIQUE constraint failed: solutions.json
				}
			}

			// remove the last piece and continue
			s[board.SolutionMax].OnBoard = false
			b.LoadSolution(s)
			b.Completed = false

			if debug {
				fmt.Print(">>")
				fmt.Scanln()
			}
			b.Completed = false
		}
	}

	fmt.Println("Completed")

	return nil
}
