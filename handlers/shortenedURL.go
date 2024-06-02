package handlers

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"short.gornikowski.pl/mongodb"
	"short.gornikowski.pl/structs"
	"short.gornikowski.pl/utils"
	"strings"
	"time"
)

func ShortenedURL(c *gin.Context) {
	var req structs.ShortenedURLRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validRunes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for _, char := range req.Code {
		if !strings.ContainsRune(validRunes, char) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid character in code"})
			return
		}
	}

	invalidCodes := []string{"create", "all"}
	for _, code := range invalidCodes {
		if req.Code == code {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid code"})
			return
		}

	}

	if req.Code == "" {
		req.Code = utils.GenerateShortCode()
	}

	var expireAt *time.Time
	if req.ExpireAt != nil {
		parsedTime, err := utils.ParseDateString(*req.ExpireAt)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		expireAt = &parsedTime
	}

	collection := mongodb.GetCollection("shortenedURLs")

	for {
		result := collection.FindOne(context.TODO(), bson.D{{"code", req.Code}})
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			break
		} else if result.Err() != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Err().Error()})
			return
		}
		req.Code = utils.GenerateShortCode()
	}

	shortURL := bson.M{
		"url":       req.URL,
		"code":      req.Code,
		"expire_at": expireAt,
	}

	if _, err := collection.InsertOne(context.TODO(), shortURL); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, shortURL)
}
