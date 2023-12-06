package event

import "github.com/samber/mo"

type Name string

type Event[T any] interface {
	EventVersion() Version
	EventName() Name
}

type Events[T any] []Event[T]

func (e Events[T]) Latest() mo.Option[Event[T]] {
	if len(e) == 0 {
		return mo.None[Event[T]]()
	}
	return mo.Some(e[len(e)-1])
}
