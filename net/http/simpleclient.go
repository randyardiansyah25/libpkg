package http

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"time"
)

func NewSimpleClient(method string, destUrl string, timeout int64) *SimpleClient {
	return &SimpleClient{
		destUrl: destUrl,
		method:  method,
		params:  url.Values{},
		header:  map[string]string{},
	}
}

type SimplePostClientAuthBasic struct {
	username string
	password string
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

func (s *SimpleClient) DoRequest() (*http.Response, error) {
	return s.do(bytes.NewBufferString(s.params.Encode()))
}

func (s *SimpleClient) DoRawRequest(body string) (*http.Response, error) {
	return s.do(bytes.NewBuffer([]byte(body)))
}

func (s *SimpleClient) do(body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(s.method, s.destUrl, body)
	if err != nil {
		return nil, err
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
		return nil, err
	}

	return resp, err
}
