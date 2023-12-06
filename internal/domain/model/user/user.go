package user

import (
	"github.com/google/uuid"
	"github.com/taniko/event-sourcing/internal/domain/event"
	"github.com/taniko/event-sourcing/internal/domain/event/user"
	"github.com/taniko/event-sourcing/internal/domain/model/user/vo"
)

type (
	User struct {
		id      vo.ID
		version event.Version
		name    vo.Name
		events  []event.Event[any]
	}
)

func Restore(events event.Events[any]) User {
	u := User{}
	for _, e := range events {
		switch e := e.(type) {
		case user.Create:
			u.id = e.ID()
			u.name = e.Name()
		case user.ChangeName:
			u.name = e.Name()
		}
		u.events = append(u.events, e)
		u.version = e.EventVersion()
	}
	return u
}

func Create(name vo.Name) event.Events[any] {
	return event.Events[any]{
		user.NewCreate(vo.ID(uuid.NewString()), name),
	}
}

func (u User) ChangeName(name vo.Name) event.Events[any] {
	return event.Events[any]{
		user.NewChangeName(u.id, name, u.version+1),
	}
}
