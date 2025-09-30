package supplier

import "tripatra-dct-service-config/database/model"

type Vendor struct {
	ID             uint        `gorm:"primaryKey" json:"id"`
	VendorCode     string      `gorm:"type:varchar(255)" json:"vendor_code"`
	VendorName     string      `gorm:"type:varchar(255)" json:"vendor_name"`
	PicName        string      `gorm:"type:varchar(255)" json:"pic_name"`
	PicEmail       string      `gorm:"type:varchar(255)" json:"pic_email"`
	Street         string      `gorm:"type:varchar(255)" json:"street"`
	PostalCode     string      `gorm:"type:varchar(255)" json:"postal_code"`
	Country        string      `gorm:"type:varchar(255)" json:"country"`
	Telephone      string      `gorm:"type:varchar(255)" json:"telephone"`
	Fax            string      `gorm:"type:varchar(255)" json:"fax"`
	EmailAddress   string      `gorm:"type:varchar(255)" json:"email_address"`
	SalesPerson    string      `gorm:"type:varchar(255)" json:"sales_person"`
	SalesTelephone string      `gorm:"type:varchar(255)" json:"sales_telephone"`
	CompanyCode    string      `gorm:"type:varchar(255)" json:"company_code"`
	CompanyName    string      `gorm:"type:varchar(255)" json:"company_name"`
	IsActive       bool        `gorm:"type:boolean" json:"is_active"`
	CreatedBy      uint        `gorm:"type:integer" json:"created_by"`
	CreatedOn      interface{} `gorm:"type:timestamp" json:"created_on"`
	ModifiedBy     uint        `gorm:"type:integer" json:"modified_by"`
	ModifiedOn     interface{} `gorm:"type:timestamp" json:"modified_on"`
}

// TableName sets the schema and table name
func (Vendor) TableName() string {
	return "user_management.vendors"
}

// Vendor with username and password
type VendorWithCredentials struct {
	ID             uint        `gorm:"primaryKey" json:"id"`
	Username       string      `json:"username"`
	Password       string      `json:"password"`
	VendorCode     string      `gorm:"type:varchar(255)" json:"vendor_code"`
	VendorName     string      `gorm:"type:varchar(255)" json:"vendor_name"`
	PicName        string      `gorm:"type:varchar(255)" json:"pic_name"`
	PicEmail       string      `gorm:"type:varchar(255)" json:"pic_email"`
	Street         string      `gorm:"type:varchar(255)" json:"street"`
	PostalCode     string      `gorm:"type:varchar(255)" json:"postal_code"`
	Country        string      `gorm:"type:varchar(255)" json:"country"`
	Telephone      string      `gorm:"type:varchar(255)" json:"telephone"`
	Fax            string      `gorm:"type:varchar(255)" json:"fax"`
	EmailAddress   string      `gorm:"type:varchar(255)" json:"email_address"`
	SalesPerson    string      `gorm:"type:varchar(255)" json:"sales_person"`
	SalesTelephone string      `gorm:"type:varchar(255)" json:"sales_telephone"`
	CompanyCode    string      `gorm:"type:varchar(255)" json:"company_code"`
	CompanyName    string      `gorm:"type:varchar(255)" json:"company_name"`
	IsActive       bool        `gorm:"type:boolean" json:"is_active"`
	CreatedBy      uint        `gorm:"type:integer" json:"created_by"`
	CreatedOn      interface{} `gorm:"type:timestamp" json:"created_on"`
	ModifiedBy     uint        `gorm:"type:integer" json:"modified_by"`
	ModifiedOn     interface{} `gorm:"type:timestamp" json:"modified_on"`
}

type VendorHeaderSAP struct {
	ID               uint                  `gorm:"primaryKey" json:"id"`
	VendorCode       string                `gorm:"type:varchar(255)" json:"vendor_code"`
	VendorName       string                `gorm:"type:varchar(255)" json:"vendor_name"`
	VendorCountry    string                `gorm:"type:varchar(255)" json:"vendor_country"`
	VendorEmail      string                `gorm:"type:varchar(255)" json:"vendor_email"`
	VendorTaxNumber1 string                `gorm:"type:varchar(255)" json:"vendor_tax_number1"`
	VendorTaxNumber2 string                `gorm:"type:varchar(255)" json:"vendor_tax_number2"`
	VendorTaxNumber3 string                `gorm:"type:varchar(255)" json:"vendor_tax_number3"`
	CreatedBy        uint                  `gorm:"type:integer" json:"created_by"`
	CreatedOn        interface{}           `gorm:"type:timestamp" json:"created_on"`
	ModifiedBy       uint                  `gorm:"type:integer" json:"modified_by"`
	ModifiedOn       interface{}           `gorm:"type:timestamp" json:"modified_on"`
	IsActive         bool                  `gorm:"type:boolean" json:"is_active"`
	VendorAccounting []VendorAccountingSAP `gorm:"foreignKey:VendorHeaderID" json:"vendor_accounting"`
}

type VendorWithCount struct {
	Data []Vendor   `json:"data"`
	Meta model.Meta `json:"meta"`
}

// TableName sets the schema and table name
func (VendorHeaderSAP) TableName() string {
	return "sap.vendor_header_integrate_sap"
}
