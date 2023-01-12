package auth

import (
	"fmt"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/krittawatcode/go-todo-clean-arch/utils/jwt"
)

// middleware
func Authentication(c *fiber.Ctx) error {
	// Get the Authorization header
	authHeader := c.Get("Authorization")
	fmt.Println("authHeader", authHeader)
	// Check if the header is set
	if authHeader == "" {
		return c.Status(401).SendString("Unauthorized")

	}

	// Split the header value into tokens
	tokens := strings.Split(authHeader, " ")

	// Check if the header value is in the correct format
	if len(tokens) != 2 || tokens[0] != "Bearer" {
		return c.Status(401).SendString("Unauthorized")
	}
	fmt.Println("authHeader2", authHeader)
	// Validate the token
	token := tokens[1]
	valid, err := jwt.ValidateToken(token, "secret")
	if err != nil {
		log.Printf("[middleware] Token validation error : %s\n", err)
		return c.Status(fiber.StatusUnauthorized).JSON(map[string]string{"error": "Unauthorized"})

	}
	if !valid {
		log.Printf("[middleware] Token is Invalide ")
		return c.Status(fiber.StatusUnauthorized).JSON(map[string]string{"error": "Unauthorized"})
	}
	log.Println("token=====>", token)
	UserId, err := jwt.ValidateAndExtractEmail(token, "secret")
	log.Println("userid===>", UserId)
	if err != nil {
		log.Printf("[middleware] Extracting UserId from Token UserId %s\n:", UserId)
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]string{"error": "Internal Server Error"})
	}
	// The token is valid, continue with the request
	return c.Next()
}
