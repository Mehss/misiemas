package supplier

type VendorAccountingSAP struct {
	ID                    uint        `gorm:"primaryKey" json:"id"`
	VendorCompanyCode     string      `gorm:"type:varchar(255)" json:"vendor_company_code"`
	VendorCreatedOn       interface{} `gorm:"type:timestamp" json:"vendor_created_on"`
	VendorReconcilAccount string      `gorm:"type:varchar(255)" json:"vendor_reconcil_account"`
	VendorPaymentTerm     string      `gorm:"type:varchar(255)" json:"vendor_payment_term"`
	VendorPaymentMethods  string      `gorm:"type:varchar(255)" json:"vendor_payment_methods"`
	VendorWhtType         string      `gorm:"type:varchar(255)" json:"vendor_wht_type"`
	VendorWhtTypeDesc     string      `gorm:"type:varchar(255)" json:"vendor_wht_type_desc"`
	VendorWhtCode         string      `gorm:"type:varchar(255)" json:"vendor_wht_code"`
	VendorHeaderID        uint        `gorm:"type:integer" json:"vendor_header_id"`
	CreatedBy             uint        `gorm:"type:integer" json:"created_by"`
	CreatedOn             interface{} `gorm:"type:timestamp" json:"created_on"`
	ModifiedBy            uint        `gorm:"type:integer" json:"modified_by"`
	ModifiedOn            interface{} `gorm:"type:timestamp" json:"modified_on"`
	IsActive              bool        `gorm:"type:boolean" json:"is_active"`
}

// TableName sets the schema and table name
func (VendorAccountingSAP) TableName() string {
	return "sap.vendor_accounting_integrate_sap"
}
