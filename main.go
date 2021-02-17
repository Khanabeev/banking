package main

import (
	"github.com/Khanabeev/banking/app"
	"github.com/Khanabeev/banking/logger"
)

func main() {
	logger.Info("Starting the application ...")
	app.Start()
}
