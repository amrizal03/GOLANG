package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open("root:H4ny4s4tu!@tcp(localhost:3306)/go_jwt_mux"))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{})

	DB = db
}
