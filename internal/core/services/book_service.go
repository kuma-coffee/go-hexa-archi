package services

import (
	"github.com/kuma-coffee/go-hexa-archi/internal/core/domain"
	"github.com/kuma-coffee/go-hexa-archi/internal/core/ports"
)

type bookService struct {
	bookRepository ports.BookRepository
}

func NewBookService(bookRepository ports.BookRepository) *bookService {
	return &bookService{bookRepository}
}

func (u *bookService) Store(book *domain.Book) error {
	return u.bookRepository.Store(book)
}

func (u *bookService) Fetch() ([]domain.Book, error) {
	return u.bookRepository.Fetch()
}

func (u *bookService) Update(id int, book *domain.Book) error {
	return u.bookRepository.Update(id, book)
}

func (u *bookService) Delete(id int) error {
	return u.bookRepository.Delete(id)
}
