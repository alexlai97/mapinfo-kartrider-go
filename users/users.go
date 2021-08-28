package users

import (
	"github.com/alexlai97/mapinfo-kartrider/database"
	"github.com/alexlai97/mapinfo-kartrider/model"
)

func RegisterUser(user model.User) error {
	result := database.GetDB().Create(&user)
	return result.Error
}
