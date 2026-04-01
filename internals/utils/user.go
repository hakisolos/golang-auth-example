package utils

import (
	"errors"
	"example/spark/internals"
	"example/spark/internals/models"
)

func Getuser(id int) (models.User, bool) {
	for _, usr := range internals.Users {
		if id == usr.ID {
			return usr, true

		}
	}
	return models.User{}, false
}

func UserExists(email string) bool {
	for _, usr := range internals.Users {
		if usr.Email == email {
			return true
		}
	}
	return false
}

func DeleteUser(email string) error {
	var res []models.User
	exits := UserExists(email)
	if !exits {
		return errors.New("user does not exist")
	}
	for _, usr := range internals.Users {
		if usr.Email != email {
			res = append(res, usr)
		}
	}
	internals.Users = res
	return nil
}
