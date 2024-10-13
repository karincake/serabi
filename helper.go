package serabi

import (
	"fmt"
	"reflect"
	"strconv"

	d "github.com/karincake/dodol"
	h "github.com/karincake/serabi/helper"
)

// just key val for the tag
type keyVal struct {
	Key string
	Val string
}

// parse tag for key - val
// turns out processing manually using slice of byte is faster than using split
// function possibly due to the fixed part that has to be searched
func parseTag(tag string) []keyVal {
	kvList := []keyVal{}
	tagByte := []byte(tag)
	lastI := 0
	// split by character ";"
	for i, v := range tagByte {
		if v == 59 { // 59 is for character ";"
			kvList = append(kvList, identifyTagRule(tagByte[lastI:i]))
			lastI = i + 1
		}
	}
	kvList = append(kvList, identifyTagRule(tagByte[lastI:]))
	return kvList
}

// parse tag using FvFunc
func checkParsedTag(parent *reflect.Value, parsedTag []keyVal, fv reflect.Value, el d.FieldErrors, key, fieldName, eNameSpace string) {
	for _, kv := range parsedTag {
		if _, ok := tagFVs[kv.Key]; ok {
			localFvType := tagFVs[kv.Key].FvType
			if localFvType == FVTBasic {
				err := tagFVs[kv.Key].FvFunc(fv, kv.Val)
				if err != nil {
					expVal := ""
					if kv.Val != "" {
						expVal = kv.Key + "(" + kv.Val + ")"
					}
					el[key] = d.FieldError{
						Source:      fieldName,
						Code:        kv.Key,
						Message:     err.Error() + " " + expVal,
						ExpectedVal: expVal,
						GivenVal:    fv.Interface(),
						EmbedSource: eNameSpace,
					}
					break
				}
			} else if localFvType == FVTField {
				err := tagFVs[kv.Key].FvFunc(fv, h.ValStringer(parent.FieldByName(kv.Val)))
				if err != nil {
					expVal := kv.Val
					if kv.Val != "" {
						expVal = kv.Key + "(" + kv.Val + ")"
					}
					el[key] = d.FieldError{
						Source:      fieldName,
						Code:        kv.Key,
						Message:     err.Error() + " " + expVal,
						ExpectedVal: expVal,
						GivenVal:    fv.Interface(),
						EmbedSource: eNameSpace,
					}
					break
				}
			} else if localFvType == FVTRegex {
				err := tagFVs["regex"].FvFunc(fv, kv.Key)
				if err != nil {
					el[key] = d.FieldError{
						Code:        kv.Key,
						Message:     err.Error(),
						ExpectedVal: kv.Key,
						GivenVal:    fv.Interface(),
						EmbedSource: eNameSpace,
					}
					break
				}
			}
		} else {
			panic(fmt.Sprintf("unregistered tag found '%v'", kv.Key))
		}
	}
}

// check slice field
func checkSliceField(pt []keyVal, fv reflect.Value, nameSpace, key string, el d.FieldErrors) { // parent *reflect.Value,
	// special case untuk required
	required := false
	for _, v := range pt {
		if string(v.Key) == "required" {
			required = true
			break
		}
	}

	// empty array
	fvLength := fv.Len()
	if fvLength == 0 {
		if required {
			fvV := reflect.ValueOf(fv.Interface())
			fvVKind := fvV.Kind()
			if fvVKind == reflect.Array || fvVKind == reflect.Slice {
				el[nameSpace+key] = d.FieldError{Code: "required", Message: ErrMessage["required"], GivenVal: "*invalid type*"}
			} else {
				el[nameSpace+key] = d.FieldError{Code: "required", Message: ErrMessage["required"], GivenVal: fv.Interface().(string)}
			}
		}
		return
	}

	// check first element
	fvZero := fv.Index(0)
	for fvZero.Kind() == reflect.Pointer || fvZero.Kind() == reflect.Interface {
		fvZero = fvZero.Elem()
	}
	if fvZero.Kind() != reflect.Struct {
		return
	}

	// loop
	if fvZero.Kind() == reflect.Struct {
		for ix := 0; ix < fvLength; ix++ {
			// cek each element
			fvIx := fv.Index(0)
			for fvIx.Kind() == reflect.Pointer || fvIx.Kind() == reflect.Interface {
				fvIx = fvIx.Elem()
			}
			if fvIx.Kind() != reflect.Struct {
				return
			}

			// validate
			err := Validate(fv.Index(ix).Interface(), key+"["+strconv.Itoa(ix)+"]")
			if err != nil {
				el.Import(err.(d.FieldErrors))
			}
		}
	} else {
		for ix := 0; ix < fv.Len(); ix++ {
			checkParsedTag(&fv, pt, fv.Index(ix), el, nameSpace+"["+strconv.Itoa(ix)+"]", key+"["+strconv.Itoa(ix)+"]", "")
		}
	}
}

// split and return
func identifyTagRule(kv []byte) keyVal {
	eqIdx := 0
	// split by =
	for i2, v2 := range kv {
		if v2 == 61 { // 61 is for character "="
			eqIdx = i2
			break
		}
	}
	if eqIdx > 0 {
		return keyVal{Key: string(kv[:eqIdx]), Val: string(kv[eqIdx+1:])}
	} else {
		return keyVal{Key: string(kv)}
	}
}

// get json tag
func keyOrJsonTag(key, jsonTag string) string {
	// jsonTag := t.Tag.Get("json")
	if jsonTag == "" {
		return key
	}
	tagByte := []byte(jsonTag)
	pos := len(tagByte)
	for i, v := range tagByte {
		if v == 44 { // 44 is for character ","
			pos = i
		}
	}
	return string(tagByte[:pos])
}
