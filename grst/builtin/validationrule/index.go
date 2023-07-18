package validationrule

import (
	"gopkg.in/validator.v2"
)

func Initialize() {
	validator.SetPrintJSON(true)
	validator.SetValidationFunc("required", ruleRequired)
	validator.SetValidationFunc("email", ruleEmail)
	validator.SetValidationFunc("url", ruleUrl)
	validator.SetValidationFunc("enum", ruleEnum)
}
