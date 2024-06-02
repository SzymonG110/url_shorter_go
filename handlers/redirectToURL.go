package handlers

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"io/ioutil"
	"log"
	"net/http"
	"short.gornikowski.pl/mongodb"
	"short.gornikowski.pl/structs"
	"time"
)

func RedirectToURL(c *gin.Context) {
	shortCode := c.Param("code")

	collection := mongodb.GetCollection("shortenedURLs")
	result := collection.FindOne(context.TODO(), bson.D{{"code", shortCode}})
	if result.Err() != nil {
		c.JSON(404, gin.H{
			"error": "URL not found",
		})
		return
	}

	shortenedURL := structs.ShortenedURL{}
	err := result.Decode(&shortenedURL)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	saveClickAnalytics(shortCode, c.Request)

	c.Redirect(301, shortenedURL.URL)
}

func saveClickAnalytics(shortURL string, r *http.Request) {
	url := "http://ip-api.com/json/" + r.RemoteAddr + "?fields=status,message,continent,country,regionName,city,district,zip,lat,lon,timezone,isp,org,as,asname,reverse,proxy,hosting&lang=en"
	response, _ := http.Get(url)
	body, _ := ioutil.ReadAll(response.Body)

	ipAPIResp := map[string]interface{}{}
	json.Unmarshal(body, &ipAPIResp)

	analytic := structs.AnalyticsURL{
		Code:          shortURL,
		Timestamp:     time.Now(),
		Referrer:      r.Referer(),
		UserAgent:     r.UserAgent(),
		IPAddress:     r.RemoteAddr,
		AcceptedLangs: r.Header.Get("Accept-Language"),
		IPData:        ipAPIResp,
	}

	collection := mongodb.GetCollection("analytics")
	_, err := collection.InsertOne(context.TODO(), analytic)
	if err != nil {
		log.Println("Error saving click analytics for code: " + analytic.Code)
		log.Println("Error content: " + err.Error())
	}
}
