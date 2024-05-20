package handles

import (
	"context"

	"github.com/Vova4o/todogrpc/internal/models"
	pb "github.com/Vova4o/todogrpc/todoproto/proto"
)

type Handlers struct {
	// вместо того чтобы впихивать по одному мы воткнули весь инерфейс Proto
	pb.TodoProtoServiceServer
	// далее прикручиваем структуру слоя сервис к Serviser interface для его использования
	serviceLevel Serviser
}

// я создаю тут интерыейс service, чтобы в него включить все методы, которые я хочу использовать в main.go
type Serviser interface {
	NextDateRequest(nowRequest string, task models.DBTask) (string, error)
	AddTaskService(ctx context.Context, task *models.DBTask) (int64, error)
	AllTasksService(ctx context.Context, in string) ([]models.DBTask, error)
	FindTaskById(ctx context.Context, id int64) (*models.DBTask, error)
	DeleteTaskService(ctx context.Context, id int64) error
	Close() string
}

// s *Handlers — это структура, которая включает в себя все методы из Proto и методы из Serviser
// вот таким образом я могу использовать методы из Proto и методы из Serviser в main.go
func NewServer(s Serviser) *Handlers {
	return &Handlers{serviceLevel: s}
}
