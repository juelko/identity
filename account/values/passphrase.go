package values

import (
	"errors"

	zxcvbn "github.com/nbutton23/zxcvbn-go"
)

var (
	ErrWeakPassphrase = errors.New("Too weak passphrase")
)

type Passphrase string

func NewPassphrase(passphrase string, userInputs []string, minScore int) (Passphrase, error) {
	if passphraseStrenght(passphrase, userInputs) < minScore {
		return "", ErrWeakPassphrase
	}
	return Passphrase(passphrase), nil
}

func passphraseStrenght(phrase string, inputs []string) int {
	return zxcvbn.PasswordStrength(phrase, inputs).Score

}
