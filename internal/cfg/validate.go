package cfg

import (
	"errors"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

var (
	trans ut.Translator
)

// tagPort validates port value
func tagPort(fl validator.FieldLevel) bool {
	if fl.Field().CanInt() {
		n := fl.Field().Int()
		if n > 0 && n < 65535 {
			return true
		}
	}
	return false
}

// tagOneOfNested validates that only one of nested fileds presents
func tagOneOfNested(fl validator.FieldLevel) bool {
	n := fl.Field().NumField()
	c := 0
	for i := 0; i < n; i++ {
		f := fl.Field().Field(i)
		if !f.IsZero() {
			c++
		}
	}
	return c == 1
}

// translations add custom translation for list of tags
func translations(v *validator.Validate) error {
	var ok bool
	e := en.New()
	uni := ut.New(e, e)
	trans, ok = uni.GetTranslator("en")
	if !ok {
		return errors.New("could not get translator")
	}
	// "required"
	if err := transRequired(v); err != nil {
		return err
	}
	// "ip"
	if err := transIp(v); err != nil {
		return err
	}
	// "port"
	if err := transPort(v); err != nil {
		return err
	}
	// "one_of_nested"
	if err := transOneOfNested(v); err != nil {
		return err
	}
	return nil
}

// transRequired translates error for "required" tag
func transRequired(v *validator.Validate) error {
	return v.RegisterTranslation(
		"required",
		trans,
		func(ut ut.Translator) error {
			return ut.Add("required", "{0} must have a value", true)
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, err := ut.T("required", fe.StructNamespace())
			if err != nil {
				return fe.(error).Error()
			}
			return t
		},
	)
}

// transIp translates error for "ip" tag
func transIp(v *validator.Validate) error {
	return v.RegisterTranslation(
		"ip",
		trans,
		func(ut ut.Translator) error {
			return ut.Add("ip", "{0} must be a valid IP address", true)
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, err := ut.T("ip", fe.StructNamespace())
			if err != nil {
				return fe.(error).Error()
			}
			return t
		},
	)
}

// transPort translates error for "port" tag
func transPort(v *validator.Validate) error {
	return v.RegisterTranslation(
		"port",
		trans,
		func(ut ut.Translator) error {
			return ut.Add("port", "{0} must be a valid port number (1-65535)", true)
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, err := ut.T("port", fe.StructNamespace())
			if err != nil {
				return fe.(error).Error()
			}
			return t
		},
	)
}

// transOneOfNested translates error for "one_of_nested" tag
func transOneOfNested(v *validator.Validate) error {
	return v.RegisterTranslation(
		"one_of_nested",
		trans,
		func(ut ut.Translator) error {
			return ut.Add("one_of_nested", "Only one value must be present in {0}", true)
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, err := ut.T("one_of_nested", fe.StructNamespace())
			if err != nil {
				return fe.(error).Error()
			}
			return t
		},
	)
}
