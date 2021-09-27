package routes

import (
	"kampus-merdeka-ca/app/presenter/books"

	"github.com/labstack/echo/v4"
)

type HandlerList struct {
	BookHandler books.Presenter
}

func (handler *HandlerList) RouteRegister(e *echo.Echo) {
	books := e.Group("books")
	books.POST("/register", handler.BookHandler.Insert)
}
