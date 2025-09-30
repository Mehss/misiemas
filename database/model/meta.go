package model

// Add a new structure to hold pagination details
type Meta struct {
	TotalCount  int `json:"totalCount"`  // Total number of records
	CurrentPage int `json:"currentPage"` // Current page number
	PerPage     int `json:"perPage"`     // Number of records per page
}
