package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	StringDBconnect = ""
	Port            = 0
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Port = 8080
	}
	// StringDBconnect = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
	StringDBconnect = fmt.Sprintf("%s:%s@tcp(mysql-container-wex:3306)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"),
	)
}
