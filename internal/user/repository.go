package user

import (
	"oms-test/database"
	"oms-test/models"
)

type UserRepository struct {}

func (repo *UserRepository) SearchUsers(usernameOrEmail *string) ([]models.User, error) {
	var user []models.User
	result := database.DB.Where("username ilike ? or email ilike ?", usernameOrEmail, usernameOrEmail).Find(&user)
	if result.Error!=nil{
		return []models.User{}, result.Error
	}
	return user, nil
}

func (repo *UserRepository) CheckUsernameEmailUniqueness(username, email *string) bool {
	var user models.User
	result := database.DB.Where("username = ? or email = ?", username, email).Find(&user)
	return result.Error != nil
}

func (repo *UserRepository) CheckUserUniquenessExcludingUserId(username, email *string, id *uint) bool {
	var user models.User
	result := database.DB.Where("(username = ? or email = ?) and id != ?", username, email, id).Find(&user)
	return result.Error != nil
}

func (repo *UserRepository) CreateUser(user *models.User) {
	database.DB.Create(&user)
}

func (repo *UserRepository) GetUser(id string) (*models.User, error) {
	var user models.User
	if result := database.DB.First(&user, id); result.Error != nil {
		return &models.User{}, result.Error
	}
	return &user, nil
}

func (repo *UserRepository) UpdateUser(newUser, oldUser *models.User) {
	database.DB.Model(oldUser).Updates(*newUser)
}

func (repo *UserRepository) DeleteUser(id string) {
	database.DB.Delete(&models.User{}, id)
}
