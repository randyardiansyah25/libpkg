package util

import "net/url"

func UrlValuesToMapString(values url.Values) map[string]string {
	var mv map[string]string
	for key, val := range values {
		var v string
		if len(val) > 0 {
			v = val[0]
		}
		mv[key] = v
	}
	return mv
}
