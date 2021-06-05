package models

import "errors"

var (
	// ErrPersonCanNotBeBil a persona no puede ser nula
	ErrPersonCanNotBeBil = errors.New("la persona no puede ser nula")
)
