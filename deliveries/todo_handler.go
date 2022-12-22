package deliveries

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/krittawatcode/go-todo-clean-arch/domains"
	"github.com/krittawatcode/go-todo-clean-arch/models"
)

// ToDoHandler use for handle framwork here and present as a controller
type ToDoHandler struct {
	todoUseCase domains.ToDoUseCase
}

// NewToDoHandler ...
func NewToDoHandler(usecase domains.ToDoUseCase) *ToDoHandler {
	return &ToDoHandler{
		todoUseCase: usecase,
	}
}

// GetAllTodo ...
func (t *ToDoHandler) GetAllTodo(c *fiber.Ctx) error {
	resp, err := t.todoUseCase.GetAllTodo()
	if err != nil {
		return c.SendStatus(http.StatusNotFound)
	} else {
		return c.Status(200).JSON(resp)
	}
}

// CreateATodo ...
func (t *ToDoHandler) CreateATodo(c *fiber.Ctx) error {
	var newToDo models.Todo
	if err := c.BodyParser(&newToDo); err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}
	err := t.todoUseCase.CreateATodo(&newToDo)
	if err != nil {
		return c.SendStatus(http.StatusNotFound)
	} else {
		return c.Status(200).JSON(newToDo)
	}
}

// GetATodo ...
func (t *ToDoHandler) GetATodo(c *fiber.Ctx) error {
	var newToDo models.Todo
	id := c.Params("id")
	err := t.todoUseCase.GetATodo(&newToDo, id)
	if err != nil {
		return c.SendStatus(http.StatusNotFound)
	} else {
		return c.Status(200).JSON(newToDo)

	}
}

// UpdateATodo ...
func (t *ToDoHandler) UpdateATodo(c *fiber.Ctx) error {
	id := c.Params("id")
	var reqToDo models.Todo
	if err := c.BodyParser(&reqToDo); err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}
	err := t.todoUseCase.UpdateATodo(&reqToDo, id)
	if err != nil {
		return c.SendStatus(http.StatusNotFound)
	} else {
		return c.Status(200).JSON(reqToDo)
	}
}

// DeleteATodo ...
func (t *ToDoHandler) DeleteATodo(c *fiber.Ctx) error {
	id := c.Params("id")
	var respToDo models.Todo
	err := t.todoUseCase.DeleteATodo(&respToDo, id)
	if err != nil {
		return c.SendStatus(http.StatusNotFound)
	} else {
		return c.Status(200).JSON(map[string]string{"id" + id: "deleted"})
	}
}
