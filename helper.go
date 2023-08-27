package serabi

import (
	"reflect"
	"strings"

	h "github.com/karincake/serabi/helper"
	te "github.com/karincake/tempe/error"
)

// just key val for the tag
type keyVal struct {
	Key []byte
	Val []byte
}

// parse tag for key - val
// turns out processing manually using slice of byte is
// faster than using split function
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
func checkParsedTag(parent *reflect.Value, parsedTag []keyVal, fv reflect.Value, el te.XErrors, key string) {
	for _, kv := range parsedTag {
		kvKey := string(kv.Key)
		kvVal := string(kv.Val)
		if _, ok := tagFVs[kvKey]; ok {
			localFvType := tagFVs[kvKey].FvType
			if localFvType == FVTBasic {
				err := tagFVs[kvKey].FvFunc(fv, kvVal)
				if err != nil {
					el[key] = te.XError{Source: key, Code: kvKey, Message: err.Error(), ExpectedVal: kvVal, GivenVal: fv.Interface()}
					break
				}
			} else if localFvType == FVTField {
				err := tagFVs[kvKey].FvFunc(fv, h.ValStringer(parent.FieldByName(kvVal)))
				if err != nil {
					el[key] = te.XError{Source: key, Code: kvKey, Message: err.Error(), ExpectedVal: kvVal, GivenVal: fv.Interface()}
					break
				}
			} else if localFvType == FVTRegex {
				err := tagFVs["regex"].FvFunc(fv, kvKey)
				if err != nil {
					el[key] = te.XError{Source: key, Code: kvKey, Message: err.Error(), ExpectedVal: kvVal, GivenVal: fv.Interface()}
					break
				}
			}
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
		return keyVal{Key: kv[:eqIdx], Val: kv[eqIdx+1:]}
	} else {
		return keyVal{Key: kv}
	}
}

// get json tag
func getJsonTag(t *reflect.StructField) string {
	tag := t.Tag.Get("json")
	tags := strings.Split(tag, ",")
	if tags[0] == "" {
		return t.Name
	}
	return tags[0]
}
