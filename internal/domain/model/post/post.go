package post

import (
	"time"

	"github.com/taniko/event-sourcing/internal/domain/event"
	"github.com/taniko/event-sourcing/internal/domain/model/post/vo"
	user "github.com/taniko/event-sourcing/internal/domain/model/user/vo"
)

type Post struct {
	id        vo.ID
	userID    user.ID
	body      vo.Body
	version   event.Version
	createdAt time.Time
}

func New(id vo.ID, userID user.ID, body vo.Body, version event.Version, createdAt time.Time) *Post {
	return &Post{
		id:        id,
		userID:    userID,
		body:      body,
		version:   version,
		createdAt: createdAt,
	}
}
