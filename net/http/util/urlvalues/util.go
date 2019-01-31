package urlvalues

import (
	"errors"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

func AddParams(url_ string, params url.Values) string {
	if len(params) == 0 {
		return url_
	}

	if !strings.Contains(url_, "?") {
		url_ += "?"
	}

	if strings.HasSuffix(url_, "?") || strings.HasSuffix(url_, "&") {
		url_ += params.Encode()
	} else {
		url_ += "&" + params.Encode()
	}

	return url_
}

func ToMapString(values url.Values) map[string]string {
	mv := make(map[string]string)
	for key, val := range values {
		var v string
		if len(val) > 0 {
			v = val[0]
		}
		mv[key] = v
	}
	return mv
}

func ToUrlValues(v interface{}) url.Values {
	switch t := v.(type) {
	case url.Values:
		return t
	case map[string][]string:
		return url.Values(t)
	case map[string]string:
		rst := make(url.Values)
		for k, v := range t {
			rst.Add(k, v)
		}
		return rst
	case nil:
		return make(url.Values)
	default:
		panic("Invalid value")
	}
}

func Marshal(in interface{}, m map[string]string, tag string) error {
	v := reflect.ValueOf(in)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return errors.New("Invalid struct")
	}

	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Type.Kind() == reflect.String {
			v.Field(i).SetString(m[f.Tag.Get(tag)])
		} else if f.Type.Kind() == reflect.Float64 || f.Type.Kind() == reflect.Float32 {
			mval := m[f.Tag.Get(tag)]
			if len(mval) > 0 {
				val, _ := strconv.ParseFloat(mval, 64)
				v.Field(i).SetFloat(val)
			}
		} else if f.Type.Kind() == reflect.Int || f.Type.Kind() == reflect.Int64 || f.Type.Kind() == reflect.Int32 {
			mval := m[f.Tag.Get(tag)]
			if len(mval) > 0 {
				val, _ := strconv.ParseInt(m[f.Tag.Get(tag)], 10, 64)
				v.Field(i).SetInt(val)
			}
		}
	}
	in = v.Interface()
	return nil
}
