package config

import (
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func connect() {
	d, err := gorm.Open("mysql", "greatgod@Oraplus@95@/inventory? charset=utf8&parseTime=true&loc=local")
	if err != nil {
		panic(err)
	}
	db = d
}
func GetDB() *gorm.DB {
	return db
}
