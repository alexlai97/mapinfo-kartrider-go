package users

import (
	"github.com/alexlai97/mapinfo-kartrider/database"
	"github.com/alexlai97/mapinfo-kartrider/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func RegisterUser(user model.User) error {
	result := database.GetDB().Create(&user)
	return result.Error
}

func GetUserByUsername(username string) model.User {
	user := model.User{}
	db := database.GetDB()
	db.Where("username = ?", username).First(&user)
	// TODO: what if not found

	return user
}

func GetLoggedInUser(c *gin.Context) (model.User, bool) {
	user := model.User{}
	session := sessions.Default(c)
	id := session.Get("account_id")
	if id == nil {
		return user, false
	} else {
		db := database.GetDB()
		db.First(&user, id)
		// TODO: what if not found

		return user, true
	}
}
