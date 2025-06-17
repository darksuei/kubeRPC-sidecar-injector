package app

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"

	config "github.com/darksuei/kubeRPC-sidecar-injector/config"
	api "github.com/darksuei/kubeRPC-sidecar-injector/internal/infrastructure/api"
)

func Run() {
	_ = godotenv.Load()

	port, err := config.ReadEnv("PORT")

	if err != nil {
		port = config.DEFAULT_PORT
	}

	// Health API
	http.HandleFunc("/health", api.Health)

	log.Println("Sidecar running on port: ", port)

	err = http.ListenAndServe(":" + port, nil)

	if err != nil {
		log.Println("Error starting sidecar:", err)
	}
}