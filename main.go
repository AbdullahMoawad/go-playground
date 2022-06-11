package main

import (
	"property/ops"
	StartServer "property/routes"
)

func main() {
	ops.Execute()
	StartServer.Routes()
}
