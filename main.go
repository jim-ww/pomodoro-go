package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/0xAX/notificator"
)

func main() {
	flag.Usage = func() {
		fmt.Println("Usage: pomodoro [flags] work_duration pause_duration\n")
		flag.PrintDefaults()
		fmt.Printf("\nExamples:\n\tpomodoro 25 5\n\tpomodoro 25m 5m\n\tpomodoro 45m 10m\n\tpomodoro 1h 15m\n\tpomodoro 90s 15s\n\n")
	}
	workText := flag.String("workText", "Back to work!", "work notification text")
	pauseText := flag.String("pauseText", "Time for a break!", "break notification text")
	workIconPath := flag.String("workIconPath", "", "work notification icon")
	pauseIconPath := flag.String("pauseIconPath", "", "break notification icon")
	defaultIconPath := flag.String("defaultIconPath", "", "default notification icon")
	flag.Parse()
	if flag.NArg() != 2 {
		flag.Usage()
		os.Exit(1)
	}

	args := flag.Args()
	parseDuration := func(s string) time.Duration {
		duration, err := time.ParseDuration(s)
		if err == nil {
			return duration
		}
		minutes, err := strconv.Atoi(s)
		if err == nil {
			return time.Minute * time.Duration(minutes)
		}
		flag.Usage()
		log.Fatalf("\n\nfailed to parse duration: %s", err)
		return 0
	}

	workDuration := parseDuration(args[0])
	pauseDuration := parseDuration(args[1])
	if workDuration <= 0 || pauseDuration <= 0 {
		log.Fatal("work_duration and pause_duration must be greater than 0")
	}

	notif := notificator.New(notificator.Options{AppName: "pomodoro", DefaultIcon: *defaultIconPath})
	notify := func(text, iconPath string) {
		if err := notif.Push(text, "", iconPath, notificator.UR_NORMAL); err != nil {
			log.Fatal(err)
		}
	}
	for {
		notify(*workText, *workIconPath)
		time.Sleep(workDuration)
		notify(*pauseText, *pauseIconPath)
		time.Sleep(pauseDuration)
	}
}
