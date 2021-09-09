package mysql_driver

import (
	"ca-amartha/drivers/mysql/category"
	"ca-amartha/drivers/mysql/news"
	"ca-amartha/drivers/mysql/users"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ConfigDB struct {
	DB_Username string
	DB_Password string
	DB_Host     string
	DB_Port     string
	DB_Database string
}

func (config *ConfigDB) InitialDB() *gorm.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func DBMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&news.News{},
		&category.Category{},
		&users.Users{},
	)
}
