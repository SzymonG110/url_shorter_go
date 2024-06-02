package structs

import (
	"time"
)

type AnalyticsURL struct {
	Code          string                 `bson:"code"`
	Timestamp     time.Time              `bson:"timestamp"`
	Referrer      string                 `bson:"referrer"`
	UserAgent     string                 `bson:"user_agent"`
	IPAddress     string                 `bson:"ip_address"`
	AcceptedLangs string                 `bson:"accepted_angs"`
	IPData        map[string]interface{} `bson:"ip_data"`
}
