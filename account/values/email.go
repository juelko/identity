package values

import (
	"errors"
	"fmt"
	"regexp"
)

// Exported errors from this file
var (
	ErrInvalidEmail   = errors.New("Invalid email address")
	ErrNotUniqueEmail = errors.New("Email address is already in use")
)

type UniqueEmailFunc func(Email) bool

type Email string

func NewEmail(address string, f UniqueEmailFunc) (Email, error) {
	return newEmail(address, f)
}

func newEmail(s string, f UniqueEmailFunc) (Email, error) {
	if !isValidEmailAddress(s) {
		return Email(""), ErrInvalidEmail
	}

	if !f(Email(s)) {
		return Email(""), ErrNotUniqueEmail
	}

	return Email(s), nil
}

func isValidEmailAddress(a string) bool {
	if len(a) < 3 || len(a) > 254 {
		return false
	}

	return emailFormatValidation(a)
}

// unexported variables for email address validation
var (
	rfc5322     = "(?i)(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|\"(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21\\x23-\\x5b\\x5d-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])*\")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\\[(?:(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9]))\\.){3}(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9])|[a-z0-9-]*[a-z0-9]:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21-\\x5a\\x53-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])+)\\])"
	emailRegexp = regexp.MustCompile(fmt.Sprintf("^%s*$", rfc5322))
)

func emailFormatValidation(a string) bool {

	if !emailRegexp.MatchString(a) {
		return false
	}
	return true
}
