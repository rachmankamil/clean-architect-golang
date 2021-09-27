package books

import (
	"kampus-merdeka-ca/app/presenter/books/request"
	"kampus-merdeka-ca/app/presenter/books/response"
	"kampus-merdeka-ca/bussiness/books"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Presenter struct {
	serviceBook books.Service
}

func NewHandler(bookServ books.Service) *Presenter {
	return &Presenter{
		serviceBook: bookServ,
	}
}

func (handler *Presenter) Insert(echoContext echo.Context) error {
	var req request.BookInsert
	if err := echoContext.Bind(&req); err != nil {
		return echoContext.JSON(http.StatusBadRequest, "something wrong")
	}

	domain := request.ToDomain(req)
	resp, err := handler.serviceBook.Append(domain)
	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, "something wrong")
	}

	return echoContext.JSON(http.StatusOK, response.FromDomain(*resp))
}
