package user

import (
	"github.com/taniko/event-sourcing/internal/domain/event"
	"github.com/taniko/event-sourcing/internal/domain/model/user/vo"
)

type ChangeName struct {
	id      vo.ID
	name    vo.Name
	version event.Version
}

var _ event.Event[any] = (*ChangeName)(nil)

const ChaneName event.Name = "vo.changeName"

func (c ChangeName) EventVersion() event.Version {
	return c.version
}

func (c ChangeName) EventName() event.Name {
	return ChaneName
}

var _ event.Event[any] = (*ChangeName)(nil)

func NewChangeName(id vo.ID, name vo.Name, version event.Version) ChangeName {
	return ChangeName{
		id:      id,
		name:    name,
		version: version,
	}
}

func (c ChangeName) Version() event.Version {
	return c.version
}

func (c ChangeName) ID() vo.ID {
	return c.id
}

func (c ChangeName) Name() vo.Name {
	return c.name
}
