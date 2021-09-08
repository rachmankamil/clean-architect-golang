package routes

import (
	middlewareApp "ca-amartha/app/middleware"
	controller "ca-amartha/controllers"
	"ca-amartha/controllers/category"
	"ca-amartha/controllers/news"
	"ca-amartha/controllers/users"
	"errors"
	"net/http"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware      middleware.JWTConfig
	UserController     users.UserController
	NewsController     news.NewsController
	CategoryController category.CategoryController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	users := e.Group("users")
	users.POST("/register", cl.UserController.Store)
	users.GET("/token", cl.UserController.CreateToken)

	category := e.Group("category")
	category.GET("/list", cl.CategoryController.GetAll, middleware.JWTWithConfig(cl.JWTMiddleware))

	news := e.Group("news", middleware.JWTWithConfig(cl.JWTMiddleware))
	news.POST("/store", cl.NewsController.Store, RoleValidation("NewsAnchor", cl.UserController))
	news.PUT("/update", cl.NewsController.Update)
}

func RoleValidation(role string, userControler users.UserController) echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := middlewareApp.GetUser(c)

			userRole := userControler.UserRole(claims.ID)

			if userRole == role {
				return hf(c)
			} else {
				return controller.NewErrorResponse(c, http.StatusForbidden, errors.New("forbidden roles"))
			}
		}
	}
}
