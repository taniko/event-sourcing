package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/taniko/event-sourcing/internal/domain/event"
	"github.com/taniko/event-sourcing/internal/domain/model/post"
	postvo "github.com/taniko/event-sourcing/internal/domain/model/post/vo"
	"github.com/taniko/event-sourcing/internal/domain/model/user"
	"github.com/taniko/event-sourcing/internal/domain/model/user/vo"
	repository "github.com/taniko/event-sourcing/internal/domain/repository/mock"
	"go.uber.org/mock/gomock"
)

func TestUser_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := repository.NewMockUser(ctrl)
	repo.EXPECT().
		Save(gomock.Any(), gomock.Any(), event.Version(0)).
		Return(nil).
		AnyTimes()
	postRepo := repository.NewMockPost(ctrl)

	tests := []struct {
		name string
		args vo.Name
	}{
		{
			name: "success create",
			args: vo.Name("test"),
		}, {
			name: "success other name",
			args: vo.Name("other"),
		},
	}
	u := NewUser(repo, postRepo)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, err := u.Create(context.Background(), tt.args)
			assert.NoError(t, err)
			assert.NotEqual(t, tt.name, v.Name())
		})
	}
}

func TestUser_Post(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	u := user.New(user.Create("userName")...)
	userRepo := repository.NewMockUser(ctrl)
	userRepo.EXPECT().
		Find(gomock.Any(), gomock.Any()).
		Return(&u, nil).AnyTimes()

	var savedPost post.Post
	postRepo := repository.NewMockPost(ctrl)
	postRepo.EXPECT().
		Save(gomock.Any(), gomock.Any()).
		DoAndReturn(func(ctx context.Context, p post.Post) error {
			savedPost = p
			return nil
		}).AnyTimes()

	type args struct {
		body postvo.Body
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{
			name: "success",
			args: args{
				body: "test message",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			savedPost = post.Post{}
			s := NewUser(userRepo, postRepo)
			err := s.Post(context.Background(), "userID", tt.args.body)
			if tt.want != nil {
				assert.Error(t, err)
				assert.ErrorIs(t, err, tt.want)
				return
			}
			assert.ErrorIs(t, err, tt.want)
			assert.NotZero(t, savedPost)
		})
	}
}
