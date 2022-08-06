package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/czM1K3/database-backupper/databases/mongodb"
	"github.com/czM1K3/database-backupper/databases/postgresql"
	"github.com/czM1K3/database-backupper/folders"
	"github.com/go-co-op/gocron"
)

const (
	MongoDB    string = "MONGODB"
	PostgreSQL        = "POSTGRES"
)

func main() {
	dbType := os.Getenv("DB_TYPE")
	switch dbType {
	case PostgreSQL:
		postgresql.CheckVariables()
	default:
		mongodb.CheckVariables()
	}
	jobType := os.Getenv("JOB_TYPE")
	if jobType == "CRON" {
		interval := os.Getenv("INTERVAL")
		re := regexp.MustCompile(`^(\*|([0-9]|1[0-9]|2[0-9]|3[0-9]|4[0-9]|5[0-9])|\*\/([0-9]|1[0-9]|2[0-9]|3[0-9]|4[0-9]|5[0-9])) (\*|([0-9]|1[0-9]|2[0-3])|\*\/([0-9]|1[0-9]|2[0-3])) (\*|([1-9]|1[0-9]|2[0-9]|3[0-1])|\*\/([1-9]|1[0-9]|2[0-9]|3[0-1])) (\*|([1-9]|1[0-2])|\*\/([1-9]|1[0-2])) (\*|([0-6])|\*\/([0-6]))$`)
		if re.MatchString(interval) {
			s := gocron.NewScheduler(time.Local)
			s.Cron(interval).Do(runBackup)
			s.StartBlocking()
		} else {
			log.Fatal("INTERVAL is not valid")
		}
	} else {
		runBackup()
	}
}

func runBackup() {
	fmt.Println("Starting backup")
	dir := folders.MakeBackupDir()
	dbType := os.Getenv("DB_TYPE")
	switch dbType {
	case PostgreSQL:
		postgresql.DoBackup(dir)
	default:
		mongodb.DoBackup(dir)
	}
	fmt.Println("Backup done")
}
