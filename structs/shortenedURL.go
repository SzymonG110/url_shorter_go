package structs

import "time"

type ShortenedURLRequest struct {
	URL      string  `json:"url" binding:"required" bson:"url"`
	Code     string  `json:"code,omitempty" bson:"code,omitempty"`
	ExpireAt *string `json:"expire_at,omitempty" bson:"expire_at,omitempty"`
}

type ShortenedURL struct {
	URL      string     `json:"url" binding:"required" bson:"url"`
	Code     string     `json:"code,omitempty" bson:"code,omitempty"`
	ExpireAt *time.Time `json:"expire_at,omitempty" bson:"expire_at,omitempty"`
}
