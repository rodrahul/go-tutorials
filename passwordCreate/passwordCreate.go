package main

import (
	"errors"
	"os"
	"strings"
)

// Creator struct
type Creator struct {
	characters string
	file       *os.File
}

const (
	letters = "abcdefghijklmnopqrstuvwxyz"
	numbers = "1234567890"
	special = "!@#$%^&*()-=_+,./<>?"
)

// NewCreator function
func NewCreator(file *os.File, lowerCase, upperCase, numerals, specialCharacters bool, userCharacters string) (creator *Creator, err error) {
	if file == nil {
		return nil, errors.New("file is nil")
	}

	characters := ""

	if lowerCase {
		characters += letters
	}

	if upperCase {
		characters += strings.ToUpper(letters)
	}

	if numerals {
		characters += numbers
	}

	if specialCharacters {
		characters += special
	}

	characters += userCharacters

	if len(characters) <= 1 {
		err = errors.New("Not enough Characters specified to generate passwords")
		return nil, err
	}

	return &Creator{characters, file}, err
}

// CreatePassword method
