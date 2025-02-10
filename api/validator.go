package api

import (
	"github.com/JulianAVG64/Simple-Bank/util"
	"github.com/go-playground/validator/v10"
)

var validCurrecy validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		// check currency is supported
		return util.IsSupportedCurrency(currency)
	}
	return false
}
