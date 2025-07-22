package main

import (
	"flag"
	"fmt"

	"github.com/barcode-api/internal/routes"
	"github.com/barcode-api/middleware/auth"
)

func main() {

	isReleaseMode := flag.Bool("release", false, "Indica se a aplicação está modo de release")
	flag.Parse()

	token, _ := auth.GenerateTokenJwt()
	fmt.Println(token)
	fmt.Println(auth.VerifyToken(token))

	routes.HandlerRoutes(*isReleaseMode)
}
