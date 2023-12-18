package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"github.com/charmbracelet/log"

	"{{ cookiecutter.app_path }}/internal/build"
	"{{ cookiecutter.app_path }}/internal/tui"
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
		printVersion := flag.Bool("v", false, fmt.Sprintf("prints version info for %s", filepath.Base(os.Args[0])))
		flag.Parse()

		if *printVersion {
			build.PrintInfo()
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
