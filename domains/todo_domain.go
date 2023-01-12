package domains
//pod ใช้ต่อ adapter
import (
	"github.com/krittawatcode/go-todo-clean-arch/models"
)

// ToDoUseCase ...
type ToDoUseCase interface {
	GetAllTodo() (t []models.Todo, err error)
	CreateATodo(t *models.Todo) (err error)
	GetATodo(t *models.Todo, id string) (err error)
	UpdateATodo(t *models.Todo, id string) (err error)
	DeleteATodo(t *models.Todo, id string) (err error)
	CreateUserTodo(t *models.LoginUser) (err error)
	LoginTodo(t *models.Login) (err error)
}

// ToDoRepository ... กำหนดว่าส่งค่าอะไรมา
type ToDoRepository interface {
	GetAllTodo(t *[]models.Todo) (err error)
	CreateATodo(t *models.Todo) (err error)
	GetATodo(t *models.Todo, id string) (err error)
	UpdateATodo(t *models.Todo, id string) (err error)
	DeleteATodo(t *models.Todo, id string) (err error)
	CreateUserTodo(t *models.LoginUser) (err error)
	LoginTodo(t *models.Login) (err error)
}
