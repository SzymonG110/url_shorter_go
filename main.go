package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"short.gornikowski.pl/handlers"
	"short.gornikowski.pl/mongodb"
	"time"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)

	// fmt.Printf("%+v\n", JSON)
	// print JSON ^^ XD

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	if err := mongodb.ConnectDB(os.Getenv("MONGODB_URI"), os.Getenv("MONGODB_DATABASE")); err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	r.GET("/all", handlers.AllShortenedURLs)
	r.POST("/create", handlers.ShortenedURL)
	r.GET("/:code", handlers.RedirectToURL)

	go func() {
		for {
			//deleteExpiredShortURLs()
			time.Sleep(10 * time.Minute)
		}
	}()

	if err := r.Run(":" + os.Getenv("PORT")); err != nil {
		log.Fatal(err)
	}
}
