// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	entities "VERITRAN/internal/domain/account/entities"

	mock "github.com/stretchr/testify/mock"
)

// AccountRepository is an autogenerated mock type for the AccountRepository type
type AccountRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: account
func (_m *AccountRepository) Create(account *entities.Account) error {
	ret := _m.Called(account)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entities.Account) error); ok {
		r0 = rf(account)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Deposit provides a mock function with given fields: nickName, account
func (_m *AccountRepository) Deposit(nickName string, account *entities.Account) (interface{}, error) {
	ret := _m.Called(nickName, account)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string, *entities.Account) interface{}); ok {
		r0 = rf(nickName, account)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *entities.Account) error); ok {
		r1 = rf(nickName, account)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: nickName
func (_m *AccountRepository) Get(nickName string) (*map[string]interface{}, error) {
	ret := _m.Called(nickName)

	var r0 *map[string]interface{}
	if rf, ok := ret.Get(0).(func(string) *map[string]interface{}); ok {
		r0 = rf(nickName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*map[string]interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(nickName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WithDrawal provides a mock function with given fields: nickName, amount
func (_m *AccountRepository) WithDrawal(nickName string, account *entities.Account) (float32, error)  {
	ret := _m.Called(nickName, account)

	var r0 float32
	if rf, ok := ret.Get(0).(func(string, *entities.Account) float32); ok {
		r0 = rf(nickName, account)
	} else {
		r0 = ret.Get(0).(float32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *entities.Account) error); ok {
		r1 = rf(nickName, account)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
