// Code generated by mockery v2.11.0. DO NOT EDIT.

package mockStorage

import (
	context "Bekzat/pkg/type/context"
	contact "Bekzat/services/contact/internal/domain/contact"

	mock "github.com/stretchr/testify/mock"

	queryParameter "Bekzat/pkg/type/queryParameter"

	testing "testing"

	uuid "github.com/google/uuid"
)

// Contact is an autogenerated mock type for the Contact type
type Contact struct {
	mock.Mock
}

// CountContact provides a mock function with given fields: ctx
func (_m *Contact) CountContact(ctx context.Context) (uint64, error) {
	ret := _m.Called(ctx)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(context.Context) uint64); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateContact provides a mock function with given fields: ctx, contacts
func (_m *Contact) CreateContact(ctx context.Context, contacts ...*contact.Contact) ([]*contact.Contact, error) {
	_va := make([]interface{}, len(contacts))
	for _i := range contacts {
		_va[_i] = contacts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []*contact.Contact
	if rf, ok := ret.Get(0).(func(context.Context, ...*contact.Contact) []*contact.Contact); ok {
		r0 = rf(ctx, contacts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*contact.Contact)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ...*contact.Contact) error); ok {
		r1 = rf(ctx, contacts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteContact provides a mock function with given fields: ctx, ID
func (_m *Contact) DeleteContact(ctx context.Context, ID uuid.UUID) error {
	ret := _m.Called(ctx, ID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, ID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListContact provides a mock function with given fields: ctx, parameter
func (_m *Contact) ListContact(ctx context.Context, parameter queryParameter.QueryParameter) ([]*contact.Contact, error) {
	ret := _m.Called(ctx, parameter)

	var r0 []*contact.Contact
	if rf, ok := ret.Get(0).(func(context.Context, queryParameter.QueryParameter) []*contact.Contact); ok {
		r0 = rf(ctx, parameter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*contact.Contact)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, queryParameter.QueryParameter) error); ok {
		r1 = rf(ctx, parameter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadContactByID provides a mock function with given fields: ctx, ID
func (_m *Contact) ReadContactByID(ctx context.Context, ID uuid.UUID) (*contact.Contact, error) {
	ret := _m.Called(ctx, ID)

	var r0 *contact.Contact
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *contact.Contact); ok {
		r0 = rf(ctx, ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*contact.Contact)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateContact provides a mock function with given fields: ctx, ID, updateFn
func (_m *Contact) UpdateContact(ctx context.Context, ID uuid.UUID, updateFn func(*contact.Contact) (*contact.Contact, error)) (*contact.Contact, error) {
	ret := _m.Called(ctx, ID, updateFn)

	var r0 *contact.Contact
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, func(*contact.Contact) (*contact.Contact, error)) *contact.Contact); ok {
		r0 = rf(ctx, ID, updateFn)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*contact.Contact)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID, func(*contact.Contact) (*contact.Contact, error)) error); ok {
		r1 = rf(ctx, ID, updateFn)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewContact creates a new instance of Contact. It also registers a cleanup function to assert the mocks expectations.
func NewContact(t testing.TB) *Contact {
	mock := &Contact{}

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
