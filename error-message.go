package serabi

import "errors"

// To make error messages easier to manage
var ErrMessage map[string]string = make(map[string]string)
var Errors map[string]error = make(map[string]error)

func init() {
	Errors["parsing-fail"] = errors.New("parsing failed for %v, error: %v")
	Errors["required"] = errors.New("value is required")
	Errors["eq"] = errors.New("must be equal to value")
	Errors["dif"] = errors.New("must be different from value")
	Errors["gt"] = errors.New("must be greater than the limit")
	Errors["gte"] = errors.New("must be greater than or equal to the limit")
	Errors["lt"] = errors.New("must be less than the limit")
	Errors["lte"] = errors.New("must be less than or equal to the limit")
	Errors["length"] = errors.New("the characters length doesn't fit the rule")
	Errors["minLength"] = errors.New("the characters length below the limit")
	Errors["maxLength"] = errors.New("the characters length exceed the limit")
	Errors["alpha"] = errors.New("must be alphabet")
}
