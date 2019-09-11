package main

import (
	"real-estate/ops"
	setupRoutes "real-estate/routes"
)

func main() {

	setupRoutes.Routes()
	ops.Execute()
}
