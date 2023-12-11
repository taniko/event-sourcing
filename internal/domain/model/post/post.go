package post

import (
	"time"

	"github.com/google/uuid"
	"github.com/taniko/event-sourcing/internal/domain/model/post/vo"
	user "github.com/taniko/event-sourcing/internal/domain/model/user/vo"
)

type Post struct {
	id        vo.ID
	userID    user.ID
	body      vo.Body
	createdAt time.Time
}

func New(id vo.ID, userID user.ID, body vo.Body, createdAt time.Time) *Post {
	return &Post{
		id:        id,
		userID:    userID,
		body:      body,
		createdAt: createdAt,
	}
}

func Create(userID user.ID, body vo.Body) *Post {
	return New(vo.ID(uuid.NewString()), userID, body, time.Now())
}
