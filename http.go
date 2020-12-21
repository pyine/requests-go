package requests

import (
	"crypto/tls"
	"net/http"
	"time"
)

func NewRequest() *Request {
	r := &Request{
		timeout: 30,
		headers: map[string]string{},
	}
	return r
}

// Debug model
func Debug(v bool) *Request {
	r := NewRequest()
	return r.Debug(v)
}

func TLSClient(v *tls.Config) *Request {
	r := NewRequest()
	return r.SetTLSClient(v)
}

func SetTLSClient(v *tls.Config) *Request {
	r := NewRequest()
	return r.SetTLSClient(v)
}

func SetHeaders(headers map[string]string) *Request {
	r := NewRequest()
	return r.SetHeaders(headers)
}

func SetBasicAuth(username, password string) *Request {
	r := NewRequest()
	return r.SetBasicAuth(username, password)
}

func SetOAuth(token string) *Request {
	r := NewRequest()
	return r.SetOAuth(token)
}

func SetPrivateAuth(token string) *Request {
	r := NewRequest()
	return r.SetPrivateAuth(token)
}

func JSON() *Request {
	r := NewRequest()
	return r.JSON()
}

func SetTimeout(d time.Duration) *Request {
	r := NewRequest()
	return r.SetTimeout(d)
}

func Transport(v *http.Transport) *Request {
	r := NewRequest()
	return r.Transport(v)
}

// Get is a get http request
func Get(url string, data ...interface{}) (*Response, error) {
	r := NewRequest()
	return r.Get(url, data...)
}

func Post(url string, data ...interface{}) (*Response, error) {
	r := NewRequest()
	return r.Post(url, data...)
}

// Put is a put http request
func Put(url string, data ...interface{}) (*Response, error) {
	r := NewRequest()
	return r.Put(url, data...)
}

// Delete is a delete http request
func Delete(url string, data ...interface{}) (*Response, error) {
	r := NewRequest()
	return r.Delete(url, data...)
}

// Head is a head http request
func Head(url string, data ...interface{}) (*Response, error) {
	r := NewRequest()
	return r.Head(url, data...)
}

// Upload file
func Upload(url, filename, fileinput string) (*Response, error) {
	r := NewRequest()
	return r.Upload(url, filename, fileinput)
}
