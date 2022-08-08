package mysql

import (
	"log"
	"os"
	"os/exec"
)

const (
	MySQLUser     = "MYSQL_USER"
	MySQLPassword = "MYSQL_PASSWORD"
	MySQLHost     = "MYSQL_HOST"
	MySQLPort     = "MYSQL_PORT"
	MySQLDatabase = "MYSQL_DATABASE"
)

func CheckVariables() {
	if os.Getenv(MySQLUser) == "" {
		log.Fatal(MySQLUser + " is not defined")
	}
	if os.Getenv(MySQLPassword) == "" {
		log.Fatal(MySQLPassword + " is not defined")
	}
	if os.Getenv(MySQLHost) == "" {
		log.Fatal(MySQLHost + " is not defined")
	}
	if os.Getenv(MySQLPort) == "" {
		log.Fatal(MySQLPort + " is not defined")
	}
	if os.Getenv(MySQLDatabase) == "" {
		log.Fatal(MySQLDatabase + " is not defined")
	}
}

func DoBackup(dir string) {
	mysqlUser := os.Getenv(MySQLUser)
	mysqlPassword := os.Getenv(MySQLPassword)
	mysqlHost := os.Getenv(MySQLHost)
	mysqlPort := os.Getenv(MySQLPort)
	mysqlDatabase := os.Getenv(MySQLDatabase)

	cmd := exec.Command("mysqldump", "--user="+mysqlUser, "--password="+mysqlPassword, "--host="+mysqlHost, "--port="+mysqlPort, "--result-file="+dir+".sql", mysqlDatabase)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
