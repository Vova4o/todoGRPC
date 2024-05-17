package server

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"

	"github.com/Vova4o/todogrpc/internal/config"
	"github.com/Vova4o/todogrpc/internal/database"
	"github.com/Vova4o/todogrpc/internal/server/handles"
	"github.com/Vova4o/todogrpc/internal/services"
	pb "github.com/Vova4o/todogrpc/nextdate/proto"

	// "github.com/Vova4o/todogrpc/internal/server/handles"
	"google.golang.org/grpc"
	// "github.com/Vova4o/todogrpc/internal/logger"
)

type ServerConfig struct {
	Addr    string
	Handler *handles.Handlers
}

func NewApp() *ServerConfig {
	addr := config.Address()

	db, err := database.New()
	if err != nil {
		log.Fatal(err)
	}

	stor := database.NewStorage(db.Db)
	serv := services.NewService(stor)
	handles := handles.NewServer(serv)

	return &ServerConfig{
		Addr:    addr,
		Handler: handles,
	}
}

func (c *ServerConfig) StartServer() {
	go func() {
		log.Printf("Starting server on %s\n", c.Addr)
		lis, err := net.Listen("tcp", c.Addr)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		if lis != nil {
			fmt.Printf("server listening at %v\n", lis.Addr())
		} else {
			log.Fatalf("Listener is nil")
		}

		s := grpc.NewServer()
		pb.RegisterNextDateServiceServer(s, c.Handler)
		pb.RegisterAddTaskToDBServiceServer(s, c.Handler)
		pb.RegisterAllTasksServiceServer(s, c.Handler)

		if err := s.Serve(lis); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	c.ShutdownServer()
}

func (c *ServerConfig) ShutdownServer() {
	log.Println("Shutting down server...")

	closed := c.Handler.Close()

	log.Printf("Server gracefully stopped %v\n", closed)
}
