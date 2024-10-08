package repository

import (
	"book/internal/core/model"
	"errors"
	"gorm.io/gorm"
	"log"
)

func (r *Repository) CreateBook(book *model.Book) error {
	return r.db.Create(&book).Error
}

func (r *Repository) GetAllBooks() ([]model.Book, error) {
	var books []model.Book
	result := r.db.Preload("Author").Find(&books)
	log.Println(result)
	if result.Error != nil {
		return nil, result.Error
	}

	return books, nil
}

func (r *Repository) FindAuthorByFullNameForCreate(authorFirstName string, authorLastName string) (*model.Author, error) {
	var author model.Author
	if authorFirstName == "" || authorLastName == "" {
		return nil, nil
	}

	result := r.db.Where("first_name = ? AND last_name = ?",
		authorFirstName,
		authorLastName).
		First(&author)

	log.Println(result.Error)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &author, nil
}

func (r *Repository) CreateAuthor(author *model.Author) error {
	return r.db.Create(&author).Error
}

func (r *Repository) UpdateBook(book *model.Book) error {
	//return r.db.Updates(&book).Error
	return r.db.Model(&book).Updates(book).Error
}

func (r *Repository) GetBookById(id string) (*model.Book, error) {
	var book model.Book
	result := r.db.Where("id = ?", id).Preload("Author").First(&book)

	log.Printf("Query result: %+v", result)

	if result.Error != nil {
		log.Printf("Query result: %+v", result)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &book, nil
}
