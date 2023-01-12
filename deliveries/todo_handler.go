package deliveries

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"

	"github.com/krittawatcode/go-todo-clean-arch/domains"
	"github.com/krittawatcode/go-todo-clean-arch/models"
	"github.com/krittawatcode/go-todo-clean-arch/utils/jwt"
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

// create login Todo
func (t *ToDoHandler) CreateUserTodo(c *fiber.Ctx) error {
	newUser := new(models.LoginUser)
	if err := c.BodyParser(&newUser); err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}
	if newUser.Name == "" || newUser.Email == "" || newUser.Password == "" {
		return fiber.NewError(fiber.StatusBadRequest, "invalid signup credentials")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	//ต้อง save struct นี้ที่ database
	loginUser := models.LoginUser{
		Name:     newUser.Name,
		Email:    newUser.Email,
		Password: string(hash),
		Role:     newUser.Role,
	}
	//import func form JWT folder
	token, exp, err := jwt.CreateJWTToken(models.LoginUser{})
	if err != nil {
		return err
	}
	err = t.todoUseCase.CreateUserTodo(&loginUser)
	if err != nil {
		return c.SendStatus(http.StatusNotFound)
	}
	return c.JSON(fiber.Map{"token": token, "exp": exp, "user": loginUser})
}

// login user
func (t *ToDoHandler) LoginTodo(c *fiber.Ctx) error {
	reqLogin := new(models.Login)
	if err := c.BodyParser(&reqLogin); err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}
	if reqLogin.Email == "" || reqLogin.Password == "" {
		return fiber.NewError(fiber.StatusBadRequest, "invalid login credentials")
	}
	err := t.todoUseCase.LoginTodo(reqLogin)
	if err != nil {
		return c.SendStatus(http.StatusNotFound)
	} else {
		return c.Status(200).JSON(fiber.Map{"status": "Login Success"})
	}
}

// // middleware
// func (t *ToDoHandler) authentication(c *fiber.Ctx) {
// 	// Get the Authorization header
// 	authHeader := c.Get("Authorization")

// 	// Check if the header is set
// 	if authHeader == "" {
// 		c.Status(401).SendString("Unauthorized")
// 		return
// 	}

// 	// Split the header value into tokens
// 	tokens := strings.Split(authHeader, " ")

// 	// Check if the header value is in the correct format
// 	if len(tokens) != 2 || tokens[0] != "Bearer" {
// 		c.Status(401).SendString("Unauthorized")
// 		return
// 	}

// 	// Validate the token
// 	token := tokens[1]
// 	if token != "valid-token" {
// 		c.Status(401).SendString("Unauthorized")
// 		return
// 	}

// 	// The token is valid, continue with the request
// 	c.Next()
// }
