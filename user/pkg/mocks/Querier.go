// Code generated by mockery v2.5.1. DO NOT EDIT.

package mocks

import (
	context "context"

	repo "github.com/chutified/booking-terminal/user/pkg/repo"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// Querier is an autogenerated mock type for the Querier type
type Querier struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: ctx, arg
func (_m *Querier) CreateUser(ctx context.Context, arg repo.CreateUserParams) (repo.User, error) {
	ret := _m.Called(ctx, arg)

	var r0 repo.User
	if rf, ok := ret.Get(0).(func(context.Context, repo.CreateUserParams) repo.User); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(repo.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, repo.CreateUserParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteUser provides a mock function with given fields: ctx, id
func (_m *Querier) DeleteUser(ctx context.Context, id uuid.UUID) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetUser provides a mock function with given fields: ctx, id
func (_m *Querier) GetUser(ctx context.Context, id uuid.UUID) (repo.User, error) {
	ret := _m.Called(ctx, id)

	var r0 repo.User
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) repo.User); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(repo.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: ctx, arg
func (_m *Querier) UpdateUser(ctx context.Context, arg repo.UpdateUserParams) (repo.User, error) {
	ret := _m.Called(ctx, arg)

	var r0 repo.User
	if rf, ok := ret.Get(0).(func(context.Context, repo.UpdateUserParams) repo.User); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(repo.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, repo.UpdateUserParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}