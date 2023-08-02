// S for serabi or struct validator
package serabi

import (
	"fmt"
	"reflect"
	"strings"

	te "github.com/karincake/tempe/error"
)

// just key val for the tag
type keyVal struct {
	Key string
	Val string
}

// viladator func interface?
// param: reflect value, string
type validator func(reflect.Value, ...string) error

// type syncvalidator func(reflect.Value, string, reflect.Value, string) error

// list of validator for the given key from tag
var tagValidator map[string]validator

// tag name to validate
const tagName = "validate"

// Validation of each field based on the registered field checkers
func Validate(input interface{}, nameSpaces ...string) te.Errors {
	// identiy value and loop if its pointer until reaches non pointer
	inputV := reflect.ValueOf(input)

	// loop until we get what kind lays behind the input
	for inputV.Kind() == reflect.Pointer || inputV.Kind() == reflect.Interface {
		inputV = inputV.Elem()
	}

	// non struct cant be validated
	if inputV.Kind() != reflect.Struct {
		return nil
	}

	// namespace will be available if it is sub validation
	nameSpace := ""
	if len(nameSpaces) > 0 {
		if len(nameSpaces) > 1 && nameSpaces[1] != "" {
			nameSpace += "(embedded:" + nameSpaces[0] + ")."
		} else {
			nameSpace += nameSpaces[0] + "."
		}
	}

	// check each field
	// inputT := reflect.TypeOf(inputV.Interface()) // keep this for now
	inputT := inputV.Type()
	errList := te.NewErrors()
	for i := 0; i < inputV.NumField(); i++ {
		// identify type and value of the field
		fieldT := inputT.Field(i)
		fieldV := inputV.Field(i)
		for fieldV.Kind() == reflect.Ptr {
			if fieldV.IsZero() {
				break
			}
			fieldV = fieldV.Elem()
		}

		// if current field is struct, validate again
		// TODO: find information about this -> || fieldT.Type.Name() == ""
		typeString := fieldT.Type.String()
		if (fieldT.Type.Kind() == reflect.Struct) && typeString != "time.Time" {
			embeddedMode := ""
			if fieldT.Anonymous {
				embeddedMode = "(embedded)"
			}
			tag := fieldT.Tag.Get("json")
			if tag == "" {
				tag = fieldT.Name
			}
			errList.Import(Validate(fieldV.Interface(), tag, embeddedMode).Get())
			// maps.Copy(errList, Validate(fieldV.Interface(), tag, embeddedMode))
			continue
		}

		tag := fieldT.Tag.Get(tagName)
		if tag != "" {
			parsedTag := parseTag(tag)
			key := fieldT.Tag.Get("json")
			if key != "" {
				keys := strings.Split(key, ",")
				if keys[0] != "" {
					key = keys[0]
				} else {
					key = fieldT.Name
				}
			} else {
				key = fieldT.Name
			}
			// based on slice or not
			if fieldV.Kind() == reflect.Slice {
				// special case untuk required
				required := false
				for _, v := range parsedTag {
					if v.Key == "required" {
						required = true
						break
					}
				}
				// empty array
				if fieldV.Len() == 0 {
					if required {
						// errList[nameSpace+key] = te.Error{Message: ErrMessage["required"], Code: "required", ExptdValue: "", GivenValue: fieldV.Interface()}
						errList[nameSpace+key] = te.NewError("required", ErrMessage["required"], "", fieldV.Interface().(string))
					}
					continue
				}
				// loop
				if fieldV.Index(0).Kind() == reflect.Struct {
					for ix := 0; ix < fieldV.Len(); ix++ {
						errList.Import(Validate(fieldV.Index(ix).Interface(), fmt.Sprintf("%v[%v]", key, ix)).Get())
						// maps.Copy(errList, Validate(fieldV.Index(ix).Interface(), fmt.Sprintf("%v[%v]", key, ix)))
					}
				} else {
					for ix := 0; ix < fieldV.Len(); ix++ {
						CheckParsedTag(parsedTag, fieldV.Index(ix), errList, fmt.Sprintf("%v[%v]", key, ix))
					}
				}
			} else {
				// non slice
				CheckParsedTag(parsedTag, fieldV, errList, nameSpace+key)
			}
		}
	}

	if len(errList) > 0 {
		return errList
	}
	return nil
}

func CheckParsedTag(parsedTag []keyVal, fv reflect.Value, el te.Errors, key string) {
	for _, kv := range parsedTag {
		if _, ok := tagValidator[kv.Key]; ok {
			err := tagValidator[kv.Key](fv, kv.Val)
			if err != nil {
				el.AddComplete(key, kv.Key, err.Error(), kv.Val, fv.Interface())
				// el[key] = te.Error{Message: err.Error(), Code: kv.Key, ExptdValue: kv.Val, GivenValue: fv.Interface()}
				break // 1 err is enough, break from error check of the current field
			}
		}
	}
}

// register a validator
func RegisterFieldChecker(tag string, validatorF validator) {
	tagValidator[tag] = validatorF
}

// unregister a validator
func UnregisterFieldChecker(tag string) {
	delete(tagValidator, tag)
}
