package helper

import (
	"strings"
	"unicode"

	"github.com/go-playground/validator/v10"
)
func NoWhitespaceOnly(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	return strings.TrimSpace(value) != ""
}
func Lowercase(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	for _, c := range value {
		if unicode.IsUpper(c) {
			return false
		}
	}
	return true
}