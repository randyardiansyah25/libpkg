package urlvalues

import (
	"fmt"
	"testing"
)

type Sample struct {
	Field1 string `param:"field1"`
	Field2 string `param:"field2"`
	Field3 string `param:"field3"`
}

func TestMarshal(t *testing.T) {
	m := make(map[string]string)

	m["FIELD1"] = "test field 1"
	m["Field2"] = "test field 2"
	m["field3"] = "test field 3"

	sample := Sample{}

	_ = Marshal(m, &sample, "param")

	fmt.Println(sample)
}
