package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDB() {
	var connStr string = "host = 127.0.0.1 port = 5432 user = postgres dbname = anime_db password = 1538 sslmode=disable"
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Не удается подключиться к Базе данных по заданным настройкам")
		panic(err)
	}
	db.AutoMigrate(&Anime{})
	DB = db
}
