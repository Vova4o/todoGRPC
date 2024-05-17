package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/Vova4o/todogrpc/internal/config"
	"github.com/Vova4o/todogrpc/internal/models"
	"github.com/Vova4o/todogrpc/pkg/datecalc"
	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver
)

// limit — количество задач, возвращаемых за один запрос из базы данных.
const limit = 10

type Database struct {
	Db *sql.DB
}

func NewStorage(db *sql.DB) *Database {
	return &Database{Db: db}
}

// NewStorage создаёт новый объект Storage.
func New() (*Database, error) {
	s := &Database{}
	err := s.InitDB()
	if err != nil {
		return nil, err
	}
	return s, nil
}

// InitDB создаёт базу данных, если она не существует, и возвращает объект sql.DB для работы с ней.
func (s *Database) InitDB() error {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	file := config.DBPath()

	dbFile := filepath.Join(currentDir, file)
	_, err = os.Stat(dbFile)

	var install bool
	if err != nil {
		install = true
	}

	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return err
	}

	log.Printf("Database file: %s\n", dbFile)

	if err = db.Ping(); err != nil {
		return err
	}

	if install {
		_, err = db.Exec(`CREATE TABLE IF NOT EXISTS scheduler (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			date TEXT NOT NULL,
			title TEXT NOT NULL,
			comment TEXT,
			repeat TEXT(128)
		);`)
		if err != nil {
			return err
		}

		_, err = db.Exec(`CREATE INDEX IF NOT EXISTS indexdate ON scheduler (date)`)
		if err != nil {
			log.Println("Не создан индекс", err)
		}
	}

	s.Db = db

	return nil
}

// CloseDB закрывает соединение с базой данных.
func (s *Database) CloseDB() {
	if s != nil {
		s.Db.Close()
	}
}

// AddTask добавляет задачу в базу данных. Возвращает идентификатор задачи.
// исходные данные: дата, заголовок, комментарий, правило повторения.
func (s *Database) AddTaskDB(ctx context.Context, task *models.DBTask) (int64, error) {
	result, err := s.Db.Exec("INSERT INTO scheduler (date, title, comment, repeat) VALUES (?, ?, ?, ?)", task.Date, task.Title, task.Comment, task.Repeat)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

// FindTask ищет задачу по идентификатору ID. Возвращает задачу или ошибку.
func (s *Database) FindTask(ctx context.Context, id string) (*models.DBTask, error) {
	task := models.DBTask{}
	if id == "" {
		return &task, errors.New("не указан id задачи")
	}

	query := "SELECT id, date, title, comment, repeat FROM scheduler WHERE id = ?"
	err := s.Db.QueryRow(query, id).Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err != nil {
		if err == sql.ErrNoRows {
			return &models.DBTask{}, errors.New("задача не найдена")
		}
		return &models.DBTask{}, err
	}

	return &task, nil
}

// UpdateTask обновляет задачу в базе данных. Возвращает ошибку.
func (s *Database) UpdateTask(ctx context.Context, task *models.DBTask) error {
	query := `UPDATE scheduler SET date = ?, title = ?, comment = ?, repeat = ? WHERE id = ?`
	_, err := s.Db.Exec(query, &task.Date, &task.Title, &task.Comment, &task.Repeat, &task.ID)
	if err != nil {
		return errors.New("задача не найдена")
	}

	return nil
}

// Tasks возвращает список задач из базы данных. Возвращает список задач или ошибку.
func (s *Database) Tasks(ctx context.Context, offset int) ([]models.DBTask, error) {
	query := fmt.Sprintf("SELECT id, date, title, comment, repeat FROM scheduler ORDER BY date LIMIT %d OFFSET %d", limit, offset)
	rows, err := s.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.DBTask
	for rows.Next() {
		var t models.DBTask
		err := rows.Scan(&t.ID, &t.Date, &t.Title, &t.Comment, &t.Repeat)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *Database) SearchTasks(ctx context.Context, search string) ([]models.DBTask, error) {
	query := "SELECT id, date, title, comment, repeat FROM scheduler WHERE title LIKE ? OR comment LIKE ?"
	rows, err := s.Db.Query(query, "%"+search+"%", "%"+search+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.DBTask
	for rows.Next() {
		var t models.DBTask
		err := rows.Scan(&t.ID, &t.Date, &t.Title, &t.Comment, &t.Repeat)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func (s *Database) TasksByDate(ctx context.Context, date string) ([]models.DBTask, error) {
	query := "SELECT id, date, title, comment, repeat FROM scheduler WHERE date = ?"
	rows, err := s.Db.Query(query, date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.DBTask
	for rows.Next() {
		var t models.DBTask
		err := rows.Scan(&t.ID, &t.Date, &t.Title, &t.Comment, &t.Repeat)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

// DoneTask помечает задачу как выполненную. Возвращает ошибку. Если задача повторяющаяся, то создаёт новую задачу на следующую дату.
func (s *Database) DoneTask(ctx context.Context, id string) error {
	var taskWeDeleting models.DBTask
	err := s.Db.QueryRow("SELECT id, date, title, comment, repeat FROM scheduler WHERE id = ?", id).Scan(&taskWeDeleting.ID, &taskWeDeleting.Date, &taskWeDeleting.Title, &taskWeDeleting.Comment, &taskWeDeleting.Repeat)
	if err != nil {
		return errors.New("задача не найдена")
	}

	if taskWeDeleting.Repeat == "" {
		_, err := s.Db.Exec("DELETE FROM scheduler WHERE id = ?", id)
		if err != nil {
			return errors.New("задача не найдена")
		}
	} else {
		taskWeDeleting.Date, err = datecalc.NextDate(time.Now(), taskWeDeleting.Date, taskWeDeleting.Repeat)
		if err != nil {
			return err
		}
		err = s.UpdateTask(ctx, &taskWeDeleting)
		if err != nil {
			return err
		}
	}

	return nil
}

// DeleteTask удаляет задачу из базы данных. Возвращает ошибку.
func (s *Database) DeleteTask(ctx context.Context,id string) error {
	var exists bool
	err := s.Db.QueryRow("SELECT exists(SELECT 1 FROM scheduler WHERE id=?)", id).Scan(&exists)
	if err != nil || !exists {
		return fmt.Errorf("задача не найдена")
	}

	_, err = s.Db.Exec("DELETE FROM scheduler WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}
