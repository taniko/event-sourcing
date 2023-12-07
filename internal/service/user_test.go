package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/taniko/event-sourcing/internal/domain/event"
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
	u := NewUser(repo)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := u.Create(context.Background(), tt.args)
			assert.NoError(t, err)
			assert.NotEqual(t, tt.name, user.Name())
		})
	}
}
