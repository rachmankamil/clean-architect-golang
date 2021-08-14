package main

import (
	_dbFactory "ca-amartha/drivers/databases"

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

	_ipLocatorDriver "ca-amartha/drivers/thirdparties/iplocator"

	_config "ca-amartha/app/config"
	_middleware "ca-amartha/app/middleware"
	_routes "ca-amartha/app/routes"

	"log"
	"time"

	echo "github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&_newsRepo.News{},
		&_categoryRepo.Category{},
		&_userRepo.Users{},
	)
}

func main() {
	configApp := _config.GetConfig()
	configDB := _dbDriver.ConfigDB{
		DB_Username: configApp.Database.User,
		DB_Password: configApp.Database.Pass,
		DB_Host:     configApp.Database.Host,
		DB_Port:     configApp.Database.Port,
		DB_Database: configApp.Database.Name,
	}
	db := configDB.InitialDB()
	dbMigrate(db)

	configJWT := _middleware.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	e := echo.New()

	iplocatorRepo := _ipLocatorDriver.NewIPLocator()

	categoryRepo := _dbFactory.NewCategoryRepository(db)
	categoryUsecase := _categoryUsecase.NewCategoryUsecase(timeoutContext, categoryRepo)
	categoryCtrl := _categoryController.NewCategoryController(categoryUsecase)

	newsRepo := _dbFactory.NewNewsRepository(db)
	newsUsecase := _newsUsecase.NewNewsUsecase(newsRepo, categoryUsecase, timeoutContext, iplocatorRepo)
	newsCtrl := _newsController.NewNewsController(newsUsecase)

	userRepo := _dbFactory.NewUserRepository(db)
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
