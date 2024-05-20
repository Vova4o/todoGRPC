package services

import (
	"context"
	"strconv"

	"github.com/Vova4o/todogrpc/internal/models"
)

func (s *Service) FindTaskById(ctx context.Context, id int64) (*models.DBTask, error) {
	idText := strconv.FormatInt(id, 10)

	task, err := s.DB.FindTask(ctx, idText)
	if err != nil {
		return &models.DBTask{}, err
	}

	return task, nil
}
