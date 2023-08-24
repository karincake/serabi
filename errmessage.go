package serabi

// To make error messages easier to manage
var ErrMessage map[string]string = make(map[string]string)

func init() {
	ErrMessage["parsing-fail"] = "parsing failed for %v, error: %v"
	ErrMessage["required"] = "required"
	ErrMessage["eq"] = "must be the same with  %v"
	ErrMessage["gt"] = "must be greater than %v"
	ErrMessage["gte"] = "must be greater than be equal to %v"
	ErrMessage["lt"] = "must be less than %v"
	ErrMessage["lte"] = "must be less than or equal to %v"
	ErrMessage["minLength"] = "the minimum length is %v characters"
	ErrMessage["maxLength"] = "the maximum length is %v characters"
	ErrMessage["alpha"] = "must be alphabet"
}
