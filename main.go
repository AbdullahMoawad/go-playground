package main

import (
	"github.com/sql-queries/ops"
	setupRoutes "github.com/sql-queries/routes"
)

func main() {
	setupRoutes.Routes()
	ops.Execute()
}
