package services

import (
	"context"

	"github.com/Vova4o/todogrpc/internal/models"
	"github.com/Vova4o/todogrpc/internal/server"
)

var taskServiceInstance = &TaskService{}

func AddTask(ctx context.Context, task *models.DBTask) (int64, error) {
	id, err := server.Config.DB(ctx, task)
	if err != nil {
		return 0, err
	}

	return id, nil
}
