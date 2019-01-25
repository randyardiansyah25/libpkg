package strutils

import (
	"testing"
)

func TestLeftPad(t *testing.T) {
	var str, padstring string
	str, padstring = "Hello world", "0"
	padlen := 15
	t.Log("String       :", str)
	t.Log("Padding with :", padstring)
	t.Log("Padding Len  :", padlen)
	t.Log("Result       :", LeftPad(str, padlen, padstring))
}

func TestRightPadPad(t *testing.T) {
	var str, padstring string
	str, padstring = "Hello world", "_"
	padlen := 15
	t.Log("String       :", str)
	t.Log("Padding with :", padstring)
	t.Log("Padding Len  :", padlen)
	t.Log("Result       :", RightPad(str, padlen, padstring))
}
