package handlers

import (
	"net/http"
	"strconv"

	"github.com/kuma-coffee/go-hexa-archi/internal/core/domain"
	"github.com/kuma-coffee/go-hexa-archi/internal/core/ports"
	"github.com/labstack/echo/v4"
)

type bookHandler struct {
	bookService ports.BookService
}

func NewBookHandler(bookService ports.BookService) *bookHandler {
	return &bookHandler{bookService}
}

func (h *bookHandler) AddBook(c echo.Context) error {
	var newBook domain.Book

	err := c.Bind(&newBook)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	err = h.bookService.Store(&newBook)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, "success create book")
}

func (h *bookHandler) GetAllBooks(c echo.Context) error {
	books, err := h.bookService.Fetch()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, books)
}

func (h *bookHandler) UpdateBook(c echo.Context) error {
	var newBook domain.Book

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	err = c.Bind(&newBook)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	err = h.bookService.Update(id, &newBook)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, "success update book")
}

func (h *bookHandler) DeleteBook(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	err = h.bookService.Delete(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, "success delete book")
}
