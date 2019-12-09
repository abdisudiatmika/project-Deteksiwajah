package database

import (
	"project-Deteksiwajah/config"
	"project-Deteksiwajah/models"
	"strconv"
)

var db = config.DB

//Geteyedb function mengambil data ketreangan mata
func Geteyedb() (interface{}, error) {
	var users []models.Usereye

	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

//CreatedbUsereye function menambah keterangan mata
func CreatedbUsereye(usereye *models.Usereye) (interface{}, error) {

	if err := db.Save(&usereye).Error; err != nil {
		return nil, err
	}
	return usereye, nil
}

//FindPersonal fungsi mencari data
func FindPersonal(jarak int) (string, string, error) {
	var users models.Usereye

	//jarak = 813
	//if err := db.Where("jarak = ?", string(jarak)).Find(&users).Error; err != nil {
	//if err := db.Find(&users, jarak).Error; err != nil {
	if err := db.Where("jarak LIKE ?", "%"+strconv.Itoa(jarak)+"%").Find(&users).Error; err != nil {
		//if err := db.First(&users).Error; err != nil {
		return "", "", err
	}
	//fmt.Println(users)
	return users.Keyword, users.Personal, nil

}

//Geteyedbbyid function mengambil keterangan mata dengan id
func Geteyedbbyid(id int) (interface{}, error) {
	var users models.Usereye
	if err := db.First(&users, id).Error; err != nil {
		return nil, err
	}
	return users, nil
}

//Updateeyedb fungstion update keretangan mata dengan id
func Updateeyedb(id int, newUser *models.Usereye) (interface{}, error) {
	var users models.Usereye

	if err := db.Where("ID = ?", id).Find(&users).Error; err != nil {
		return nil, err
	}
	if err := db.Save(&newUser).Error; err != nil {
		return nil, err
	}
	return newUser, nil

}

//Deleteeyedb hapus data keterangan mata
func Deleteeyedb(id int) (interface{}, error) {
	var users models.Usereye

	if err := db.Where("ID = ?", id).Find(&users).Error; err != nil {
		return nil, err
	}
	if err := db.Delete(&users).Error; err != nil {
		return nil, err
	}
	return users, nil

}
