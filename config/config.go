package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var (
	//DB Database
	DB *gorm.DB
)

func init() {
	initDB()

}

func initDB() {
	db, err := gorm.Open("sqlite3", "./db/dbusereye.db")
	if err != nil {
		panic(err)
	}
	DB = db

}
