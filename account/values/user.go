package values

import (
	"errors"
)

var (
	ErrInvalidUsername   = errors.New("Invalid username")
	ErrNotUniqueUsername = errors.New("Username already in use")
)

type UniqueUsernameFunc func(Username) bool

type Username string

func NewUsername(name string, f UniqueUsernameFunc) (Username, error) {
	return newUsername(name, f)
}

type Userinfo struct {
	GivenName string
	LastName  string
}

func newUsername(u string, f UniqueUsernameFunc) (Username, error) {
	if isValidUsername(u) {
		return Username(""), ErrInvalidUsername
	}

	if !f(Username(u)) {
		return Username(""), ErrNotUniqueUsername
	}

	return Username(u), nil
}

func isValidUsername(name string) bool {
	return true
}
