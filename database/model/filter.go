package model

// Filter represents a single filter condition
type Filter struct {
	Field    string
	Operator string
	Value    interface{}
	Type     string // "int", "string", "bool", etc.
}
