package main

import (
	"fmt"

	"github.com/eserilev/migration.winc.services/server"
)

func main() {
	fmt.Println("sandbox.winc.services")
	server.Start()
}
