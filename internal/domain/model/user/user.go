package user

import (
	"context"

	"github.com/google/uuid"
	command "github.com/taniko/event-sourcing/internal/command/user"
	"github.com/taniko/event-sourcing/internal/domain/event"
	"github.com/taniko/event-sourcing/internal/domain/event/user"
	post "github.com/taniko/event-sourcing/internal/domain/model/post/vo"
	"github.com/taniko/event-sourcing/internal/domain/model/user/vo"
)

type (
	User struct {
		id      vo.ID
		version event.Version
		name    vo.Name
		events  event.Events[any]
	}
)

func New(events ...event.Event[any]) User {
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

// 名前を変更するイベントを作成する
func (u User) changeName(name vo.Name) event.Events[any] {
	return event.Events[any]{
		user.NewChangeName(u.id, name, u.version+1),
	}
}

func (u User) Apply(events ...event.Event[any]) User {
	return New(append(u.events, events...)...)
}

func (u User) ID() vo.ID {
	return u.id
}

func (u User) Version() event.Version {
	return u.version
}

func (u User) Name() vo.Name {
	return u.name
}

// Execute コマンドを実行する
func (u User) Execute(cmd any) event.Events[any] {
	var events event.Events[any]
	switch cmd := cmd.(type) {
	case command.ChangeProfile:
		if name, ok := cmd.Name.Get(); ok {
			events = append(events, u.changeName(name)...)
		}
	}
	return events
}

func (u User) Post(_ context.Context, body post.Body) (event.Events[any], error) {
	var events event.Events[any]
	latestVersion := event.Version(0)
	if lastEvent, ok := u.events.Latest().Get(); ok {
		latestVersion = lastEvent.EventVersion()
	}
	events = append(events, user.NewCreatePost(u.id, body, latestVersion.Next()))
	return events, nil
}
