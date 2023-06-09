package models

import (
	"github.com/jinzhu/gorm"
	"github.com/RahulManu79/go-bookstore/pkg/config"
)

var db *gorm.DB

type Book struct{
	gorm.Model
	Name string `gorm:"json:name"`
	Author string `gorm:"json:author"`
	Publication string `gorm:"json:publication"`
}

func init() {
	config.Connect()
	db = config.GetDatabase()
	db.AutoMigrate(&Book{})
}

 