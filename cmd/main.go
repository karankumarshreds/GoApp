package main

import (
	"github.com/karankumarshreds/GoApp/internal/app"
	"github.com/karankumarshreds/GoApp/internal/pkg/logger"
)

func main() {
	logger.InfoLog("Starting the application...")
	app.Start()
}

