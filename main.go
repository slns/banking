package main

import (
	"github.com/slns/banking/app"
	"github.com/slns/banking/app/logger"
)

func main() {

	// log.Println("Starting our application...")
	logger.Info("Starting the application...")
	app.Start()
}
