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
		NamaFieldString        string  `map:"field_1"`
		NamaFieldFloat         float64 `map:"field_2"`
		NamaFieldPointerstring *string `map:"field_3"`
	}

	m.NamaFieldString = "nilai 1"
	m.NamaFieldFloat = 10.2
	var a string = "test"
	m.NamaFieldPointerstring = &a
	nm, _ := SimpleStructToMap(m)
	fmt.Println("STRUCT TO MAP : ", nm)
}

func TestScanMapToStruct(t *testing.T) {
	m := map[string]string{
		"field_1": "field 1",
		"field_2": "1000",
		"field_3": "5",
	}
	var st = struct {
		NamaFieldString string  `map:"field_1"`
		NamaFieldFloat  float64 `map:"field_2"`
		NamaFieldInt    int     `map:"field_3"`
	}{}

	err := ScanMapToStruct(&st, m, "map")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("NamaFieldString : ", st.NamaFieldString)
	fmt.Println("NamaFieldFloet : ", st.NamaFieldFloat)
	fmt.Println("NamaFieldInt : ", st.NamaFieldInt)
}
