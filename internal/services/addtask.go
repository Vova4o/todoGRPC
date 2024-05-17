package services

import (
	"context"

	"github.com/Vova4o/todogrpc/internal/models"
)

func (s *Service) AddTaskService(ctx context.Context, task *models.DBTask) (int64, error) {
	res, err := s.DB.AddTaskDB(ctx, task)
	if err != nil {
		return 0, err
	}
	
	return res, nil
}
