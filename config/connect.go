package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func MethodDB() *gorm.DB {

	errEnv := godotenv.Load()

	if errEnv != nil {
		fmt.Println(errEnv)
	}

	host := os.Getenv("HOST")
	pass := os.Getenv("DB_PASS")
	user := os.Getenv("USER")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("PORT_DB")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, pass, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	}

	return db
}
