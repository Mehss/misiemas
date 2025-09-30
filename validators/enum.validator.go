package validators

import (
	"tripatra-dct-service-config/database/model/enum"
)

func ValidateModulePosition(mt enum.EnumPosition) bool {
	return mt.IsValid()
}
