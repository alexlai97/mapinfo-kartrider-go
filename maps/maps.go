package maps

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/alexlai97/mapinfo-kartrider/database"
	"github.com/alexlai97/mapinfo-kartrider/model"
)

// json file for map details
const MAPS_JSON_FILE = "./maps.json"

// construct and return Maps from a json file
func getMapsFromJsonFile(filename string) (maps model.Maps) {
	var err error
	jsonFile, err := os.Open(filename)
	if err != nil {
		log.Fatalln("os.Open", MAPS_JSON_FILE, "failed", err.Error())
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
func InsertDefaultMapsToDB() {
	database.GetDB().AutoMigrate(&model.Map{})
	maps := getMapsFromJsonFile(MAPS_JSON_FILE)
	for _, m := range maps.Maps {
		result := database.GetDB().Create(&m)
		if err := result.Error; err != nil {
			log.Fatalln("could not create map: ", err.Error())
		}
	}
}

// request DB and return list of all Map detail
func GetAllMapsFromDB() (maps []model.Map) {
	result := database.GetDB().Find(&maps)
	if err := result.Error; err != nil {
		log.Fatalln("get all maps from db failed")
	}
	return
}

// request DB and return a single Map detail
func GetSingleMapFromDB(id int) (m model.Map) {
	database.GetDB().First(&m, id)
	return
}
