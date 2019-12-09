package main

import (
	"os"
	"project-Deteksiwajah/config"
	"project-Deteksiwajah/models"
	"project-Deteksiwajah/routes"
)

func main() {
	e := routes.New()
	port := os.Getenv("PORT")
	InitialMigration()

	e.Logger.Fatal(e.Start(":" + port))
}

//InitialMigration migrasi database
func InitialMigration() {
	var db = config.DB
	if !db.HasTable(&models.Usereye{}) {
		db.AutoMigrate(&models.Usereye{})
	}
}
