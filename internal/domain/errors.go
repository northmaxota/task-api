package domain

import "errors"

var (
	ErrNotFound      = errors.New("not found")
	ErrTitleRequired = errors.New("title is required")
)
