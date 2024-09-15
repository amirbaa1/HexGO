package service

import (
	"book/internal/core/model"
	"errors"
	"github.com/google/uuid"
	"time"
)

func (s *Service) CreateServiceBook(book *model.RequestBook) error {
	if book == nil {
		return errors.New("book is nil")
	}
	var err error
	var authorId uuid.UUID

	findAuthor, err := s.bookRepository.FindAuthorByFullNameForCreate(book.FirstName, book.LastName)
	if err != nil {
		return err
	}
	if findAuthor == nil {
		newAuthor := model.Author{
			Id:        uuid.New(),
			FirstName: book.FirstName,
			LastName:  book.LastName,
		}
		err = s.bookRepository.CreateAuthor(&newAuthor)
		if err != nil {
			return err
		}
		authorId = newAuthor.Id
	} else {
		authorId = findAuthor.Id
	}

	newBook := model.Book{
		Id:         uuid.New(),
		Title:      book.Title,
		AuthorId:   authorId,
		CreateTime: time.Now(),
	}

	err = s.bookRepository.CreateBook(&newBook)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetAllBooks() (*[]model.ResponseBook, error) {
	getAll, err := s.bookRepository.GetAllBooks()
	if err != nil {
		return nil, err
	}

	var responseBooks []model.ResponseBook
	for _, book := range getAll {
		responseBooks = append(responseBooks, model.ResponseBook{
			Id:       book.Id,
			Title:    book.Title,
			AuthorId: book.Author.Id,
			Author: model.Author{
				Id:        book.Author.Id,
				FirstName: book.Author.FirstName,
				LastName:  book.Author.LastName,
			},
			CreatedAt: book.CreateTime,
			UpdatedAt: book.UpdateTime,
		})
	}

	return &responseBooks, nil
}

func (s *Service) UpdateBook(book *model.RequestBookUpdate, bookId string) (*model.RequestBookUpdate, error) {

	existingBook, err := s.bookRepository.GetBookById(bookId)

	if err != nil {
		return nil, err
	}

	if existingBook == nil {
		return nil, errors.New("book not found")
	}

	existingBook.Title = book.Title
	existingBook.UpdateTime = time.Now()

	err = s.bookRepository.UpdateBook(existingBook)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (s *Service) GetBookById(bookId string) (*model.ResponseBook, error) {
	book, err := s.bookRepository.GetBookById(bookId)
	if err != nil {
		return nil, err
	}

	if book == nil {
		return nil, errors.New("book not found")
	}

	responseBook := &model.ResponseBook{
		Id:       book.Id,
		Title:    book.Title,
		AuthorId: book.Author.Id,
		Author: model.Author{
			Id:        book.Author.Id,
			FirstName: book.Author.FirstName,
			LastName:  book.Author.LastName,
		},
		CreatedAt: book.CreateTime,
		UpdatedAt: book.UpdateTime,
	}

	return responseBook, nil
}
