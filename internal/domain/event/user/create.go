package user

import (
	"github.com/taniko/event-sourcing/internal/domain/event"
	"github.com/taniko/event-sourcing/internal/domain/model/user/vo"
)

const CreateName event.Name = "user.create"

type Create struct {
	id   vo.ID
	name vo.Name
}

func (e Create) EventName() event.Name {
	return CreateName
}

var _ event.Event[any] = (*Create)(nil)

func (e Create) EventVersion() event.Version {
	return 1
}

func NewCreate(id vo.ID, name vo.Name) Create {
	return Create{
		id:   id,
		name: name,
	}
}

func (e Create) ID() vo.ID {
	return e.id
}

func (e Create) Name() vo.Name {
	return e.name
}
