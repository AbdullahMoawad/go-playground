package main

import (
	"fmt"
	"property/ops"
	StartServer "property/routes"
)

func main() {
	ops.Execute()
	StartServer.Routes()
}
