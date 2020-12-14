package values

import "golang.org/x/crypto/bcrypt"

func validCost(cost int) int {
	if cost < bcrypt.MinCost || cost > bcrypt.MaxCost {
		cost = bcrypt.DefaultCost
	}
	return cost
}
