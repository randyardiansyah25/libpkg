package strutils

import (
	"fmt"
	"testing"
)

func TestGenerateNumber(t *testing.T) {
	s := GenerateNumber(12)

	fmt.Println("first: ", s)

	s = GenerateNumber(12)

	fmt.Println("second : ", s)

}

func TestGenerateChars(t *testing.T) {
	s := GenerateChars(6)

	fmt.Println("first: ", s)

	s = GenerateNumber(6)

	fmt.Println("second : ", s)

}
