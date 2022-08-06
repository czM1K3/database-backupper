package mongodb

import (
	"log"
	"os"
	"os/exec"
)

const (
	MongoDBURI string = "MONGODB_URI"
)

func CheckVariables() {
	if os.Getenv(MongoDBURI) == "" {
		log.Fatal("MONGODB_URI is not defined")
	}
}

func DoBackup(dir string) {
	mongouri := os.Getenv(MongoDBURI)
	cmd := exec.Command("mongodump", "--uri", mongouri, "--out", dir)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
