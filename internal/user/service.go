package user

import (
	"errors"
	"oms-test/models"
)

type UserService struct {
	repo UserRepository
}

func (ser *UserService) searchUsers(usernameOrEmail *string) ([]models.User, error) {
	user, err := ser.repo.SearchUsers(usernameOrEmail)
	if err != nil {
		return []models.User{}, err
	}
	return user, nil
}

func (ser *UserService) createUser(user *models.User) (*models.User, error) {
	username, email := user.Username, user.Email
	if ser.repo.CheckUsernameEmailUniqueness(&username, &email) {
		userNotUnique := errors.New("user id or email already exists")
		return &models.User{}, userNotUnique
	}
	ser.repo.CreateUser(user)
	return user, nil
}

func (ser *UserService) getUser(id string) (*models.User, error) {
	user, err := ser.repo.GetUser(id)

	if err != nil {
		return &models.User{}, err
	}

	return user, nil
}

func (ser *UserService) updateUser(old, new *models.User) error {
	is_duplicate := ser.repo.CheckUserUniquenessExcludingUserId(&new.Username, &new.Email, &old.ID)
	
	if is_duplicate {
		return errors.New("user with the same username or email already exists")
	}

	ser.repo.UpdateUser(new, old)
	return nil
}

func (ser *UserService) deleteUser(id string) {
	ser.repo.DeleteUser(id)
}