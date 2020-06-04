package logzio

import (
	"encoding/json"
	"reflect"
	"strings"
	"unicode"
)

const (
	BASE_10            int    = 10
	BITSIZE_64         int    = 64
	VALIDATE_URL_REGEX string = "^http(s):\\/\\/"
)

func findStringInArray(v string, values []string) bool {
	for i := 0; i < len(values); i++ {
		value := values[i]
		if strings.EqualFold(v, value) {
			return true
		}
	}
	return false
}

func stripAllWhitespace(inputString string) string {
	var b strings.Builder
	b.Grow(len(inputString))
	for _, ch := range inputString {
		if !unicode.IsSpace(ch) {
			b.WriteRune(ch)
		}
	}
	return b.String()
}

func jsonEqual(old, new string) bool {
	if old == new {
		return true
	}

	var expected, actual interface{}
	oldString := stripAllWhitespace(old)
	newString := stripAllWhitespace(new)

	if err := json.Unmarshal([]byte(oldString), &expected); err != nil {
		return oldString == newString
	}

	if err := json.Unmarshal([]byte(newString), &actual); err != nil {
		return false
	}

	return reflect.DeepEqual(expected, actual)
}
