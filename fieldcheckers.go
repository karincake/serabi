// built-in field checker
package serabi

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"

	"github.com/google/uuid"
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
	RegisterFieldChecker("numeric", numericTagValidator)
	RegisterFieldChecker("alpha", alphabetTagValidator)
	RegisterFieldChecker("alphaSpace", alphabetSTagValidator)
	RegisterFieldChecker("alphaNumeric", alphaNumericTagValidator)
	RegisterFieldChecker("alphaNumericU", alphaNumericUTagValidator)
	RegisterFieldChecker("email", emailTagValidator)
	RegisterFieldChecker("base64", base64TagValidator)
	RegisterFieldChecker("b64size", b64SizeKbTagValidator)
	RegisterFieldChecker("uuid", uuidTagValidator)
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

func numericTagValidator(val reflect.Value, opts ...string) error {
	if (val.Kind() == reflect.Pointer && val.IsNil()) || val.String() == "" {
		return nil
	}
	re := regexp.MustCompile("^[0-9]+$")

	if !re.MatchString(valStringer(val)) {
		return errors.New(ErrMessage["numeric"])
	}

	return nil
}

func alphabetTagValidator(val reflect.Value, opts ...string) error {
	if (val.Kind() == reflect.Pointer && val.IsNil()) || val.String() == "" {
		return nil
	}
	re := regexp.MustCompile("^[a-zA-Z]+$")

	if !re.MatchString(valStringer(val)) {
		return errors.New(ErrMessage["alpha"])
	}

	return nil
}

func alphabetSTagValidator(val reflect.Value, opts ...string) error {
	if (val.Kind() == reflect.Pointer && val.IsNil()) || val.String() == "" {
		return nil
	}
	re := regexp.MustCompile("^[ a-zA-Z]+$")

	if !re.MatchString(valStringer(val)) {
		return errors.New(ErrMessage["alphaSpace"])
	}

	return nil
}

func alphaNumericTagValidator(val reflect.Value, opts ...string) error {
	if (val.Kind() == reflect.Pointer && val.IsNil()) || val.String() == "" {
		return nil
	}
	re := regexp.MustCompile("^[a-zA-Z0-9]+$")

	if !re.MatchString(valStringer(val)) {
		return errors.New(ErrMessage["alphaNumeric"])
	}

	return nil
}

func alphaNumericUTagValidator(val reflect.Value, opts ...string) error {
	if (val.Kind() == reflect.Pointer && val.IsNil()) || val.String() == "" {
		return nil
	}
	re := regexp.MustCompile("^[a-zA-Z0-9_]+$")

	if !re.MatchString(valStringer(val)) {
		return errors.New(ErrMessage["alphaNumericU"])
	}

	return nil
}

func emailTagValidator(val reflect.Value, opts ...string) error {
	if (val.Kind() == reflect.Pointer && val.IsNil()) || val.String() == "" {
		return nil
	}
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if !re.MatchString(valStringer(val)) {
		return errors.New(ErrMessage["email"])
	}

	return nil
}

// check base64 string validation
func base64TagValidator(val reflect.Value, opts ...string) error {
	if (val.Kind() == reflect.Pointer && val.IsNil()) || val.String() == "" {
		return nil
	}
	re := regexp.MustCompile(`^(data:)([\w\/\+-.]*)(;charset=[\w-]+|;base64){0,1},([A-Za-z0-9+/]{4})*([A-Za-z0-9+/]{3}=|[A-Za-z0-9+/]{2}==)?$`)

	b64RawString := valStringer(val)
	if !re.MatchString(b64RawString) {
		return errors.New(ErrMessage["base64"])
	}

	return nil
}

// validate base64 approx size file
// using kb for parameter value eg: 1024 means 1024KB or 1MB or 1024000 B max allowed size file
func b64SizeKbTagValidator(val reflect.Value, opts ...string) error {
	if (val.Kind() == reflect.Pointer && val.IsNil()) || val.String() == "" {
		return nil
	}

	opts0Int, err := strconv.Atoi(opts[0])
	if err != nil {
		return errors.New(ErrMessage["numeric"])
	}

	datas := valStringer(val)

	l := len(datas)

	// count how many trailing '=' there are (if any)
	eq := 0
	if l >= 2 {
		if datas[l-1] == '=' {
			eq++
		}
		if datas[l-2] == '=' {
			eq++
		}

		l -= eq
	}

	// basically:
	// eq == 0 :	bits-wasted = 0
	// eq == 1 :	bits-wasted = 2
	// eq == 2 :	bits-wasted = 4

	// so orig length ==  (l*6 - eq*2) / 8

	// if bytes size > max bytes allowed then
	if (l*3-eq)/4 > opts0Int*1000 {
		return fmt.Errorf(ErrMessage["fileSizeBase64"], opts0Int)
	}

	return nil
}

func uuidTagValidator(val reflect.Value, opts ...string) error {
	if (val.Kind() == reflect.Pointer && val.IsNil()) || val.String() == "" {
		return nil
	}

	_, err := uuid.Parse(valStringer(val))
	if err != nil {
		return errors.New(ErrMessage["numeric"])
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
