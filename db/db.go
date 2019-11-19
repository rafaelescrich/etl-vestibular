package db

import (
	"github.com/jinzhu/gorm"
	"github.com/rafaelescrich/etl-vestibular/config"

	// postgres dialect
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB instantiate a global variable with the database connection
var DB *gorm.DB

// Connect to postgres database
func Connect() (err error) {

	cfg := config.Cfg.Database

	dbConfig := cfg.User + ":" + cfg.Password + "@/" + cfg.DBName + "?charset=utf8&parseTime=True&loc=Local"

	DB, err = gorm.Open("mysql", dbConfig)

	return
}
