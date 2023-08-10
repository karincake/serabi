package serabi

var ErrMessage map[string]string = make(map[string]string)

func init() {
	ErrMessage["parsing-fail"] = "parsing failed for %v, error: %v"
	ErrMessage["required"] = "required"
	ErrMessage["equalTo"] = "value must be the same with  %v"
	ErrMessage["numeric"] = "value must be numeric"
	ErrMessage["min"] = "the minimum value is %v"
	ErrMessage["max"] = "the maximum value is %v"
	ErrMessage["minLength"] = "the minimum length is %v characters"
	ErrMessage["maxLength"] = "the maximum length is %v characters"
	ErrMessage["alpha"] = "value must be alphabet"
	ErrMessage["alphaSpace"] = "value must be alphabet or space"
	ErrMessage["alphaNumeric"] = "value must be alphabet or number"
	ErrMessage["alphaNumericU"] = "value must be alphabet, number, or underscore"
	ErrMessage["email"] = "value must be an valid email addres format"
	ErrMessage["base64"] = "value must be in base64 format"
	ErrMessage["fileSizeBase64"] = "file size must be less than %d KB"
	ErrMessage["uuid"] = "value must be a valid uuid"
}
