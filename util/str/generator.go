package strutils

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// noinspection GoErrorStringFormat
func GenerateNumber(length int) (string, error) {
	src := rand.NewSource(time.Now().UnixNano())
	random := rand.New(src)
	num := make([]string, 0)

	for i := 0; i < length; i++ {
		n := random.Intn(9)
		num = append(num, strconv.Itoa(n))
	}

	return strings.Join(num, ""), nil

}

func GenerateChars(length int) (string, error) {
	src := rand.NewSource(time.Now().UnixNano())
	random := rand.New(src)
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[random.Intn(len(chars))])
	}
	return b.String(), nil

}
