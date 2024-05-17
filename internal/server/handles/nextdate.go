package handles

import (
	"context"
	"fmt"
	"log"

	"github.com/Vova4o/todogrpc/internal/models"
	pb "github.com/Vova4o/todogrpc/nextdate/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)



func (s *Handlers) NextDate(ctx context.Context, in *pb.NextDateRequest) (*pb.NextDateResponse, error) {
	log.Printf("Received: %v", in)

	task := models.DBTask{
		Date:   in.Date,
		Repeat: in.Repeat,
	}

	newDate, err := s.serviceLevel.NextDateRequest(in.Now, task)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Error: %v", err),
		)
	}

	return &pb.NextDateResponse{Date: newDate}, nil
}
