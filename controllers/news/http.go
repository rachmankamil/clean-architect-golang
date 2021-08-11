package news

import (
	"ca-amartha/businesses/news"
	controller "ca-amartha/controllers"
	"ca-amartha/controllers/news/request"
	"net/http"

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

	err := ctrl.newsUseCase.Store(ctx, ip, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, "Successfully inserted")
}
