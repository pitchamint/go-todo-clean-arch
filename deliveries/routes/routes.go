package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/krittawatcode/go-todo-clean-arch/databases"
	"github.com/krittawatcode/go-todo-clean-arch/deliveries"
	"github.com/krittawatcode/go-todo-clean-arch/repositories"
	"github.com/krittawatcode/go-todo-clean-arch/usecases"
)

// SetupRouter ...
func SetupRouter() *fiber.App {

	todoRepo := repositories.NewToDoRepository(databases.DB)
	todoUseCase := usecases.NewToDoUseCase(todoRepo)
	todoHandler := deliveries.NewToDoHandler(todoUseCase)

	// r := gin.Default()
	app := fiber.New()
	api := app.Group("api")
	v1 := api.Group("/v1")
	{
		v1.Get("todo", todoHandler.GetAllTodo)
		v1.Post("todo", todoHandler.CreateATodo)
		v1.Get("todo/:id", todoHandler.GetATodo)
		v1.Put("todo/:id", todoHandler.UpdateATodo)
		v1.Delete("todo/:id", todoHandler.DeleteATodo)
	}
	return app
}
