package model

type QueryParams struct {
	Skip    int      `json:"skip"`
	Take    int      `json:"take"`
	Filters []Filter `json:"filters"`
	Search  string   `json:"search"`
	SortBy  string   `json:"sortBy"`
	Order   string   `json:"order"`
}
