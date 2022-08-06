package main

import (
	"log"
	"os"
)

func main() {
	mongouri := os.Getenv("MONGODB_URI")
	if mongouri == "" {
		log.Fatal("MONGODB_URI is not defined")
	}
}
