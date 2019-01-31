package http

import (
	"bytes"
	"fmt"
	"io/ioutil"
	winnet "libpkg/net"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"syscall"
	"time"
)

type PostClientResponse struct {
	message string
	code    int
}

func (e *PostClientResponse) Message() string {
	return e.message
}

func (e *PostClientResponse) StatusCode() int {
	return e.code
}

func newPostResponse(errCode int, errMessage string) *PostClientResponse {
	return &PostClientResponse{errMessage, errCode}
}

type Params map[string]interface{}

type PostClient struct {
	Url         string
	ContentType string
	Timeout     int64
}

func (pc *PostClient) encode(p Params) string {
	var buf strings.Builder
	for key, val := range p {
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}
		buf.WriteString(url.QueryEscape(key))
		buf.WriteByte('=')
		sval := fmt.Sprint(val)
		buf.WriteString(url.QueryEscape(sval))
	}
	return buf.String()
}

func (pc *PostClient) doPost(body *bytes.Buffer) (response *PostClientResponse) {
	if pc.Timeout == 0 {
		pc.Timeout = 30
	}

	to := time.Duration(time.Duration(pc.Timeout) * time.Second)
	httpClient := http.Client{
		Transport: &http.Transport{
			IdleConnTimeout: to,
		},
		Timeout: to,
	}

	resp, err := httpClient.Post(pc.Url, pc.ContentType, body)
	if err != nil {
		switch errType := err.(type) {
		case *url.Error:
			if _, ok := errType.Err.(net.Error); ok && errType.Timeout() {
				return newPostResponse(http.StatusRequestTimeout, fmt.Sprint("Request timeout for ", pc.Timeout, " second..."))
			} else if opErr, ok := errType.Err.(*net.OpError); ok {
				if sysErr, ok := opErr.Err.(*os.SyscallError); ok {
					if errno, ok := sysErr.Err.(syscall.Errno); ok {
						if errno == syscall.ECONNABORTED || errno == winnet.WSAECONNABORTED {
							return newPostResponse(http.StatusNotFound, "Connection abort")
						} else if errno == syscall.ECONNRESET || errno == winnet.WSAECONNRESET {
							return newPostResponse(http.StatusBadGateway, "Connection reset by peer")
						} else if errno == syscall.ECONNREFUSED || errno == winnet.WSAECONNREFUSED {
							return newPostResponse(http.StatusNotFound, "Connection refused")
						} else if errno == syscall.ENETUNREACH || errno == winnet.WSAEHOSTUNREACH {
							return newPostResponse(http.StatusBadGateway, "Connection unreachable")
						} else if errno == winnet.WSAEHOSTDOWN {
							return newPostResponse(http.StatusBadGateway, "Host is down")
						} else if errno == winnet.WSAESHUTDOWN {
							return newPostResponse(http.StatusBadGateway, "Cannot send after socket shutdown.")
						} else if errno == winnet.WSAETIMEDOUT {
							return newPostResponse(http.StatusBadGateway, "Connection timed out.")
						} else {
							return newPostResponse(http.StatusBadGateway, err.Error())
						}
					} else {
						return newPostResponse(http.StatusBadGateway, "Closed : "+err.Error())
					}
				} else {
					return newPostResponse(http.StatusBadGateway, err.Error())
				}
			} else {
				errs := fmt.Sprint(err.(*url.Error).Err)
				if errs == "EOF" {
					return newPostResponse(http.StatusBadRequest, "Connection reset with : "+errs)
				} else {
					return newPostResponse(http.StatusBadRequest, err.Error())
				}
			}

		case net.Error:
			if errType.Timeout() {
				return newPostResponse(http.StatusRequestTimeout, fmt.Sprint("Request timeout for ", pc.Timeout, " second..."))
			}
		default:

		}
		//}
		return newPostResponse(http.StatusBadGateway, "else : "+err.Error())

	}
	defer pc.closeBody(resp)
	bodyByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return newPostResponse(http.StatusLengthRequired, string(bodyByte))
	}

	return newPostResponse(resp.StatusCode, string(bodyByte))
}

func (pc *PostClient) closeBody(response *http.Response) {
	_ = response.Body.Close()
}

func (pc *PostClient) PostJson(json []byte) *PostClientResponse {
	return pc.doPost(bytes.NewBuffer(json))
}

func (pc *PostClient) PostParam(param map[string]interface{}) *PostClientResponse {
	sparam := pc.encode(param)
	return pc.doPost(bytes.NewBufferString(sparam))
}
