package services

import (
	"encoding/json"
	"fmt"
	model "tripatra-dct-service-config/database/model"
	projectModel "tripatra-dct-service-config/database/model/project"
	resolver "tripatra-dct-service-config/resolver/project"
	"tripatra-dct-service-config/utils"

	"github.com/gofiber/fiber/v2"
)

// ParseTodoRequest parses and validates the request body for creating or updating a todo.
func ParseProjectService(c *fiber.Ctx) (*projectModel.ProjectModel, error) {
	data := new(projectModel.ProjectModel)
	if err := c.BodyParser(data); err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		fmt.Printf("Request body: %s\n", string(c.Body()))
		return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())

	}
	return data, nil
}

func GetProjectService(c *fiber.Ctx, res *resolver.ProjectResolver) error {
	// Parse query parameters
	skip := c.QueryInt("skip")
	take := c.QueryInt("take")

	var filters []model.Filter
	filterStr := c.Query("filter")
	if filterStr != "" {
		err := json.Unmarshal([]byte(filterStr), &filters)
		if err != nil {
			return utils.RespondWithError(c, 400, "Invalid 'filter' parameter")
		}
	}

	if len(filters) == 0 {
		slug := c.Query("slug")
		if slug != "" {
			filters = append(filters, model.Filter{
				Field:    "slug",
				Operator: "=",
				Value:    slug,
				Type:     "string",
			})
		}
		types := c.Query("position")
		if types != "" {
			filters = append(filters, model.Filter{
				Field:    "position",
				Operator: "=",
				Value:    types,
				Type:     "string",
			})
		}
		// Add more individual query parameters as needed
	}

	datas, totalRecords, err := res.GetProjectList(c, skip, take, filters)
	if err != nil {
    fmt.Println((err))
		return utils.RespondWithError(c, 500, "Error when getting Project")
	}

	fmt.Printf("total records : %v ", totalRecords)

	return utils.RespondWithSuccess(c, 200, datas)
}

func GetProjectPbiService(c *fiber.Ctx, res *resolver.ProjectResolver) error {
	// Parse query parameters
	skip := c.QueryInt("skip")
	take := c.QueryInt("take")

	var filters []model.Filter
	filterStr := c.Query("filter")
	if filterStr != "" {
		err := json.Unmarshal([]byte(filterStr), &filters)
		if err != nil {
			return utils.RespondWithError(c, 400, "Invalid 'filter' parameter")
		}
	}

	if len(filters) == 0 {
		slug := c.Query("slug")
		if slug != "" {
			filters = append(filters, model.Filter{
				Field:    "slug",
				Operator: "=",
				Value:    slug,
				Type:     "string",
			})
		}
		types := c.Query("position")
		if types != "" {
			filters = append(filters, model.Filter{
				Field:    "position",
				Operator: "=",
				Value:    types,
				Type:     "string",
			})
		}
		// Add more individual query parameters as needed
	}

	datas, totalRecords, err := res.GetProjectPbi(c, skip, take, filters)
	if err != nil {
    fmt.Println(err)
		return utils.RespondWithError(c, 500, "Error when getting Project")
	}

	fmt.Printf("total records : %v ", totalRecords)

	return utils.RespondWithSuccess(c, 200, datas)
}
