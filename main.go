package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"

	api "github.com/darksuei/kubeRPC-sidecar-injector/api"
	util "github.com/darksuei/kubeRPC-sidecar-injector/util"
)

func main() {
	godotenv.Load()

	port, err := util.ReadEnv("PORT")

	if err != nil {
		port = util.DEFAULT_PORT
	}

	// Health API
	http.HandleFunc("/health", api.Health)

	log.Println("Sidecar running on port: ", port)

	err = http.ListenAndServe(":" + port, nil)

	if err != nil {
		log.Println("Error starting sidecar:", err)
	}
}