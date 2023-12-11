package user

import (
	"time"

	"github.com/taniko/event-sourcing/internal/domain/event"
	"github.com/taniko/event-sourcing/internal/domain/model/post/vo"
	user "github.com/taniko/event-sourcing/internal/domain/model/user/vo"
)

const CreatePostName event.Name = "post.create"

type CreatePost struct {
	id        vo.ID
	userID    user.ID
	version   event.Version
	body      vo.Body
	createdAt time.Time
}

var (
	_ event.Event[any] = (*CreatePost)(nil)
	_ Event            = (*CreatePost)(nil)
)

func (e CreatePost) EventVersion() event.Version {
	return e.version
}

func (e CreatePost) EventName() event.Name {
	return CreatePostName
}

func NewCreatePost(userID user.ID, body vo.Body, version event.Version) CreatePost {
	return CreatePost{
		userID:    userID,
		body:      body,
		version:   version,
		createdAt: time.Now(),
	}
}

func (e CreatePost) UserID() user.ID {
	return e.userID
}

func (e CreatePost) Body() vo.Body {
	return e.body
}

func (e CreatePost) CreatedAt() time.Time {
	return e.createdAt
}

func (e CreatePost) ID() vo.ID {
	return e.id
}
