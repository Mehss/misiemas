package user

import (
	"tripatra-dct-service-config/database/model"
)

// type User struct {
// 	ID         uint      `gorm:"primaryKey" json:"id"`
// 	Role       *Role     `gorm:"foreignKey:UserID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;" json:"role,omitempty"`
// 	Name       string    `gorm:"type:varchar(255)" json:"name"`
// 	Email      string    `gorm:"type:varchar(255);unique" json:"email"`
// 	Password   string    `gorm:"type:varchar(255)" json:"password"`
// 	IsActive   bool      `gorm:"type:boolean" json:"is_active"`
// 	Type       string    `gorm:"type:varchar(255)" json:"type"`
// 	CreatedId  *uint     `json:"created_id,omitempty"`
// 	ModifiedId *uint     `json:"modified_id,omitempty"`
// 	CreatedBy  *User     `gorm:"foreignKey:CreatedId" json:"created_by,omitempty"`
// 	ModifiedBy *User     `gorm:"foreignKey:ModifiedId" json:"modified_by,omitempty"`
// 	ModifiedAt time.Time `gorm:"type:timestamp" json:"modified_at"`
// 	CreatedAt  time.Time `gorm:"type:timestamp" json:"created_at"`
// }

type User struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Email      string    `gorm:"type:varchar(255);unique" json:"email"`
  Permissions []Permission
}

type UserWithCount struct {
	Data []User     `json:"data"`
	Meta model.Meta `json:"meta"`
}

// TableName sets the schema and table name
func (User) TableName() string {
	return "django_core.authuser_customuser"
}
