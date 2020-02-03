package database

import (
	"database/sql"
	"time"
)

func GetAllBooks() ([]Book, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, title, published_at FROM books")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Book

	for rows.Next() {
		var book Book
		err = rows.Scan(&book.ID, &book.Title, &book.PublishedAt)
		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, err
}

func GetBookById(id int64) (*Book, error) {
	book := &Book{
		ID: id,
	}

	db, err := Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	err = db.QueryRow("SELECT title, published_at FROM books WHERE id = ? LIMIT 1", book.ID).Scan(&book.Title, &book.PublishedAt)

	if err != nil {
		return nil, err
	}

	return book, err
}

func CreateBook(book Book) (sql.Result, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	result, err := db.Exec(
		"INSERT INTO books (title, published_at, created_at, updated_at) VALUES (?, ?, ?, ?)",
		book.Title,
		book.PublishedAt,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return nil, err
	}

	return result, err
}

func UpdateBook(book Book) (sql.Result, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	result, err := db.Exec(
		"UPDATE books SET title = ?, published_at = ?, updated_at = ? WHERE id = ?",
		book.Title,
		book.PublishedAt,
		time.Now(),
		book.ID,
	)

	if err != nil {
		return nil, err
	}

	return result, err
}

func DeleteBook(id int64) (sql.Result, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	result, err := db.Exec("DELETE FROM books WHERE id = ?", id)

	if err != nil {
		return nil, err
	}

	return result, err
}
