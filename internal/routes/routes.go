package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func HandlerRoutes(isReleaseMode bool) {
	port := "1111"
	if isReleaseMode {
		gin.SetMode(gin.ReleaseMode)
		fmt.Printf("Servidor iniciado na porta %s", port)
	}

	r := gin.Default()

	UserRoutes(r)
	ProductRoutes(r)

	r.Run(":" + port)
}
