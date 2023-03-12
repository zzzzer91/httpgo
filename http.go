package httpgo

import (
	"net/http"
	"time"
)

var (
	defaultClient = NewClient(5*time.Second, nil)
)

func Request(method, url string, body []byte, headers ...Header) (*http.Response, error) {
	return defaultClient.Request(method, url, body, headers...)
}

func RequestJSON(method, url string, data interface{}, headers ...Header) (*http.Response, error) {
	return defaultClient.RequestJSON(method, url, data, headers...)
}

func Get(url string, headers ...Header) (*http.Response, error) {
	return defaultClient.Get(url, headers...)
}

func Post(url string, body []byte, headers ...Header) (*http.Response, error) {
	return defaultClient.Post(url, body, headers...)
}

func Put(url string, body []byte, headers ...Header) (*http.Response, error) {
	return defaultClient.Put(url, body, headers...)
}

func Delete(url string, headers ...Header) (*http.Response, error) {
	return defaultClient.Delete(url, headers...)
}

func GetJSON(url string, headers ...Header) (*http.Response, error) {
	return defaultClient.GetJSON(url, headers...)
}

func PostJSON(url string, data interface{}, headers ...Header) (*http.Response, error) {
	return defaultClient.PostJSON(url, data, headers...)
}

func PutJSON(url string, data interface{}, headers ...Header) (*http.Response, error) {
	return defaultClient.PutJSON(url, data, headers...)
}

func DeleteJSON(url string, headers ...Header) (*http.Response, error) {
	return defaultClient.DeleteJSON(url, headers...)
}

func GetJsonWithAuth(url string, token string, headers ...Header) (*http.Response, error) {
	return defaultClient.GetJsonWithAuth(url, token, headers...)
}

func PostJsonWithAuth(url string, data interface{}, token string, headers ...Header) (*http.Response, error) {
	return defaultClient.PostJsonWithAuth(url, data, token, headers...)
}

func PutJsonWithAuth(url string, data interface{}, token string, headers ...Header) (*http.Response, error) {
	return defaultClient.PutJsonWithAuth(url, data, token, headers...)
}

func DeleteJsonWithAuth(url string, token string, headers ...Header) (*http.Response, error) {
	return defaultClient.DeleteJsonWithAuth(url, token, headers...)
}
