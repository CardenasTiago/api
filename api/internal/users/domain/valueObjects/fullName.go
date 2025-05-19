package valueobjects

import (
	"errors"
	"unicode"
)

type (
	FullName struct {
		Name     string
		Lastname string
	}
)

func NewFullName(name, lastname string) (*FullName, error) {
	if name == "" || !isOnlyLetters(name) {
		return nil, errors.New("Nombre invalido")
	}

	if lastname == "" || !isOnlyLetters(lastname) {
		return nil, errors.New("Apellido invalido")
	}

	return &FullName{
		Name:     name,
		Lastname: lastname,
	}, nil
}

func isOnlyLetters(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
