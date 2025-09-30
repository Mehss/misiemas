package projectResolver

import (
	"tripatra-dct-service-config/database/model"
	"tripatra-dct-service-config/database/model/project"
	projectRepository "tripatra-dct-service-config/database/repository/project"

	"github.com/gofiber/fiber/v2"
)

type ProjectResolver struct {
	repo *projectRepository.ProjectRepository
}

func NewProjectResolver(repo *projectRepository.ProjectRepository) *ProjectResolver {
	return &ProjectResolver{repo}
}

func (r *ProjectResolver) GetProjectList(ctx *fiber.Ctx, skip int, take int, filters []model.Filter) ([]*project.ProjectModel, int64, error) {
	return r.repo.GetProjectList(ctx, skip, take, filters)
}

func (r *ProjectResolver) GetProjectPbi(ctx *fiber.Ctx, skip int, take int, filters []model.Filter) ([]*project.ProjectPbiModel, int64, error) {
	return r.repo.GetProjectPbi(ctx, skip, take, filters)
}