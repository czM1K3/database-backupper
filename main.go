package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/go-co-op/gocron"
)

func main() {
	jobType := os.Getenv("JOB_TYPE")
	if jobType == "CRON" {
		interval := os.Getenv("INTERVAL")
		re := regexp.MustCompile(`^(\*|([0-9]|1[0-9]|2[0-9]|3[0-9]|4[0-9]|5[0-9])|\*\/([0-9]|1[0-9]|2[0-9]|3[0-9]|4[0-9]|5[0-9])) (\*|([0-9]|1[0-9]|2[0-3])|\*\/([0-9]|1[0-9]|2[0-3])) (\*|([1-9]|1[0-9]|2[0-9]|3[0-1])|\*\/([1-9]|1[0-9]|2[0-9]|3[0-1])) (\*|([1-9]|1[0-2])|\*\/([1-9]|1[0-2])) (\*|([0-6])|\*\/([0-6]))$`)
		if re.MatchString(interval) {
			s := gocron.NewScheduler(time.Local)
			s.Cron(interval).Do(run)
			s.StartBlocking()
		} else {
			log.Fatal("INTERVAL is not valid")
		}
	} else {
		run()
	}
}

func run() {
	fmt.Println("JOB")
}
