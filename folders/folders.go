package folders

import (
	"fmt"
	"os/exec"
	"time"
)

func makeBackupDir() string {
	currentTime := time.Now().Format("2006-01-02-15-04-05")
	fmt.Println("Creating folder: " + currentTime + "\"")
	cmd := exec.Command("mkdir", "-p", "backup/"+currentTime)
	cmd.Start()
	return currentTime
}
