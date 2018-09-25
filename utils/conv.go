package utils

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// Seconds-based time units
const (
	Minute = 60
	Hour   = 60 * Minute
	Day    = 24 * Hour
	Week   = 7 * Day
	Month  = 30 * Day
	Year   = 12 * Month
)

func IntPtrTo64(ptr interface{}) (value int64) {
	if v := reflect.ValueOf(ptr); v.Kind() == reflect.Ptr {
		p := v.Elem()
		switch p.Kind() {
		case reflect.Int:
			value = int64(*ptr.(*int))
		case reflect.Int8:
			value = int64(*ptr.(*int8))
		case reflect.Int16:
			value = int64(*ptr.(*int16))
		case reflect.Int32:
			value = int64(*ptr.(*int32))
		case reflect.Int64:
			value = *ptr.(*int64)
		}
	}
	return
}

func UintPtrTo64(ptr interface{}) (value uint64) {
	if v := reflect.ValueOf(ptr); v.Kind() == reflect.Ptr {
		p := v.Elem()
		switch p.Kind() {
		case reflect.Uint:
			value = uint64(*ptr.(*uint))
		case reflect.Uint8:
			value = uint64(*ptr.(*uint8))
		case reflect.Uint16:
			value = uint64(*ptr.(*uint16))
		case reflect.Uint32:
			value = uint64(*ptr.(*uint32))
		case reflect.Uint64:
			value = *ptr.(*uint64)
		}
	}
	return
}

func computeTimeDiff(diff int64) (int64, string) {
	diffStr := ""
	switch {
	case diff < 1*Minute:
		diffStr = "不足 1 分钟"
		diff = 0

	case diff < 1*Hour:
		diffStr = fmt.Sprintf("%d 分钟", diff/Minute)
		diff -= diff / Minute * Minute

	case diff < 1*Day:
		diffStr = fmt.Sprintf("%d 小时", diff/Hour)
		diff -= diff / Hour * Hour

	default:
		diffStr = fmt.Sprintf("%d 天", diff/Day)
		diff -= diff / Day * Day
		diff = 0
	}
	return diff, diffStr
}

// TimeSincePro calculates the time interval and generate full user-friendly string.
func TimeSincePro(then time.Time) string {
	now := time.Now()
	diff := now.Unix() - then.Unix()

	if then.After(now) {
		return "future"
	}

	var timeStr, diffStr string
	for {
		if diff < 1*Minute && diffStr != "" {
			break
		}

		diff, diffStr = computeTimeDiff(diff)
		timeStr += ", " + diffStr
	}
	return strings.TrimPrefix(timeStr, ", ")
}

// Subtract deals with subtraction of all types of number.
func Subtract(left interface{}, right interface{}) interface{} {
	var rleft, rright int64
	var fleft, fright float64
	var isInt bool = true
	switch left.(type) {
	case int:
		rleft = int64(left.(int))
	case int8:
		rleft = int64(left.(int8))
	case int16:
		rleft = int64(left.(int16))
	case int32:
		rleft = int64(left.(int32))
	case int64:
		rleft = left.(int64)
	case float32:
		fleft = float64(left.(float32))
		isInt = false
	case float64:
		fleft = left.(float64)
		isInt = false
	}

	switch right.(type) {
	case int:
		rright = int64(right.(int))
	case int8:
		rright = int64(right.(int8))
	case int16:
		rright = int64(right.(int16))
	case int32:
		rright = int64(right.(int32))
	case int64:
		rright = right.(int64)
	case float32:
		fright = float64(left.(float32))
		isInt = false
	case float64:
		fleft = left.(float64)
		isInt = false
	}

	if isInt {
		return rleft - rright
	} else {
		return fleft + float64(rleft) - (fright + float64(rright))
	}
}

// EllipsisString returns a truncated short string,
// it appends '...' in the end of the length of string is too large.
func EllipsisString(str string, length int) string {
	if len(str) < length {
		return str
	}
	return str[:length-3] + "..."
}

// TruncateString returns a truncated string with given limit,
// it returns input string if length is not reached limit.
func TruncateString(str string, limit int) string {
	if len(str) < limit {
		return str
	}
	return str[:limit]
}

// StringsToInt64s converts a slice of string to a slice of int64.
func StringsToInt64s(strs []string) []int64 {
	ints := make([]int64, len(strs))
	for i := range strs {
		v, _ := strconv.ParseInt(strs[i], 10, 64)
		ints[i] = v
	}
	return ints
}

// Int64sToStrings converts a slice of int64 to a slice of string.
func Int64sToStrings(ints []int64) []string {
	strs := make([]string, len(ints))
	for i := range ints {
		strs[i] = strconv.FormatInt(ints[i], 10)
	}
	return strs
}

// Int64sToMap converts a slice of int64 to a int64 map.
func Int64sToMap(ints []int64) map[int64]bool {
	m := make(map[int64]bool)
	for _, i := range ints {
		m[i] = true
	}
	return m
}

// IsLetter reports whether the rune is a letter (category L).
// https://github.com/golang/go/blob/master/src/go/scanner/scanner.go#L257
func IsLetter(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch >= 0x80 && unicode.IsLetter(ch)
}
