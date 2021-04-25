package thanos

import (
	"testing"
	"time"
)

const (
	URL          = "https://thanos-query.com"
	TIMEOUT_TEST = 20
)

func Test_Query(t *testing.T) {
	ts := time.Now().Unix()
	param := ThanosQueryParams{
		Ql: `sum(metric{})`,
		Ts: ts - 120,
	}
	c := NewThanosClient(URL, TIMEOUT)
	data, err := c.Query(param)
	t.Log(err)
	for i := 0; i < len(data); i++ {
		t.Logf("%v", data[i].Label)
		for _, item := range data[i].Value {
			item := item
			t.Logf("%v", item)
		}
	}
}

func Test_QueryRange(t *testing.T) {
	ts := time.Now().Unix()
	param := ThanosQueryRangeParams{
		Ql:    `sum(metric{})`,
		Start: ts - 600,
		End:   ts - 580,
		Step:  10,
	}
	c := NewThanosClient(URL, TIMEOUT_TEST)
	data, err := c.QueryRange(param)
	t.Log(err)
	for i := 0; i < len(data); i++ {
		t.Logf("%v", data[i].Label)
		for _, item := range data[i].Value {
			item := item
			t.Logf("%v", item)
		}
	}
}
