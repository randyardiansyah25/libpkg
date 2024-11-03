package env

import (
	"os"
	"strconv"
)

func GetString(key string) string {
	return os.Getenv(key)
}

/*
Mengambil nilai int dari environment dengan mengirimkan default value.
Jika value pada environment tidak ditemukan, maka akan mengembalikan default value
*/
func GetInt(key string, def int) int {
	value := GetString(key)
	if value == "" {
		return def
	}

	if nvalue, err := strconv.Atoi(value); err != nil {
		return def
	} else {
		return nvalue
	}
}

/*
Mengambil nilai int32 dari environment dengan mengirimkan default value.
Jika value pada environment tidak ditemukan, maka akan mengembalikan default value
*/
func GetInt32(key string, def int32) int32 {
	value := GetString(key)
	if value == "" {
		return def
	}

	if nvalue, err := strconv.ParseInt(value, 10, 32); err != nil {
		return def
	} else {
		return int32(nvalue)
	}
}

/*
Mengambil nilai int64 dari environment dengan mengirimkan default value.
Jika value pada environment tidak ditemukan, maka akan mengembalikan default value
*/
func GetInt64(key string, def int64) int64 {
	value := GetString(key)
	if value == "" {
		return def
	}

	if nvalue, err := strconv.ParseInt(value, 10, 64); err != nil {
		return def
	} else {
		return nvalue
	}
}

/*
Mengambil nilai float32 dari environment dengan mengirimkan default value.
Jika value pada environment tidak ditemukan, maka akan mengembalikan default value
*/
func GetFloat32(key string, def float32) float32 {
	value := GetString(key)
	if value == "" {
		return def
	}

	if nvalue, err := strconv.ParseFloat(value, 32); err != nil {
		return def
	} else {
		return float32(nvalue)
	}
}

/*
Mengambil nilai float64 dari environment dengan mengirimkan default value.
Jika value pada environment tidak ditemukan, maka akan mengembalikan default value
*/
func GetFloat64(key string, def float64) float64 {
	value := GetString(key)
	if value == "" {
		return def
	}

	if nvalue, err := strconv.ParseFloat(value, 64); err != nil {
		return def
	} else {
		return nvalue
	}
}

/*
Mengambil nilai bool dari environment dengan mengirimkan default value.
Jika value pada environment tidak ditemukan, maka akan mengembalikan default value
*/
func GetBool(key string, def bool) bool {
	value := GetString(key)
	if value == "" {
		return def
	}

	if bvalue, err := strconv.ParseBool(value); err != nil {
		return def
	} else {
		return bvalue
	}
}
