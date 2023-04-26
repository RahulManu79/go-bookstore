package config

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


var (
	db *gorm.DB
)

func Connect(){
	d, err := gorm.Open("mysql","rahul:rahul2742@/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err!= nil {
        panic(err)
    } 
	db = d
}

func GetDatabase() *gorm.DB {
	return db
}