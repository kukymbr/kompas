package domain

import "errors"

// Package errors
var (
	ErrInvalidFilePath = errors.New("kompas file path is not valid")
	ErrInvalidFileType = errors.New("unknown or disabled kompas file type given")
)
