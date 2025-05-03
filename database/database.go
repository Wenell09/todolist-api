package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var PORT string

func ConnectDB() {
	if err := godotenv.Load(); err != nil {
		panic("Error get env")
	}
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	PORT = os.Getenv("PORT")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, PORT, dbName)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("cant connect to DB 🚀")
	}
	fmt.Printf("success connect database 🚀\nhost:%s/%s", host, PORT)
}
