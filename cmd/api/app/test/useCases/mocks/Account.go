// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Account is an autogenerated mock type for the Account type
type Account struct {
	mock.Mock
}

// Deposit provides a mock function with given fields: nickName, amount
func (_m *Account) Deposit(nickName string, amount float64) (float64, error) {
	ret := _m.Called(nickName, amount)

	var r0 float64
	if rf, ok := ret.Get(0).(func(string, float64) float64); ok {
		r0 = rf(nickName, amount)
	} else {
		r0 = ret.Get(0).(float64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, float64) error); ok {
		r1 = rf(nickName, amount)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
