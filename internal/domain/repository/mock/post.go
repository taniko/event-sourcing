// Code generated by MockGen. DO NOT EDIT.
// Source: post.go
//
// Generated by this command:
//
//	mockgen -source=post.go -destination=mock/post.go -package=repository
//
// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	reflect "reflect"

	post "github.com/taniko/event-sourcing/internal/domain/model/post"
	vo "github.com/taniko/event-sourcing/internal/domain/model/user/vo"
	gomock "go.uber.org/mock/gomock"
)

// MockPost is a mock of Post interface.
type MockPost struct {
	ctrl     *gomock.Controller
	recorder *MockPostMockRecorder
}

// MockPostMockRecorder is the mock recorder for MockPost.
type MockPostMockRecorder struct {
	mock *MockPost
}

// NewMockPost creates a new mock instance.
func NewMockPost(ctrl *gomock.Controller) *MockPost {
	mock := &MockPost{ctrl: ctrl}
	mock.recorder = &MockPostMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPost) EXPECT() *MockPostMockRecorder {
	return m.recorder
}

// FindByID mocks base method.
func (m *MockPost) FindByID(ctx context.Context, userID vo.ID) ([]*post.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", ctx, userID)
	ret0, _ := ret[0].([]*post.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockPostMockRecorder) FindByID(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockPost)(nil).FindByID), ctx, userID)
}

// Save mocks base method.
func (m *MockPost) Save(ctx context.Context, post post.Post) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, post)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockPostMockRecorder) Save(ctx, post any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockPost)(nil).Save), ctx, post)
}
