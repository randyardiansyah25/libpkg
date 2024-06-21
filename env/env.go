package env

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetString(key string) string {
	return os.Getenv(key)
}

func GetInt(key string) (int, error) {
	v := os.Getenv(key)
	return strconv.Atoi(v)
}

func GetInt64(key string) (int64, error) {
	v := os.Getenv(key)
	return strconv.ParseInt(v, 10, 64)
}

func GetInt32(key string) (int32, error) {
	v := os.Getenv(key)
	if res, er := strconv.ParseInt(v, 10, 32); er != nil {
		return 0, er
	} else {
		return int32(res), nil
	}

}

func GetFloat64(key string) (float64, error) {
	v := os.Getenv(key)
	return strconv.ParseFloat(v, 64)
}

func GetFloat32(key string) (float32, error) {
	v := os.Getenv(key)
	if res, er := strconv.ParseFloat(v, 32); er != nil {
		return 0, er
	} else {
		return float32(res), nil
	}
}

type xslice []string

func (x xslice) Has(a string) bool {
	for _, v := range x {
		if v == a {
			return true
		}
	}
	return false
}

func GetBool(key string) (bool, error) {
	v := os.Getenv(key)
	v = strings.ToLower(v)
	var allowedKeyword = xslice{"yes", "y", "1", "true", "ya", "t", "no", "n", "tidak", "0", "false"}
	var trueConditionLIst = xslice{"yes", "y", "1", "true", "ya", "t"}

	if !allowedKeyword.Has(v) {
		recognized := strings.Join(allowedKeyword, "|")
		return false, fmt.Errorf("invalid boolean value for key %s. Recognized value is [%s]", key, recognized)
	}

	return trueConditionLIst.Has(v), nil

}
