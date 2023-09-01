package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type writer struct {
	io.Writer
	timeFormat string
}

func (w writer) Write(b []byte) (n int, err error) {
	return w.Writer.Write(append([]byte(time.Now().Format(w.timeFormat)), b...))
}

func main() {
	logger := log.New(&writer{os.Stdout, "2006/01/02 15:04:05 "}, "[info] ", 0)

	// Set client options
	mongoUri := os.Getenv("database_address")
	if mongoUri == "" {
		port := os.Getenv("MONGODB_PORT_NUMBER")
		if port == "" {
			log.Fatalln("MONGODB_PORT_NUMBER MISSING!")
		}
		mongoUri = fmt.Sprintf("mongodb://localhost:%s", port)
	}

	clientOptions := options.Client().ApplyURI(mongoUri)

	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*500)
	defer cancel()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		logger.Fatal(err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)

	if err != nil {
		logger.Fatal(err)
	}

	logger.Println("Connected to MongoDB!")
}
