package dbsqlite3

import (
	"encoding/json"

	"github.com/arcenik/go-curiositybox-calendar-puzzle-solver/board"
)

func (db *DB) InsertSolution(month, day int, json string) error {
	query := "INSERT INTO solutions(month,day,json) VALUES(?,?,?);"
	stmt, err := db.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(month, day, json)

	return err
}

// Return one solution, randomly selected
func (db *DB) GetOneSolution(month, day int) (s board.Solution, err error) {
	var (
		jsonStr string
		query   string = `
		SELECT json
		FROM solutions
		WHERE id >= (abs(random()) % (SELECT max(id) FROM solutions))
		AND month=?
		AND day=?
		LIMIT 1;`
	)

	stmt, err := db.db.Prepare(query)
	if err != nil {
		return board.Solution{}, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(month, day)
	if row.Err() != nil {
		return board.Solution{}, row.Err()
	}

	err = row.Scan(&jsonStr)
	if err != nil {
		return board.Solution{}, err
	}

	err = json.Unmarshal([]byte(jsonStr), &s)
	if err != nil {
		return board.Solution{}, err
	}

	return s, err
}

// Return the count of solution for the given month,day
func (db *DB) GetSolutionsCount(month, day int) (c int, err error) {
	var (
		query string = `
		SELECT COUNT(id)
		FROM solutions
		WHERE month=?
		AND day=?;`
	)

	stmt, err := db.db.Prepare(query)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(month, day)
	if row.Err() != nil {
		return -1, row.Err()
	}

	err = row.Scan(&c)
	if err != nil {
		return -1, err
	}

	return c, err
}
