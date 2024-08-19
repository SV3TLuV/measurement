package model

import "github.com/pkg/errors"

var (
	NotFound = errors.New("not found")
	Empty    = errors.New("empty")
)
