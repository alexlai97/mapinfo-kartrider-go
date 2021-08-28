package model

import "gorm.io/gorm"

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

type User struct {
	gorm.Model
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}
