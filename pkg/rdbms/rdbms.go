package rdbms

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

type RDBMS interface {
	Execute(query string, in []interface{}) error

	QueryRaw(query string, in []interface{}, out []interface{}) error

	Query(query string, in []interface{}, out [][]interface{}) error
}

type rdbms struct {
	db *sql.DB
}

var (
	ErrPrepareStatement = "error when to prepare statement"
	ErrNotFound         = "error no entry with given arguments"
	ErrDuplicate        = "error operation canceled due to the duplication entry"
)

func (db *rdbms) Execute(query string, in []interface{}) error {
	stmt, err := db.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("%s\n%s", ErrPrepareStatement, err)
	}
	defer stmt.Close()

	if _, err = stmt.Exec(in...); err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return errors.New(ErrDuplicate)
		}
		return fmt.Errorf("%s\n%s", "error when trying to execute statement", err)
	}

	return nil
}

func (db *rdbms) QueryRaw(query string, in []interface{}, out []interface{}) error {
	stmt, err := db.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("%s\n%s", ErrPrepareStatement, err)
	}
	defer stmt.Close()

	if err = stmt.QueryRow(in...).Scan(out...); err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return errors.New(ErrDuplicate)
		} else if errors.Is(err, sql.ErrNoRows) {
			return errors.New(ErrNotFound)
		}
		return fmt.Errorf("%s\n%s", "error while executing the query or scanning the row", err)
	}

	return nil
}

func (db *rdbms) Query(query string, in []interface{}, out [][]interface{}) error {
	stmt, err := db.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("%s\n%s", ErrPrepareStatement, err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(in...)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return errors.New(ErrDuplicate)
		}
		return fmt.Errorf("%s\n%s", "error while executing the query", err)
	}

	var index = 0
	for ; rows.Next(); index++ {
		if err = rows.Scan(out[index]...); err != nil {
			return fmt.Errorf("%s\n%s", "error while scanning the row", err)
		}
	}

	out = out[:index+1]

	if err = rows.Err(); err != nil {
		return fmt.Errorf("%s\n%s", "there's an error in result of the query", err)
	}

	return nil
}
