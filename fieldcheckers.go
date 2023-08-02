// built-in field checker
package serabi

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

// register the field checkers
func init() {
	tagValidator = make(map[string]validator)
	RegisterFieldChecker("required", requiredTagValidator)
	RegisterFieldChecker("equalto", eqTagValidator)
	RegisterFieldChecker("min", minTagValidator)
	RegisterFieldChecker("max", maxTagValidator)
	RegisterFieldChecker("minLength", minLengthTagValidator)
	RegisterFieldChecker("maxLength", maxLengthTagValidator)
}

///// Field checkers

func requiredTagValidator(val reflect.Value, opts ...string) error {
	if (val.Kind() == reflect.String && val.String() == "") || (val.Kind() == reflect.Ptr && val.IsNil()) {
		return errors.New(ErrMessage["required"])
	}
	return nil
}

func eqTagValidator(val reflect.Value, opts ...string) error {
	return nil
}

func minTagValidator(val reflect.Value, opts ...string) error {
	if val.Kind() == reflect.Pointer && val.IsNil() {
		return nil
	}
	if err := valLimiter(val, opts[0], "<"); err != nil {
		return err
	}
	return nil
}

func maxTagValidator(val reflect.Value, opts ...string) error {
	if val.Kind() == reflect.Pointer && val.IsNil() {
		return nil
	}
	if err := valLimiter(val, opts[0], ">"); err != nil {
		return err
	}
	return nil
}

func minLengthTagValidator(val reflect.Value, opts ...string) error {
	if val.Kind() == reflect.Pointer && val.IsNil() {
		return nil
	}
	opts0Int, err := strconv.Atoi(opts[0])
	if err != nil {
		return errors.New(ErrMessage["numeric"])
	}

	valC := valStringer(val) // value converted
	if len(valC) < opts0Int {
		return fmt.Errorf(ErrMessage["minLength"], opts[0])
	}
	return nil
}

func maxLengthTagValidator(val reflect.Value, opts ...string) error {
	if val.Kind() == reflect.Pointer && val.IsNil() {
		return nil
	}
	opts0Int, err := strconv.Atoi(opts[0])
	if err != nil {
		return errors.New(ErrMessage["numeric"])
	}

	valC := valStringer(val) // value converted
	if len(valC) > opts0Int {
		return fmt.Errorf(ErrMessage["maxLength"], opts[0])
	}
	return nil
}

// //// some helper for the default field checker
func valLimiter(val reflect.Value, exptVal string, mode string) error {
	exptValFloat, err := strconv.ParseFloat(exptVal, 64)
	if err != nil {
		return err
	}

	valC := 0.0 // converted value
	valK := val.Kind()
	if valK == reflect.String {
		valCT, err := strconv.ParseFloat(val.String(), 64)
		if err != nil {
			return errors.New("nilai harus berupa angka/numerik")
		}
		valC = valCT
	} else if valK >= reflect.Int && valK <= reflect.Int64 {
		valC = float64(val.Int())
	} else if valK >= reflect.Uint && valK <= reflect.Uint64 {
		valC = float64(val.Uint())
	} else if valK <= reflect.Float32 && valK <= reflect.Float64 {
		valC = val.Float()
	}

	if mode == "<" {
		if exptValFloat > valC {
			return fmt.Errorf(ErrMessage["min"], exptVal)
		}
	} else {
		if exptValFloat < valC {
			return fmt.Errorf(fmt.Sprintf(ErrMessage["max"], exptVal))
		}
	}
	return nil
}
