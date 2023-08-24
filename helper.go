package serabi

import (
	"reflect"
	"strings"

	te "github.com/karincake/tempe/error"
)

// just key val for the tag
type keyVal struct {
	Key string
	Val string
}

// parse tag for key - val
func parseTag(tag string) []keyVal {
	kvList := []keyVal{}
	for _, item := range strings.Split(tag, ";") {
		pair := strings.SplitN(strings.TrimSpace(item), "=", 2)
		if len(pair) == 0 {
			continue
		}
		if len(pair) == 1 {
			kvList = append(kvList, keyVal{pair[0], ""})
		}
		if len(pair) == 2 {
			kvList = append(kvList, keyVal{pair[0], pair[1]})
		}
	}
	return kvList
}

// parse tag using fvFunc
func checkParsedTag(parent *reflect.Value, parsedTag []keyVal, fv reflect.Value, el te.Errors, key string) {
	for _, kv := range parsedTag {
		if _, ok := tagFVs[kv.Key]; ok {
			localFvType := tagFVs[kv.Key].fvType
			if localFvType == FVTBasic {
				err := tagFVs[kv.Key].fvFunc(fv, kv.Val)
				if err != nil {
					el.AddComplete(key, kv.Key, err.Error(), kv.Val, fv.Interface())
					break
				}
			} else if localFvType == FVTFieldComparison {
				err := tagFVs[kv.Key].fvFunc(fv, parent.FieldByName(kv.Val).String())
				if err != nil {
					el.AddComplete(key, kv.Key, err.Error(), kv.Val, fv.Interface())
					break
				}
			} else if localFvType == FVTRegex {
				err := tagFVs["regex"].fvFunc(fv, parent.FieldByName(kv.Val).String())
				if err != nil {
					el.AddComplete(key, kv.Key, err.Error(), kv.Val, fv.Interface())
					break
				}
			}
		}
	}
}
