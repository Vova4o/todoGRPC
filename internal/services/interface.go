package services

import (
	"context"

	"github.com/Vova4o/todogrpc/internal/models"
)

type Service struct {
	DB DBworker
}

type DBworker interface {
	AddTaskDB(ctx context.Context, task *models.DBTask) (int64, error)
	FindTask(ctx context.Context, id string) (*models.DBTask, error)
	UpdateTask(ctx context.Context, task *models.DBTask) error
	Tasks(ctx context.Context, offset int) ([]models.DBTask, error)
	SearchTasks(ctx context.Context, title string) ([]models.DBTask, error)
	TasksByDate(ctx context.Context, date string) ([]models.DBTask, error)
	DoneTask(ctx context.Context, id string) error
	DeleteTask(ctx context.Context, id string) error
	CloseDB()
}

func NewService(db DBworker) *Service {
	return &Service{DB: db}
}
