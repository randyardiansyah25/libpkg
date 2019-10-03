package strutils

import (
	"errors"
	"math/rand"
	"strconv"
	"strings"
)

var prepared bool

func PrepareNumberGenerator() {
	prepared = true
}

//noinspection GoErrorStringFormat
func GenerateNumber(length int) (string, error) {
	if prepared {
		num := make([]string, 0)

		for i := 0; i < length; i++ {
			n := rand.Intn(9)
			num = append(num, strconv.Itoa(n))
		}

		return strings.Join(num, ""), nil
	} else {
		return "", errors.New("Please prepare by adding rand.Seed(time.Now().UTC().UnixNano()) on Main Module")
	}
}

func GenerateChars(length int) (string, error) {
	if prepared {
		chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
		var b strings.Builder
		for i := 0; i < length; i++ {
			b.WriteRune(chars[rand.Intn(len(chars))])
		}
		return b.String(), nil
	} else {
		return "", errors.New("Please prepare by adding rand.Seed(time.Now().UTC().UnixNano()) on Main Module")
	}
}
