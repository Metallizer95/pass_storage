package app

import "os"

func getMongoUrl() string {
	var result string
	url := os.Getenv("MONGODB_URL")
	port := os.Getenv("MONGODB_PORT")

	if url == "" || port == "" {
		result = "mongodb://localhost:27017"
	} else {
		result = "mongodb://" + url + ":" + port
	}
	return result
}

func getServerPort() string {
	result := os.Getenv("APP_PORT")
	if result == "" {
		result = ":80"
	}
	return result
}
