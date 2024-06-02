package utils

import "time"

func ParseDateString(dateString string) (time.Time, error) {
	layout := "2006-01-02 15:04"
	parsedTime, err := time.Parse(layout, dateString)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}
