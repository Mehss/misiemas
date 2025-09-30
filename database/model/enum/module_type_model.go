package enum

// Define a custom type for the enum

// Define constants for each enum value
// const (
// 	TOP_URL    ModuleType = "TOP_URL"
// 	BOTTOM_URL ModuleType = "BOTTOM_URL"
// 	LEFT_URL   ModuleType = "LEFT_URL"
// 	RIGHT_URL  ModuleType = "RIGHT_URL"
// 	MIDDLE_URL ModuleType = "MIDDLE_URL"
// 	BUTTON     ModuleType = "BUTTON"
// 	INPUT      ModuleType = "INPUT"
// 	TEXT       ModuleType = "TEXT"
// )

type EnumPosition string

const (
	TOP_URL    EnumPosition = "TOP_URL"
	BOTTOM_URL EnumPosition = "BOTTOM_URL"
	LEFT_URL   EnumPosition = "LEFT_URL"
	RIGHT_URL  EnumPosition = "RIGHT_URL"
	MIDDLE_URL EnumPosition = "MIDDLE_URL"
	BUTTON     EnumPosition = "BUTTON"
	INPUT      EnumPosition = "INPUT"
	TEXT       EnumPosition = "TEXT"
)

func (mt EnumPosition) IsValid() bool {
	switch mt {
	case TOP_URL, BOTTOM_URL, LEFT_URL, RIGHT_URL, MIDDLE_URL, BUTTON, INPUT, TEXT:
		return true
	}
	return false
}
