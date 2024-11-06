package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseConnection(){
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/go_jwt_mux"))

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{})

	DB = db
}