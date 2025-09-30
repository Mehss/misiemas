package utils

import "github.com/gofiber/fiber/v2"

// RespondWithError handles error responses with appropriate status codes.
func RespondWithError(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(fiber.Map{"error": message})
}

// RespondWithSuccess handles successful responses with appropriate status codes.
func RespondWithSuccess(c *fiber.Ctx, statusCode int, data interface{}) error {
	return c.Status(statusCode).JSON(data)
}

// RespondInternalServerError handles 500 Internal Server Error responses.
func RespondInternalServerError(c *fiber.Ctx, message string) error {
	return RespondWithError(c, fiber.StatusInternalServerError, message)
}

// RespondBadRequest handles 400 Bad Request responses.
func RespondBadRequest(c *fiber.Ctx, message string) error {
	return RespondWithError(c, fiber.StatusBadRequest, message)
}

// RespondCreated handles 201 Created responses.
func RespondCreated(c *fiber.Ctx, data interface{}) error {
	return RespondWithSuccess(c, fiber.StatusCreated, data)
}

// RespondOK handles 200 OK responses.
func RespondOK(c *fiber.Ctx, data interface{}) error {
	return RespondWithSuccess(c, fiber.StatusOK, data)
}
