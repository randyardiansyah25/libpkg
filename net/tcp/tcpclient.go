package tcp

import (
	"fmt"
	winnet "github.com/randyardiansyah25/libpkg/net"
	"net"
	"os"
	"syscall"
	"time"
)

const (
	CONNOK        = 0
	ECONNREFUSED  = 1
	ECONNABORT    = 2
	ECONNRESET    = 3
	ECONNTIMEOUT  = 4
	EOTHER        = 5
	INVALIDHEADER = 6
)

type HandlerReader func(net.Conn) (string, error)

type ConnectionState struct {
	Code    int
	Message string
}

type TCPClient struct {
	Addr    string
	Port    int
	Timeout int64
	HandlerReader
}

func (t *TCPClient) Send(message string) ConnectionState {
	to := time.Duration(t.Timeout) * time.Second
	d := net.Dialer{
		Timeout: to,
	}
	sAddr := fmt.Sprint(t.Addr, ":", t.Port)

	c, err := d.Dial("tcp", sAddr)
	if err != nil {
		return t.errType(err)
	}
	defer t.close(c)

	err = c.SetReadDeadline(time.Now().Add(to))
	if err != nil {
		return ConnectionState{EOTHER, err.Error()}
	}
	_, err = fmt.Fprint(c, message)
	if err != nil {
		return t.errType(err)
	}

	msg, err := t.HandlerReader(c)
	if err != nil {
		return t.errType(err)
	}

	return ConnectionState{CONNOK, msg}

}

func (t *TCPClient) close(conn net.Conn) {
	_ = conn.Close()
}

func (t *TCPClient) errType(err error) ConnectionState {

	if opErr, ok := err.(*net.OpError); ok {
		if sysErr, ok := opErr.Err.(*os.SyscallError); ok {
			if errno, ok := sysErr.Err.(syscall.Errno); ok {
				if errno == syscall.ECONNABORTED || errno == winnet.WSAECONNABORTED {
					return ConnectionState{ECONNABORT, "Connectioen aborted"}
				} else if errno == syscall.ECONNRESET || errno == winnet.WSAECONNRESET {
					return ConnectionState{ECONNRESET, "Connection reset by peer"}
				} else if errno == syscall.ECONNREFUSED || errno == winnet.WSAECONNREFUSED {
					return ConnectionState{ECONNREFUSED, "Connection refused"}
				} else if errno == syscall.ENETUNREACH || errno == winnet.WSAEHOSTUNREACH {
					return ConnectionState{EOTHER, "Connection unreachable"}
				} else if errno == winnet.WSAEHOSTDOWN {
					return ConnectionState{EOTHER, "Host is down"}
				} else if errno == winnet.WSAESHUTDOWN {
					return ConnectionState{EOTHER, "Cannot send after socket shutdown."}
				} else if errno == winnet.WSAETIMEDOUT {
					return ConnectionState{EOTHER, "Connection timed out."}
				} else {
					return ConnectionState{EOTHER, err.Error()}
				}
			}
		} else if opErr.Timeout() {
			return ConnectionState{ECONNTIMEOUT, fmt.Sprint("Read timeout for ", t.Timeout, " second")}
		}
	} else {
		errs := fmt.Sprint(err.Error())
		if errs == "EOF" {
			return ConnectionState{ECONNRESET, "Connection reset by peer"}
		} else if errs == "ATOICONVERR" {
			return ConnectionState{INVALIDHEADER, "Invalid header value"}
		}
	}
	return ConnectionState{EOTHER, err.Error()}
}

func NewTCPClient(addr string, port int, timeout int64) TCPClient {
	return TCPClient{
		Timeout:       timeout,
		Addr:          addr,
		Port:          port,
		HandlerReader: BasicIOHandlerReader,
	}
}

func NewTCPClientWithCustomHandler(addr string, port int, timeout int64, customReaderHandler HandlerReader) TCPClient {
	return TCPClient{
		Timeout:       timeout,
		Addr:          addr,
		Port:          port,
		HandlerReader: customReaderHandler,
	}
}
