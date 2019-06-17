package strutils

import (
	"fmt"
	"testing"
)

func TestGenerateNumber(t *testing.T) {
	Prepare()
	s, err := GenerateNumber(12)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("first: ", s)
	}

	s, err = GenerateNumber(12)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("second : ", s)
	}
}
