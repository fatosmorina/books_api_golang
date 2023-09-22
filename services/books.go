package services

import (
	"books_api/models"
	"books_api/storage"
	"books_api/types"
	"fmt"
	"github.com/labstack/gommon/random"
)

func GetBook(bookId string) (interface{}, error) {
	book := models.Book{}
	err := storage.DB.Where("id = ?", bookId).Select("id, isbn, title, author_id").First(&book).Error
	if err != nil {
		fmt.Println("Error while getting a book", err)
		return nil, err
	}
	return book, nil
}

func GetBooksForAuthor(authorId string) (interface{}, error) {
	books := []models.Book{}
	err := storage.DB.Where("author_id = ?", authorId).Select("id, isbn, title, author_id").Find(&books).Error
	if err != nil {
		fmt.Println("Error while getting books for an author", err)
		return nil, err
	}
	return books, nil
}

func AddBook(bookInputBody *types.BookInputBody) error {
	author := models.Author{
		ID:        random.String(30),
		FirstName: bookInputBody.Author.FirstName,
		LastName:  bookInputBody.Author.LastName,
	}

	err := storage.DB.Create(&author).Error
	if err != nil {
		fmt.Println("Error while creating an author", err)
		return err
	}
	book := models.Book{
		ID:       random.String(30),
		Isbn:     bookInputBody.Isbn,
		Title:    bookInputBody.Title,
		AuthorId: author.ID,
	}

	errorWhileCreatingABook := storage.DB.Create(&book).Error
	if errorWhileCreatingABook != nil {
		fmt.Println("Error while adding a book", errorWhileCreatingABook)
		return errorWhileCreatingABook
	}
	return nil
}

func UpdateBook(bookInputBody *types.BookInputBody) error {
	book := &models.Book{}
	error := storage.DB.Where("id = ?", bookInputBody.ID).First(&book).Error

	if error != nil {
		fmt.Println("Error while updating a book", error)
		return error
	}

	book.Isbn = bookInputBody.Isbn
	book.Title = bookInputBody.Title

	err := storage.DB.Save(&book).Error

	if err != nil {
		fmt.Println("Error while saving an updated book in the database", err)
		return err
	}
	return nil
}

func DeleteBook(bookId string) error {
	book := models.Book{}
	err := storage.DB.Where("id = ?", bookId).Delete(&book).Error
	if err != nil {
		fmt.Println("Error while deleting a book", err)
		return err
	}
	return nil
}
