package model

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// json file for map details
const MAPS_JSON_FILE = "./maps.json"

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
	GetDB().AutoMigrate(&Map{})
	maps := getMapsFromJsonFile(MAPS_JSON_FILE)
	for _, m := range maps.Maps {
		GetDB().Create(&m)
	}
}

// request DB and return list of all Map detail
func GetAllMapsFromDB() (maps []Map) {
	result := GetDB().Find(&maps)
	if err := result.Error; err != nil {
		log.Fatalln("get all maps from db failed")
	}
	return
}

// request DB and return a single Map detail
func GetSingleMapFromDB(id int) (m Map) {
	GetDB().First(&m, id)
	return
}
