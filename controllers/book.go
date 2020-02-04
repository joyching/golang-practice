package controllers

import (
	"database/sql"
	"fmt"
	"strconv"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joyching/golang-practice/database"
)

type BookController struct{}

// 取得所有書籍
func (h BookController) GetAllBooks(c *gin.Context) {
	books, err := database.GetAllBooks()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Query select failed: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}

// 取得特定書籍
func (h BookController) GetBook(c *gin.Context) {
	bid, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Parse string to int64 failed: %s", err.Error()),
		})
		return
	}

	book, err := database.GetBookById(bid)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": fmt.Sprintf("Resource not found"),
			})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Query select failed: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, book)
}

// 新增書籍
func (h BookController) CreateBook(c *gin.Context) {
	book := database.Book{}
	book.Title = c.PostForm("title")
	book.PublishedAt = c.PostForm("published_at")

	if len(book.Title) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("title field is required"),
		})
		return
	}

	if len(book.PublishedAt) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("published_at field is required"),
		})
		return
	}

	result, err := database.CreateBook(book)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Create failed: %s", err.Error()),
		})
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Create failed: %s", err.Error()),
		})
		return
	}
	book.ID = id

	c.JSON(http.StatusOK, book)
}

// 更新書籍資料
func (h BookController) UpdateBook(c *gin.Context) {
	bid, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Parse string to int64 failed: %s", err.Error()),
		})
		return
	}

	book := database.Book{
		ID: bid,
		Title: c.PostForm("title"),
		PublishedAt: c.PostForm("published_at"),
	}

	if len(book.Title) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("title field is required"),
		})
		return
	}

	if len(book.PublishedAt) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("published_at field is required"),
		})
		return
	}

	_, err = database.UpdateBook(book)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Update failed: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, book)
}

// 刪除書籍
func (h BookController) DeleteBook(c *gin.Context) {
	bid, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Parse string to int64 failed: %s", err.Error()),
		})
		return
	}

	_, err = database.DeleteBook(bid)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Delete failed: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Delete Book ID:%d Success", bid),
	})
}
