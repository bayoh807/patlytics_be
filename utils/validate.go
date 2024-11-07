package utils

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type CustomTransfer interface {
	GetKeys() map[string]string
	GetMessages(tag string) *string
}

type ValidateErr struct {
	Key     string
	Value   interface{}
	Tag     string
	Message string
}

type request struct {
	Validate *validator.Validate
}

var (
	Request request
	uni     *ut.UniversalTranslator
)

func init() {
	Request.Validate = validator.New()

	en := en.New()

	uni = ut.New(en)

}

func (r *request) ToValidate(req interface{}) (*validator.Validate, *ValidateErr) {

	if err := Request.Validate.Struct(req); err != nil {
		var errStruct *ValidateErr
		for _, e := range err.(validator.ValidationErrors) {
			if errStruct != nil {
				break
			}
			errStruct = &ValidateErr{
				Key:     e.Field(),
				Value:   e.Value(),
				Tag:     e.Tag(),
				Message: r.getMessage(e),
			}
			return nil, &ValidateErr{
				Key:     e.Field(),
				Value:   e.Value(),
				Tag:     e.Tag(),
				Message: r.getMessage(e),
			}
		}
	}

	return Request.Validate, nil
}
func (r *request) getMessage(err validator.FieldError) string {

	return err.Error()
}
