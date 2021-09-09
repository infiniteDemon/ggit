package axios

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"strings"
)

// demon TODO: 返回请求结果和错误
func GET(url string) func() ([]byte, error) {
	var body []byte
	var err error

	c := make(chan struct{}, 1)
	go func() {
		defer close(c)

		var res *http.Response
		res, err = http.Get(url)
		if err != nil {
			return
		}

		defer res.Body.Close()
		body, err = ioutil.ReadAll(res.Body)
	}()

	return func() ([]byte, error) {
		<-c
		return body, err
	}
}

func ClientPost(url string) func() ([]byte, error) {
	var body []byte
	var err error

	c := make(chan struct{}, 1)
	go func() {
		defer close(c)

		var res *http.Response

		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}
		res, err = client.Post(url, "application/json", strings.NewReader("{\n\t\"path\": \"pages/index/index\",\n\t\"scene\": \"a=1\"\n}"))

		if err != nil {
			return
		}

		defer res.Body.Close()
		body, err = ioutil.ReadAll(res.Body)
	}()
	return func() ([]byte, error) {
		<-c
		return body, err
	}
}

func ClientGet(url string) func() ([]byte, error) {
	var body []byte
	var err error

	c := make(chan struct{}, 1)
	go func() {
		defer close(c)

		var res *http.Response

		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}
		res, err = client.Get(url)

		if err != nil {
			return
		}

		defer res.Body.Close()
		body, err = ioutil.ReadAll(res.Body)
	}()
	return func() ([]byte, error) {
		<-c
		return body, err
	}
}

func HttpGet(url string) (content string, statusCode int, err error) {
	var (
		res  *http.Response
		data []byte
	)
	if res, err = http.Get(url); err != nil {
		statusCode = -100
		return
	}
	defer res.Body.Close()
	if data, err = ioutil.ReadAll(res.Body); err != nil {
		statusCode = -200
		return
	}
	statusCode = res.StatusCode
	content = string(data)
	return

}
