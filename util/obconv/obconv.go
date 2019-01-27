package obconv

import (
	"echannelgateway/model/errs"
	"encoding/json"
	"reflect"
)

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
		return nil, errs.ErrInvalidStruct
	}

	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		// gets us a StructField
		fi := typ.Field(i)
		if tagv := fi.Tag.Get(tag); tagv != "" {
			out[tagv] = v.Field(i).Interface()
		}
	}
	return out, nil
}

func SimpleStructToMap(in interface{}) (map[string]interface{}, error) {
	return SimpleStructToMapCustomTag(in, "map")
}
