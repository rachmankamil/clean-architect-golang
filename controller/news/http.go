package news

import (
	"ca-amartha/bussiness/news"
	"ca-amartha/controller"
	"ca-amartha/controller/news/request"
	"net/http"

	"github.com/labstack/echo"
)

type NewsController struct {
	newsUseCase news.Usecase
}

func NewNewsController(e *echo.Echo, newsUC news.Usecase) {
	controller := &NewsController{
		newsUseCase: newsUC,
	}

	news := e.Group("news")
	news.POST("/store", controller.Store)
}

func (ctrl *NewsController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	ip := c.QueryParam("ip")

	req := request.News{}
	if err := c.Bind(req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err := ctrl.newsUseCase.Store(ctx, ip, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, "Successfully inserted")
}
