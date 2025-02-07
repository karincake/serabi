package serabi

// included default tag validation for basic validation

import (
	"errors"
	"reflect"
	"strconv"
	"time"

	h "github.com/karincake/serabi/helper"
)

// register the field checkers
func init() {
	AddTag("required", requiredTagValidator)
	AddTag("gt", gtTagValidator)
	AddTag("gte", gteTagValidator)
	AddTag("lt", ltTagValidator)
	AddTag("lte", lteTagValidator)
	AddTag("length", lengthTagValidator)
	AddTag("minLength", minLengthTagValidator)
	AddTag("maxLength", maxLengthTagValidator)
}

// /// Field checkers
func requiredTagValidator(val reflect.Value, expectVal string) error {
	kind := val.Kind()
	if kind == reflect.Ptr && val.IsNil() {
		return Errors["required"]
	} else if kind == reflect.String && val.String() == "" {
		return Errors["required"]
	} else if kind >= reflect.Int && kind <= reflect.Int64 && expectVal != "allowZero" && val.Int() == 0 {
		return Errors["required"]
	} else if kind >= reflect.Uint && kind <= reflect.Uint64 && expectVal != "allowZero" && val.Uint() == 0 {
		return Errors["required"]
	} else if kind >= reflect.Float32 && kind <= reflect.Float64 && expectVal != "allowZero" && val.Float() == 0 {
		return Errors["required"]
	} else if kind == reflect.Struct {
		dtype := val.Type().String()
		if dtype == "time.Time" {
			original := val.Interface().(time.Time)
			if original.IsZero() {
				return Errors["required"]
			}
		}
	}
	return nil
}

func eqTagValidator(val reflect.Value, expectVal string) error {
	if val.Kind() == reflect.Pointer && val.IsNil() {
		return nil
	}

	if val.String() != expectVal {
		return Errors["eq"]
	}
	return nil
}

func difTagValidator(val reflect.Value, expectVal string) error {
	if val.Kind() == reflect.Pointer && val.IsNil() {
		return nil
	}

	if val.String() == expectVal {
		return Errors["dif"]
	}
	return nil
}

func gtTagValidator(val reflect.Value, expectVal string) error {
	if val.Kind() == reflect.Pointer && val.IsNil() {
		return nil
	}

	val1, val2, err := resourceToNumval(val, expectVal)
	if err != nil {
		return err
	} else if val1 <= val2 {
		return Errors["gt"]
	}
	return nil
}

func gteTagValidator(val reflect.Value, expectVal string) error {
	if val.Kind() == reflect.Pointer && val.IsNil() {
		return nil
	}

	val1, val2, err := resourceToNumval(val, expectVal)
	if err != nil {
		return err
	} else if val1 < val2 {
		return Errors["gte"]
	}
	return nil
}

func ltTagValidator(val reflect.Value, expectVal string) error {
	if val.Kind() == reflect.Pointer && val.IsNil() {
		return nil
	}

	val1, val2, err := resourceToNumval(val, expectVal)
	if err != nil {
		return err
	} else if val1 >= val2 {
		return Errors["lt"]
	}
	return nil
}

func lteTagValidator(val reflect.Value, expectVal string) error {
	if val.Kind() == reflect.Pointer && val.IsNil() {
		return nil
	}

	val1, val2, err := resourceToNumval(val, expectVal)
	if err != nil {
		return err
	} else if val1 > val2 {
		return Errors["lte"]
	}
	return nil
}

func lengthTagValidator(val reflect.Value, expectVal string) error {
	if val.Kind() == reflect.Pointer && val.IsNil() {
		return nil
	}
	opts0Int, err := strconv.Atoi(expectVal)
	if err != nil {
		panic(Errors["numeric"])
	}

	valC := h.ValStringer(val) // value converted
	if len(valC) != opts0Int {
		return Errors["length"]
	}
	return nil
}
func minLengthTagValidator(val reflect.Value, expectVal string) error {
	if val.Kind() == reflect.Pointer && val.IsNil() {
		return nil
	}
	opts0Int, err := strconv.Atoi(expectVal)
	if err != nil {
		panic(Errors["numeric"])
	}

	valC := h.ValStringer(val) // value converted
	length := len(valC)
	if length != 0 && length < opts0Int {
		return Errors["minLength"]
	}
	return nil
}

func maxLengthTagValidator(val reflect.Value, expectVal string) error {
	if val.Kind() == reflect.Pointer && val.IsNil() {
		return nil
	}
	opts0Int, err := strconv.Atoi(expectVal)
	if err != nil {
		panic(Errors["numeric"])
	}

	valC := h.ValStringer(val) // value converted
	length := len(valC)
	if length == 0 {
		return nil
	}
	if len(valC) > opts0Int {
		return Errors["maxLength"]
	}
	return nil
}

// //// some helper for the default field checker
func resourceToNumval(val reflect.Value, exptVal string) (float64, float64, error) { // , mode string
	exptValFloat, err := strconv.ParseFloat(exptVal, 64)
	if err != nil {
		return 0, 0, err
	}

	valK := val.Kind()
	if valK == reflect.String {
		valCT, err := strconv.ParseFloat(val.String(), 64)
		if err != nil {
			return 0, 0, errors.New(ErrMessage["numeric"])
		}
		return valCT, exptValFloat, nil
	} else if valK >= reflect.Int && valK <= reflect.Int64 {
		return float64(val.Int()), exptValFloat, nil
	} else if valK >= reflect.Uint && valK <= reflect.Uint64 {
		return float64(val.Uint()), exptValFloat, nil
	} else if valK <= reflect.Float32 && valK <= reflect.Float64 {
		return val.Float(), exptValFloat, nil
	}

	return 0, 0, errors.New("unconvertable value")
}
