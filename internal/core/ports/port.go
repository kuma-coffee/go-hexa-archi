package ports

import (
	"github.com/kuma-coffee/go-hexa-archi/internal/core/domain"
	"github.com/labstack/echo"
)

type BookRepository interface {
	Store(book *domain.Book) error
	Fetch() ([]domain.Book, error)
	Update(id int, book *domain.Book) error
	Delete(id int) error
}

type BookService interface {
	Store(book *domain.Book) error
	Fetch() ([]domain.Book, error)
	Update(id int, book *domain.Book) error
	Delete(id int) error
}

type BookHandler interface {
	AddBook(c echo.Context) error
	GetAllBooks(c echo.Context) error
	UpdateBook(c echo.Context) error
	DeleteBook(c echo.Context) error
}
