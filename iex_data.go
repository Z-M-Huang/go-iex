package iex

import (
	"fmt"
	"strconv"
	"time"
)

//HistoricalOption for https://iexcloud.io/docs/api/#historical-prices
type HistoricalOption struct {
	Symbol          string
	ChartCloseOnly  bool
	ChartSimplify   bool
	ChartInterval   int
	ChangeFromClose bool
	ChartLast       int
	Range           string
	ExactDate       string
	Sort            string
	includeToday    bool
}

//IntradayOption for https://iexcloud.io/docs/api/#intraday-prices
type IntradayOption struct {
	Symbol           string
	ChartIEXOnly     bool
	ChartReset       bool
	ChartSimplify    bool
	ChartInterval    int
	ChangeFromClose  bool
	ChartLast        int
	ExactDate        string
	ChartIEXWhenNull bool
}

//SystemEvent models a system event for a quote.
type SystemEvent struct {
	SystemEvent string    `json:"systemEvent"`
	Timestamp   EpochTime `json:"timestamp"`
}

// EpochTime refers to unix timestamps used for some fields in the API
type EpochTime time.Time

// MarshalJSON implements the Marshaler interface for EpochTime.
func (e EpochTime) MarshalJSON() ([]byte, error) {
	v := time.Time(e).UnixNano() / int64(time.Millisecond)
	if v < 0 {
		return []byte(`null`), nil
	}
	return []byte(fmt.Sprint(v)), nil
}

// UnmarshalJSON implements the Unmarshaler interface for EpochTime.
func (e *EpochTime) UnmarshalJSON(data []byte) (err error) {
	s := string(data)
	if s == "null" {
		return
	}
	ts, err := strconv.Atoi(string(data))
	if err != nil {
		return err
	}
	// Per docs: If the value is -1, IEX has not quoted the symbol in the trading day.
	if ts == -1 {
		return
	}

	*e = EpochTime(time.Unix(int64(ts)/1000, 0))
	return
}

// String implements the Stringer interface for EpochTime.
func (e EpochTime) String() string { return time.Time(e).String() }
