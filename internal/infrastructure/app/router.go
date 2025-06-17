package app

import (
	"github.com/darksuei/kubeRPC-sidecar-injector/internal/infrastructure/handlers"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	router.GET("/health", handlers.Health)
	router.POST("/mutate", handlers.Mutate)
	
	return router
}