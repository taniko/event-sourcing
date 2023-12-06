package user

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/taniko/event-sourcing/internal/domain/event"
	"github.com/taniko/event-sourcing/internal/domain/event/user"
	"github.com/taniko/event-sourcing/internal/domain/model/user/vo"
)

func TestCreate(t *testing.T) {
	tests := []struct {
		name string
		in   vo.Name
		want event.Event[any]
	}{
		{
			name: "success",
			in:   vo.Name(uuid.NewString()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			events := Create(tt.in)
			assert.Len(t, events, 1)
			e, ok := events[0].(user.Create)
			assert.True(t, ok)
			assert.NotZero(t, e)
			assert.Equal(t, user.CreateName, e.EventName())
			assert.Equal(t, event.Version(1), e.EventVersion())
			assert.Equal(t, tt.in, e.Name())
		})
	}
}

func TestRestore(t *testing.T) {
	type want struct {
		id      vo.ID
		name    vo.Name
		version event.Version
	}

	tests := []struct {
		name   string
		events event.Events[any]
		want   want
	}{
		{
			name: "create",
			events: event.Events[any]{
				user.NewCreate("user-1", "name-1"),
			},
			want: want{
				id:      "user-1",
				name:    "name-1",
				version: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := Restore(tt.events)
			assert.Equal(t, tt.want.id, u.id)
			assert.Equal(t, tt.want.name, u.name)
			assert.Equal(t, tt.want.version, u.version)
		})

	}
}
