package user

import "github.com/taniko/event-sourcing/internal/domain/event"

type (
	Event interface {
		event.Event[any]
	}
)
