package handles

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/Vova4o/todogrpc/nextdate/proto"
)

func NextDate(c pb.NextDateServiceClient) string {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := c.NextDate(ctx, &pb.NextDateRequest{
		Now:    "20210304",
		Date:   "20210302",
		Repeat: "d 2",
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Printf("Next date: %s\n", res.Date)
	return res.Date
}
