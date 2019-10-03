package strutils

import (
	"fmt"
	"testing"
)

func TestGenerateNumber(t *testing.T) {
	PrepareNumberGenerator()
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

func TestGenerateChars(t *testing.T) {
	PrepareNumberGenerator()
	s, err := GenerateChars(6)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("first: ", s)
	}

	s, err = GenerateNumber(6)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("second : ", s)
	}
}
