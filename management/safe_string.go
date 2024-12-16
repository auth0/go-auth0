package management

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// SafeString converts various input types to string with maximum safety
func SafeString(v interface{}) string {
	// Immediate nil check
	if v == nil {
		return ""
	}

	// Handle reflection-based type conversion with multiple safety checks
	rv := reflect.ValueOf(v)

	// Dereference pointer if needed
	if rv.Kind() == reflect.Ptr {
		// Nil pointer check
		if rv.IsNil() {
			return ""
		}
		// Get the actual value
		rv = rv.Elem()
	}

	// Try Stringer interface first
	if stringer, ok := v.(fmt.Stringer); ok {
		return safeStringerConversion(stringer)
	}

	// Type-specific conversion with safety checks
	switch rv.Kind() {
	case reflect.String:
		return safeStringConversion(rv.String())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return safeIntConversion(rv.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return safeUintConversion(rv.Uint())
	case reflect.Float32, reflect.Float64:
		return safeFloatConversion(rv.Float())
	case reflect.Bool:
		return safeBoolConversion(rv.Bool())
	case reflect.Struct:
		return safeStructConversion(rv)
	case reflect.Slice, reflect.Array:
		return safeSliceConversion(rv)
	case reflect.Map:
		return safeMapConversion(rv)
	default:
		return safeFallbackConversion(v)
	}
}

// Specialized conversion functions with additional safety

func safeStringerConversion(s fmt.Stringer) string {
	if s == nil {
		return ""
	}
	defer func() {
		recover() // Catch any panic from String() method
	}()
	return strings.TrimSpace(s.String())
}

func safeStringConversion(s string) string {
	return strings.TrimSpace(s)
}

func safeIntConversion(i int64) string {
	return strconv.FormatInt(i, 10)
}

func safeUintConversion(u uint64) string {
	return strconv.FormatUint(u, 10)
}

func safeFloatConversion(f float64) string {
	// Use 'f' format, precision of 6, 64-bit float
	return strconv.FormatFloat(f, 'f', 6, 64)
}

func safeBoolConversion(b bool) string {
	return strconv.FormatBool(b)
}

func safeStructConversion(rv reflect.Value) string {
	// Try to find a String() method
	method := rv.MethodByName("String")
	if method.IsValid() {
		result := method.Call(nil)
		if len(result) > 0 {
			return strings.TrimSpace(result[0].String())
		}
	}

	// Fallback to fmt.Sprintf with recovery
	defer func() {
		recover()
	}()
	return strings.TrimSpace(fmt.Sprintf("%v", rv.Interface()))
}

func safeSliceConversion(rv reflect.Value) string {
	if rv.Len() == 0 {
		return ""
	}

	// Convert first element, fallback to empty string
	defer func() {
		recover()
	}()

	firstElem := rv.Index(0)
	return SafeString(firstElem.Interface())
}

func safeMapConversion(rv reflect.Value) string {
	if rv.Len() == 0 {
		return ""
	}

	// Try to get first key's string representation
	defer func() {
		recover()
	}()

	keys := rv.MapKeys()
	if len(keys) > 0 {
		return SafeString(keys[0].Interface())
	}

	return ""
}

func safeFallbackConversion(v interface{}) string {
	defer func() {
		recover() // Catch any unexpected panics
	}()

	return strings.TrimSpace(fmt.Sprintf("%v", v))
}
