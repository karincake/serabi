// Go struct validator
package serabi

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"regexp"

	s "github.com/karincake/semprit"
	te "github.com/karincake/tempe/error"

	h "github.com/karincake/serabi/helper"
)

// viladator func interface?
// param: reflect value, string
type FvType string
type FvFunc func(reflect.Value, string) error
type fv struct {
	FvType
	FvFunc
}

const (
	FVTBasic FvType = "func"
	FVTRegex FvType = "regex"
	FVTField FvType = "fieldCompare"
)

// tag name to validate
const tagName = "validate"

// list of validator for the given key from tag
var tagFVs map[string]fv = map[string]fv{}

// special case, regex and field comparison
// var tagRegexes map[string]string
var regexes map[string]*regexp.Regexp = map[string]*regexp.Regexp{}
var fields map[string]string = map[string]string{}

var required = "required"

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

	// get the type
	inputT := inputV.Type()
	errList := te.XErrors{}

	// check cache
	className := inputT.Name()
	if CacheEnabled && !cache.classExists(className) {
		rc := registeredClass{
			inputVNFC: inputV.NumField(),
			fieldT:    []reflect.StructField{},
			parsedTag: [][]keyVal{},
		}
		for i := 0; i < rc.inputVNFC; i++ {
			// 	// identify type and value of the field
			fieldV := inputV.Field(i)
			rc.fieldT = append(rc.fieldT, inputT.Field(i))
			rc.tag = append(rc.tag, rc.fieldT[i].Tag.Get(tagName))
			rc.key = append(rc.key, keyOrJsonTag(rc.fieldT[i].Name, rc.fieldT[i].Tag.Get("json")))

			for fieldV.Kind() == reflect.Ptr {
				if fieldV.IsZero() {
					break
				}
				fieldV = fieldV.Elem()
			}

			// if current field is struct, validate again
			rc.typeString = append(rc.typeString, rc.fieldT[i].Type.String())
			if (rc.fieldT[i].Type.Kind() == reflect.Struct) && rc.typeString[i] != "time.Time" {
				embeddedMode := ""
				if rc.fieldT[i].Anonymous {
					embeddedMode = "(embedded)"
				}
				errList.Import(Validate(fieldV.Interface(), rc.key[i], embeddedMode).(te.XErrors))
				continue
			}

			if rc.tag[i] != "" {
				rc.parsedTag = append(rc.parsedTag, parseTag(rc.tag[i]))
				// based on slice or not
				if fieldV.Kind() == reflect.Slice {
					// special case untuk required
					required := false
					for _, v := range rc.parsedTag[i] {
						if string(v.Key) == "required" {
							required = true
							break
						}
					}
					// empty array
					if fieldV.Len() == 0 {
						if required {
							errList[nameSpace+rc.key[i]] = te.XError{Code: "required", Message: ErrMessage["required"], GivenVal: fieldV.Interface().(string)}
						}
						continue
					}
					// loop
					if fieldV.Index(0).Kind() == reflect.Struct {
						for ix := 0; ix < fieldV.Len(); ix++ {
							errList.Import(Validate(fieldV.Index(ix).Interface(), fmt.Sprintf("%v[%v]", rc.key[i], ix)).(te.XErrors))
						}
					} else {
						for ix := 0; ix < fieldV.Len(); ix++ {
							// fv :=
							checkParsedTag(&inputV, rc.parsedTag[i], fieldV.Index(ix), errList, fmt.Sprintf("%v[%v]", rc.key[i], ix))
						}
					}
				} else {
					// non slice
					checkParsedTag(&inputV, rc.parsedTag[i], fieldV, errList, nameSpace+rc.key[i])
				}
			} else {
				rc.parsedTag = append(rc.parsedTag, nil)
			}
		}
		cache.push(className, rc)
	} else if CacheEnabled {
		rc := cache.get(className)
		for i := 0; i < rc.inputVNFC; i++ {
			// 	// identify type and value of the field
			fieldV := inputV.Field(i)
			for fieldV.Kind() == reflect.Ptr {
				if fieldV.IsZero() {
					break
				}
				fieldV = fieldV.Elem()
			}

			// if current field is struct, validate again
			if (rc.fieldT[i].Type.Kind() == reflect.Struct) && rc.typeString[i] != "time.Time" {
				embeddedMode := ""
				if rc.fieldT[i].Anonymous {
					embeddedMode = "(embedded)"
				}
				errList.Import(Validate(fieldV.Interface(), rc.key[i], embeddedMode).(te.XErrors))
				continue
			}

			if rc.tag[i] != "" {
				// based on slice or not
				if fieldV.Kind() == reflect.Slice {
					// special case untuk required
					required := false
					for _, v := range rc.parsedTag[i] {
						if string(v.Key) == "required" {
							required = true
							break
						}
					}
					// empty array
					if fieldV.Len() == 0 {
						if required {
							errList[nameSpace+rc.key[i]] = te.XError{Code: "required", Message: ErrMessage["required"], GivenVal: h.ValStringer(fieldV)}
						}
						continue
					}
					// loop
					if fieldV.Index(0).Kind() == reflect.Struct {
						for ix := 0; ix < fieldV.Len(); ix++ {
							errList.Import(Validate(fieldV.Index(ix).Interface(), fmt.Sprintf("%v[%v]", rc.key, ix)).(te.XErrors))
						}
					} else {
						for ix := 0; ix < fieldV.Len(); ix++ {
							// fv := fieldV.Index(ix)
							checkParsedTag(&inputV, rc.parsedTag[i], fieldV.Index(ix), errList, fmt.Sprintf("%v[%v]", rc.key, ix))
						}
					}
				} else {
					// non slice
					checkParsedTag(&inputV, rc.parsedTag[i], fieldV, errList, nameSpace+rc.key[i])
				}
			}
		}
	} else {
		// check each field
		for i := 0; i < inputV.NumField(); i++ {
			// identify type and value of the field
			fieldV := inputV.Field(i)
			fieldT := inputT.Field(i)
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
				errList.Import(Validate(fieldV.Interface(), keyOrJsonTag(fieldT.Name, fieldT.Tag.Get("json")), embeddedMode).(te.XErrors))
				continue
			}

			tag := fieldT.Tag.Get(tagName)
			if tag != "" {
				key := keyOrJsonTag(fieldT.Name, fieldT.Tag.Get("json"))
				// key := keyOrJsonTag(fieldT.Name, fieldT.Tag.Get("json"))
				parsedTag := parseTag(tag)
				// based on slice or not
				if fieldV.Kind() == reflect.Slice {
					// special case untuk required
					for _, v := range parsedTag {
						if string(v.Key) == "required" && fieldV.Len() == 0 {
							errList[nameSpace+key] = te.XError{Code: "required", Message: ErrMessage["required"], GivenVal: fieldV.Interface().(string)}
							break
						}
					}
					// empty array
					if fieldV.Len() == 0 {
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
	}

	if len(errList) > 0 {
		return errList
	}
	return nil
}

// Validation for form-data
func ValidateFormData(container any, input *http.Request) error {
	err := s.HttpFormData(container, input)
	if err != nil {
		return err.(te.XErrors)
	}

	return Validate(container)
}

// Validation for url
// caveat: url's structure makes it impossible to do deep parsing
func ValidateURL(container any, input url.URL) error {
	err := s.UrlQueryParam(container, input)
	if err != nil {
		return err.(te.XErrors)
	}

	return Validate(container)
}

// Validation for IO Reader to help validate, for example, payload of http request
func ValidateIoReader(container any, input io.Reader) error {
	err := s.IOReaderJson(container, input)
	if err != nil {
		return err.(te.XError)
	}

	// same process with normal validation
	return Validate(container)
}

// Add tag validator
// Requires tag name and validation function for the parameters
func AddTag(tag string, f FvFunc) {
	tagFVs[tag] = fv{FVTBasic, f}
}

// Add tag validator for field comparison
// Field comparison validator is the same with basic valicator, except it uses
// tag value as target field to be compared. Therefore, it can utilize the
// existing function. Please note the difference is in its usage
// i.e: gtField=age, gtField is the tag, age is the target field
func AddTagForField(tag string, f FvFunc) {
	tagFVs[tag] = fv{FVTField, f}
}

// Add a tag validator for regex
// Regex validator requires tag, regex, and message for the parameters
// Note: the message is stated here since it utilizes single function for all
// of the validation.
func AddTagForRegex(tag string, r string, msg string) {
	tagFVs[tag] = fv{FVTRegex, regexTagValidator}
	regexes[tag] = regexp.MustCompile(r)
	Errors[tag] = errors.New(msg)
}

// Remove a tag validator
func RemoveTag(tag string) {
	// forbidden tag to remove
	if tag == "regex" {
		return
	}
	delete(tagFVs, tag)
}
