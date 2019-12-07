package main

import (
	"go-training-restful/config"
	"go-training-restful/models"
	"go-training-restful/routes"
)

func main() {
	e := routes.New()

	InitialMigration()

	e.Logger.Fatal(e.Start(":8000"))
}

func InitialMigration() {
	var db = config.DB
	if !db.HasTable(&models.User{}) {
		db.AutoMigrate(&models.User{})
	}
}
