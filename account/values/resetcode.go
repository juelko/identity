package values

import "github.com/google/uuid"

func EmptyResetCode() ResetCode {
	return ResetCode("")
}

func GenerateResetCode() ResetCode {
	return ResetCode(uuid.New().String())
}

type ResetCode string

func (c ResetCode) IsSet() bool {
	if c == "" {
		return false
	}
	return true
}
