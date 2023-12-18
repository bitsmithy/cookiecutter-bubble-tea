package build

import (
	"fmt"
	"strings"
)

var (
	Version string
	Commit  string
	Time    string
	User    string
)

func isReleaseBuild() bool {
	return len(strings.TrimSpace(Version)) != 0
}

func PrintInfo() {
	if isReleaseBuild() {
		fmt.Println("Version:\t", Version)
		fmt.Println("Commit SHA:\t", Commit)
		fmt.Println("Built At:\t", Time)
		fmt.Println("Built By:\t", User)
	} else {
		fmt.Println("Version: ", "development")
	}
}
