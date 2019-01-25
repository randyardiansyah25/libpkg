package http

import (
	"fmt"
	"testing"
)

func TestGetClient_Get(t *testing.T) {
	client := GetClient{
		Url:     "http://192.176.249.10:80/sendsms",
		Timeout: 5,
	}
	var m = map[string]string{
		"username":    "admin",
		"password":    "admin",
		"phonenumber": "082121916325",
		"message":     "ok bro,, sms sudah keterima, thanks..!!",
	}
	resp := client.Get(m)
	fmt.Println("Status : ", resp.StatusCode(), ", Message : ", resp.Message())
}
