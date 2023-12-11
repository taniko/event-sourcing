package user

import (
	"context"
	"slices"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/taniko/event-sourcing/internal/domain/event"
	"github.com/taniko/event-sourcing/internal/domain/event/user"
	postvo "github.com/taniko/event-sourcing/internal/domain/model/post/vo"
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
			u := New(tt.events...)
			assert.Equal(t, tt.want.id, u.id)
			assert.Equal(t, tt.want.name, u.name)
			assert.Equal(t, tt.want.version, u.version)
		})
	}
}

func TestUser_Apply(t *testing.T) {
	type args struct {
		events event.Events[any]
	}
	tests := []struct {
		name      string
		args      args
		want      []User
		wantPosts int
	}{
		{
			name: "create and change name",
			args: args{
				events: event.Events[any]{
					user.NewCreate("user-1", "name-1"),
					user.NewChangeName("user-1", "name-2", 2),
				},
			},
			want: []User{
				{
					id:      "user-1",
					name:    "name-1",
					version: 1,
				}, {
					id:      "user-1",
					name:    "name-2",
					version: 2,
				},
			},
		}, {
			name: "create and twice change name",
			args: args{
				events: event.Events[any]{
					user.NewCreate("user-1", "name-1"),
					user.NewChangeName("user-1", "name-2", 2),
					user.NewChangeName("user-1", "name-3", 3),
				},
			},
			want: []User{
				{
					id:      "user-1",
					name:    "name-1",
					version: 1,
				}, {
					id:      "user-1",
					name:    "name-2",
					version: 2,
				}, {
					id:      "user-1",
					name:    "name-3",
					version: 3,
				},
			},
		}, {
			name: "create post",
			args: args{
				events: event.Events[any]{
					user.NewCreate("user-1", "name-1"),
					user.NewCreatePost("user-1", "post-1", 2),
				},
			},
			want: []User{
				{
					id:      "user-1",
					name:    "name-1",
					version: 1,
				}, {
					id:      "user-1",
					name:    "name-1",
					version: 2,
				},
			},
			wantPosts: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := New()
			for i, e := range tt.args.events {
				u = u.Apply(e)
				assert.Equal(t, tt.want[i].id, u.id, "i=%d", i)
				assert.Equal(t, tt.want[i].name, u.name, "i=%d", i)
				assert.Equal(t, tt.want[i].version, u.version, "i=%d", i)
			}
			posts := slices.DeleteFunc[event.Events[any], event.Event[any]](u.events, func(e event.Event[any]) bool {
				_, ok := e.(user.CreatePost)
				t.Log(ok)
				return !ok
			})
			assert.Len(t, posts, tt.wantPosts)
		})
	}
}

func TestUser_Execute(t *testing.T) {
	type args struct {
		event.Events[any]
		name vo.Name
	}
	tests := []struct {
		name string
		args args
		want event.Events[any]
	}{
		{
			name: "change name",
			args: args{
				Events: event.Events[any]{
					user.NewCreate("user-1", "original-name"),
				},
				name: "changed-name",
			},
			want: event.Events[any]{
				user.NewChangeName("user-1", "changed-name", 2),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := New(tt.args.Events...)
			events := u.changeName(tt.args.name)
			assert.Equal(t, tt.want, events)
		})
	}
}

func TestUser_Post(t *testing.T) {
	tests := []struct {
		name string
		args postvo.Body
	}{
		{
			name: "success",
			args: "text",
		}, {
			name: "success another text",
			args: "another text",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := New(user.NewCreate("user-1", "name-1"))
			assert.Greater(t, u.version, event.Version(0))
			baseVersion := u.version
			events, err := u.Post(context.Background(), tt.args)
			assert.NoError(t, err)
			assert.Len(t, events, 1)
			e, ok := events[0].(user.CreatePost)
			assert.True(t, ok)
			assert.NotZero(t, e)
			assert.Equal(t, user.CreatePostName, e.EventName())
			assert.Equal(t, baseVersion.Next(), e.EventVersion())
			assert.Equal(t, tt.args, e.Body())
		})
	}
}
