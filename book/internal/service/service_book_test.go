package service

import (
	"book/internal/core/model"
	"book/mocks"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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
	err:= service.CreateServiceBook(bookCreate)

	assert.NoError(t, err)

	mockBookRepository.AssertExpectations(t)

}
