package validator

import (
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type InputValidator struct {
	v     *validator.Validate
	trans ut.Translator
}

type InvalidData struct {
	Field        string `json:"field"`
	ErrorMessage string `json:"error"`
}

func (val *InputValidator) initTranslator() error {
	translator := en.New()
	uni := ut.New(translator, translator)

	val.trans, _ = uni.GetTranslator("en")

	en_translations.RegisterDefaultTranslations(val.v, val.trans)

	return nil
}

func (val *InputValidator) registerMessage() {
	val.v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	val.v.RegisterTranslation(
		"required",
		val.trans,
		func(ut ut.Translator) error {
			return ut.Add("required", "{0} is a required field", true)
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("required", fe.Field())
			return t
		},
	)

	val.v.RegisterTranslation(
		"email",
		val.trans,
		func(ut ut.Translator) error {
			return ut.Add("email", "{0} must be a valid email", true)
			// return ut.Add("numeric", "{0} dqmkdqwmdqo", true)
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("email", fe.Field())
			return t
		},
	)
}

func (val *InputValidator) Validate(s interface{}) ([]InvalidData, bool) {
	err := val.v.Struct(s)

	if err != nil {
		invalids := make([]InvalidData, 0)
		for _, e := range err.(validator.ValidationErrors) {
			invalids = append(invalids, InvalidData{
				Field:        e.Field(),
				ErrorMessage: e.Translate(val.trans),
			})
		}

		return invalids, false
	}

	return nil, true
}

var Validator *InputValidator

func InitValidator() {
	coreValidator := newCoreValidator()
	Validator = &InputValidator{v: coreValidator}

	Validator.initTranslator()
	Validator.registerMessage()
}

func newCoreValidator() *validator.Validate {
	v := validator.New()

	return v
}
