package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

var buildTime = ""
var gitHash = ""
var goVersion = ""

func main() {
	args := os.Args
	if len(args) == 2 && (args[1] == "--version" || args[1] == "-v") {
		fmt.Printf("Git Commit Hash: %s\n", gitHash)
		fmt.Printf("UTC Build Time : %s\n", buildTime)
		fmt.Printf("Golang Version : %s\n", goVersion)
		return
	}
	fmt.Println("Hello, World!")
	log.Info("Hello, World! logrus")
	time.Sleep(9999 * time.Hour)
}
