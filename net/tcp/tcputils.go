package tcp

import (
	"fmt"
	"github.com/randyardiansyah25/libpkg/util/str"
	"strconv"
)

func SetHeader(message string, headerLen int) string {
	nHead := len(message)
	sHead := strutils.LeftPad(strconv.Itoa(nHead), headerLen, "0")
	return fmt.Sprint(sHead, message)

}
