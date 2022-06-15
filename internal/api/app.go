package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stefanowolf/go-expense-manager/internal/infra/logger"
	"log"
	"net/http"
	"os"
)

func Start() {
	sanityCheck()
	router := mux.NewRouter()

	ah := ActuatorHandler{}
	router.HandleFunc("/actuator/health", ah.healthCheck).Methods(http.MethodGet)

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
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
