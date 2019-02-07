package http

import (
	"fmt"
	"github.com/randyardiansyah25/libpkg/net/http/util/urlvalues"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"time"
)

type GetClientResponse struct {
	message string
	code    int
}

func (e *GetClientResponse) Message() string {
	return e.message
}

func (e *GetClientResponse) StatusCode() int {
	return e.code
}

func newGetResponse(errCode int, errMessage string) *GetClientResponse {
	return &GetClientResponse{errMessage, errCode}
}

type GetClient struct {
	Url     string
	Timeout int64
}

func (gc *GetClient) Get(param map[string]string) (response *GetClientResponse) {
	if gc.Timeout == 0 {
		gc.Timeout = 30
	}

	to := time.Duration(time.Duration(gc.Timeout) * time.Second)
	httpClient := http.Client{
		Transport: &http.Transport{
			IdleConnTimeout: to,
		},
		Timeout: to,
	}

	newUrl := urlvalues.AddParams(gc.Url, urlvalues.ToUrlValues(param))

	resp, err := httpClient.Get(newUrl)
	if err != nil {
		switch err := err.(type) {
		case net.Error:
			if err.Timeout() {
				return newGetResponse(http.StatusRequestTimeout, fmt.Sprint("Request timeout for ", gc.Timeout, " second..."))
			}
		case *url.Error:
			if err, ok := err.Err.(net.Error); ok && err.Timeout() {
				return newGetResponse(http.StatusRequestTimeout, fmt.Sprint("Request timeout for ", gc.Timeout, " second..."))
			}
		}
		return newGetResponse(http.StatusBadGateway, err.Error())

	}
	//defer gc.closeBody(resp)
	bodyByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return newGetResponse(http.StatusLengthRequired, string(bodyByte))
	}

	return newGetResponse(resp.StatusCode, string(bodyByte))
}
