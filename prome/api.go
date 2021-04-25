package prome

import (
	"fmt"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/parnurzeal/gorequest"
)

var jsoner = jsoniter.ConfigCompatibleWithStandardLibrary

const (
	TIMEOUT = 10
)

type ThanosClient struct {
	BaseUrl string
	Timeout int
}

func NewThanosClient(url string, timeout int) ThanosClient {
	return ThanosClient{BaseUrl: url, Timeout: timeout}
}

func (p *ThanosClient) Query(t ThanosQueryParams) ([]*TimeSeries, error) {
	retSlice := []*TimeSeries{}
	timeout := TIMEOUT
	if p.Timeout > 0 {
		timeout = p.Timeout
	}
	timeoutDuration := time.Duration(timeout) * time.Second

	v := QueryResponse{}
	_, body, errs := gorequest.New().Get(p.BaseUrl + "/api/v1/query").Query(t).Timeout(timeoutDuration).EndBytes()
	if errs != nil {
		return retSlice, fmt.Errorf("%v", errs)
	}

	err := jsoner.Unmarshal(body, &v)
	if err != nil {
		return retSlice, err
	}

	if v.Status != "success" {
		return retSlice, fmt.Errorf("%s", v.ErrorStr)
	}

	items := v.Data.Result
	for i := 0; i < len(items); i++ {
		item := items[i].TranstoTimeSeries()
		if item != nil {
			retSlice = append(retSlice, items[i].TranstoTimeSeries())
		}
	}

	return retSlice, nil
}

func (p *ThanosClient) QueryRange(t ThanosQueryRangeParams) ([]*TimeSeries, error) {
	retSlice := []*TimeSeries{}
	timeout := TIMEOUT
	if p.Timeout > 0 {
		timeout = p.Timeout
	}
	timeoutDuration := time.Duration(timeout) * time.Second

	v := QueryRangeResponse{}
	_, body, errs := gorequest.New().Post(p.BaseUrl + "/api/v1/query_range").Send(t).Type("multipart").Timeout(timeoutDuration).EndBytes()
	if errs != nil {
		return retSlice, fmt.Errorf("%v", errs)
	}

	err := jsoner.Unmarshal(body, &v)
	if err != nil {
		return retSlice, err
	}

	if v.Status != "success" {
		return retSlice, fmt.Errorf("%s", v.ErrorStr)
	}

	items := v.Data.Result
	for i := 0; i < len(items); i++ {
		item := items[i].TranstoTimeSeries()
		if item != nil {
			retSlice = append(retSlice, items[i].TranstoTimeSeries())
		}
	}

	return retSlice, nil
}
