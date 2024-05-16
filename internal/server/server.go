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
	Handler *grpc.Server
	DB      services.TaskService
	// Storage *database.Storage
	// Log     *logger.Logger
}

var Config *ServerConfig

func NewApp() {
	addr := config.Address()

	db, err := database.New()
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()

	pb.RegisterNextDateServiceServer(s, &handles.Server{DB: db})

	Config = &ServerConfig{
		Addr:    addr,
		Handler: s,
		DB:      db,
		// Storage: storage,
		// Log:     log,
	}
}

func (c *ServerConfig) NewServer() *http.Server {
	return &http.Server{
		Addr:    c.Addr,
		Handler: c.Handler,
	}
}

func (c *ServerConfig) StartServer() {
	go func() {
		log.Printf("Starting server on %s\n", c.Addr)
		lis, err := net.Listen("tcp", c.Addr)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		fmt.Printf("server listening at %v\n", lis.Addr())

		if err := c.Handler.Serve(lis); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	c.ShutdownServer(c.NewServer())
}

func (c *ServerConfig) ShutdownServer(srv *http.Server) {
	log.Println("Shutting down server...")

	c.Handler.GracefulStop()

	log.Println("Server gracefully stopped")
}
