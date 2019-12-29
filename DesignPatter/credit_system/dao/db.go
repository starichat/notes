package dao

import (
	_ "github.com/go-sql-driver/mysql" // mysql
	"github.com/jinzhu/gorm"
	"github.com/starichat/notes/DesignPatter/credit_sys/config"
	"log"
)

// DB gorm
var DB *gorm.DB

func InitDB() *gorm.DB {
	log.Println( config.DBConfig.URL)

	db, err := gorm.Open(config.DBConfig.Connection, config.DBConfig.URL)
	if err != nil {
		log.Fatal("Database connection failed. Database url : "+config.DBConfig.URL+" error: ",err)
	}
	log.Println("gorm!!!")
	db = db .Set("gorm:table_options", "ENGINE=InnoDB  DEFAULT CHARSET=utf8;").AutoMigrate()

	db.LogMode(config.DBConfig.Debug)
	DB = db

	return db


}