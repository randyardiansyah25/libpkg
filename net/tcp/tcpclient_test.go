package tcp

import (
	"fmt"
	"net"
	"testing"
)

//BASIC IO HANDLER READER
func TestTCPClient_Send(t *testing.T) {
	client := NewTCPClient("localhost", 8081, 10)
	st := client.Send("test")
	fmt.Println(st.Code, " : ", st.Message)
}

//===================================================================================================
//CUSTOM IO HANDLER READER
func CustomHandler(conn net.Conn) (string, error) {
	return "custom okey", nil
}

func TestNewTCPClientWithCustomHandler(t *testing.T) {
	client := NewTCPClientWithCustomHandler("localhost", 8081, 10, CustomHandler)
	st := client.Send("test")
	fmt.Println(st.Code, " : ", st.Message)
}
