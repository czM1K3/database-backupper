package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/czM1K3/database-backupper/databases/mongodb"
	"github.com/czM1K3/database-backupper/databases/mysql"
	"github.com/czM1K3/database-backupper/databases/postgresql"
	"github.com/czM1K3/database-backupper/path"
	"github.com/go-co-op/gocron"
)

const (
	MongoDB    = "MONGODB"
	PostgreSQL = "POSTGRES"
	MySQL      = "MYSQL"
)

func main() {
	dbType := os.Getenv("DB_TYPE")
	switch dbType {
	case PostgreSQL:
		postgresql.CheckVariables()
	case MySQL:
		mysql.CheckVariables()
	default:
		mongodb.CheckVariables()
	}
	interval := os.Getenv("CRON_INTERVAL")
	if interval != "" {
		re := regexp.MustCompile(`^(\*|([0-9]|1[0-9]|2[0-9]|3[0-9]|4[0-9]|5[0-9])|\*\/([0-9]|1[0-9]|2[0-9]|3[0-9]|4[0-9]|5[0-9])) (\*|([0-9]|1[0-9]|2[0-3])|\*\/([0-9]|1[0-9]|2[0-3])) (\*|([1-9]|1[0-9]|2[0-9]|3[0-1])|\*\/([1-9]|1[0-9]|2[0-9]|3[0-1])) (\*|([1-9]|1[0-2])|\*\/([1-9]|1[0-2])) (\*|([0-6])|\*\/([0-6]))$`)
		if re.MatchString(interval) {
			fmt.Println("Starting cron with interval: " + interval)
			s := gocron.NewScheduler(time.Local)
			s.Cron(interval).Do(runBackup)
			s.StartBlocking()
		} else {
			log.Fatal("CRON_INTERVAL is not valid")
		}
	} else {
		runBackup()
	}
}

func runBackup() {
	fmt.Println("Starting backup")

	dbType := os.Getenv("DB_TYPE")
	switch dbType {
	case PostgreSQL:
		dir := path.MakeBackupPath(false)
		postgresql.DoBackup(dir)
	case MySQL:
		dir := path.MakeBackupPath(false)
		mysql.DoBackup(dir)
	default:
		dir := path.MakeBackupPath(true)
		mongodb.DoBackup(dir)
	}
	fmt.Println("Backup done")
}
