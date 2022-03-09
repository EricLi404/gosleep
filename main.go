package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"
)

func main() {
	fmt.Println("Hello, World!")
	log.Info("Hello, World! logrus")
	time.Sleep(9999 * time.Hour)
}
