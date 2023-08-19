package repository

import (
	"golang-unit-test/entity"

	"github.com/stretchr/testify/mock"
)

type CategoryRepositoryMock struct{
	Mock mock.Mock
}

func (repository *CategoryRepositoryMock) FindById(id string) *entity.Category{
	arguments := repository.Mock.Called(id)
	if arguments.Get(0) == nil{
		return nil
	}
	category := arguments.Get(0).(entity.Category) //because only get the first returned, must convert to the needed return type
	return &category
}