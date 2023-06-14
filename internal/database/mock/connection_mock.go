package mock

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/mock"
)

type Connection struct {
	mock.Mock
}

func (_m *Connection) Close() {
	_m.Called()
}

type mockConstructorTestingTNewPGConnection interface {
	mock.TestingT
	Cleanup(func())
}

func NewPgConnection(t mockConstructorTestingTNewPGConnection) *Connection {
	mock := &Connection{}
	mock.Mock.Test(t)
	t.Cleanup(func() { mock.AssertExpectations(t) })
	return mock
}
func (_m *Connection) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	var _ca []interface{}
	_ca = append(_ca, ctx, sql)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 pgx.Rows
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) pgx.Rows); ok {
		r0 = rf(ctx, sql, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pgx.Rows)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, ...interface{}) error); ok {
		r1 = rf(ctx, sql, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
