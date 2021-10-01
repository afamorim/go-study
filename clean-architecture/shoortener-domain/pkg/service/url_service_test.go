package service

import (
	"github.com/stretchr/testify/mock"
)

var (
	ulRepositoryMock *ulRepositoryMock
)

type ulRepositoryMock struct {
	mock.Mock
}
