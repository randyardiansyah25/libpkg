package tcp

import (
	"bufio"
	"errors"
	"io"
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
	nHeader, err := strconv.Atoi(sHeader)
	if err != nil {
		return "", errors.New("ATOICONVERR")
	}

	buffBody := make([]byte, nHeader)

	reader := bufio.NewReader(conn)
	n, err = io.ReadFull(reader, buffBody)

	if err != nil {
		return "", err
	}

	if n != nHeader {
		return "", errors.New("premature body")
	}

	return string(buffBody), nil
	// n = 0
	// for {
	// 	n, err = conn.Read(buffBody)
	// 	if err != nil {
	// 		return "", err
	// 	}

	// 	if n == nHeader {
	// 		break
	// 	}
	// }
	// return string(buffBody[:n]), nil
}
