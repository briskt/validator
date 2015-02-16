package validator

import (
	"fmt"
	"reflect"
	"strconv"
)

// BakedInValidators is the map of ValidationFunc used internally
// but can be used with any new Validator if desired
var BakedInValidators = map[string]ValidationFunc{
	"required":    required,
	"len":         length,
	"min":         min,
	"max":         max,
	"lt":          lt,
	"lte":         lte,
	"gt":          gt,
	"gte":         gte,
	"alpha":       alpha,
	"alphanum":    alphanum,
	"numeric":     numeric,
	"number":      number,
	"hexadecimal": hexadecimal,
	"hexcolor":    hexcolor,
	"rgb":         rgb,
	"rgba":        rgba,
	"hsl":         hsl,
	"hsla":        hsla,
	"email":       email,
}

func email(field interface{}, param string) bool {

	st := reflect.ValueOf(field)

	switch st.Kind() {

	case reflect.String:
		return emailRegex.MatchString(field.(string))
	default:
		panic(fmt.Sprintf("Bad field type %T", field))
	}
}

func hsla(field interface{}, param string) bool {

	st := reflect.ValueOf(field)

	switch st.Kind() {

	case reflect.String:
		return hslaRegex.MatchString(field.(string))
	default:
		panic(fmt.Sprintf("Bad field type %T", field))
	}
}

func hsl(field interface{}, param string) bool {

	st := reflect.ValueOf(field)

	switch st.Kind() {

	case reflect.String:
		return hslRegex.MatchString(field.(string))
	default:
		panic(fmt.Sprintf("Bad field type %T", field))
	}
}

func rgba(field interface{}, param string) bool {

	st := reflect.ValueOf(field)

	switch st.Kind() {

	case reflect.String:
		return rgbaRegex.MatchString(field.(string))
	default:
		panic(fmt.Sprintf("Bad field type %T", field))
	}
}

func rgb(field interface{}, param string) bool {

	st := reflect.ValueOf(field)

	switch st.Kind() {

	case reflect.String:
		return rgbRegex.MatchString(field.(string))
	default:
		panic(fmt.Sprintf("Bad field type %T", field))
	}
}

func hexcolor(field interface{}, param string) bool {

	st := reflect.ValueOf(field)

	switch st.Kind() {

	case reflect.String:
		return hexcolorRegex.MatchString(field.(string))
	default:
		panic(fmt.Sprintf("Bad field type %T", field))
	}
}

func hexadecimal(field interface{}, param string) bool {

	st := reflect.ValueOf(field)

	switch st.Kind() {

	case reflect.String:
		return hexadecimalRegex.MatchString(field.(string))
	default:
		panic(fmt.Sprintf("Bad field type %T", field))
	}
}

func number(field interface{}, param string) bool {

	st := reflect.ValueOf(field)

	switch st.Kind() {

	case reflect.String:
		return numberRegex.MatchString(field.(string))
	default:
		panic(fmt.Sprintf("Bad field type %T", field))
	}
}

func numeric(field interface{}, param string) bool {

	st := reflect.ValueOf(field)

	switch st.Kind() {

	case reflect.String:
		return numericRegex.MatchString(field.(string))
	default:
		panic(fmt.Sprintf("Bad field type %T", field))
	}
}

func alphanum(field interface{}, param string) bool {

	st := reflect.ValueOf(field)

	switch st.Kind() {

	case reflect.String:
		return alphaNumericRegex.MatchString(field.(string))
	default:
		panic(fmt.Sprintf("Bad field type %T", field))
	}
}

func alpha(field interface{}, param string) bool {

	st := reflect.ValueOf(field)

	switch st.Kind() {

	case reflect.String:
		return alphaRegex.MatchString(field.(string))
	default:
		panic(fmt.Sprintf("Bad field type %T", field))
	}
}

func required(field interface{}, param string) bool {

	st := reflect.ValueOf(field)

	switch st.Kind() {

	case reflect.Slice, reflect.Map, reflect.Array:
		return field != nil && int64(st.Len()) > 0

	default:
		return field != nil && field != reflect.Zero(reflect.TypeOf(field)).Interface()
	}
}

func gte(field interface{}, param string) bool {

	st := reflect.ValueOf(field)

	switch st.Kind() {

	case reflect.String:
		p := asInt(param)

		return int64(len(st.String())) >= p

	case reflect.Slice, reflect.Map, reflect.Array:
		p := asInt(param)

		return int64(st.Len()) >= p

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		p := asInt(param)

		return st.Int() >= p

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		p := asUint(param)

		return st.Uint() >= p

	case reflect.Float32, reflect.Float64:
		p := asFloat(param)

		return st.Float() >= p

	default:
		panic(fmt.Sprintf("Bad field type for Input Param %s for %s\n", param, field))
	}
}

func gt(field interface{}, param string) bool {

	st := reflect.ValueOf(field)

	switch st.Kind() {

	case reflect.String:
		p := asInt(param)

		return int64(len(st.String())) > p

	case reflect.Slice, reflect.Map, reflect.Array:
		p := asInt(param)

		return int64(st.Len()) > p

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		p := asInt(param)

		return st.Int() > p

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		p := asUint(param)

		return st.Uint() > p

	case reflect.Float32, reflect.Float64:
		p := asFloat(param)

		return st.Float() > p

	default:
		panic(fmt.Sprintf("Bad field type for Input Param %s for %s\n", param, field))
	}
}

// length tests whether a variable's length is equal to a given
// value. For strings it tests the number of characters whereas
// for maps and slices it tests the number of items.
func length(field interface{}, param string) bool {

	st := reflect.ValueOf(field)

	switch st.Kind() {

	case reflect.String:
		p := asInt(param)

		return int64(len(st.String())) == p

	case reflect.Slice, reflect.Map, reflect.Array:
		p := asInt(param)

		return int64(st.Len()) == p

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		p := asInt(param)

		return st.Int() == p

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		p := asUint(param)

		return st.Uint() == p

	case reflect.Float32, reflect.Float64:
		p := asFloat(param)

		return st.Float() == p

	default:
		panic(fmt.Sprintf("Bad field type for Input Param %s for %s\n", param, field))
	}
}

// min tests whether a variable value is larger or equal to a given
// number. For number types, it's a simple lesser-than test; for
// strings it tests the number of characters whereas for maps
// and slices it tests the number of items.
func min(field interface{}, param string) bool {

	return gte(field, param)
}

func lte(field interface{}, param string) bool {

	st := reflect.ValueOf(field)

	switch st.Kind() {

	case reflect.String:
		p := asInt(param)

		return int64(len(st.String())) <= p

	case reflect.Slice, reflect.Map, reflect.Array:
		p := asInt(param)

		return int64(st.Len()) <= p

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		p := asInt(param)

		return st.Int() <= p

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		p := asUint(param)

		return st.Uint() <= p

	case reflect.Float32, reflect.Float64:
		p := asFloat(param)

		return st.Float() <= p

	default:
		panic(fmt.Sprintf("Bad field type for Input Param %s for %s\n", param, field))
	}
}

func lt(field interface{}, param string) bool {

	st := reflect.ValueOf(field)

	switch st.Kind() {

	case reflect.String:
		p := asInt(param)

		return int64(len(st.String())) < p

	case reflect.Slice, reflect.Map, reflect.Array:
		p := asInt(param)

		return int64(st.Len()) < p

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		p := asInt(param)

		return st.Int() < p

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		p := asUint(param)

		return st.Uint() < p

	case reflect.Float32, reflect.Float64:
		p := asFloat(param)

		return st.Float() < p

	default:
		panic(fmt.Sprintf("Bad field type for Input Param %s for %s\n", param, field))
	}
}

// max tests whether a variable value is lesser than a given
// value. For numbers, it's a simple lesser-than test; for
// strings it tests the number of characters whereas for maps
// and slices it tests the number of items.
func max(field interface{}, param string) bool {

	return lte(field, param)
}

// asInt retuns the parameter as a int64
// or panics if it can't convert
func asInt(param string) int64 {

	i, err := strconv.ParseInt(param, 0, 64)

	if err != nil {
		panic(fmt.Sprintf("Bad Input Param %s\n", param))
	}

	return i
}

// asUint returns the parameter as a uint64
// or panics if it can't convert
func asUint(param string) uint64 {

	i, err := strconv.ParseUint(param, 0, 64)

	if err != nil {
		panic(fmt.Sprintf("Bad Input Param %s\n", param))
	}

	return i
}

// asFloat returns the parameter as a float64
// or panics if it can't convert
func asFloat(param string) float64 {

	i, err := strconv.ParseFloat(param, 64)

	if err != nil {
		panic(fmt.Sprintf("Bad Input Param %s\n", param))
	}

	return i
}