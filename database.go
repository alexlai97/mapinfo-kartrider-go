package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const INSTANCE_DIR = "instance"

// sqlite3 file for database
const DATABASE_FILE = "db.sqlite3"

// json file for map details
const MAPS_JSON_FILE = "./model/maps.json"

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
func initDB() {
	database := GetDB()
	// updates the schema using structs
	database.AutoMigrate(&Map{})
	// insert_default_maps_to_db()
}

// Maps contains [] Map
// it is used for parse from json file
type Maps struct {
	Maps []Map `json:"maps"`
}

// Map contains all information for a certain map
type Map struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	MapName     string `json:"mapname"`
	KoreanName  string `json:"koreanname"`
	ChineseName string `json:"chinesename"`
	Difficulty  string `json:"difficulty"`
	Tierpro     string `json:"tierpro"`
	Tier1       string `json:"tier1"`
	Tier2       string `json:"tier2"`
	Tier3       string `json:"tier3"`
	Tier4       string `json:"tier4"`
}

// construct and return Maps from a json file
func getMapsFromJsonFile(filename string) (maps Maps) {
	var err error
	jsonFile, err := os.Open(filename)
	if err != nil {
		log.Fatalln("load", MAPS_JSON_FILE, "failed", err.Error())
	}
	log.Println(MAPS_JSON_FILE, "successfully loaded")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	// load it to maps
	err = json.Unmarshal([]byte(byteValue), &maps)
	if err != nil {
		log.Fatalln("unmarshal json failed", err.Error())
	}
	return
}

// insert the defail json map file into database
func insertDefaultMapsToDB() {
	GetDB().AutoMigrate(&Map{})
	maps := getMapsFromJsonFile(MAPS_JSON_FILE)
	for _, m := range maps.Maps {
		GetDB().Create(&m)
	}
}

// request DB and return list of all Map detail
func getAllMapsFromDB() (maps []Map) {
	result := GetDB().Find(&maps)
	if err := result.Error; err != nil {
		log.Fatalln("get all maps from db failed")
	}
	return
}

// request DB and return a single Map detail
func getSingleMapFromDB(id int) (m Map) {
	GetDB().First(&m, id)
	return
}
