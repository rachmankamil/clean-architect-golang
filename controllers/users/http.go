package users

import (
	"ca-amartha/businesses/users"
	controller "ca-amartha/controllers"
	"ca-amartha/controllers/users/request"
	"context"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type UserController struct {
	userUseCase users.Usecase
}

func NewUserController(uc users.Usecase) *UserController {
	return &UserController{
		userUseCase: uc,
	}
}

func (ctrl *UserController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Users{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err := ctrl.userUseCase.Store(ctx, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, "Successfully inserted")
}

func (ctrl *UserController) CreateToken(c echo.Context) error {
	ctx := c.Request().Context()

	username := c.QueryParam("username")
	password := c.QueryParam("password")

	token, err := ctrl.userUseCase.CreateToken(ctx, username, password)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	response := struct {
		Token string `json:"token"`
	}{Token: token}

	return controller.NewSuccessResponse(c, response)
}

func (ctrl *UserController) UserRole(id int) string {
	role := ""
	user, err := ctrl.userUseCase.GetByID(context.Background(), id)
	if err == nil {
		role = user.Name
	}
	return role
}
