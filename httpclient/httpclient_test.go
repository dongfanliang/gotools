package httpclient

import (
	"encoding/json"
	"testing"
)

// func TestGet(t *testing.T) {
// 	url := "https://baidu.com"
// 	client := NewHttpClient()
// 	client.SetTimeout(10)
// 	body, err := client.Do("GET", url, nil, nil)
// 	t.Log(string(body), err)
// }

func TestPost(t *testing.T) {
	url := "https://baidu.com"
	client := NewHttpClient()
	client.SetTimeout(10)
	reqBody := struct {
		A string
		B string
	}{
		"a",
		"b",
	}
	reqBodyBytes, _ := json.Marshal(reqBody)
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	body, err := client.Do("GET", url, reqBodyBytes, headers)
	t.Log(string(body), err)
}
