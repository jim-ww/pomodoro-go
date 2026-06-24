package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/0xAX/notificator"
)

const appName = "pomodoro"

func main() {
	flag.Usage = func() {
		fmt.Printf("Usage:\n%s [flags] work_duration pause_duration\n\n", appName)
		flag.PrintDefaults()
		fmt.Printf("\nExamples:\n\t%s 25 5\n\t%s 25m 5m\n\t%s 45m 10m\n\t%s 1h 15m\n\t%s 90s 15s\n\n", appName, appName, appName, appName, appName)
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
		fmt.Printf("failed to parse duration: %s\n\n", err)
		flag.Usage()
		os.Exit(1)
		return 0
	}

	workDuration := parseDuration(args[0])
	pauseDuration := parseDuration(args[1])

	if workDuration <= 0 || pauseDuration <= 0 {
		fmt.Println("work_duration and pause_duration must be greater than 0")
		os.Exit(1)
	}

	if err := startPomodoro(config{
		WorkDuration:  workDuration,
		PauseDuration: pauseDuration,
		WorkText:      *workText,
		PauseText:     *pauseText,
		DefaultIcon:   *defaultIconPath,
		WorkIcon:      *workIconPath,
		PauseIcon:     *pauseIconPath,
	}); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type config struct {
	WorkDuration  time.Duration
	PauseDuration time.Duration

	WorkText  string
	PauseText string

	DefaultIcon string
	WorkIcon    string
	PauseIcon   string
}

func startPomodoro(cfg config) error {
	notif := notificator.New(notificator.Options{AppName: appName, DefaultIcon: cfg.DefaultIcon})
	pushWork := func() error {
		return notif.Push(cfg.WorkText, "", cfg.WorkIcon, notificator.UR_NORMAL)
	}
	pushPause := func() error {
		return notif.Push(cfg.PauseText, "", cfg.PauseIcon, notificator.UR_NORMAL)
	}

	if err := pushWork(); err != nil {
		return fmt.Errorf("push notification: %w", err)
	}

	working := true
	ticker := time.NewTicker(cfg.WorkDuration)
	defer ticker.Stop()

	for range ticker.C {
		if working {
			if err := pushPause(); err != nil {
				return fmt.Errorf("push notification: %w", err)
			}
			ticker.Reset(cfg.PauseDuration)
		} else {
			if err := pushWork(); err != nil {
				return fmt.Errorf("push notification: %w", err)
			}
			ticker.Reset(cfg.WorkDuration)
		}
		working = !working
	}
	return nil
}
