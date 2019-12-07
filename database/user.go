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

//GetUserByID function mengambil data user
func GetUserByID(id int) (interface{}, error) {
	var users models.User
	if err := db.First(&users, id).Error; err != nil {
		return nil, err
	}
	return users, nil
}

//UpdateUserByID fungstion update dengan id
func UpdateUserByID(id int, newUser *models.User) (interface{}, error) {
	var users models.User

	if err := db.Where("ID = ?", id).Find(&users).Error; err != nil {
		return nil, err
	}
	if err := db.Save(&newUser).Error; err != nil {
		return nil, err
	}
	return newUser, nil

}

//DeleteUser hapus data
func DeleteUser(id int) (interface{}, error) {
	var users models.User

	if err := db.Where("ID = ?", id).Find(&users).Error; err != nil {
		return nil, err
	}
	if err := db.Delete(&users).Error; err != nil {
		return nil, err
	}
	return users, nil

}
