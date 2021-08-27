package database

import (
	"log"
	"os"

	"github.com/alexlai97/mapinfo-kartrider/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const INSTANCE_DIR = "instance"

// sqlite3 file for database
const DATABASE_FILE = "db.sqlite3"

// db is a pointer to the stored gorm.DB
var db *gorm.DB

// GetDB() returns a pointer to the global variable gorm.DB
// if it's empty, try to connect with the database
func GetDB() *gorm.DB {
	// db hasn't been initialized, then try to initialize db
	if db == nil {
		if _, err := os.Stat(INSTANCE_DIR); os.IsNotExist(err) {
			log.Println(INSTANCE_DIR, "directory does not exist")

			err := os.Mkdir(INSTANCE_DIR, 0755)
			if err != nil {
				log.Panicln("creating", INSTANCE_DIR, "failed: ", err.Error())
			}
			log.Println("mkdir", INSTANCE_DIR, "successfully")
		}

		database, err := gorm.Open(sqlite.Open(INSTANCE_DIR+"/"+DATABASE_FILE), &gorm.Config{})
		if err != nil {
			log.Fatalln("GetDB failed, gorm.Open error: ", err.Error())
		}
		log.Println("connect to", INSTANCE_DIR+"/"+DATABASE_FILE, "successfully")

		db = database
		return database
	}
	return db
}

// initDB() does the following:
// - connects to database
// - updates the schema using structs
func InitDB() {
	// try to get database
	database := GetDB()

	// updates the schema using structs
	database.AutoMigrate(&model.Map{}, &model.User{})
}
