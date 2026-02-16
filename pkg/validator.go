package pkg

import (
	"errors"
	"reflect"

	"github.com/NARUBROWN/spine/pkg/header"
)

func ValidateHeaders[T interface{}](headers *header.Values, st T) (*T, error) {
	s := reflect.ValueOf(st)

	for i := 0; i < s.NumField(); i++ {
		field := s.Type().Field(i)
		var val reflect.Value

		if field.Tag.Get("validate") == "" || field.Tag.Get("validate") == "optional" {
			continue
		}

		if (*headers).Has(field.Name) {
			val = reflect.ValueOf((*headers).Get(field.Name))
		} else if (*headers).Has(field.Tag.Get("header")) {
			val = reflect.ValueOf((*headers).Get(field.Tag.Get("header")))
		} else {
			return nil, errors.New("Validation failed: missing header " + field.Name)
		}

		if !s.Type().ConvertibleTo(val.Type()) {
			s.Field(i).Set(val.Convert(s.Type()))
		}
	}

	return s.Interface().(*T), nil
}
