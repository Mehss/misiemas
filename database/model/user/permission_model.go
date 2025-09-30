package user

import (
	"tripatra-dct-service-config/database/model"
)

// type Permission struct {
// 	ID uint `gorm:"primaryKey" json:"id"`
// 	// RoleId      int       `gorm:"type:integer" json:"role_id"`
// 	RoleId      uint      `gorm:"type:integer" json:"role_id"`
// 	Role        *Role     `gorm:"foreignKey:RoleId" json:"role,omitempty"`
// 	Name        string    `gorm:"type:varchar(255)" json:"name"`
// 	Description string    `gorm:"type:text" json:"description"`
// 	Constant    string    `gorm:"type:text" json:"constant"`
// 	IsActive    bool      `gorm:"type:boolean" json:"is_active"`
// 	CreatedId   *uint     `json:"created_id,omitempty"`
// 	ModifiedId  *uint     `json:"modified_id,omitempty"`
// 	CreatedBy   *User     `gorm:"foreignKey:CreatedId" json:"created_by,omitempty"`
// 	ModifiedBy  *User     `gorm:"foreignKey:ModifiedId" json:"modified_by,omitempty"`
// 	CreatedAt   interface{} `gorm:"type:timestamp" json:"created_at"`
// 	ModifiedAt  time.Time `gorm:"type:timestamp" json:"modified_at"`
// }

type Permission struct {
	ContentModel  string `gorm:"type:text" json:"content_model"`
	ContentLabel  string `gorm:"type:text" json:"content_label"`
	Name         string `gorm:"type:text" json:"name"`
}

type PermissionWithCount struct {
	Data []Permission `json:"data"`
	Meta model.Meta   `json:"meta"`
}
