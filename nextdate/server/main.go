package main

import (
	"github.com/Vova4o/todogrpc/internal/server"
)

func main() {
	server.NewApp()

	server.Config.StartServer()
}
