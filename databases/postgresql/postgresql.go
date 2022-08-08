package postgresql

import (
	"log"
	"os"
	"os/exec"
)

const (
	PostgreSQLURI = "POSTGRESQL_URI"
)

func CheckVariables() {
	if os.Getenv(PostgreSQLURI) == "" {
		log.Fatal(PostgreSQLURI + " is not defined")
	}
}

func DoBackup(dir string) {
	postgresqluri := os.Getenv(PostgreSQLURI)
	cmd := exec.Command("pg_dump", "--dbname="+postgresqluri, "--file="+dir+".sql")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
