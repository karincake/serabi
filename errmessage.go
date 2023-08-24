package serabi

var ErrMessage map[string]string = make(map[string]string)

func init() {
	ErrMessage["parsing-fail"] = "parsing failed for %v, error: %v"
	ErrMessage["required"] = "required"
	ErrMessage["equalTo"] = "must be the same with  %v"
	ErrMessage["gt"] = "must be greater than %v"
	ErrMessage["gte"] = "must greater than be equal to %v"
	ErrMessage["lt"] = "must be less than %v"
	ErrMessage["lte"] = "must be less than or equal to %v"
	ErrMessage["minLength"] = "the minimum length is %v characters"
	ErrMessage["maxLength"] = "the maximum length is %v characters"
	ErrMessage["alpha"] = "must be alphabet"
	ErrMessage["alphaSpace"] = "must be alphabet or space"
	ErrMessage["alphaUnder"] = "must be alphabet or spaceunderscores"
	ErrMessage["alphaNumeric"] = "must be alphabet or number"
	ErrMessage["alphaNumericUnder"] = "must be alphabet, number, or underscore"
	ErrMessage["numeric"] = "must be numeric"
	ErrMessage["numval"] = "must be number"
	ErrMessage["email"] = "must be a valid email addres format"
	ErrMessage["base64"] = "must be a base64 format"
	ErrMessage["fileSizeBase64"] = "file size must be less than %d KB"
}
