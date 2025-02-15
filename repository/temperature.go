package repository

import (
	"encoding/json"
	"github.com/mahdi-cpp/api-go-emqx/config"
	"github.com/mahdi-cpp/api-go-emqx/model"
	"time"
)

// ReadTemperatures retrieves all temperature records from the database
func ReadTemperatures() (string, error) {
	var temperatures []model.Temperature
	if err := config.DB.Limit(60).Find(&temperatures).Error; err != nil {
		return "", err
	}

	// Convert the slice of Temperature structs to JSON
	jsonData, err := json.Marshal(temperatures)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

// ReadRecentTemperatures retrieves temperature records from the last hour and returns them as a JSON string
func ReadRecentTemperatures() (string, error) {
	var temperatures []model.Temperature
	oneHourAgo := time.Now().Add(-1 * time.Hour) // Get the time one hour ago

	// Query to find records created in the last hour
	if err := config.DB.Where("created_at >= ?", oneHourAgo).Find(&temperatures).Error; err != nil {
		return "", err
	}

	// Convert the slice of Temperature structs to JSON
	jsonData, err := json.Marshal(temperatures)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}
