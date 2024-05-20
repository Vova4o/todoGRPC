package handles

import (
	"context"

	"github.com/Vova4o/todogrpc/internal/models"
	pb "github.com/Vova4o/todogrpc/todoproto/proto"
)

func (h *Handlers) FindId(ctx context.Context, in *pb.FindIdRequest) (*pb.FindIdResponse, error) {
	// Find task by ID
	var task *models.DBTask
	var err error
	task, err = h.serviceLevel.FindTaskById(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	protoTask := &pb.FindIdResponse{
		Id:      task.ID,
		Title:   task.Title,
		Date:    task.Date,
		Comment: task.Comment,
		Repeat:  task.Repeat,
	}

	return protoTask, nil
}
