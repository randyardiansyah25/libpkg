package tcp

import (
	"errors"
	"net"
	"strconv"
)

func BasicIOHandlerReader(conn net.Conn) (string, error) {
	buffHeader := make([]byte, 4)
	var n int
	var err error
	for {
		n, err = conn.Read(buffHeader)
		if err != nil {
			return "", err
		}

		if n == 4 {
			break
		}
	}
	sHeader := string(buffHeader[:n])
	//fmt.Println("HEADER : ", sHeader)
	nHeader, err := strconv.Atoi(sHeader)
	if err != nil {
		return "", errors.New("ATOICONVERR")
	}
	buffBody := make([]byte, nHeader)
	n = 0
	for {
		n, err = conn.Read(buffBody)
		if err != nil {
			return "", err
		}

		if n == nHeader {
			break
		}
	}
	//fmt.Println("Body : ", string(buffBody[:n]))
	return string(buffBody[:n]), nil
}
