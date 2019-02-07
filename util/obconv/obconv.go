package obconv

import (
	"encoding/json"
	"errors"
	"reflect"
	"strconv"
)

var ErrInvalidStruct = errors.New("invalid struct : only accepts structs")

func JsonToMap(msg string) (map[string]interface{}, error) {
	var msgTemplate interface{}
	err := json.Unmarshal([]byte(msg), &msgTemplate)
	if err == nil {
		return msgTemplate.(map[string]interface{}), nil
	} else {
		return nil, err
	}
}

func SimpleStructToMapCustomTag(in interface{}, tag string) (map[string]interface{}, error) {
	out := make(map[string]interface{})

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// we only accept structs
	if v.Kind() != reflect.Struct {
		return nil, ErrInvalidStruct
	}

	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		// gets us a StructField

		fi := typ.Field(i)
		//println("field name : ", fi.Name)
		//println("field type : ", fi.Type.Name())
		//println("field type is string : ", fi.Type.Kind() == reflect.String)
		//println("field type is pointer : ", fi.Type.Kind() == reflect.Ptr)
		//println("field type is float64 : ", fi.Type.Kind() == reflect.Float64)
		//println("==============================")
		if tagv := fi.Tag.Get(tag); tagv != "" {
			out[tagv] = v.Field(i).Interface()
		}
	}
	return out, nil
}

func MapStringToStruct(in interface{}, m map[string]string, tag string) error {
	v := reflect.ValueOf(in)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return errs.ErrInvalidStruct
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

func SimpleStructToMap(in interface{}) (map[string]interface{}, error) {
	return SimpleStructToMapCustomTag(in, "map")
}
