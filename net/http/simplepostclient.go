package http

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
)

func NewSimplePostClient(method string, destUrl string, timeout int64) (*SimplePostClient){
	return &SimplePostClient{
		destUrl:destUrl,
		method:method,
		params:url.Values{},
		header: map[string]string{},
	}
}

type SimplePostClientAuthBasic struct {
	username string
	password string
}

type SimplePostClient struct {
	destUrl string
	method string
	request *http.Request
	client *http.Client
	params url.Values
	timeout int64
	header map[string]string
	authBasic SimplePostClientAuthBasic
}

func (s *SimplePostClient) SetHeader(key string, value string){
	s.header[key] = value
}

func (s *SimplePostClient) SetAuthBasic(username string, password string){
	s.authBasic = SimplePostClientAuthBasic{
		username: username,
		password: password,
	}
}

func (s *SimplePostClient) SetAuthorization(value string){
	s.SetHeader("Authorization", value)
}

func (s *SimplePostClient) AddParam(key string, value string){
	if len(s.params) == 0 {
		s.params.Set(key, value)
	}else{
		s.params.Add(key, value)
	}
}

func (s *SimplePostClient) DoRequest()(*http.Response, error){
	return s.do(bytes.NewBufferString(s.params.Encode()))
}

func (s *SimplePostClient) DoRawRequest(body string)(*http.Response, error){
	return s.do(bytes.NewBuffer([]byte(body)))
}

func (s *SimplePostClient) do(body io.Reader) (*http.Response, error) {
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

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, err
}
