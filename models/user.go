package models

import (
	"fmt"
	Config "simple-api/config"
)

//GetAllUsers Fetch all user data
func GetAllUsers(user *[]Users) (err error) {
	if err = Config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}

//CreateUser ... Insert New data
func CreateUser(user *Users) (err error) {
	if err = Config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

//GetUserByID ... Fetch only one user by Id
func GetUserByID(user *Users, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

//UpdateUser ... Update user
func UpdateUser(user *Users, id string) (err error) {
	fmt.Println(user)
	Config.DB.Save(user)
	return nil
}

//DeleteUser ... Delete user
func DeleteUser(user *Users, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(user)
	return nil
}
