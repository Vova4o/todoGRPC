package handles

import (
	"context"
	"fmt"
	"log"

	"github.com/Vova4o/todogrpc/internal/models"
	"github.com/Vova4o/todogrpc/internal/services"
	pb "github.com/Vova4o/todogrpc/nextdate/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) AddTask(ctx context.Context, in *pb.AddTaskRequest) (*pb.AddTaskResponse, error) {
	log.Printf("Received: %v", in)

	task := models.DBTask{
		Date:    in.Date,
		Title:   in.Title,
		Comment: in.Comment,
		Repeat:  in.Repeat,
	}

	id, err := services.AddTask(ctx, &task)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Error: %v", err),
		)
	}

	return &pb.AddTaskResponse{Id: id}, err
}
