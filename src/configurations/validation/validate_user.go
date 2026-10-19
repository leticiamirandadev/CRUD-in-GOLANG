package validation

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/configurations/rest_err"
)

var (
	validate = validator.New()
	transl   ut.Translator
)
func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		uni := ut.New(en, en)
		transl, _ = uni.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}
func ValidateUserError(validation_err error) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidation *validator.InvalidValidationError
	var validationErrors validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return rest_err.NewBadRequestError("Erro ao converter os campos.")
	}

	if errors.As(validation_err, &jsonValidation) {
		return rest_err.NewBadRequestError("Erro na validação dos campos.")
	}

	if errors.As(validation_err, &validationErrors) {
		var causes []rest_err.Causes
              for _, e := range validationErrors {
			cause := rest_err.Causes{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}
			causes = append(causes, cause)
		}
              return rest_err.NewBadRequestValidationError("Alguns campos estão incorretos.", causes)
	}

	return rest_err.NewBadRequestError("Erro ao tentar converter os campos.")
}
