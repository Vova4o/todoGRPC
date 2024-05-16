package main

import (
	"fmt"
	"log"

	"github.com/Vova4o/todogrpc/internal/client/handles"
	pb "github.com/Vova4o/todogrpc/nextdate/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {
	creds := insecure.NewCredentials()
	opt := grpc.WithTransportCredentials(creds)
	conn, err := grpc.NewClient(addr, opt)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewNextDateServiceClient(conn)

	newDate := handles.NextDate(c)

	fmt.Println(newDate)
}
