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
