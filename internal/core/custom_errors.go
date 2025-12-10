package core

import (
	"errors"
)

var (
	ErrEmptyDescription error = errors.New("description cannot be empty")
	ErrTaskNotFound     error = errors.New("task not found")
	ErrInvalidId        error = errors.New("id must be greater than zero")
)
