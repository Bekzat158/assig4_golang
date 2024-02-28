// Code generated by mockery v2.11.0. DO NOT EDIT.

package mockStorage

import (
	context "Bekzat/pkg/type/context"
	contact "Bekzat/services/contact/internal/domain/contact"

	group "Bekzat/services/contact/internal/domain/group"

	mock "github.com/stretchr/testify/mock"

	queryParameter "Bekzat/pkg/type/queryParameter"

	testing "testing"

	uuid "github.com/google/uuid"
)

// Group is an autogenerated mock type for the Group type
type Group struct {
	mock.Mock
}

// AddContactsToGroup provides a mock function with given fields: ctx, groupID, contactIDs
func (_m *Group) AddContactsToGroup(ctx context.Context, groupID uuid.UUID, contactIDs ...uuid.UUID) error {
	_va := make([]interface{}, len(contactIDs))
	for _i := range contactIDs {
		_va[_i] = contactIDs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, groupID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, ...uuid.UUID) error); ok {
		r0 = rf(ctx, groupID, contactIDs...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CountGroup provides a mock function with given fields: ctx
func (_m *Group) CountGroup(ctx context.Context) (uint64, error) {
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

// CreateContactIntoGroup provides a mock function with given fields: ctx, groupID, contacts
func (_m *Group) CreateContactIntoGroup(ctx context.Context, groupID uuid.UUID, contacts ...*contact.Contact) ([]*contact.Contact, error) {
	_va := make([]interface{}, len(contacts))
	for _i := range contacts {
		_va[_i] = contacts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, groupID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []*contact.Contact
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, ...*contact.Contact) []*contact.Contact); ok {
		r0 = rf(ctx, groupID, contacts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*contact.Contact)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID, ...*contact.Contact) error); ok {
		r1 = rf(ctx, groupID, contacts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateGroup provides a mock function with given fields: ctx, _a1
func (_m *Group) CreateGroup(ctx context.Context, _a1 *group.Group) (*group.Group, error) {
	ret := _m.Called(ctx, _a1)

	var r0 *group.Group
	if rf, ok := ret.Get(0).(func(context.Context, *group.Group) *group.Group); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*group.Group)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *group.Group) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteContactFromGroup provides a mock function with given fields: ctx, groupID, contactID
func (_m *Group) DeleteContactFromGroup(ctx context.Context, groupID uuid.UUID, contactID uuid.UUID) error {
	ret := _m.Called(ctx, groupID, contactID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, uuid.UUID) error); ok {
		r0 = rf(ctx, groupID, contactID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteGroup provides a mock function with given fields: ctx, ID
func (_m *Group) DeleteGroup(ctx context.Context, ID uuid.UUID) error {
	ret := _m.Called(ctx, ID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, ID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListGroup provides a mock function with given fields: ctx, parameter
func (_m *Group) ListGroup(ctx context.Context, parameter queryParameter.QueryParameter) ([]*group.Group, error) {
	ret := _m.Called(ctx, parameter)

	var r0 []*group.Group
	if rf, ok := ret.Get(0).(func(context.Context, queryParameter.QueryParameter) []*group.Group); ok {
		r0 = rf(ctx, parameter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*group.Group)
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

// ReadGroupByID provides a mock function with given fields: ctx, ID
func (_m *Group) ReadGroupByID(ctx context.Context, ID uuid.UUID) (*group.Group, error) {
	ret := _m.Called(ctx, ID)

	var r0 *group.Group
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *group.Group); ok {
		r0 = rf(ctx, ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*group.Group)
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

// UpdateGroup provides a mock function with given fields: ctx, ID, updateFn
func (_m *Group) UpdateGroup(ctx context.Context, ID uuid.UUID, updateFn func(*group.Group) (*group.Group, error)) (*group.Group, error) {
	ret := _m.Called(ctx, ID, updateFn)

	var r0 *group.Group
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, func(*group.Group) (*group.Group, error)) *group.Group); ok {
		r0 = rf(ctx, ID, updateFn)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*group.Group)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID, func(*group.Group) (*group.Group, error)) error); ok {
		r1 = rf(ctx, ID, updateFn)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewGroup creates a new instance of Group. It also registers a cleanup function to assert the mocks expectations.
func NewGroup(t testing.TB) *Group {
	mock := &Group{}

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
