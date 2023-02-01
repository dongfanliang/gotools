package httpclient

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"net/http"
	"sync"
	"time"
)

type HttpClient struct {
	client *http.Client
	pool   sync.Pool
}

func NewHttpClient() *HttpClient {
	c := &HttpClient{
		pool: sync.Pool{
			New: func() interface{} {
				return bytes.NewBuffer(make([]byte, 1024))
			},
		},
		client: &http.Client{
			Timeout: time.Duration(5) * time.Second, // 默认超时时间5s
			Transport: &http.Transport{
				DialContext: (&net.Dialer{
					Timeout:   time.Duration(3) * time.Second,  // 建链超时3s
					KeepAlive: time.Duration(60) * time.Second, // 连接保持
				}).DialContext,
				MaxIdleConns:        100, // 最大空闲链接数
				MaxIdleConnsPerHost: 100, // 单host最大空闲链接数
				IdleConnTimeout:     time.Duration(60) * time.Second,
			},
		},
	}
	return c
}

func (c *HttpClient) SetTimeout(timeout int) {
	c.client.Timeout = time.Duration(timeout) * time.Second
}

func (c *HttpClient) SetTransport(transport *http.Transport) {
	c.client.Transport = transport
}

func (c *HttpClient) Do(method, url string, body []byte, headers map[string]string) ([]byte, error) {
	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return []byte{}, err
	}

	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 || resp.StatusCode < 200 {
		return nil, fmt.Errorf("request fail: StatusCode=%d", resp.StatusCode)
	}

	buf := c.pool.Get().(*bytes.Buffer)
	buf.Reset()
	defer c.pool.Put(buf)

	_, err = io.Copy(buf, resp.Body)
	if err != nil {
		return []byte{}, err
	}
	return buf.Bytes(), nil
}
