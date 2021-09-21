package main

import (
	"log"

	"github.com/joho/godotenv"
	"gizumon.com/life-manager-api/routes"
)

func main() {
	LoadEnv()
	e := routes.Init()
	e.Logger.Fatal(e.Start(":1323"))
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
