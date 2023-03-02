package controllers

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/segmentio/kafka-go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	kafkaHandler "pos/api/kafka"
	"pos/api/models"
)

// Server struct
type Server struct {
	db          *gorm.DB
	kafkaReader *kafka.Reader
	router      *mux.Router
}

// Init function
func (server *Server) Init(DBConnectionString, kafkaURL, kafkaGroupId, kafkaTopic string) {
	var err error
	// Connect to the database
	server.db, err = gorm.Open(postgres.Open(DBConnectionString), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database", err)
	}
	// Auto migrate the database
	err = server.db.AutoMigrate(&models.Collection{}, &models.Item{}, &models.Order{}, &models.LineItem{})
	if err != nil {
		return
	}
	// Create the kafka reader
	server.kafkaReader = kafkaHandler.KafkaReader(kafkaURL, kafkaGroupId, kafkaTopic)

	server.router = mux.NewRouter()
	server.initializeRoutes()
}

func (server *Server) Run() {
	// Start the kafka reader
	defer func(kafkaReader *kafka.Reader) {
		err := kafkaReader.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(server.kafkaReader)
	// Run the kafka reader on a separate go routine
	go server.KafkaWorker()
	// Run the server
	var port string = os.Getenv("PORT")
	fmt.Println("Listening to port " + port)
	log.Fatal(http.ListenAndServe(":"+port, server.router))
}

// KafkaWorker function
func (server *Server) KafkaWorker() {
	// Listen to the kafka reader
	fmt.Println("Listening to kafka")
	for {
		// Read the message from the kafka reader
		msg, err := server.kafkaReader.ReadMessage(context.Background())
		fmt.Println("Received message: ", string(msg.Value))
		if err != nil {
			fmt.Println(err)
			continue
		}
		// Unmarshal the message into a payload

		// Persist the transaction to the database

		if err != nil {
			fmt.Println(err)
			continue
		}
		// Confirm the transaction

	}
}
