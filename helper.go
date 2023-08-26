package serabi

import (
	"reflect"

	h "github.com/karincake/serabi/helper"
	te "github.com/karincake/tempe/error"
)

// just key val for the tag
type keyVal struct {
	Key []byte
	Val []byte
}

// parse tag for key - val
func parseTag(tag string) []keyVal {
	kvList := []keyVal{}
	tagByte := []byte(tag)
	lastI := 0
	for i, v := range tagByte {
		if v == 59 {
			kvByte := tagByte[lastI:i]
			lastI = i + 1
			eqIdx := 0
			for i2, v2 := range kvByte {
				if v2 == 61 {
					eqIdx = i2
					break
				}
			}
			if eqIdx > 0 {
				kvList = append(kvList, keyVal{Key: kvByte[:eqIdx], Val: kvByte[eqIdx+1:]})
			} else {
				kvList = append(kvList, keyVal{Key: kvByte})
			}
		}
	}
	kvByte := tagByte[lastI:]
	eqIdx := 0
	for i2, v2 := range kvByte {
		if v2 == 61 {
			eqIdx = i2
			break
		}
	}
	if eqIdx > 0 {
		kvList = append(kvList, keyVal{Key: kvByte[:eqIdx], Val: kvByte[eqIdx+1:]})
	} else {
		kvList = append(kvList, keyVal{Key: kvByte})
	}
	return kvList
}

// parse tag using FvFunc
func checkParsedTag(parent *reflect.Value, parsedTag []keyVal, fv reflect.Value, el te.XErrors, key string) {
	for _, kv := range parsedTag {
		kvKey := string(kv.Key)
		kvVal := string(kv.Val)
		if _, ok := tagFVs[kvKey]; ok {
			localFvType := tagFVs[kvKey].fvType
			if localFvType == FVTBasic {
				err := tagFVs[kvKey].FvFunc(fv, kvVal)
				if err != nil {
					el[key] = te.XError{Code: kvKey, Message: err.Error(), ExpectedVal: kvVal, GivenVal: fv.Interface()}
					break
				}
			} else if localFvType == FVTField {
				err := tagFVs[kvKey].FvFunc(fv, h.ValStringer(parent.FieldByName(kvVal)))
				if err != nil {
					el[key] = te.XError{Code: kvKey, Message: err.Error(), ExpectedVal: kvVal, GivenVal: fv.Interface()}
					break
				}
			} else if localFvType == FVTRegex {
				err := tagFVs["regex"].FvFunc(fv, kvKey)
				if err != nil {
					el[key] = te.XError{Code: kvKey, Message: err.Error(), ExpectedVal: kvVal, GivenVal: fv.Interface()}
					break
				}
			}
		}
	}
}
