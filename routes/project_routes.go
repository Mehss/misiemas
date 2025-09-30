package routes

import (
	"github.com/gofiber/fiber/v2"

	resolver "tripatra-dct-service-config/resolver/project"
	"tripatra-dct-service-config/services"
)

func ProjectApi(api fiber.Router, res *resolver.ProjectResolver) {
	api.Get("/projects", func(c *fiber.Ctx) error {
		return services.GetProjectService(c, res)
	})
	api.Get("/projects/:projectDef", func(c *fiber.Ctx) error {
		return services.GetProjectPbiService(c, res)
	})
}
