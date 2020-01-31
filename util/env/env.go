package env

import (
	"os"
	"strconv"
)

func GetString(key string) string {
	return os.Getenv(key)
}

func GetInt(key string) int {
	value := GetString(key)
	if value == "" {
		return 0
	}

	if nvalue, err := strconv.Atoi(value); err != nil {
		return 0
	} else {
		return nvalue
	}
}

func GetInt32(key string) int32 {
	value := GetString(key)
	if value == "" {
		return 0
	}

	if nvalue, err := strconv.ParseInt(value, 10, 32); err != nil {
		return 0
	} else {
		return int32(nvalue)
	}
}

func GetInt64(key string) int64 {
	value := GetString(key)
	if value == "" {
		return 0
	}

	if nvalue, err := strconv.ParseInt(value, 10, 64); err != nil {
		return 0
	} else {
		return nvalue
	}
}

func GetFloat32(key string) float32 {
	value := GetString(key)
	if value == "" {
		return 0
	}

	if nvalue, err := strconv.ParseFloat(value, 32); err != nil {
		return 0
	} else {
		return float32(nvalue)
	}
}

func GetFloat64(key string) float64 {
	value := GetString(key)
	if value == "" {
		return 0
	}

	if nvalue, err := strconv.ParseFloat(value, 64); err != nil {
		return 0
	} else {
		return nvalue
	}
}

func GetBool(key string) bool {
	value := GetString(key)
	if value == "" {
		return false
	}

	if bvalue, err := strconv.ParseBool(value); err != nil {
		return false
	} else {
		return bvalue
	}
}
