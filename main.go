package main

import (
	"github.com/joho/godotenv"
	"log"
)

func main() {
	// fmt.Printf("%+v\n", JSON)
	// ^^ print json XD

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}
