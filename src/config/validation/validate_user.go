package validation

import (
	"encoding/json"
	"errors"

	err_rest "github.com/ale-neto/golang/src/config/err"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		un := ut.New(en, en)
		transl, _ = un.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(validation_error error) *err_rest.Err {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_error, &jsonErr) {
		return err_rest.NewBadRequestErr("Invalid field type")
	} else if errors.As(validation_error, &jsonValidationError) {
		errorsCauses := []err_rest.Causes{}
		for _, fieldErr := range validation_error.(validator.ValidationErrors) {
			cause := err_rest.Causes{
				Field:   fieldErr.Field(),
				Message: fieldErr.Translate(transl),
			}

			errorsCauses = append(errorsCauses, cause)
		}
		return err_rest.NewBadRequestValidationErr("Validation error", errorsCauses)
	} else {
		return err_rest.NewBadRequestErr("Error trying to convert the field type")
	}

}
