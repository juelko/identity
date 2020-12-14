package account

import "github.com/juelko/identity/account/values"

type EmailProjection interface {
	UniqueFunc() values.UniqueEmailFunc
}

type UsernameProjection interface {
	UniqueFunc() values.UniqueUsernameFunc
}
