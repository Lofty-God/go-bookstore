package config
import(
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/mysql"
)
var db *gorm.DB
d, err:=gorm.Open("mysql", "greatgod@Oraplus@95@/inventory? charset=utf8&parseTime=true&loc=local")
if err!=nil{
	panic(err)
}
db=d
func getDB() *gorm.DB{
	return db
}