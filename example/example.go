package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/codepumper/go-grafana-stack-helpers/logging"
)

func init() {
	path := os.Args[0]
	appName := filepath.Base(path)
	labels := "{application=\"" + appName + "\"}"
	err := logging.InitLogger("json", labels, nil)
	if err != nil {
		fmt.Println("Failed to initialize Logger")
	}
}

func main() {
	defer cleanup()

	// Log some messages
	logging.Debugf("This is a debug message")
	logging.Infof("This is an info message")
	logging.Warnf("This is a warning message")
	logging.Errorf("This is an error message")
}

func cleanup() {
	logging.Shutdown()
}
