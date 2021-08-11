package news

import (
	"ca-amartha/businesses/news"
	controller "ca-amartha/controllers"
	"ca-amartha/controllers/news/request"
	"ca-amartha/controllers/news/response"
	"errors"
	"net/http"
	"strconv"
	"strings"

	echo "github.com/labstack/echo/v4"
)

type NewsController struct {
	newsUseCase news.Usecase
}

func NewNewsController(newsUC news.Usecase) *NewsController {
	return &NewsController{
		newsUseCase: newsUC,
	}
}

func (ctrl *NewsController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	ip := c.QueryParam("ip")

	req := request.News{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	resp, err := ctrl.newsUseCase.Store(ctx, ip, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp))
}

func (ctrl *NewsController) Update(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.QueryParam("id")
	if strings.TrimSpace(id) == "" {
		return controller.NewErrorResponse(c, http.StatusBadRequest, errors.New("missing required id"))
	}

	req := request.News{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	domainReq := req.ToDomain()
	idInt, _ := strconv.Atoi(id)
	domainReq.ID = idInt
	resp, err := ctrl.newsUseCase.Update(ctx, domainReq)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(*resp))
}
