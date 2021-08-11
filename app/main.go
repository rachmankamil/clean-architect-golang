package main

import (
	_newsUsecase "ca-amartha/businesses/news"
	_newsController "ca-amartha/controllers/news"
	_newsRepo "ca-amartha/drivers/databases/news"

	_categoryUsecase "ca-amartha/businesses/category"
	_categoryController "ca-amartha/controllers/category"
	_categoryRepo "ca-amartha/drivers/databases/category"

	_userUsecase "ca-amartha/businesses/users"
	_userController "ca-amartha/controllers/users"
	_userRepo "ca-amartha/drivers/databases/users"

	_dbDriver "ca-amartha/drivers/mysql"

	_middleware "ca-amartha/app/middleware"
	_routes "ca-amartha/app/routes"

	"log"
	"time"

	echo "github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`app/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	configDB := _dbDriver.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}
	db := configDB.InitialDB()

	configJWT := _middleware.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	e := echo.New()
	categoryRepo := _categoryRepo.NewCategoryRepository(db)
	categoryUsecase := _categoryUsecase.NewCategoryUsecase(timeoutContext, categoryRepo)
	categoryCtrl := _categoryController.NewCategoryController(categoryUsecase)

	newsRepo := _newsRepo.NewMySQLNewsRepository(db)
	newsUsecase := _newsUsecase.NewNewsUsecase(newsRepo, categoryUsecase, timeoutContext)
	newsCtrl := _newsController.NewNewsController(newsUsecase)

	userRepo := _userRepo.NewMySQLUserRepository(db)
	userUsecase := _userUsecase.NewUserUsecase(userRepo, &configJWT, timeoutContext)
	userCtrl := _userController.NewUserController(userUsecase)

	routesInit := _routes.ControllerList{
		JWTMiddleware:      configJWT.Init(),
		UserController:     *userCtrl,
		NewsController:     *newsCtrl,
		CategoryController: *categoryCtrl,
	}
	routesInit.RouteRegister(e)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
