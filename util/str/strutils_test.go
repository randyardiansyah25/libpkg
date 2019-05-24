package strutils

import (
	"fmt"
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

func TestCenterPad(t *testing.T) {
	var str string
	str = "hello world!"

	t.Log("String       :", str)
	t.Log("Center :", CenterPad(str, 32, "*"))
}

func TestFormatCenter(t *testing.T) {
	var str string
	str = "untuk check balance, dikarenakan di app Dainan ada fitur pembiayaan , atau tabungan saldo kurang , sehingga meminta satu informasi saldo_akhir dari saldo tasakur-nya."

	t.Log("String       :", str)
	fmt.Print("Center :", "\n", FormatCenter(str, true, 32))
}
