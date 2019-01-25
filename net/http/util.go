package http

import (
	"net/url"
	"strings"
)

func AddParams(url_ string, params url.Values) string {
	if len(params) == 0 {
		return url_
	}

	if !strings.Contains(url_, "?") {
		url_ += "?"
	}

	if strings.HasSuffix(url_, "?") || strings.HasSuffix(url_, "&") {
		url_ += params.Encode()
	} else {
		url_ += "&" + params.Encode()
	}

	return url_
}

func ToUrlValues(v interface{}) url.Values {
	switch t := v.(type) {
	case url.Values:
		return t
	case map[string][]string:
		return url.Values(t)
	case map[string]string:
		rst := make(url.Values)
		for k, v := range t {
			rst.Add(k, v)
		}
		return rst
	case nil:
		return make(url.Values)
	default:
		panic("Invalid value")
	}
}
