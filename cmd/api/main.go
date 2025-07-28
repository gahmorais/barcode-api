package main

import (
	"flag"
	"fmt"

	"github.com/barcode-api/config"
	"github.com/barcode-api/db"
	"github.com/barcode-api/internal/routes"
)

func main() {

	isReleaseMode := flag.Bool("release", false, "Indica se a aplicação está modo de release")
	flag.Parse()
	env := config.NewEnv()
	strCon := fmt.Sprintf("mongodb://%s:%s@%s:%d", env.User, env.Password, env.Address, env.Port)
	if err := db.InitDb(strCon, env.DatabaseName); err != nil {
		panic(err)
	}
	routes.HandlerRoutes(*isReleaseMode)
}
