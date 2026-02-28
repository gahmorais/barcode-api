package main

import (
	"flag"
	"log"

	"github.com/barcode-api/config"
	"github.com/barcode-api/internal/database"
	"github.com/barcode-api/internal/routes"
)

type application struct {
	jwtSecret string
}

func main() {
	isReleaseMode := flag.Bool("release", false, "Indica se a aplicação está modo de release")
	flag.Parse()

	env := config.NewEnv()
	if err := database.InitDb(env.ConnectionString(), env.DatabaseName); err != nil {
		log.Fatalf("erro ao iniciar banco de dados: %v", err)
	}

	routes.HandlerRoutes(*isReleaseMode, env.APIPort, env.JWTSecret)
}
