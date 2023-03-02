package api

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"pos/api/controllers"
)

var server = controllers.Server{}

// Run the server
func Run() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	server.Init(os.Getenv("DB_CONNECTION_STRING"), os.Getenv("KAFKA_URL"), os.Getenv("KAFKA_GROUP_ID"), os.Getenv("KAFKA_TOPIC"))
	server.Run()
}
