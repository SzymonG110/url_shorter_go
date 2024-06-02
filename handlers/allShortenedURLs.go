package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"short.gornikowski.pl/mongodb"
	"short.gornikowski.pl/structs"
)

func AllShortenedURLs(c *gin.Context) {
	collection := mongodb.GetCollection("shortenedURLs")

	shortenedURLsQuery, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	shortenedURLs := []structs.ShortenedURL{}
	err = shortenedURLsQuery.All(context.TODO(), &shortenedURLs)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message":        "All shortened URLs",
		"shortened_urls": shortenedURLs,
	})
}
