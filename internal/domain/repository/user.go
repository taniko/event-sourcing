//go:generate mockgen -source=$GOFILE -destination=mock/$GOFILE -package=$GOPACKAGE
package repository

import (
	"context"

	"github.com/taniko/event-sourcing/internal/domain/event"
	"github.com/taniko/event-sourcing/internal/domain/model/user/vo"
)

type User interface {
	Save(ctx context.Context, user *User, version event.Version) error
	Find(ctx context.Context, id vo.ID) (*User, error)
}
