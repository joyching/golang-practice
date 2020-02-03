package server

import (
	"github.com/gin-gonic/gin"
	"github.com/joyching/golang-practice/controllers"
)

// NewRouter 定義 RESTful API 的所有路由
func NewRouter() *gin.Engine {
	router := gin.Default()

	health := new(controllers.HealthController)
	router.GET("/health", health.Status)

	book := new(controllers.BookController)
	router.GET("/books", book.GetAllBooks)
	router.GET("/books/:id", book.GetBook)
	router.POST("/books", book.CreateBook)
	router.PUT("/books/:id", book.UpdateBook)
	router.DELETE("/books/:id", book.DeleteBook)

	return router
}
