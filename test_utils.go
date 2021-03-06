package iex

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

//****************Unit Testing**************************
type roundTripFunc func(req *http.Request) *http.Response

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func getRoundTripFunc(path string, statusCode int, respObj interface{}) roundTripFunc {
	return func(req *http.Request) *http.Response {
		resp := &http.Response{
			StatusCode: http.StatusNotFound,
			Body:       ioutil.NopCloser(bytes.NewBufferString("Not found")),
		}
		if strings.Contains(req.URL.Path, path) {
			resp.StatusCode = statusCode
			resp.Header = make(http.Header)
			if respObj != nil {
				if reflect.TypeOf(respObj).Kind() == reflect.Float64 {
					resp.Body = ioutil.NopCloser(bytes.NewBufferString(fmt.Sprintf("%v", respObj)))
				} else {
					b, _ := json.Marshal(respObj)
					resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
				}
			}
		}
		return resp
	}
}

func getTestData(data string, out interface{}) {
	e := json.Unmarshal([]byte(data), &out)
	if e != nil {
		panic(e)
	}
}
