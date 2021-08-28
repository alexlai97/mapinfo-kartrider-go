package users

import (
	"github.com/alexlai97/mapinfo-kartrider/database"
	"github.com/alexlai97/mapinfo-kartrider/model"
)

func RegisterUser(user model.User) error {
	result := database.GetDB().Create(&user)
	return result.Error
}

func GetUserByUsername(username string) model.User {
	user := model.User{}
	db := database.GetDB()
	db.Find(&user, db.Where("username = ?", username))
	// TODO: what if not found

	return user
}
