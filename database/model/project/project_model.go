package project

import (
	"tripatra-dct-service-config/database/model/user"
)

// Project struct
type ProjectModel struct {
	ProjectNo                 string  `gorm:"column:project_no" json:"projectNo,omitempty"`
	ProjectName               string  `gorm:"type:varchar(255)" json:"projectName,omitempty"`
	ProjectGroup              string  `gorm:"type:varchar(255)" json:"projectGroup,omitempty"`
	ProjectType               string  `gorm:"type:varchar(255)" json:"projectType,omitempty"`
	Contract                  string  `gorm:"type:varchar(255)" json:"contract,omitempty"`
	Partnership               string  `gorm:"type:varchar(255)" json:"partnership,omitempty"`
	Facility                  string  `gorm:"type:varchar(255)" json:"facility,omitempty"`
	Client                    string  `gorm:"type:varchar(255)" json:"client,omitempty"`
	PlantLocation             string  `gorm:"type:varchar(255)" json:"plantLocation,omitempty"`
	PlantType                 string  `gorm:"type:varchar(255)" json:"plantType,omitempty"`
	Scope                     string  `gorm:"type:varchar(255)" json:"scope,omitempty"`
	PlantArea                 string  `gorm:"type:varchar(255)" json:"plantArea,omitempty"`
	ContractEffDate           string  `gorm:"type:varchar(255)" json:"contractEffDate,omitempty"`
	ContractCompletionDate    string  `gorm:"type:varchar(255)" json:"contractCompletionDate,omitempty"`
	ContractDuration          string  `gorm:"type:varchar(255)" json:"contractDuration,omitempty"`
	ProjectPlanStartDate      string  `gorm:"type:varchar(255)" json:"projectPlanStartDate,omitempty"`
	ProjectPlanCompletionDate string  `gorm:"type:varchar(255)" json:"projectPlan,omitempty"`
	ProjectDuration           string  `gorm:"type:varchar(255)" json:"projectDuration,omitempty"`
	ProjectDirector           string  `gorm:"type:varchar(255)" json:"projectDirector,omitempty"`
	ProjectManager            string  `gorm:"type:varchar(255)" json:"projectManager,omitempty"`
	ContractValue             string  `gorm:"type:varchar(255)" json:"contractValue,omitempty"`
	ClientLogo                string  `gorm:"type:varchar(255)" json:"clientLogo,omitempty"`
	IDPsgrql                  int32   `gorm:"primaryKey" json:"id_psgrql"`
	Project                   string  `gorm:"type:varchar(255)" json:"project,omitempty"`
	ProjectDefinition         string  `gorm:"type:varchar(255)" json:"projectDefinition,omitempty"`
	IsActive                  string  `gorm:"type:varchar(255)" json:"isActive,omitempty"`
	DaysLeft                  string  `json:"daysLeft,omitempty" `
}

// Project struct
type ProjectPbiModel struct {
	Code        string `json:"code,omitempty"`
	ProjectName string `json:"projectName,omitempty"`
	Category    string `json:"projectCategory,omitempty"`
	SubCategory string `json:"projectSubCategory,omitempty"`
	PageName    string `json:"pageName,omitempty"`
	Url         string `json:"url,omitempty"`
	IsActive    bool   `json:"isActive,omitempty"`
}

type UserProjectModel struct {
  ID int32  `gorm:"primaryKey" json:"id,omitempty"`
  ProjectNo string  `gorm:"type:varchar(255)" json:"projectNo,omitempty"`
  Project ProjectModel `gorm:"foreignKey:ProjectNo;associationForeignKey:ProjectNo,omitempty"`
  Member string  `gorm:"type:varchar(255)" json:"member,omitempty"`
  User user.User `gorm:"foreignKey:Member;associationForeignKey:ID,omitempty"`
  AccessLevel string
}

// TableName sets the schema and table name
func (ProjectModel) TableName() string {
	return "misi_emas.a0_01_corp_project_historical_database"
}

// TableName sets the schema and table name
func (ProjectPbiModel) TableName() string {
	return "misi_emas.pbi_embedded"
}

// TableName sets the schema and table name
func (UserProjectModel) TableName() string {
	return "django_core.authuser_userproject"
}
