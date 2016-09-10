package sys

import (
	"crypto/sha512"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func HasValue(source string) bool {
	return len(source) > 0
}

func AsDate(source string) (time.Time, bool) {
	t, err := time.Parse(DATE_FMT, source)
	return t, err == nil
}

func AsTime(source string) (time.Time, bool) {
	t, err := time.Parse(DATETIME_FMT, source)
	return t, err == nil
}

func AsFloat64(source string, defaultValue float64) float64 {
	i, err := strconv.ParseFloat(source, 64)
	if err != nil {
		return defaultValue
	} else {
		return i
	}
}

func AsInt64(source string, defaultValue int64) int64 {
	i, err := strconv.ParseInt(source, 10, 64)
	if err != nil {
		return defaultValue
	} else {
		return i
	}
}

func AsInt(source string, defaultValue int) int {
	i, err := strconv.Atoi(source)
	if err != nil {
		return defaultValue
	} else {
		return i
	}
}

// 将个位号码统一成带十位的号 如: "9" -> "09"
func AsZeroNo(no string) string {
	if len(no) == 1 {
		return "0" + no
	}

	return no
}

func Atoi64(v string, d int64) int64 {
	if i, err := strconv.ParseInt(v, 10, 64); err != nil {
		return d
	} else {
		return i
	}
}

func Sha512(s string) string {
	data := []byte(s)
	return strings.ToLower(fmt.Sprintf("%x", sha512.Sum512(data)))
}
