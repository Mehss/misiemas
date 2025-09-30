package model

import (
	"database/sql"
)

type ProjectModel struct {
	ProjectNo                 sql.NullString `json:"ProjectNo"`
	ProjectName               sql.NullString `json:"ProjectName"`
	ProjectGroup              sql.NullString `json:"ProjectGroup"`
	ProjectType               sql.NullString `json:"ProjectType"`
	Contract                  sql.NullString `json:"Contract"`
	Partnership               sql.NullString `json:"Partnership"`
	Facility                  sql.NullString `json:"Facility"`
	Client                    sql.NullString `json:"Client"`
	PlantLocation             sql.NullString `json:"PlantLocation"`
	PlantType                 sql.NullString `json:"PlantType"`
	Scope                     sql.NullString `json:"Scope"`
	PlantArea                 sql.NullString `json:"PlantArea"`
	ContractEffDate           sql.NullString `json:"ContractEffDate"`
	ContractCompletionDate    sql.NullString `json:"ContractCompletionDate"`
	ContractDuration          sql.NullString `json:"ContractDuration"`
	ProjectPlanStartDate      sql.NullString `json:"ProjectPlanStartDate"`
	ProjectPlanCompletionDate sql.NullString `json:"ProjectPlanCompletionDate"`
	ProjectDuration           sql.NullString `json:"ProjectDuration"`
	ProjectDirector           sql.NullString `json:"ProjectDirector"`
	ProjectManager            sql.NullString `json:"ProjectManager"`
	ContractValue             sql.NullString `json:"ContractValue"`
	ClientLogo                sql.NullString `json:"ClientLogo"`
	IDPsgrql                  int32          `json:"IDPsgrql"`
	Project                   sql.NullString `json:"Project"`
	ProjectDefinition         sql.NullString `json:"ProjectDefinition"`
	IsActive                  sql.NullBool   `json:"IsActive"`
}