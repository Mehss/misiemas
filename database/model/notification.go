package model

type Notifications struct {
	ID             uint        `gorm:"primaryKey" json:"id"`
	Title          string      `gorm:"type:varchar(255)" json:"title"`
	Description    string      `gorm:"type:text" json:"description"`
	Date           interface{} `gorm:"type:timestamp" json:"date"`
	Source         string      `gorm:"type:varchar(255)" json:"source"`
	Status         string      `gorm:"type:varchar(255)" json:"status"`
	ProjectCode    string      `gorm:"type:varchar(255)" json:"project_code"`
	ProjectName    string      `gorm:"type:varchar(255)" json:"project_name"`
	CompanyCode    string      `gorm:"type:varchar(255)" json:"company_code"`
	CompanyName    string      `gorm:"type:varchar(255)" json:"company_name"`
	CreatedBy      uint        `gorm:"type:integer" json:"created_by"`
	CreatedOn      interface{} `gorm:"type:timestamp" json:"created_on"`
	ModifiedBy     uint        `gorm:"type:integer" json:"modified_by"`
	ModifiedOn     interface{} `gorm:"type:timestamp" json:"modified_on"`
	IsActive       bool        `gorm:"type:boolean" json:"is_active"`
	UserReceiverID uint        `gorm:"type:integer" json:"user_receiver_id"`
	RoleReceiverID uint        `gorm:"type:integer" json:"role_receiver_id"`
}

type NotificationWithCount struct {
	Data []Notifications `json:"data"`
	Meta Meta            `json:"meta"`
}

// TableName sets the schema and table name
func (Notifications) TableName() string {
	return "settings.notifications"
}
