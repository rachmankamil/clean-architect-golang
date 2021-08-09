package category

import (
	"ca-amartha/bussiness/category"
	"ca-amartha/controller/category/response"
	"net/http"

	"ca-amartha/controller"

	"github.com/labstack/echo"
)

type CategoryController struct {
	categoryUsecase category.Usecase
}

func NewCategoryController(e *echo.Echo, cu category.Usecase) {
	controller := &CategoryController{
		categoryUsecase: cu,
	}

	category := e.Group("category")
	category.GET("/list", controller.GetAll)
}

func (ctrl *CategoryController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	resp, err := ctrl.categoryUsecase.GetAll(ctx)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	responseController := []response.Category{}
	for _, value := range resp {
		responseController = append(responseController, response.FromDomain(value))
	}

	return controller.NewSuccessResponse(c, responseController)
}
