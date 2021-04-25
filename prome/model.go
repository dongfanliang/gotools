package prome

import (
	"math"
	"strconv"
)

type ThanosQueryParams struct {
	Ql              string `json:"query"`
	Ts              int64  `json:"time"`
	Dedup           bool   `json:"dedup"`
	PartialResponse bool   `json:"partial_response"`
}

type ThanosQueryRangeParams struct {
	Ql              string `json:"query"`
	Start           int64  `json:"start"`
	End             int64  `json:"end"`
	Step            int    `json:"step"`
	Dedup           bool   `json:"dedup"`
	PartialResponse bool   `json:"partial_response"`
}

type PromeData struct {
	Timestamp int64   `json:"timestamp"`
	Value     float64 `json:"value"`
}

type TimeSeries struct {
	Counter string       `json:"counter"`
	Value   []*PromeData `json:"value"`
}

// query
type QueryResponse struct {
	Status   string            `json:"status"`
	ErrorStr string            `json: "error"`
	Data     QueryResponseData `json:"data"`
}

type QueryResponseData struct {
	Result     []QueryResponseTimeSeries `json:"result"`
	ResultType string                    `json:"resultType"`
}

type QueryResponseTimeSeries struct {
	Metric map[string]string `json:"metric"`
	Value  []interface{}     `json:"value"`
}

func (series QueryResponseTimeSeries) TranstoTimeSeries() *TimeSeries {
	if len(series.Value) != 2 {
		return nil
	}

	ts, ok := series.Value[0].(float64)
	if !ok {
		return nil
	}
	value, ok := series.Value[1].(string)
	if !ok {
		return nil
	}

	valueFloat, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return nil
	}

	if math.IsNaN(valueFloat) {
		return nil
	}

	promeDataSlice := []*PromeData{&PromeData{Value: valueFloat, Timestamp: int64(ts*1000) / 1000}}
	ret := &TimeSeries{
		Counter: SortedTags(series.Metric),
		Value:   promeDataSlice,
	}
	return ret
}

// query range
type QueryRangeResponse struct {
	Status   string                 `json:"status"`
	ErrorStr string                 `json: "error"`
	Data     QueryRangeResponseData `json:"data"`
}

type QueryRangeResponseData struct {
	Result     []QueryRangeResponseTimeSeries `json:"result"`
	ResultType string                         `json:"resultType"`
}

type QueryRangeResponseTimeSeries struct {
	Metric map[string]string `json:"metric"`
	Values [][]interface{}   `json:"values"`
}

func (series QueryRangeResponseTimeSeries) TranstoTimeSeries() *TimeSeries {
	if len(series.Values) == 0 {
		return nil
	}

	promeDataSlice := []*PromeData{}
	for i, _ := range series.Values {
		v := series.Values[i]
		if len(v) != 2 {
			continue
		}

		ts, ok := v[0].(float64)
		if !ok {
			continue
		}

		value, ok := v[1].(string)
		if !ok {
			continue
		}

		valueFloat, err := strconv.ParseFloat(value, 64)
		if err != nil {
			continue
		}

		if math.IsNaN(valueFloat) {
			continue
		}

		promeDataSlice = append(promeDataSlice, &PromeData{Value: valueFloat, Timestamp: int64(ts*1000) / 1000})
	}

	if len(promeDataSlice) > 0 {
		ret := &TimeSeries{
			Counter: SortedTags(series.Metric),
			Value:   promeDataSlice,
		}
		return ret
	}

	return nil
}
