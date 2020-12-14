package events

import (
	"reflect"
	"time"
)

type Event interface {
	Meta() Meta
}

type Meta struct {
	EventType     string
	OccuredAt     time.Time
	StreamVersion int
}

func NewMeta(e Event, version int) Meta {
	return Meta{
		EventType:     reflect.TypeOf(e).String(),
		OccuredAt:     time.Now(),
		StreamVersion: version,
	}
}
