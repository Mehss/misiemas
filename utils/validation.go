package utils

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// IsAdminOnly checks if the user has admin privileges
func IsAdminOnly(c *fiber.Ctx) error {
	claims, ok := c.UserContext().Value("claims").(*Claims)
	if !ok {
		return fmt.Errorf("unauthorized access") // Custom error message
	}

	if claims.Role != "Admin" {
		return fmt.Errorf("unauthorized access") // Custom error message
	}

	return nil
}

func IsIDSame(c *fiber.Ctx, id string) error {
	claims, ok := c.UserContext().Value("claims").(*Claims)
	if !ok {
		return fmt.Errorf("unauthorized access") // Custom error message
	}

	if claims.Role == "Admin" || claims.UserID == id {
		return nil
	}

	return fmt.Errorf("unauthorized access") // Custom error message
}
