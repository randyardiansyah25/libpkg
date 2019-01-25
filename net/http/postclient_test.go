package http

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestPostClient_PostJson(t *testing.T) {
	p := PostClient{"http://localhost:8081/", "application/json", 5}
	var m = map[string]interface{}{}
	kodeBank := "0002"
	m["kode_bank"] = kodeBank
	m["no_rekening"] = "001.01.00001"
	m["no_hp"] = "082121916325"
	m["amount"] = 15000

	mJson, _ := json.Marshal(m)
	resp := p.PostJson(mJson)
	fmt.Println("Status : ", resp.StatusCode(), ", Message : ", resp.Message())
}

func TestPostClient_PostParam(t *testing.T) {
	p := PostClient{"http://localhost:8081/", "application/x-www-form-urlencoded", 5}
	var m = map[string]interface{}{}
	kodeBank := "0002"
	m["kode_bank"] = kodeBank
	m["no_rekening"] = "001.01.00001"
	m["no_hp"] = "082121916325"
	m["amount"] = 15000

	resp := p.PostParam(m)
	fmt.Println("Status : ", resp.StatusCode(), ", Message : ", resp.Message())
}
