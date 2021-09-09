package axios

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

/**
 * @Author demon
 * @Description //TODO post请求封装
 * @Date 2020-7-15 20:53:26
 **/

func POST(url string, contentType string, data map[string]interface{}) (body []byte, err error) {
	bytesData, _ := json.Marshal(data)
	resp, err := http.Post(url,
		contentType,
		bytes.NewReader(bytesData))
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		return body, err
	}
	return body, err
}
