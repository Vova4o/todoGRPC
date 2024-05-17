package handles

import (
	pb "github.com/Vova4o/todogrpc/nextdate/proto"
)

func (h *Handlers) AllTasks(in *pb.TaskRequest, stream pb.AllTasksService_AllTasksServer) error {
	tasks, err := h.serviceLevel.AllTasksService(stream.Context(), in.Search)
	if err != nil {
		return err
	}

	// Move to models as converter!
	// convert tasks to protoTasks
	var protoTasks []*pb.AllTasksResponse
	for _, task := range tasks {
		protoTask := &pb.AllTasksResponse{
			Id:      task.ID,
			Title:   task.Title,
			Date:    task.Date,
			Comment: task.Comment,
			Repeat:  task.Repeat,
			// continue mapping all other fields
		}
		protoTasks = append(protoTasks, protoTask)
	}

	for _, task := range protoTasks {
		if err := stream.Send(task); err != nil {
			return err
		}
	}

	return nil
}
