package values

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

func NewSecret(params SecretParameters) (Secret, error) {
	return newSecret(params)
}

type SecretParameters struct {
	Passphrase Passphrase
	Cost       int
	Expires    time.Duration
}

type Secret struct {
	Hash      []byte
	Cost      int
	ResetCode ResetCode
	Expires   time.Time
}

func newSecret(p SecretParameters) (Secret, error) {
	var s Secret

	s.Cost = validCost(p.Cost)

	var err error

	s.Hash, err = bcrypt.GenerateFromPassword([]byte(p.Passphrase), s.Cost)
	if err != nil {
		return Secret{}, err
	}

	s.ResetCode = EmptyResetCode()

	s.Expires = time.Now().Add(p.Expires)

	return s, nil
}
