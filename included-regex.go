package serabi

// included default tag validation for regex

import (
	"reflect"

	h "github.com/karincake/serabi/helper"
)

// register regexes
func init() {
	AddTag("regex", regexTagValidator)
	AddTagForRegex("alpha", "^[a-zA-Z]+$", "must be alphabet")
	AddTagForRegex("alphaSpace", "^[a-zA-Z ]+$", "must be alphabet, and allowed to place space within it")
	AddTagForRegex("alphaNumeric", "^[a-zA-Z0-9]+$", "must be alphabet or number")
	AddTagForRegex("alphaUnder", "^[a-zA-Z_]+$", "must be alphabet or underscore")
	AddTagForRegex("alphaNumericSpace", "^[a-zA-Z0-9 ]+$", "must be alphabet, number, or underscore")
	AddTagForRegex("alphaNumericUnder", "^[a-zA-Z0-9_]+$", "must be alphabet, number, or underscore")
	AddTagForRegex("numeric", "^[0-9]+$", "must be numeric")
	AddTagForRegex("numval", "[-+]?[0-9]+(?:\\.[0-9]+)?$", "must be value of numeric")
	AddTagForRegex("email", "^(?:(?:(?:(?:[a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(?:\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|(?:(?:\\x22)(?:(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(?:\\x20|\\x09)+)?(?:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(\\x20|\\x09)+)?(?:\\x22))))@(?:(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$", "must be a valid email addres format")
}

// /// Field checkers
func regexTagValidator(val reflect.Value, code string) error {
	if (val.Kind() == reflect.Pointer && val.IsNil()) || val.Interface() == "" {
		return nil
	}

	if !regexes[code].MatchString(h.ValStringer(val)) {
		return Errors[code]
	}

	return nil
}
