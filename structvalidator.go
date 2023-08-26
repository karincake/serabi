// Go struct validator
package serabi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	h "github.com/karincake/serabi/helper"
	te "github.com/karincake/tempe/error"
	"gorm.io/datatypes"
)

// viladator func interface?
// param: reflect value, string
type fvType string
type fvFunc func(reflect.Value, string) error
type fv struct {
	fvType
	fvFunc
}

const (
	FVTBasic fvType = "func"
	FVTRegex fvType = "regex"
	FVTField fvType = "fieldCompare"
)

// tag name to validate
const tagName = "validate"

// list of validator for the given key from tag
var tagFVs map[string]fv = map[string]fv{}

// special case, regex and field comparison
// var tagRegexes map[string]string
var regexes map[string]*regexp.Regexp = map[string]*regexp.Regexp{}
var fields map[string]string = map[string]string{}

var bRequired = []byte("required")

// Validation of each field based on the registered tag
func Validate(input any, nameSpaces ...string) error {
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
			nameSpace += "(" + nameSpaces[0] + ")."
		} else {
			nameSpace += nameSpaces[0] + "."
		}
	}

	// check each field
	// inputV.Type()
	inputT := inputV.Type()
	errList := te.XErrors{}
	inputVNFC := inputV.NumField()
	for i := 0; i < inputVNFC; i++ {
		// 	// identify type and value of the field
		fieldT := inputT.Field(i)
		fieldV := inputV.Field(i)
		for fieldV.Kind() == reflect.Ptr {
			if fieldV.IsZero() {
				break
			}
			fieldV = fieldV.Elem()
		}

		// if current field is struct, validate again
		typeString := fieldT.Type.String()
		if (fieldT.Type.Kind() == reflect.Struct) && typeString != "time.Time" {
			embeddedMode := ""
			if fieldT.Anonymous {
				embeddedMode = "(embedded)"
			}
			tag := fieldT.Tag.Get("json")
			tags := strings.Split(tag, ",")
			if tags[0] == "" {
				tag = fieldT.Name
			}
			tag = tags[0]

			errList.Import(Validate(fieldV.Interface(), tag, embeddedMode).(te.XErrors))
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
					if bytes.Equal(v.Key, bRequired) {
						required = true
						break
					}
				}
				// empty array
				if fieldV.Len() == 0 {
					if required {
						errList[nameSpace+key] = te.XError{Code: "required", Message: ErrMessage["required"], GivenVal: fieldV.Interface().(string)}
					}
					continue
				}
				// loop
				if fieldV.Index(0).Kind() == reflect.Struct {
					for ix := 0; ix < fieldV.Len(); ix++ {
						errList.Import(Validate(fieldV.Index(ix).Interface(), fmt.Sprintf("%v[%v]", key, ix)).(te.XErrors))
					}
				} else {
					for ix := 0; ix < fieldV.Len(); ix++ {
						checkParsedTag(&inputV, parsedTag, fieldV.Index(ix), errList, fmt.Sprintf("%v[%v]", key, ix))
					}
				}
			} else {
				// non slice
				checkParsedTag(&inputV, parsedTag, fieldV, errList, nameSpace+key)
			}
		}
	}

	if len(errList) > 0 {
		return errList
	}
	return nil
}

// Validation for IO Reader to help validate, for example, payload of http request
func ValidateIoReader(container interface{}, input io.Reader) error {
	decoder := json.NewDecoder(input)
	err := decoder.Decode(&container)
	if err != nil {
		cV := reflect.ValueOf(container)
		for cV.Kind() == reflect.Pointer || cV.Kind() == reflect.Interface {
			cV = cV.Elem()
		}
		structName := cV.Type().Name()
		return te.XErrors{
			"payload-bad": te.XError{
				Code:        "payload-bad",
				Message:     fmt.Sprintf(ErrMessage["parsing-fail"], structName, err),
				ExpectedVal: fmt.Sprintf("value of %v", structName),
			},
		}
	}

	// same process with normal validation
	return Validate(container)
}

// Validation for url
// caveat: url's structure makes it impossible to do deep parsing
func ValidateURL(container any, url url.URL) error {
	cV := reflect.ValueOf(container).Elem()
	for cV.Kind() == reflect.Pointer || cV.Kind() == reflect.Interface {
		cV = cV.Elem()
	}

	cT := cV.Type()
	values := url.Query()
	errList := te.XErrors{}
	for i := 0; i < cV.NumField(); i++ {
		fieldV := cV.Field(i)
		fieldT := cT.Field(i)

		if !fieldV.CanSet() {
			continue
		}

		key := fieldT.Tag.Get("json")
		if key == "" {
			key = fieldT.Name
		}

		vals, ok := values[key]
		if !ok {
			continue
		}

		switch fieldV.Interface().(type) {
		case bool, *bool:
			var v bool
			if strings.ToLower(vals[0]) == "true" || vals[0] == "1" {
				v = true
			} else {
				v = false
			}
			if fieldV.Kind() == reflect.Ptr {
				fieldV.Set(reflect.ValueOf(&v))
			} else {
				fieldV.Set(reflect.ValueOf(v))
			}
		case string, *string:
			if fieldV.Kind() == reflect.Ptr {
				fieldV.Set(reflect.ValueOf(&vals[0]))
			} else {
				fieldV.Set(reflect.ValueOf(vals[0]))
			}
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, *int, *int8, *int16, *int32, *int64, *uint, *uint8, *uint16, *uint32, *uint64:
			if valInt, err := strconv.Atoi(vals[0]); err != nil {
				errList[key] = te.XError{Code: key, Message: err.Error(), ExpectedVal: vals[0], GivenVal: fieldV.Interface()}
			} else {
				fieldV.Set(h.IntToVal(valInt, fieldV))
			}
		case float64, *float64:
			strFloat, err := strconv.ParseFloat(vals[0], 64)
			if err != nil {
				errList[key] = te.XError{Code: key, Message: err.Error(), ExpectedVal: vals[0], GivenVal: fieldV.Interface()}
			}
			if fieldV.Kind() == reflect.Ptr {
				fieldV.Set(reflect.ValueOf(&strFloat))
			} else {
				fieldV.Set(reflect.ValueOf(strFloat))
			}
		case float32, *float32:
			strFloat, err := strconv.ParseFloat(vals[0], 32)
			if err != nil {
				errList[key] = te.XError{Code: key, Message: err.Error(), ExpectedVal: vals[0], GivenVal: fieldV.Interface()}
			}
			strFloat32 := float32(strFloat)
			if fieldV.Kind() == reflect.Ptr {
				fieldV.Set(reflect.ValueOf(&strFloat32))
			} else {
				fieldV.Set(reflect.ValueOf(strFloat32))
			}
		case []string, *[]string:
			fieldV.Set(reflect.ValueOf(&vals))
		case datatypes.Date, *datatypes.Date:
			time, err := time.Parse("2006-01-02T15:04:05.000Z", vals[0])
			if err != nil {
				errList[key] = te.XError{Code: key, Message: err.Error(), ExpectedVal: vals[0], GivenVal: fieldV.Interface()}
			}
			date := datatypes.Date(time)
			if fieldV.Kind() == reflect.Ptr {
				fieldV.Set(reflect.ValueOf(&date))
			} else {
				fieldV.Set(reflect.ValueOf(date))
			}
		case time.Time, *time.Time:
			time, err := time.Parse("2006-01-02T15:04:05.000Z", vals[0])
			if err != nil {
				errList[key] = te.XError{Code: key, Message: err.Error(), ExpectedVal: vals[0], GivenVal: fieldV.Interface()}
			}
			if fieldV.Kind() == reflect.Ptr {
				fieldV.Set(reflect.ValueOf(&time))
			} else {
				fieldV.Set(reflect.ValueOf(time))
			}
		// TODO: make any *[]int as a function
		case *[]int8:
			failed := false
			valX := []int8{}
			for _, val := range vals {
				if valInt, err := strconv.Atoi(val); err != nil {
					failed = true
					errList[key] = te.XError{Code: key, Message: err.Error(), ExpectedVal: vals[0], GivenVal: fieldV.Interface()}
				} else {
					valX = append(valX, int8(valInt))
				}
			}
			if !failed {
				fieldV.Set(reflect.ValueOf(valX))
			}
			// case []int16:
			// 	failed := false
			// 	valX := []int16{}
			// 	for _, val := range vals {
			// 		if valInt, err := strconv.Atoi(val); err != nil {
			// 			failed = true
			// 			errList[key] = t.Error{err.Error(), key, val, fieldV.Interface()}
			// 		} else {
			// 			valX = append(valX, int16(valInt))
			// 		}
			// 	}
			// 	if !failed {
			// 		fieldV.Set(reflect.ValueOf(valX))
			// 	}
			// case []int32:
			// 	failed := false
			// 	valX := []int32{}
			// 	for _, val := range vals {
			// 		if valInt, err := strconv.Atoi(val); err != nil {
			// 			failed = true
			// 			errList[key] = t.Error{err.Error(), key, val, fieldV.Interface()}
			// 		} else {
			// 			valX = append(valX, int32(valInt))
			// 		}
			// 	}
			// 	if !failed {
			// 		fieldV.Set(reflect.ValueOf(valX))
			// 	}
			// case []int64:
			// 	failed := false
			// 	valX := []int64{}
			// 	for _, val := range vals {
			// 		if valInt, err := strconv.Atoi(val); err != nil {
			// 			failed = true
			// 			errList[key] = t.Error{err.Error(), key, val, fieldV.Interface()}
			// 		} else {
			// 			valX = append(valX, int64(valInt))
			// 		}
			// 	}
			// 	if !failed {
			// 		fieldV.Set(reflect.ValueOf(valX))
			// 	}
		}
	}

	if len(errList) > 0 {
		return errList
	}

	return Validate(container)
}

// Add tag validator
// Requires tag name and validation function for the parameters
func AddTag(tag string, f fvFunc) {
	tagFVs[tag] = fv{FVTBasic, f}
}

// Add tag validator for field comparison
// Field comparison validator is the same with basic valicator, except it uses
// tag value as target field to be compared. Therefore, it can utilize the
// existing function. Please note the difference is in its usage
// i.e: gtField=age, gtField is the tag, age is the target field
func AddTagForField(tag string, f fvFunc) {
	tagFVs[tag] = fv{FVTField, f}
}

// Add a tag validator for regex
// Regex validator requires tag, regex, and message for the parameters
// Note: the message is stated here since it utilizes single function for all
// of the validation.
func AddTagForRegex(tag string, r string, msg string) {
	tagFVs[tag] = fv{FVTRegex, regexTagValidator}
	regexes[tag] = regexp.MustCompile(r)
	ErrMessage[tag] = msg
}

// Remove a tag validator
func RemoveTag(tag string) {
	// forbidden tag to remove
	if tag == "regex" {
		return
	}
	delete(tagFVs, tag)
}
