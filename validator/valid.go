package validator

import (
	"regexp"

	v "github.com/go-playground/validator/v10"
)

var Validate *v.Validate
func Init() {
	Validate = v.New()

	Validate.RegisterValidation("bustype", func(fl v.FieldLevel) bool {
		btype := fl.Field().String()
		accept := []string{"IT", "Food", "Transport", "Building"}
		for _, a := range accept {
            if btype == a {
                return true
            }
        }
        return false
    })

	Validate.RegisterValidation("username", func(fl v.FieldLevel) bool {
	    re := regexp.MustCompile(`^[_A-Za-z0-9\\-]+$`)
        return re.MatchString(fl.Field().String())
    })

	// Validate.RegisterValidation("website", func(fl v.FieldLevel) bool {
    //     re := regexp.MustCompile(`^[a-z0-9\-]+$`)
    //     return re.MatchString(fl.Field().String())
	// })
}

