// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import filestore "github.com/filecoin-project/go-fil-components/filestore"
import mock "github.com/stretchr/testify/mock"

// File is an autogenerated mock type for the File type
type File struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *File) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Path provides a mock function with given fields:
func (_m *File) Path() filestore.Path {
	ret := _m.Called()

	var r0 filestore.Path
	if rf, ok := ret.Get(0).(func() filestore.Path); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(filestore.Path)
	}

	return r0
}

// Read provides a mock function with given fields: p
func (_m *File) Read(p []byte) (int, error) {
	ret := _m.Called(p)

	var r0 int
	if rf, ok := ret.Get(0).(func([]byte) int); ok {
		r0 = rf(p)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Seek provides a mock function with given fields: offset, whence
func (_m *File) Seek(offset int64, whence int) (int64, error) {
	ret := _m.Called(offset, whence)

	var r0 int64
	if rf, ok := ret.Get(0).(func(int64, int) int64); ok {
		r0 = rf(offset, whence)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, int) error); ok {
		r1 = rf(offset, whence)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Size provides a mock function with given fields:
func (_m *File) Size() int64 {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// Write provides a mock function with given fields: p
func (_m *File) Write(p []byte) (int, error) {
	ret := _m.Called(p)

	var r0 int
	if rf, ok := ret.Get(0).(func([]byte) int); ok {
		r0 = rf(p)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
