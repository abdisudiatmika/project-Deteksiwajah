package database

import (
	"go-training-restful/config"
	"go-training-restful/models"
)

var db = config.DB

//GetUsers function mengambil data user
func GetUsers() (interface{}, error) {
	var users []models.User
	//DB.Find(&users) select all from users

	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

//CreateUser function menambah user
func CreateUser(user *models.User) (interface{}, error) {

	if err := db.Save(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
