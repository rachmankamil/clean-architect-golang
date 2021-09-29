package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	_handlerBooks "kampus-merdeka-ca/app/presenter/books"
	_servBooks "kampus-merdeka-ca/bussiness/books"
	_repoBooks "kampus-merdeka-ca/repository/mysql/books"

	_servAuthor "kampus-merdeka-ca/bussiness/author"
	_repoAuthor "kampus-merdeka-ca/repository/mysql/author"

	_routes "kampus-merdeka-ca/app/routes"
)

const JWT_SECRET string = "testmvc"
const JWT_EXP int = 1

func InitDB(status string) *gorm.DB {
	db := "kampusmerdeka"
	if status == "testing" {
		db = "kampusmerdeka_test"
	}
	connectionString := fmt.Sprintf("root:masukaja@tcp(mysql-kampus-merdeka:3306)/%s?parseTime=True", db)

	var err error
	DB, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(
		&_repoBooks.Books{},
		&_repoAuthor.Author{},
	)

	return DB
}

func main() {
	db := InitDB("")
	e := echo.New()

	authorRepo := _repoAuthor.NewRepoMySQL(db)
	authorServ := _servAuthor.NewService(authorRepo)

	// factory of domain
	booksRepo := _repoBooks.NewRepoMySQL(db)
	booksService := _servBooks.NewService(booksRepo, nil, authorServ)
	booksHandler := _handlerBooks.NewHandler(booksService)

	// initial of routes
	routesInit := _routes.HandlerList{
		BookHandler: *booksHandler,
	}
	routesInit.RouteRegister(e)

	log.Fatal(e.Start(":8080"))
}
