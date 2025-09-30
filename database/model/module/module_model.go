package module

import (
	"time"
	"tripatra-dct-service-config/database/model/enum"
)

// Module struct
type Module struct {
	ID          uint              `gorm:"primaryKey" json:"id"`
	ParentID    *uint             `gorm:"type:integer" json:"parent_id,omitempty"` // 0 for parent
	Name        string            `gorm:"type:varchar(255)" json:"name"`
	Slug        string            `gorm:"type:varchar(255)" json:"slug"`
	Description string            `gorm:"type:text" json:"description"`
	ImageURL    string            `gorm:"type:text" json:"image_url"`
	PathName    string            `gorm:"type:varchar(255)" json:"path_name"`
	Type        string            `gorm:"type:varchar(255)" json:"type"`
	Position    enum.EnumPosition `gorm:"type:varchar(255)" json:"position"`
	Order       int               `gorm:"type:integer" json:"order"`
	Constant    string            `gorm:"type:varchar(255)" json:"constant"`
	IsActive    bool              `gorm:"type:boolean" json:"is_active"`
	CreatedBy   int               `gorm:"type:integer" json:"created_by"`
	CreatedOn   time.Time         `gorm:"type:timestamp" json:"created_on"`
	ModifiedBy  int               `gorm:"type:integer" json:"modified_by"`
	ModifiedOn  time.Time         `gorm:"type:timestamp" json:"modified_on"`
	Modules     []*Module         `gorm:"foreignKey:ParentID" json:"modules,omitempty"` // Added for self-referential relationship
}

// TableName sets the schema and table name
func (Module) TableName() string {
	return "user_management.modules"
}
