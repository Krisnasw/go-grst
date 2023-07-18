package validationrule

import (
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"regexp"
	"strings"
	"unicode/utf8"
)

func ruleRequired(in interface{}, param string) error {
	st := reflect.ValueOf(in)
	valid := true
	switch st.Kind() {
	case reflect.String:
		valid = utf8.RuneCountInString(st.String()) != 0
	case reflect.Ptr, reflect.Interface:
		valid = !st.IsNil()
	case reflect.Slice, reflect.Map, reflect.Array:
		valid = st.Len() != 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		valid = st.Int() != 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		valid = st.Uint() != 0
	case reflect.Float32, reflect.Float64:
		valid = st.Float() != 0
	case reflect.Bool:
		valid = st.Bool()
	case reflect.Invalid:
		valid = false // always invalid
	case reflect.Struct:
		valid = true // always valid since only nil pointers are empty
	default:
		return errors.New("unsupported type")
	}

	if !valid {
		return errors.New("param `{field}` is required")
	}
	return nil
}

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func ruleEmail(in interface{}, param string) error {
	inStr := ""
	if v, ok := in.(string); ok {
		inStr = v
	} else {
		return errors.New("rule `email` can only be used by string data type")
	}

	if len(inStr) < 3 && len(inStr) > 254 {
		return errors.New("invalid email length")
	}
	if !emailRegex.MatchString(inStr) {
		return errors.New("invalid email format")
	}
	return nil
}

func ruleUrl(in interface{}, param string) (err error) {
	inStr := ""
	if v, ok := in.(string); ok {
		inStr = v
	} else {
		return errors.New("rule `email` can only be used by string data type")
	}

	if inStr == "" {
		return nil
	}

	u, err := url.Parse(inStr)
	if err != nil {
		return fmt.Errorf("invalid url format: %s", err.Error())
	}

	if (u.Scheme == "" || u.Host == "") && inStr != "localhost" {
		return fmt.Errorf("invalid url format: scheme or host is empty")
	}

	return nil
}

func ruleEnum(in interface{}, param string) (err error) {
	inStr := ""
	if v, ok := in.(string); ok {
		inStr = v
	} else {
		return errors.New("rule `enum` can only be used by string data type")
	}

	enumList := strings.Split(param, "|")
	for _, e := range enumList {
		if e == inStr {
			return nil
		}
	}

	return errors.New("rule `enum` require value " + strings.Join(enumList, "|") + ". Got: " + inStr)

}
