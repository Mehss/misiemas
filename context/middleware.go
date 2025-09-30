// context/middleware.go
package context

import (
	"context"
	"strings"

	"tripatra-dct-service-config/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// MiddlewareUser sets common headers and passes the context to the next handler
func MiddlewareUser(c *fiber.Ctx) error {
  return c.Next()
	// Check if token is on midleware
	// authHeader := c.Get("Authorization")
	// if authHeader == "" {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Missing Authorization header"})
	// }
  // agent := fiber.Get(os.Getenv("AUTH_HOST")+"/user/oauth2/user/fetch_user/")
  // fmt.Println(os.Getenv("AUTH_HOST"))
  // agent.Set("Authorization", authHeader)
  // status, data, _ := agent.Bytes()
  // if status == 401 {
  //   var err_msg string
  //   json.Unmarshal(data, &err_msg)
  //   return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err_msg})
  // }

  // userData := &user.User{}
  // err := json.Unmarshal(data, &userData)

  // if err != nil {
	// 	log.Println("Error unmarshalling data", err)
	// }
	// ctx := context.WithValue(c.UserContext(), "userData", userData)
	// c.SetUserContext(ctx)

	// return c.Next()
}

// MiddlewareMicroservice - Middleware for microservices
func MiddlewareMicroservice(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing Authorization header"})
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid Authorization header format"})
	}

	token := parts[1]
	claims, err := utils.ValidateJWTMicroservice(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token", "details": err.Error()})
	}

	// Store claims in Fiber context
	ctx := context.WithValue(c.UserContext(), "claims", claims)
	c.SetUserContext(ctx)

	return c.Next()
}

func DatabaseMiddleware(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	}
}
