package main

import (
	"encoding/json"
	"fmt"

	"github.com/randyardiansyah25/libpkg/net/http"
)

func main() {

	//client := tcp.NewTCPClient("localhost", 8081, 10)
	//st := client.Send("test")
	//fmt.Println(st.Code, " : ", st.Message)

	p := http.PostClient{
		Url:         "http://localhost:8081/",
		ContentType: "application/json",
		Timeout:     5,
	}
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
