package obconv

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonToMap(t *testing.T) {
	var j struct {
		Field1 string `json:"field1"`
		Field2 string `json:"field2"`
		Field3 string `json:"field3"`
	}

	j.Field1 = "value1"
	j.Field2 = "value2"
	j.Field3 = "value3"

	jo, _ := json.Marshal(j)
	js := string(jo)
	fmt.Println("JSON STRING : ", js)

	var m map[string]interface{}
	m, _ = JsonToMap(js)
	fmt.Println("MAP : ", m)
}

func TestSimpleStructToMap(t *testing.T) {
	var m struct {
		Field1 string  `map:"field_1"`
		Field2 string  `map:"field_2"`
		Field3 *string `map:"field_3"`
	}

	m.Field1 = "nilai 1"
	m.Field2 = "nilai 2"
	var a string = "test"
	m.Field3 = &a
	nm, _ := SimpleStructToMap(m)
	fmt.Println("STRUCT TO MAP : ", nm)
}
