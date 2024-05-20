package services

import (
	"context"
	"strconv"
)

func (s *Service) DeleteTaskService(ctx context.Context, id int64) error {
	// Delete task by ID
	idText := strconv.FormatInt(id, 10)

	err := s.DB.DeleteTask(ctx, idText)
	if err != nil {
		return err
	}

	return nil
}
