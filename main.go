package main

import (
	"project-Deteksiwajah/config"
	"project-Deteksiwajah/models"
	"project-Deteksiwajah/routes"
)

func main() {
	e := routes.New()

	InitialMigration()

	e.Logger.Fatal(e.Start(":8000"))
}

//InitialMigration migrasi database
func InitialMigration() {
	var db = config.DB
	if !db.HasTable(&models.User{}) {
		db.AutoMigrate(&models.User{})
	}
}
