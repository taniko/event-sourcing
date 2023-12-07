package service

import (
	"context"

	model "github.com/taniko/event-sourcing/internal/domain/model/user"
	"github.com/taniko/event-sourcing/internal/domain/model/user/vo"
	"github.com/taniko/event-sourcing/internal/domain/repository"
)

type User struct {
	repo repository.User
}

func NewUser(repo repository.User) *User {
	return &User{repo: repo}
}

func (u *User) Create(ctx context.Context, name vo.Name) (*model.User, error) {
	events := model.Create(name)
	user := model.New(events...)
	if err := u.repo.Save(ctx, user, 0); err != nil {
		return nil, err
	}
	return &user, nil
}
