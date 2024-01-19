package config

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect() {
	d, err := gorm.Open("mysql", "root@Oraplus@95@/inventory?charset=utf8&parseTime=true&loc=Africa%2FLagos")
	if err != nil {
		log.Fatal(err)
	}
	db = d
}
func GetDB() *gorm.DB {
	return db
}
