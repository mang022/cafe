package action

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

const phoneRegex = `^01([0|1|6|7|8|9])(-)([0-9]{3,4})(-)([0-9]{4})$`

func ValidateRegex(regex, value string) bool {
	reg := regexp.MustCompile(regex)
	return reg.Match([]byte(value))
}

var ValidatePhone validator.Func = func(fl validator.FieldLevel) bool {
	if value, ok := fl.Field().Interface().(string); ok {
		return ValidateRegex(phoneRegex, value)
	}
	return false
}

type SignUpOnwerDto struct {
	Phone    string `form:"phone" json:"phone" binding:"required,phone"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=128"`
}
