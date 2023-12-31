package service

import (
	"context"

	"github.com/samber/mo"
	command "github.com/taniko/event-sourcing/internal/command/user"
	"github.com/taniko/event-sourcing/internal/domain/model/post"
	postvo "github.com/taniko/event-sourcing/internal/domain/model/post/vo"
	model "github.com/taniko/event-sourcing/internal/domain/model/user"
	"github.com/taniko/event-sourcing/internal/domain/model/user/vo"
	"github.com/taniko/event-sourcing/internal/domain/repository"
)

type User struct {
	repo repository.User
	p    repository.Post
}

func NewUser(repo repository.User, p repository.Post) *User {
	return &User{
		repo: repo,
		p:    p,
	}
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
func (u *User) Post(ctx context.Context, userID vo.ID, body postvo.Body) error {
	user, err := u.repo.Find(ctx, userID)
	if err != nil {
		return err
	}
	p := post.Create(user.ID(), body)
	if err := u.p.Save(ctx, *p); err != nil {
		return err
	}
	return nil
}
