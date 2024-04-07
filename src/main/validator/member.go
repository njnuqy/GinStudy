package member

import (
	"github.com/go-playground/validator/v10"
)

func NameValid(fl validator.FieldLevel) bool {
	// 获取字段值并转换为string
	value := fl.Field().Interface()
	if strValue, ok := value.(string); ok {
		// 如果字段值是"admin"，则返回false
		if strValue == "admin" {
			return false
		}
	}
	return true
}
