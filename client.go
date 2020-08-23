//Package iex provides a http client to access IEX data
package iex

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/Z-M-Huang/go-iex/enum/chartrange"
)

//Client IEX http client
type Client struct {
	baseURL string
	sseURL  string
	sk      string
	client  *http.Client
}

//NewClient new IEX http client
func NewClient(sk string, sandbox bool) *Client {
	c := &Client{
		sk:     sk,
		client: &http.Client{},
	}

	if sandbox {
		c.baseURL = "https://sandbox.iexapis.com/stable"
		c.sseURL = "https://sandbox-sse.iexapis.com/stable"
	} else {
		c.baseURL = "https://cloud.iexapis.com/stable"
		c.sseURL = "https://cloud-sse.iexapis.com/stable"
	}
	return c
}

//AccountMessageBudget https://iexcloud.io/docs/api/#message-budget
// API is returning 400 for some reason
// func (o *Client) AccountMessageBudget(total uint64) error {
// 	params := url.Values{}
// 	params.Add("token", o.sk)
// 	params.Add("totalMessages", fmt.Sprintf("%d", total))
// 	req, err := http.NewRequest(http.MethodPost, o.getEndpoint("/account/messagebudget", params.Encode()), nil)
// 	if err != nil {
// 		return err
// 	}
// 	err = o.getJSON(req, nil)
// 	return err
// }

//Metadata https://iexcloud.io/docs/api/#metadata
func (o *Client) Metadata() (*Metadata, error) {
	params := url.Values{}
	params.Add("token", o.sk)
	req, err := http.NewRequest(http.MethodGet, o.getEndpoint("/account/metadata", params.Encode()), nil)
	if err != nil {
		return nil, err
	}
	ret := &Metadata{}
	err = o.getJSON(req, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

//Book https://iexcloud.io/docs/api/#book
func (o *Client) Book(symbol string) (*Book, error) {
	params := url.Values{}
	params.Add("token", o.sk)
	req, err := http.NewRequest(http.MethodGet, o.getEndpoint(fmt.Sprintf("/stock/%s/book", symbol), params.Encode()), nil)
	if err != nil {
		return nil, err
	}
	ret := &Book{}
	err = o.getJSON(req, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

//HistoricalPrice https://iexcloud.io/docs/api/#historical-prices
func (o *Client) HistoricalPrice(option HistoricalOption) ([]*HistoricalPrice, error) {
	option.Range = strings.ToLower(option.Range)
	params := url.Values{}
	params.Add("token", o.sk)
	params.Add("sort", "desc")
	endpoint := fmt.Sprintf("/stock/%s/chart", option.Symbol)
	if option.Range != "" {
		endpoint += fmt.Sprintf("/%s", option.Range)
		if option.Range == chartrange.Date && option.ExactDate != "" {
			endpoint += fmt.Sprintf("/%s", option.ExactDate)
			params.Add("chartByDate", "true")
		}
	}
	if option.ChartCloseOnly {
		params.Add("chartCloseOnly", "true")
	}
	if option.ChartSimplify {
		params.Add("chartSimplify", "true")
	}
	if option.ChartInterval > 0 {
		params.Add("chartInterval", strconv.Itoa(option.ChartInterval))
	}
	if option.ChangeFromClose {
		params.Add("changeFromClose", "true")
	}
	if option.ChartLast > 0 {
		params.Add("chartLast", strconv.Itoa(option.ChartInterval))
	}
	if option.Sort != "" {
		params.Add("sort", option.Sort)
	}
	if option.includeToday {
		params.Add("includeToday", "true")
	}
	req, err := http.NewRequest(http.MethodGet, o.getEndpoint(endpoint, params.Encode()), nil)
	if err != nil {
		return nil, err
	}
	var ret []*HistoricalPrice
	err = o.getJSON(req, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

//IntradayPrice for https://iexcloud.io/docs/api/#intraday-prices
func (o *Client) IntradayPrice(option IntradayOption) ([]*IntradayPrice, error) {
	params := url.Values{}
	params.Add("token", o.sk)
	if option.ChartIEXOnly {
		params.Add("chartIEXOnly", "true")
	}
	if option.ChartReset {
		params.Add("chartReset", "true")
	}
	if option.ChartSimplify {
		params.Add("chartSimplify", "true")
	}
	if option.ChartInterval > 0 {
		params.Add("chartInterval", strconv.Itoa(option.ChartInterval))
	}
	if option.ChangeFromClose {
		params.Add("changeFromClose", "true")
	}
	if option.ChartLast > 0 {
		params.Add("chartLast", strconv.Itoa(option.ChartInterval))
	}
	if option.ExactDate != "" {
		params.Add("exactDate", option.ExactDate)
	}
	if option.ChartIEXWhenNull {
		params.Add("chartIEXWhenNull", "true")
	}
	req, err := http.NewRequest(http.MethodGet, o.getEndpoint(fmt.Sprintf("/stock/%s/intraday-prices", option.Symbol), params.Encode()), nil)
	if err != nil {
		return nil, err
	}
	var ret []*IntradayPrice
	err = o.getJSON(req, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

//DelayedQuote https://iexcloud.io/docs/api/#delayed-quote
func (o *Client) DelayedQuote(symbol string) (*DelayedQuote, error) {
	params := url.Values{}
	params.Add("token", o.sk)
	req, err := http.NewRequest(http.MethodGet, o.getEndpoint(fmt.Sprintf("/stock/%s/delayed-quote", symbol), params.Encode()), nil)
	if err != nil {
		return nil, err
	}
	ret := &DelayedQuote{}
	err = o.getJSON(req, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

//LargestTrades https://iexcloud.io/docs/api/#largest-trades
func (o *Client) LargestTrades(symbol string) ([]*LargestTrade, error) {
	params := url.Values{}
	params.Add("token", o.sk)
	req, err := http.NewRequest(http.MethodGet, o.getEndpoint(fmt.Sprintf("/stock/%s/largest-trades", symbol), params.Encode()), nil)
	if err != nil {
		return nil, err
	}
	var ret []*LargestTrade
	err = o.getJSON(req, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

//OHLC https://iexcloud.io/docs/api/#open-close-price
func (o *Client) OHLC(symbol string) (*OHLC, error) {
	params := url.Values{}
	params.Add("token", o.sk)
	req, err := http.NewRequest(http.MethodGet, o.getEndpoint(fmt.Sprintf("/stock/%s/ohlc", symbol), params.Encode()), nil)
	if err != nil {
		return nil, err
	}
	ret := &OHLC{}
	err = o.getJSON(req, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

//PreviousDayPrice https://iexcloud.io/docs/api/#previous-day-price
func (o *Client) PreviousDayPrice(symbol string) (*PreviousDayPrice, error) {
	params := url.Values{}
	params.Add("token", o.sk)
	req, err := http.NewRequest(http.MethodGet, o.getEndpoint(fmt.Sprintf("/stock/%s/previous", symbol), params.Encode()), nil)
	if err != nil {
		return nil, err
	}
	ret := &PreviousDayPrice{}
	err = o.getJSON(req, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

//PriceOnly https://iexcloud.io/docs/api/#price-only
func (o *Client) PriceOnly(symbol string) (*float64, error) {
	params := url.Values{}
	params.Add("token", o.sk)
	req, err := http.NewRequest(http.MethodGet, o.getEndpoint(fmt.Sprintf("/stock/%s/price", symbol), params.Encode()), nil)
	if err != nil {
		return nil, err
	}
	ret, err := o.getFloat64(req)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

//Quote https://iexcloud.io/docs/api/#quote
func (o *Client) Quote(symbol string, displayPercent bool) (*Quote, error) {
	params := url.Values{}
	params.Add("token", o.sk)
	if displayPercent {
		params.Add("displayPercent", "true")
	}
	req, err := http.NewRequest(http.MethodGet, o.getEndpoint(fmt.Sprintf("/stock/%s/quote", symbol), params.Encode()), nil)
	if err != nil {
		return nil, err
	}
	ret := &Quote{}
	err = o.getJSON(req, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

//VolumeByVenue https://iexcloud.io/docs/api/#volume-by-venue
func (o *Client) VolumeByVenue(symbol string) ([]*VolumeByVenue, error) {
	params := url.Values{}
	params.Add("token", o.sk)
	req, err := http.NewRequest(http.MethodGet, o.getEndpoint(fmt.Sprintf("/stock/%s/volume-by-venue", symbol), params.Encode()), nil)
	if err != nil {
		return nil, err
	}
	var ret []*VolumeByVenue
	err = o.getJSON(req, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (o *Client) getJSON(req *http.Request, out interface{}) error {
	resp, err := o.doRequest(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	jsonBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonBytes, &out)
	if err != nil {
		return err
	}
	return nil
}

func (o *Client) getFloat64(req *http.Request) (*float64, error) {
	resp, err := o.doRequest(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	floatBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	ret, err := strconv.ParseFloat(string(floatBytes), 64)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

func (o *Client) doRequest(req *http.Request) (*http.Response, error) {
	req.Header.Add("Host", "https://github.com/Z-M-Huang/go-iex")
	resp, err := o.client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		contentBytes, err := ioutil.ReadAll(resp.Body)
		msg := ""
		if err == nil {
			msg = string(contentBytes)
		}
		return nil, APIError{
			StatusCode: resp.StatusCode,
			Message:    msg,
		}
	}
	return resp, nil
}

func (o *Client) getEndpoint(path, queryString string) string {
	urlStr := fmt.Sprintf("%s%s", o.baseURL, path)
	if queryString != "" {
		urlStr += fmt.Sprintf("?%s", queryString)
	}
	return urlStr
}

func (o *Client) setTestTransport(fn roundTripFunc) {
	o.client = &http.Client{
		Transport: roundTripFunc(fn),
	}
}
