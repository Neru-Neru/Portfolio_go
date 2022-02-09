package main

import (
	"fmt"
	"os"
	"log"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

type Work struct {
	gorm.Model
	Title   string    `json:"title"`
	Image   string    `json:"image"`
	Url     string    `json:"url"`
}

func loadEnv() {
	// .envファイル全体を読み込む関数
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := os.Getenv("DBUSER")
	PASS := os.Getenv("PASSWORD")
	PROTOCOL := "tcp(" + os.Getenv("DOMAIN") + ":" + os.Getenv("PORT") + ")"
	DBNAME := os.Getenv("DBNAME") + "?parseTime=true&loc=Asia%2FTokyo"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME

	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	fmt.Println("db connected: ", &db)
	return db
}

func main() {
	loadEnv()
	db := gormConnect()

	defer db.Close()
	db.LogMode(true)
}