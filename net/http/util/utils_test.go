package util

import (
	"fmt"
	"net/url"
	"testing"
)

func TestUrlValuesToMapString(t *testing.T) {
	values := url.Values{
		"test1": {"value test1"},
		"test2": {"value test2"},
	}
	mval := UrlValuesToMapString(values)
	for k, v := range mval {
		fmt.Println(k, " : ", v[0])
	}
}
