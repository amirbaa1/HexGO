package service

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"book/internal/core/model"
	"book/mocks"
)

func TestCreateBookService(t *testing.T) {
	// mockeService := new(mocks.BookService)
	mockBookRepository := new(mocks.BookRepository) // استفاده از mock برای repository

	bookCreate := &model.RequestBook{
		Title:     "Test",
		FirstName: "TestMamad",
		LastName:  "Taster",
	}

	author := &model.Author{
		Id:        uuid.New(),
		FirstName: bookCreate.FirstName,
		LastName:  bookCreate.LastName,
	}

	mockBookRepository.On("FindAuthorByFullNameForCreate", bookCreate.FirstName, bookCreate.LastName).Return(author, nil).Once()
	mockBookRepository.On("CreateBook", mock.Anything).Return(nil).Once()

	service := &Service{
		bookRepository: mockBookRepository,
	}

	// err := mockeService.CreateServiceBook(bookCreate)
	err := service.CreateServiceBook(bookCreate)

	assert.NoError(t, err)

	mockBookRepository.AssertExpectations(t)

}

func TestCreateService_error(t *testing.T) {
	mockRespository := new(mocks.BookRepository)

	bookCreate := &model.RequestBook{
		Title:     "Test",
		FirstName: "TestMamad",
		LastName:  "Taster",
	}

	mockRespository.On("FindAuthorByFullNameForCreate", bookCreate.FirstName, bookCreate.LastName).Return(nil, errors.New("database error")).Once()

	service := &Service{bookRepository: mockRespository}

	err := service.CreateServiceBook(bookCreate)

	assert.Error(t, err)

	assert.EqualError(t, err, "database error")

	mockRespository.AssertExpectations(t)
}
