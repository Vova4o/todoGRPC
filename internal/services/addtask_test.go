package services

import (
	"context"
	"testing"
	"time"

	"github.com/Vova4o/todogrpc/internal/models"
)

type MockDB struct{}

func (db *MockDB) AddTaskDB(ctx context.Context, task *models.DBTask) (int64, error) {
	// Return a mock value and nil error
	return 1, nil
}

func (db *MockDB) FindTask(ctx context.Context, id string) (*models.DBTask, error) {
	// Return a mock task and nil error
	return &models.DBTask{}, nil
}

func (db *MockDB) UpdateTask(ctx context.Context, task *models.DBTask) error {
	// Return nil error
	return nil
}

func (db *MockDB) Tasks(ctx context.Context, offset int) ([]models.DBTask, error) {
	// Return a slice of mock tasks and nil error
	return []models.DBTask{}, nil
}

func (db *MockDB) SearchTasks(ctx context.Context, title string) ([]models.DBTask, error) {
	// Return a slice of mock tasks and nil error
	return []models.DBTask{}, nil
}

func (db *MockDB) TasksByDate(ctx context.Context, date string) ([]models.DBTask, error) {
	// Return a slice of mock tasks and nil error
	return []models.DBTask{}, nil
}

func (db *MockDB) DoneTask(ctx context.Context, id string) error {
	// Return nil error
	return nil
}

func (db *MockDB) DeleteTask(ctx context.Context, id string) error {
	// Return nil error
	return nil
}

func (db *MockDB) CloseDB() {
	// Do nothing
}

func TestCheckTask(t *testing.T) {
	tests := []struct {
		name    string
		task    *models.DBTask
		wantErr bool
	}{
		{
			name: "Valid task",
			task: &models.DBTask{
				Title:  "Test Task",
				Date:   time.Now().Format("20060102"),
				Repeat: "d 1",
			},
			wantErr: false,
		},
		{
			name: "Task with no title",
			task: &models.DBTask{
				Date:   time.Now().Format("20060102"),
				Repeat: "d 1",
			},
			wantErr: true,
		},
		{
			name: "No date",
			task: &models.DBTask{
				Title:  "Test Task 3",
				Date:   "",
				Repeat: "d 1",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkTask(tt.task); (err != nil) != tt.wantErr {
				t.Errorf("checkTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAddTaskService(t *testing.T) {
	tests := []struct {
		name    string
		task    *models.DBTask
		wantErr bool
	}{
		{
			name: "Valid task",
			task: &models.DBTask{
				Title:  "Test Task",
				Date:   "20220101",
				Repeat: "d 1",
			},
			wantErr: false,
		},
		{
			name: "Task with no title",
			task: &models.DBTask{
				Date:   "20220101",
				Repeat: "d 1",
			},
			wantErr: true,
		},
		// Add more test cases as needed
	}

	s := &Service{
		DB: &MockDB{},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := s.AddTaskService(context.Background(), tt.task)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddTaskService() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
