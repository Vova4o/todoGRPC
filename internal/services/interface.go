package services

import (
	"context"
	"database/sql"

	"github.com/Vova4o/todogrpc/internal/models"
)

type TaskService interface {
	AddTaskDB(ctx context.Context, task *models.DBTask) (int64, error)
	// FindTask(ctx context.Context, id int64) (*models.DBTask, error)
	// UpdateTask(ctx context.Context, task *models.DBTask) error
	// Tasks(ctx context.Context) ([]*models.DBTask, error)
	// SearchTasks(ctx context.Context, title string, comment string) ([]*models.DBTask, error)
	// TasksByDate(ctx context.Context, date string) ([]*models.DBTask, error)
	// DoneTask(ctx context.Context, id int64, done bool) error
	// DeleteTask(ctx context.Context, id int64) error
}

type DBTaskService struct {
	DB *sql.DB
}

func NewDBTaskService(db *sql.DB) TaskService {
	return &DBTaskService{DB: db}
}
