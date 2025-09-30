package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	DateLayout      = "2006-01-02" // YYYY-MM-DD
	DatetimeLayout  = "2006-01-02 15:04:05"
	DatetimeRFC3339 = "2006-01-02T15:04:05Z07:00"
)

// GetCurrentTimeUTC returns the current time in UTC formatted as a string.
func GetCurrentTimeUTC() string {
	return time.Now().UTC().Format("2006-01-02 15:04:05")
}

// String Utilities

// TrimSpace trims leading and trailing spaces from a string.
func TrimSpace(s string) string {
	return strings.TrimSpace(s)
}

// ToUpper converts a string to uppercase.
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// Join joins a slice of strings into a single string with a separator.
func Join(slice []string, separator string) string {
	return strings.Join(slice, separator)
}

// Validation Utilities

// IsValidEmail checks if the given email is valid.
func IsValidEmail(email string) bool {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(regex)
	return re.MatchString(email)
}

// File Utilities

// ReadFile reads the content of a file and returns it as a string.
func ReadFile(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// FileExists checks if a file exists.
func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

// JSON Utilities

// SerializeToJSON converts a struct to JSON.
func SerializeToJSON(data interface{}) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

// DeserializeFromJSON converts JSON string to a struct.
func DeserializeFromJSON(jsonStr string, data interface{}) error {
	return json.Unmarshal([]byte(jsonStr), data)
}

// Error Handling Utilities

// LogError logs an error with a specific message.
func LogError(err error, msg string) {
	if err != nil {
		log.Printf("ERROR: %s: %s", msg, err)
	}
}

func GetUniqueArrayString(dataInput []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range dataInput {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// Escape single and double quote in sql syntax
func EscapeQuote(str string) string {
	var buf bytes.Buffer
	for _, char := range str {
		switch char {
		case '\'':
			buf.WriteRune('\'')
			// case '"':
			// 	buf.WriteRune('\\')
		}
		buf.WriteRune(char)
	}
	return buf.String()
}

func ConvertArrayIntToString(a []int, delim string) string {
	//return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
	//return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), delim), "[]")
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
}

func CleanNullEmailAddress(emailAddresses []string) []string {
	//Fungsi ini digunakan untuk memfilter alamat email yang null/kosong
	cleanList := []string{}
	for _, address := range emailAddresses {
		if strings.Trim(address, " ") != "" {
			cleanList = append(cleanList, address)
		}
	}
	return cleanList
}

func Contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func StringToUint32(input string) uint32 {
	output, err := strconv.ParseUint(input, 10, 32)
	if err != nil {
		log.Println("[TAP-debug] [err] [utils] [StringToUint32] [ParseUint]: ", err)
	}
	return uint32(output)
}

// ExtractWBS extracts the 4-digit WBS number from a WBS code string.
func ExtractWBS(wbsCode string) string {
	// Compile the regular expression
	re := regexp.MustCompile(`\bS(\d{4})\b`)
	// Find the first match
	match := re.FindStringSubmatch(wbsCode)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}

