package main

import (
	_newsUsecase "ca-amartha/bussiness/news"
	_newsController "ca-amartha/controller/news"
	_newsRepo "ca-amartha/driver/database/news"

	_categoryUsecase "ca-amartha/bussiness/category"
	_categoryController "ca-amartha/controller/category"
	_categoryRepo "ca-amartha/driver/database/category"

	_ipLocatorRepo "ca-amartha/driver/thirdparty/iplocator"

	_dbHelper "ca-amartha/helper/database"
	"log"
	"time"

	"github.com/labstack/echo"
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
	configdb := _dbHelper.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}
	db := configdb.InitialDB()
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	e := echo.New()

	ipLocator := _ipLocatorRepo.NewIPLocator()

	categoryRepo := _categoryRepo.NewCategoryRepository(db)
	categoryUsecase := _categoryUsecase.NewCategoryUsecase(timeoutContext, categoryRepo)
	_categoryController.NewCategoryController(e, categoryUsecase)

	newsRepo := _newsRepo.NewMySQLNewsRepository(db)
	newsUsecase := _newsUsecase.NewNewsUsecase(newsRepo, categoryUsecase, ipLocator, timeoutContext)
	_newsController.NewNewsController(e, newsUsecase)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
