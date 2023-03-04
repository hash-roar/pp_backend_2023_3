package userservice

import (
	"pp-bakcend/models"
)

func Update(user *models.User) error {
	return models.UpdateUser(user)
}

func CreateIfNotExist(user *models.User) error {

	findUser, err := models.GetUserByMid(user.Mid)
	if err != nil {
		return err
	}
	if findUser != nil {
		return nil
	}
	return models.CreateUser(user)
}
