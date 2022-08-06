package folders

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func MakeBackupDir() string {
	currentTime := time.Now().Format("2006-01-02-15-04-05")
	fmt.Println("Creating folder: " + currentTime)
	path := os.Getenv("BACKUP_PATH")
	if path == "" {
		path = "backup/" + currentTime
	} else {
		path = path + "/" + currentTime
	}
	cmd := exec.Command("mkdir", "-p", path)
	cmd.Start()
	return path
}
