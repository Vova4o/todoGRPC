package services

import (
	"context"
	"log"
	"time"

	"github.com/Vova4o/todogrpc/internal/models"
)

func (s *Service) AllTasksService(ctx context.Context, in string) ([]models.DBTask, error) {
	log.Printf("Service level All tasks was invoked with search query: %s\n", in)
	var tasks []models.DBTask
	var err error

	if in == "" {
		tasks, err = s.DB.Tasks(ctx, 0)
		if err != nil {
			log.Panicln(err)
			return tasks, err
		}
	} else {
		parsedDate, err := time.Parse("02.01.2006", in)
		if err != nil {
			// The search query is not a date, so perform a string search.
			tasks, err = s.DB.SearchTasks(ctx, in)
			if err != nil {
				log.Println(err)
				return tasks, err
			}
		} else {
			// The search query is a date.
			tasks, err = s.DB.TasksByDate(ctx, parsedDate.Format("20060102"))
			if err != nil {
				log.Println(err)
				return tasks, err
			}
		}
	}

	if tasks == nil {
		tasks = []models.DBTask{}
	}

	return tasks, nil
}
