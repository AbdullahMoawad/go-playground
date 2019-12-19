package main

import (
	"real-estate/ops"
	StartServer "real-estate/routes"
)

func main() {
	StartServer.Routes()
	ops.Execute()
}
