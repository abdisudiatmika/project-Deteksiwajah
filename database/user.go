package database

import (
	"project-Deteksiwajah/config"
	"project-Deteksiwajah/models"
	"strconv"
)

var db = config.DB

//GetUsers function mengambil data user
func GetUsers() (interface{}, error) {
	var users []models.Usereye
	//DB.Find(&users) select all from users

	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

//CreateUsereye function menambah user
func CreateUsereye(usereye *models.Usereye) (interface{}, error) {

	if err := db.Save(&usereye).Error; err != nil {
		return nil, err
	}
	return usereye, nil
}

//FindPersonal fungsi mencari data
func FindPersonal(jarak int) (string, error) {
	var users models.Usereye

	//jarak = 813
	//if err := db.Where("jarak = ?", string(jarak)).Find(&users).Error; err != nil {
	//if err := db.Find(&users, jarak).Error; err != nil {
	if err := db.Where("jarak LIKE ?", "%"+strconv.Itoa(jarak)+"%").Find(&users).Error; err != nil {
		return "", err
	}
	//fmt.Println(users)
	return users.Personal, nil

}

//GetUserByID function mengambil data user
func GetUserByID(id int) (interface{}, error) {
	var users models.Usereye
	if err := db.First(&users, id).Error; err != nil {
		return nil, err
	}
	return users, nil
}

//UpdateUserByID fungstion update dengan id
func UpdateUserByID(id int, newUser *models.Usereye) (interface{}, error) {
	var users models.Usereye

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
	var users models.Usereye

	if err := db.Where("ID = ?", id).Find(&users).Error; err != nil {
		return nil, err
	}
	if err := db.Delete(&users).Error; err != nil {
		return nil, err
	}
	return users, nil

}
