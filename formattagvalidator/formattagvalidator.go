package idtagvalidator

import (
	"errors"
	"reflect"

	"github.com/google/uuid"
	h "github.com/karincake/serabi/helper"
)

func UUIDTagValidator(val reflect.Value, opts ...string) error {
	if (val.Kind() == reflect.Pointer && val.IsNil()) || val.String() == "" {
		return nil
	}

	_, err := uuid.Parse(h.ValStringer(val))
	if err != nil {
		return errors.New("value must be a valid uuid")
	}

	return nil
}
