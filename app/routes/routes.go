package routes

import (
	middlewareApp "ca-amartha/app/middleware"
	"ca-amartha/controllers/category"
	"ca-amartha/controllers/news"
	"ca-amartha/controllers/users"

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
	news.POST("/store", cl.NewsController.Store, middlewareApp.RoleValidation("NewsAnchor"))
	news.PUT("/update", cl.NewsController.Update)
}
