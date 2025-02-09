package helpers

import (
	"fmt"

	appConfig "LoganXav/sori/configs"
)

func ConnectionUrlBuilder(n string) (string, error){
	var url string

	// Driver options
	switch n {
	case "postgres":		
		// URL for a postgres connection
		url = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			appConfig.GetEnv("DB_HOST"),
			appConfig.GetEnv("DB_PORT"),
			appConfig.GetEnv("DB_USER"),
			appConfig.GetEnv("DB_PASSWORD"),
			appConfig.GetEnv("DB_NAME"),
		)
	case "mysql":		
		// URL for Mysql connection.
		url = fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			appConfig.GetEnv("DB_USER"),
			appConfig.GetEnv("DB_PASSWORD"),
			appConfig.GetEnv("DB_HOST"),
			appConfig.GetEnv("DB_PORT"),
			appConfig.GetEnv("DB_NAME"),
		)
	case "fiber":
		// URL for Fiber connection.
		url = fmt.Sprintf(
			"%s:%s",
			appConfig.GetEnv("SERVER_HOST"),
			appConfig.GetEnv("SERVER_PORT"),
		)
	default:
		// Return error message.
		return "", fmt.Errorf("connection name '%v' is not supported", n)
	}

	// Return connection URL.
	return url, nil
}