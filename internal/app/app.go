package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stefanowolf/go-expense-manager/internal/app/handler"
	"github.com/stefanowolf/go-expense-manager/internal/core/dto"
	"github.com/stefanowolf/go-expense-manager/internal/infra/connector"
	"github.com/stefanowolf/go-expense-manager/internal/infra/logger"
	"log"
	"net/http"
	"os"
)

var (
	router = gin.Default()
)

func Start() {
	sanityCheck()
	connector.StartMySQLConn()
	initRouter()
}

func sanityCheck() {
	envProps := []string{
		"SERVER_ADDRESS",
		"SERVER_PORT",
	}
	for _, k := range envProps {
		if os.Getenv(k) == "" {
			logger.Error(fmt.Sprintf("Environment variable %s not defined. Terminating application...", k))
		}
	}
}

func initRouter() {
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")

	router.GET("/actuator/health", healthCheck)

	rg := router.Group("/api")
	handler.AddUserRoutes(rg)

	log.Fatal(router.Run(fmt.Sprintf("%s:%s", address, port)))
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, dto.HealthCheckResponse{Status: "OK"})
}
