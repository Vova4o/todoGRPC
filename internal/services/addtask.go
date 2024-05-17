package services

import (
	"context"

	"github.com/Vova4o/todogrpc/internal/models"
)

func (s *Service) AddTaskService(ctx context.Context, task *models.DBTask) (int64, error) {
	return 0, nil
}
