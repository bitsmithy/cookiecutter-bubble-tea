package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/charmbracelet/log"

	"{{ cookiecutter.app_path }}/internal/tui"
	"{{ cookiecutter.app_path }}/internal/version"
)

func main() {
	handler := log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller:    true,
		ReportTimestamp: true,
		TimeFormat:      time.DateTime,
	})
	logger := slog.New(handler)
	slog.SetDefault(logger)

	if len(os.Args[1:]) >= 1 {
		printVersion := flag.Bool("v", false, "prints version info")
		flag.Parse()

		if *printVersion {
			fmt.Printf("version: %s\n", version.Get())
			os.Exit(0)
		}
	} else {
		log.Info("Starting TUI")
		if err := tui.Start(); err != nil {
			log.Fatalf("Error: %s", err)
			os.Exit(1)
		}
	}
}
