package utils

import (
	"fmt"
	"strings"
	"time"
	"tripatra-dct-service-config/database/model"

	"gorm.io/gorm"
)

func GenerateFilter(query *gorm.DB, filters []model.Filter) (*gorm.DB, error) {
	for _, filter := range filters {
		if filter.Value == nil {
			continue
		}

		// Handle multiple values for filters
		values, ok := filter.Value.([]interface{})
		if !ok || len(values) == 0 {
			continue
		}

		switch filter.Type {
		case "string":
			if filter.Operator == "LIKE" {
				for _, val := range values {
					query = query.Where(fmt.Sprintf("%s LIKE ?", filter.Field), "%"+val.(string)+"%")
				}
			} else {
				query = query.Where(fmt.Sprintf("%s %s ?", filter.Field, filter.Operator), values)
			}

		case "int", "bool":
			for i, val := range values {
				if strVal, ok := val.(string); ok {
					values[i] = strings.ReplaceAll(strVal, ",", ".")
				}
			}
			query = query.Where(fmt.Sprintf("%s %s ?", filter.Field, filter.Operator), values)

		case "date":
			if len(values) == 2 {
				// Handle BETWEEN for date range
				startDate, err1 := parseDate(values[0])
				endDate, err2 := parseDate(values[1])
				if err1 != nil || err2 != nil {
					return nil, fmt.Errorf("invalid date range: %v, %v", err1, err2)
				}
				query = query.Where(fmt.Sprintf("DATE(%s) BETWEEN ? AND ?", filter.Field), startDate, endDate)
			} else if len(values) == 1 {
				// Handle single date filter
				singleDate, err := parseDate(values[0])
				if err != nil {
					return nil, fmt.Errorf("invalid date: %v", err)
				}
				switch filter.Operator {
				case "=":
					query = query.Where(fmt.Sprintf("DATE(%s) = ?", filter.Field), singleDate)
				case ">=":
					query = query.Where(fmt.Sprintf("DATE(%s) >= ?", filter.Field), singleDate)
				case "<=":
					query = query.Where(fmt.Sprintf("DATE(%s) <= ?", filter.Field), singleDate)
				default:
					return nil, fmt.Errorf("unsupported operator for date: %s", filter.Operator)
				}
			}

		case "dropdown":
			query = query.Where(fmt.Sprintf("%s IN (?)", filter.Field), values)

		default:
			return nil, fmt.Errorf("unsupported filter type: %s", filter.Type)
		}
	}

	return query, nil
}

// Helper function to parse dates in "DD/MM/YYYY" format
func parseDate(value interface{}) (string, error) {
	dateStr, ok := value.(string)
	if !ok {
		return "", fmt.Errorf("invalid date value: %v", value)
	}
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return "", fmt.Errorf("invalid date format: %v", err)
	}
	return date.Format("2006-01-02"), nil
}

// Mapping Filter
func FilterMapping(queryParams map[string]string) []model.Filter {
	var filters []model.Filter

	// Iterate over all query parameters
	for key, values := range queryParams {
		// Check if the parameter has values
		if values != "" {
			filterValues := strings.Split(values, ",") // Split by comma for multiple values
			filters = append(filters, model.Filter{
				Field:    key,  // Use the query parameter key directly
				Operator: "IN", // Assuming we want to check if the value is in the filter values
				Value:    filterValues,
				Type:     "string", // Assuming all values are strings; adjust as needed
			})
		}
	}

	return filters
}

// AddUserFilter adds a filter for the created_by field based on user claims
func AddUserFilter(filters []model.Filter, userIDUint uint64) []model.Filter {
	createdByFilter := model.Filter{
		Field:    "created_by",              // Adjust this according to your model
		Operator: "=",                       // Assuming you want to filter by equality
		Value:    []interface{}{userIDUint}, // Wrap the uint in a slice of interface{}
		Type:     "int",                     // Type of the filter
	}
	filters = append(filters, createdByFilter)
	return filters
}

// AddUNotificationFilter adds a filter for the userReceiverID field based on user claims
func AddNotificationFilter(filters []model.Filter, userIDUint uint64) []model.Filter {
	createdByFilter := model.Filter{
		Field:    "user_receiver_id",        // Adjust this according to your model
		Operator: "=",                       // Assuming you want to filter by equality
		Value:    []interface{}{userIDUint}, // Wrap the uint in a slice of interface{}
		Type:     "int",                     // Type of the filter
	}
	filters = append(filters, createdByFilter)
	return filters
}
