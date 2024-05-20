package main

import (
	"github.com/Vova4o/todogrpc/internal/server"
)

func main() {
	serv := server.NewApp()

	serv.StartServer()
}
