package routes

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func HandlerRoutes(isReleaseMode bool, port int, jwtSecret string) {
	if isReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	address := fmt.Sprintf(":%d", port)
	fmt.Printf("Servidor iniciado na porta %d\n", port)

	r := gin.Default()

	UserRoutes(r, jwtSecret)
	ProductRoutes(r, jwtSecret)

	if err := r.Run(address); err != nil {
		log.Fatalf("erro ao iniciar servidor: %v", err)
	}
}
