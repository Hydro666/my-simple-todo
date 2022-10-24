package internal

import "errors"

var (
	ErrListNotFound     = errors.New("list not found")
	ErrDatabaseNotValid = errors.New("database is in an inconsistent state")
)
