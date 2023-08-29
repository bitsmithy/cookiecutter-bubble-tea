package main

import (
  "flag"
	"fmt"
  "log"
	"os"

  "{{ cookiecutter.app_path }}/internal/tui"
)

const version = "{{ cookiecutter.version }}"

func main() {
  if len(os.Args[1:]) >= 1 {
    printVersion := flag.Bool("v", false, "prints version number of {{ cookiecutter.app_name }}")
    flag.Parse()

    if (*printVersion == true) {
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
