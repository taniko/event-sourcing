//go:generate mockgen -source=$GOFILE -destination=mock/$GOFILE -package=$GOPACKAGE
package repository

import (
	"context"

	"github.com/taniko/event-sourcing/internal/domain/model/post"
	user "github.com/taniko/event-sourcing/internal/domain/model/user/vo"
)

type Post interface {
	Save(ctx context.Context, post post.Post) error
	FindByID(ctx context.Context, userID user.ID) ([]*post.Post, error)
}
