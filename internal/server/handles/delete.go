package handles

import (
	"context"

	pb "github.com/Vova4o/todogrpc/todoproto/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *Handlers) DeleteTask(ctx context.Context, in *pb.DeleteTaskRequest) (*empty.Empty, error) {
	// Delete task by ID
	err := h.serviceLevel.DeleteTaskService(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
