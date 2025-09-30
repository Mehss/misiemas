package projectRepository

import (
	"slices"

	"github.com/gofiber/fiber/v2"

	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
	"tripatra-dct-service-config/database/model"
	"tripatra-dct-service-config/database/model/project"
	"tripatra-dct-service-config/database/model/user"
	"tripatra-dct-service-config/utils"

	"gorm.io/gorm"
)

type ProjectRepository struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{db}
}

func (r *ProjectRepository) GetProjectList(ctx *fiber.Ctx, skip, take int, filters []model.Filter) ([]*project.ProjectModel, int64, error) {
  userData, ok := ctx.UserContext().Value("userData").(*user.User)
  if !ok {
		// Handle the error appropriately, e.g., by returning an error
		return nil, 0, fmt.Errorf("missing or invalid claims")
	}

  permissions := userData.Permissions
  var projectIDs []string
  if len(permissions) > 0 {
    for _, perm := range permissions {
      if perm.ContentLabel == "DCT" && !slices.Contains(projectIDs, perm.ContentModel){
          projectIDs = append(projectIDs, perm.ContentModel)
      }
    }
  }

  if len(projectIDs) == 0 {
    return nil, 0, fmt.Errorf("missing or invalid permission")
  }

  hidden_groups := [...]string{"TPE", "TPEC"}

  projects := []*project.ProjectModel{}
  r.db.Model(&project.ProjectModel{}).
      Where("project_definition IN ?", projectIDs).
      Where("project_group NOT IN ?", hidden_groups).
      Find(&projects)

	items := make([]*project.ProjectModel, 0)
	if len(projects) > 0 {
		for _, v := range projects {
			var (
				item                              project.ProjectModel
				daysLeft                          float64
				err                               error
				contractCompDate, contractEffDate time.Time
				duration                          int
				isAnyDuration                     bool
			)
			today := time.Now().UTC().Add(time.Hour * 7)
			if v.ContractEffDate != "" && v.ContractDuration != "" {
				duration, err = strconv.Atoi(strings.Split(v.ContractDuration, " ")[0])
				if err != nil {
					log.Panicln("Error on parsing contract duration ", err)
				}
				contractEffDate, err = time.Parse(utils.DateLayout, strings.Split(v.ContractEffDate, " ")[0])
				if err != nil {
					log.Panicln("Error on parsing date ", err)
				}
				durationPlusEffdate := contractEffDate.AddDate(0, duration, 0)
				daysLeft = durationPlusEffdate.Sub(today).Hours() / 24
				isAnyDuration = true
				contractCompDate = durationPlusEffdate
			} else if v.ContractCompletionDate != "" {
				contractCompDate, err = time.Parse(utils.DateLayout, v.ContractCompletionDate)
				if err != nil {
					log.Panicln("Error on parsing date ", err)
				}
				daysLeft = contractCompDate.Sub(today).Hours() / 24
				isAnyDuration = true
			} else {
				daysLeft = 0
				isAnyDuration = false
			}

			item.ProjectNo = v.ProjectNo
			item.ProjectName = v.ProjectName
			item.ProjectGroup = v.ProjectGroup
			item.ProjectType = v.ProjectType
			item.Contract = v.Contract
			item.Partnership = v.Partnership
			item.Facility = v.Facility
			item.Client = v.Client
			item.PlantLocation = v.PlantLocation
			item.PlantType = v.PlantType
			item.Scope = v.Scope
			item.PlantArea = v.PlantArea
			item.ContractEffDate = v.ContractEffDate
			item.ContractCompletionDate = v.ContractCompletionDate
			item.ContractDuration = v.ContractDuration
			item.ProjectPlanStartDate = v.ProjectPlanStartDate
			item.ProjectPlanCompletionDate = v.ProjectPlanCompletionDate
			item.ProjectDuration = v.ProjectDuration
			item.ProjectDirector = v.ProjectDirector
			item.ProjectManager = v.ProjectManager
			item.ContractValue = v.ContractValue
			item.ClientLogo = v.ClientLogo
			item.IDPsgrql = v.IDPsgrql
			item.Project = v.Project
			item.ProjectDefinition = v.ProjectDefinition
			item.IsActive = v.IsActive
			if daysLeft > 30 && isAnyDuration {
				daysLeft = daysLeft / 30
				item.DaysLeft = fmt.Sprintf("%.2f Months Left", daysLeft)
			} else if daysLeft > 0 && daysLeft < 30 && isAnyDuration {
				item.DaysLeft = fmt.Sprintf("%.f Days Left", daysLeft)
			} else if isAnyDuration && daysLeft < 1 {
				item.DaysLeft = fmt.Sprintf("Completed %s", contractCompDate.Format("02/01/2006"))
			} else {
				item.DaysLeft = "-"
			}

			items = append(items, &item)
		}
	}

	return items, int64(len(items)), nil
}


func (r *ProjectRepository) GetProjectPbi(ctx *fiber.Ctx, skip, take int, filters []model.Filter) ([]*project.ProjectPbiModel, int64, error) {
	// userData, ok := ctx.UserContext().Value("userData").(*user.User)
  // if !ok {
	// 	// Handle the error appropriately, e.g., by returning an error
	// 	return nil, 0, fmt.Errorf("missing or invalid claims")
	// }
  project_def := ctx.Params("projectDef")

  // var permissions []string
  // if len(userData.Permissions) > 0 {
  //   for _, perm := range userData.Permissions {
  //     if perm.ContentLabel == "DCT" && perm.ContentModel == project_def {
  //       permissions = append(permissions, perm.Name)
  //     }
  //   }
  // }
  
  pbiList := []*project.ProjectPbiModel{}
  query := r.db.Model(&project.ProjectPbiModel{}).Where("code = ?", project_def[4:])
  
  query.Find(&pbiList)
	return pbiList, int64(len(pbiList)), nil
}
