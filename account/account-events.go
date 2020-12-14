package account

import (
	"github.com/juelko/identity/account/values"
	"github.com/juelko/identity/events"
)

type Created struct {
	Username  values.Username
	Email     values.Email
	Secret    values.Secret
	Userinfo  values.Userinfo
	EventMeta events.Meta
}

func (e Created) Meta() events.Meta {
	return e.EventMeta
}

type Deleted struct {
	Username  values.Username
	Email     values.Email
	EventMeta events.Meta
}

func (e Deleted) Meta() events.Meta {
	return e.EventMeta
}
