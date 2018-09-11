package http

import (
	"bytes"
	"context"
	"crypto/tls"
	"github.com/vlorc/lua-vm/base"
	vmnet "github.com/vlorc/lua-vm/net"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type HTTPFactory struct {
	client *http.Client
}

type Request struct {
	Method string
	Url    string
	Type   string
	Header map[string]string
	Query  map[string]string
	Cookie map[string]string
	Body   io.Reader
}

func NewHTTPFactory(driver vmnet.NetDriver) *HTTPFactory {
	return &HTTPFactory{
		client: __client(driver, &tls.Config{InsecureSkipVerify: true}),
	}
}

func __client(driver vmnet.NetDriver, config *tls.Config) *http.Client {
	return &http.Client{
		Timeout: 45 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: config,
			Dial: func(network, addr string) (net.Conn, error) {
				return driver.Dial(context.Background(), network, addr)
			},
			DialContext:           driver.Dial,
			TLSHandshakeTimeout:   10 * time.Second,
			ResponseHeaderTimeout: 10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}
}

func (f *HTTPFactory) Get(rawurl string) (*http.Response, error) {
	return f.client.Get(rawurl)
}

func (f *HTTPFactory) Post(rawurl, contentType string, body io.Reader) (*http.Response, error) {
	return f.client.Post(rawurl, contentType, body)
}

func (f *HTTPFactory) PostForm(rawurl string, values url.Values, args ...string) (*http.Response, error) {
	contentType := "application/x-www-form-urlencoded"
	if len(args) > 0 {
		contentType = args[0]
	}
	return f.client.Post(rawurl, contentType, strings.NewReader(values.Encode()))
}

func (f *HTTPFactory) Head(rawurl string) (*http.Response, error) {
	return f.client.Head(rawurl)
}

func (f *HTTPFactory) Do(r *Request) (*http.Response, error) {
	req, err := http.NewRequest(r.Method, r.Url, r.Body)
	if nil != err {
		return nil, err
	}
	if len(r.Query) > 0 {
		query := url.Values{}
		for k, v := range r.Query {
			query[k] = []string{v}
		}
		req.URL.RawQuery = query.Encode()
	}
	if len(r.Header) > 0 {
		for k, v := range r.Header {
			req.Header.Set(k, v)
		}
	}
	if len(r.Cookie) > 0 {
		cookie := &http.Cookie{}
		for k, v := range r.Cookie {
			cookie.Name = k
			cookie.Value = v
			req.AddCookie(cookie)
		}
	}

	return f.client.Do(req)
}

func (f *HTTPFactory) GetString(rawurl string) (string, error) {
	buf, err := f.GetBuffer(rawurl)
	if nil != err {
		return "", err
	}
	return buf.ToString("raw"), nil
}

func (f *HTTPFactory) GetBuffer(rawurl string) (base.Buffer, error) {
	resp, err := f.client.Get(rawurl)
	if nil != err {
		return nil, err
	}
	defer resp.Body.Close()

	buf, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		return nil, err
	}
	return base.Buffer(buf), nil
}

func (f *HTTPFactory) PostString(rawurl string, values string, args ...string) (*http.Response, error) {
	contentType := "text/plain"
	if len(args) > 0 {
		contentType = args[0]
	}
	return f.client.Post(rawurl, contentType, strings.NewReader(values))
}

func (f *HTTPFactory) PostBuffer(rawurl string, values base.Buffer, args ...string) (*http.Response, error) {
	contentType := "application/octet-stream"
	if len(args) > 0 {
		contentType = args[0]
	}
	return f.client.Post(rawurl, contentType, bytes.NewReader(values))
}
