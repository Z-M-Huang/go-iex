//Package iex provides a http client to access IEX data
package iex

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

//Client IEX http client
type Client struct {
	baseURL string
	sseURL  string
	pk      string
	sk      string
	client  *http.Client
}

//NewClient new IEX http client
func NewClient(pk, sk string, sandbox bool) *Client {
	c := &Client{
		pk:     pk,
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
	req, err := http.NewRequest(http.MethodPost, o.getEndpoint("/account/metadata", params.Encode()), nil)
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

func (o *Client) doRequest(req *http.Request) (*http.Response, error) {
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
