package main

import (
	"database/sql"
	"fmt"

	"github.com/kuma-coffee/go-hexa-archi/internal/core/services"
	"github.com/kuma-coffee/go-hexa-archi/internal/handlers"
	"github.com/kuma-coffee/go-hexa-archi/internal/repositories"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	username = "postgres"
	password = "postgres"
	dbName   = "test"
	port     = 5432
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, username, password, dbName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	e := echo.New()

	bookRepo := repositories.NewBookRepo(db)
	bookService := services.NewBookService(bookRepo)
	bookHandler := handlers.NewBookHandler(bookService)

	e.POST("/books", bookHandler.AddBook)
	e.GET("/books", bookHandler.GetAllBooks)
	e.PUT("/books/:id", bookHandler.UpdateBook)
	e.DELETE("/books/:id", bookHandler.DeleteBook)

	e.Start(":8080")
}
