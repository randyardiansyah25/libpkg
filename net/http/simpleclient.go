package http

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"syscall"
	"time"

	winnet "github.com/randyardiansyah25/libpkg/net"
	"github.com/randyardiansyah25/libpkg/net/http/util/urlvalues"
)

func NewSimpleClient(method string, destUrl string, timeout int64) *SimpleClient {
	return &SimpleClient{
		destUrl: destUrl,
		method:  method,
		params:  url.Values{},
		header:  map[string]string{},
		timeout: timeout,
	}
}

type SimplePostClientAuthBasic struct {
	username string
	password string
}

type SimpleClientResponse struct {
	message string
	code    int
}

func newByteFromStringResponse(statusCode int, message string) *SimpleClientResponseByte {
	msg := []byte(message)
	return &SimpleClientResponseByte{msg, statusCode}
}

func newByteResponse(statusCode int, message []byte) *SimpleClientResponseByte {
	return &SimpleClientResponseByte{message, statusCode}
}

type SimpleClientResponseByte struct {
	message []byte
	code    int
}

func (e *SimpleClientResponse) Message() string {
	return e.message
}

func (e *SimpleClientResponse) StatusCode() int {
	return e.code
}

type SimpleClient struct {
	destUrl   string
	method    string
	params    url.Values
	timeout   int64
	header    map[string]string
	authBasic SimplePostClientAuthBasic
}

func (s *SimpleClient) SetHeader(key string, value string) {
	s.header[key] = value
}

func (s *SimpleClient) SetContentType(value string) {
	s.header["Content-Type"] = value
}

func (s *SimpleClient) SetContentTypeFormUrlEncoded() {
	s.header["Content-Type"] = "application/x-www-form-urlencoded"
}

func (s *SimpleClient) SetAuthBasic(username string, password string) {
	s.authBasic = SimplePostClientAuthBasic{
		username: username,
		password: password,
	}
}

func (s *SimpleClient) SetAuthorization(value string) {
	s.SetHeader("Authorization", value)
}

func (s *SimpleClient) AddParam(key string, value string) {
	if len(s.params) == 0 {
		s.params.Set(key, value)
	} else {
		s.params.Add(key, value)
	}
}

func (s *SimpleClient) AddParams(params map[string]string) {
	for key, val := range params {
		s.params.Add(key, val)
	}
}

func (s *SimpleClient) GetParams() url.Values {
	return s.params
}

func (s *SimpleClient) DoRequest() *SimpleClientResponse {
	if s.method == "GET" {
		s.destUrl = urlvalues.AddParams(s.destUrl, s.params)
	}
	res := s.Do(bytes.NewBufferString(s.params.Encode()))
	return &SimpleClientResponse{string(res.message), res.code}
}

func (s *SimpleClient) DoRawRequest(body string) *SimpleClientResponse {
	res := s.Do(bytes.NewBuffer([]byte(body)))
	return &SimpleClientResponse{string(res.message), res.code}
}

func (s *SimpleClient) Do(body io.Reader) *SimpleClientResponseByte {
	req, err := http.NewRequest(s.method, s.destUrl, body)
	if err != nil {
		return newByteFromStringResponse(http.StatusBadGateway, err.Error())
	}

	if s.authBasic.username != "" {
		req.SetBasicAuth(s.authBasic.username, s.authBasic.password)
	}

	for key, value := range s.header {
		req.Header.Add(key, value)
	}
	to := time.Duration(s.timeout) * time.Second
	tr := &http.Transport{
		IdleConnTimeout:    to,
		DisableCompression: true,
	}

	client := &http.Client{Transport: tr, Timeout: to}

	resp, err := client.Do(req)
	if err != nil {
		switch errType := err.(type) {
		case *url.Error:
			if _, ok := errType.Err.(net.Error); ok && errType.Timeout() {
				return newByteFromStringResponse(http.StatusRequestTimeout, fmt.Sprint("Request timeout for ", to, " second..."))
			} else if opErr, ok := errType.Err.(*net.OpError); ok {
				if sysErr, ok := opErr.Err.(*os.SyscallError); ok {
					if errno, ok := sysErr.Err.(syscall.Errno); ok {
						if errno == syscall.ECONNABORTED || errno == winnet.WSAECONNABORTED {
							return newByteFromStringResponse(http.StatusNotFound, "Connection abort")
						} else if errno == syscall.ECONNRESET || errno == winnet.WSAECONNRESET {
							return newByteFromStringResponse(http.StatusBadGateway, "Connection reset by peer")
						} else if errno == syscall.ECONNREFUSED || errno == winnet.WSAECONNREFUSED {
							return newByteFromStringResponse(http.StatusNotFound, "Connection refused")
						} else if errno == syscall.ENETUNREACH || errno == winnet.WSAEHOSTUNREACH {
							return newByteFromStringResponse(http.StatusBadGateway, "Connection unreachable")
						} else if errno == winnet.WSAEHOSTDOWN {
							return newByteFromStringResponse(http.StatusBadGateway, "Host is down")
						} else if errno == winnet.WSAESHUTDOWN {
							return newByteFromStringResponse(http.StatusBadGateway, "Cannot send after socket shutdown.")
						} else if errno == winnet.WSAETIMEDOUT {
							return newByteFromStringResponse(http.StatusBadGateway, "Connection timed out.")
						} else {
							return newByteFromStringResponse(http.StatusBadGateway, err.Error())
						}
					} else {
						return newByteFromStringResponse(http.StatusBadGateway, "Closed : "+err.Error())
					}
				} else {
					return newByteFromStringResponse(http.StatusBadGateway, err.Error())
				}
			} else {
				errs := fmt.Sprint(err.(*url.Error).Err)
				if errs == "EOF" {
					return newByteFromStringResponse(http.StatusBadRequest, "Connection reset with : "+errs)
				} else {
					return newByteFromStringResponse(http.StatusBadRequest, err.Error())
				}
			}

		case net.Error:
			if errType.Timeout() {
				return newByteFromStringResponse(http.StatusRequestTimeout, fmt.Sprint("Request timeout for ", to, " second..."))
			}
		default:

		}
		//}
		return newByteFromStringResponse(http.StatusBadGateway, "else : "+err.Error())
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	bodyByte, err := io.ReadAll(resp.Body)
	if err != nil {
		return newByteResponse(http.StatusLengthRequired, bodyByte)
	}

	return newByteResponse(resp.StatusCode, bodyByte)
}
