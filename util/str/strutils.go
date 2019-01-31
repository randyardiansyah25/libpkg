package strutils

import (
	"encoding/hex"
	"fmt"
	"strings"
)

func LeftPad(s string, length int, pad string) string {
	if len(s) >= length {
		return s
	}
	padding := strings.Repeat(pad, length-len(s))
	return fmt.Sprintf("%s%s", padding, s)
}

func RightPad(s string, length int, pad string) string {
	if len(s) >= length {
		return s
	}
	padding := strings.Repeat(pad, length-len(s))
	return fmt.Sprintf("%s%s", s, padding)
}

func ByteToHexString(src []byte) string {
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)
	test := string(dst)
	return strings.ToUpper(test)
}

func HexStringToByte(src string) ([]byte, error) {
	srcb := []byte(src)
	dst := make([]byte, hex.DecodedLen(len(srcb)))
	n, err := hex.Decode(dst, srcb)
	if err != nil {
		return nil, err
	}
	return dst[:n], nil
}

func PhoneNumberEscape(phoneNumber string) string {
	phoneNumber = strings.TrimSpace(phoneNumber)
	if phoneNumber == "" {
		return phoneNumber
	} else if strings.HasPrefix(phoneNumber, "0") {
		np := []string{"62", phoneNumber[1:len(phoneNumber)]}
		return strings.Join(np, "")
	} else if strings.HasPrefix(phoneNumber, "+") {
		return phoneNumber[1:len(phoneNumber)]
	} else {
		return phoneNumber
	}
}
