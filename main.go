package main

import (
	"github.com/stefanowolf/go-expense-manager/internal/app"
	"github.com/stefanowolf/go-expense-manager/internal/infra/logger"
)

func main() {
	logger.Info("Starting application...")
	app.Start()
}
