package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"{{ cookiecutter.app_path }}/internal/tui"
)

const (
	version = "{{ cookiecutter.version }}"
	appName = "{{ cookiecutter.app_name }}"
)

func main() {
	if len(os.Args[1:]) >= 1 {
		printVersion := flag.Bool("v", false, fmt.Sprintf("prints version number of %s", appName))
		flag.Parse()

		if *printVersion {
			fmt.Println(version)
			os.Exit(0)
		}
	} else {
		if err := tui.Start(); err != nil {
			log.Fatalf("Error: %s", err)
			os.Exit(1)
		}
	}
}
