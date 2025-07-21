package main

import (
	"flag"

	"github.com/barcode-api/internal/routes"
)

func main() {

	isReleaseMode := flag.Bool("release", false, "Indica se a aplicação está modo de release")
	flag.Parse()

	routes.HandlerRoutes(*isReleaseMode)
}
