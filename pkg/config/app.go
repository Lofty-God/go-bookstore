package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "root:Oraplus@123@/cinema?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}
func GetDB() *gorm.DB {
	return db
}
