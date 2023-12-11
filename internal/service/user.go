package service

import (
	"context"

	"github.com/samber/mo"
	command "github.com/taniko/event-sourcing/internal/command/user"
	post "github.com/taniko/event-sourcing/internal/domain/model/post/vo"
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

// ChangeName 名前を変更する
func (u *User) ChangeName(ctx context.Context, id vo.ID, name vo.Name) (*model.User, error) {
	user, err := u.repo.Find(ctx, id)
	if err != nil {
		return nil, err
	}
	originalVersion := user.Version()
	events := user.Execute(command.ChangeProfile{Name: mo.Some(name)})
	newUser := user.Apply(events...)
	if err := u.repo.Save(ctx, newUser, originalVersion); err != nil {
		return nil, err
	}
	return &newUser, nil
}

// Post 投稿する
func (u *User) Post(ctx context.Context, userID vo.ID, body post.Body) (*model.User, error) {
	user, err := u.repo.Find(ctx, userID)
	if err != nil {
		return nil, err
	}
	originalVersion := user.Version()
	events, err := user.Post(ctx, body)
	if err != nil {
		return nil, err
	}
	newUser := user.Apply(events...)
	if err := u.repo.Save(ctx, newUser, originalVersion); err != nil {
		return nil, err
	}
	return &newUser, nil
}
