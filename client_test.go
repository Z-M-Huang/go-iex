//Package iex provides a http client to access IEX data
package iex

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	"github.com/Z-M-Huang/go-iex/enum/chartrange"
)

func TestNewClient(t *testing.T) {
	type args struct {
		sk      string
		sandbox bool
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Sandbox",
			args: args{
				sk:      "",
				sandbox: true,
			},
		},
		{
			name: "Live",
			args: args{
				sk:      "",
				sandbox: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClient(tt.args.sk, tt.args.sandbox); got == nil {
				t.Errorf("NewClient() is nil for %s", tt.name)
			}
		})
	}
}

func TestClient_getJSON(t *testing.T) {
	req1, _ := http.NewRequest(http.MethodGet, "https://example.com/req1", nil)

	type args struct {
		req *http.Request
		out interface{}
	}
	tests := []struct {
		name      string
		o         *Client
		args      args
		roundTrip roundTripFunc
		wantErr   bool
	}{
		{
			name: "Failed to unmarshal",
			o:    NewClient("", true),
			args: args{
				req: req1,
				out: nil,
			},
			roundTrip: func(req *http.Request) *http.Response {
				return &http.Response{
					StatusCode: http.StatusOK,
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		if tt.roundTrip != nil {
			tt.o.setTestTransport(tt.roundTrip)
		}
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.o.getJSON(tt.args.req, tt.args.out); (err != nil) != tt.wantErr {
				t.Errorf("Client.getJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_getFloat64(t *testing.T) {
	req1, _ := http.NewRequest(http.MethodGet, "https://example.com/req1", nil)

	type args struct {
		req *http.Request
	}
	tests := []struct {
		name      string
		o         *Client
		args      args
		want      *float64
		roundTrip roundTripFunc
		wantErr   bool
	}{
		{
			name: "Failed to parse",
			o:    NewClient("", true),
			args: args{
				req: req1,
			},
			roundTrip: func(req *http.Request) *http.Response {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewBufferString("abcd")),
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		if tt.roundTrip != nil {
			tt.o.setTestTransport(tt.roundTrip)
		}
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.o.getFloat64(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.getFloat64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.getFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func TestClient_AccountMessageBudget(t *testing.T) {
// 	type args struct {
// 		total uint64
// 	}
// 	tests := []struct {
// 		name    string
// 		o       *Client
// 		args    args
// 		wantErr bool
// 	}{
// 		{
// 			name: "Success",
// 			o:    NewClient("", true),
// 			args: args{
// 				total: 1000000000,
// 			},
// 			wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if err := tt.o.AccountMessageBudget(tt.args.total); (err != nil) != tt.wantErr {
// 				t.Errorf("Client.AccountMessageBudget() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

func TestClient_Metadata(t *testing.T) {
	d := &Metadata{}
	getTestData(`{"payAsYouGoEnabled":false,"effectiveDate":1576643807066,"subscriptionTermType":"lnunaa","tierName":"_atlsdrot","messageLimit":504798,"messagesUsed":0,"circuitBreaker":null}`, &d)
	tests := []struct {
		name      string
		o         *Client
		want      *Metadata
		roundTrip roundTripFunc
		wantErr   bool
	}{
		{
			name:      "Success",
			o:         NewClient("", true),
			want:      d,
			roundTrip: getRoundTripFunc("/account/metadata", http.StatusOK, *d),
			wantErr:   false,
		},
		{
			name: "Failed to create request",
			o: &Client{
				baseURL: "://",
			},
			want:      nil,
			roundTrip: nil,
			wantErr:   true,
		},
		{
			name:      "Request Failed",
			o:         NewClient("", true),
			want:      nil,
			roundTrip: getRoundTripFunc("/account/metadata", http.StatusBadRequest, APIError{StatusCode: http.StatusBadRequest}),
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.roundTrip != nil {
				tt.o.setTestTransport(tt.roundTrip)
			}
			got, err := tt.o.Metadata()
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Metadata() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.Metadata() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Book(t *testing.T) {
	d := &Book{}
	getTestData(`{"quote":{"symbol":"AAPL","companyName":"Apple, Inc.","primaryExchange":"SDANQA","calculationPrice":"close","open":494.25,"openTime":1675436555197,"openSource":"lfcfiiao","close":501.75,"closeTime":1650784639416,"closeSource":"lcoiffia","high":514.609,"highTime":1604853423081,"highSource":"iditep am1n5le ue dcrey","low":496,"lowTime":1616492925482,"lowSource":"n  mieu taree51eclipddy","latestPrice":505.43,"latestSource":"Close","latestTime":"August 21, 2020","latestUpdate":1623827950265,"latestVolume":85834768,"iexRealtimePrice":null,"iexRealtimeSize":null,"iexLastUpdated":null,"delayedPrice":512.5,"delayedPriceTime":1607680185943,"oddLotDelayedPrice":323.89,"oddLotDelayedPriceTime":1676708949197,"extendedPrice":512.16,"extendedChange":1.51,"extendedChangePercent":0.00297,"extendedPriceTime":1656469083963,"previousClose":491.2,"previousVolume":32544056,"change":25.34,"changePercent":0.05257,"volume":88306106,"iexMarketPercent":null,"iexVolume":null,"avgTotalVolume":42425326,"iexBidPrice":null,"iexBidSize":null,"iexAskPrice":null,"iexAskSize":null,"iexOpen":null,"iexOpenTime":null,"iexClose":499.92,"iexCloseTime":1626501157809,"marketCap":2149873865922,"peRatio":38.52,"week52High":519.224,"week52Low":207,"ytdChange":0.75003,"lastTradeTime":1670822875005,"isUSMarketOpen":false},"bids":[],"asks":[],"systemEvent":{}}`, &d)
	type args struct {
		symbol string
	}
	tests := []struct {
		name      string
		o         *Client
		args      args
		want      *Book
		roundTrip roundTripFunc
		wantErr   bool
	}{
		{
			name: "Success",
			o:    NewClient("", true),
			args: args{
				symbol: "AAPL",
			},
			want:      d,
			roundTrip: getRoundTripFunc("/stock/AAPL/book", http.StatusOK, *d),
			wantErr:   false,
		},
		{
			name: "Failed to create request",
			o: &Client{
				baseURL: "://",
			},
			args: args{
				symbol: "AAPL",
			},
			want:      nil,
			roundTrip: nil,
			wantErr:   true,
		},
		{
			name: "Request Failed",
			o:    NewClient("", true),
			args: args{
				symbol: "AAPL",
			},
			want:      nil,
			roundTrip: getRoundTripFunc("/stock/AAPL/book", http.StatusBadRequest, APIError{StatusCode: http.StatusBadRequest}),
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		if tt.roundTrip != nil {
			tt.o.setTestTransport(tt.roundTrip)
		}
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.o.Book(tt.args.symbol)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Book() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.Book() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_HistoricalPrice(t *testing.T) {
	var d []*HistoricalPrice
	getTestData(`[{"date":"2020-08-17","close":470.7,"volume":30671436,"change":0,"changePercent":0,"changeOverTime":0}]`, &d)
	type args struct {
		option HistoricalOption
	}
	tests := []struct {
		name      string
		o         *Client
		args      args
		want      []*HistoricalPrice
		roundTrip roundTripFunc
		wantErr   bool
	}{
		{
			name: "Case 1",
			o:    NewClient("", true),
			args: args{
				option: HistoricalOption{
					Symbol:          "AAPL",
					ChartCloseOnly:  true,
					ChartInterval:   5,
					ChangeFromClose: true,
					ChartSimplify:   true,
					ChartLast:       1,
					Sort:            "asc",
					includeToday:    true,
				},
			},
			want:      d,
			roundTrip: getRoundTripFunc("/stock/AAPL/chart", http.StatusOK, d),
			wantErr:   false,
		},
		{
			name: "By Date",
			o:    NewClient("", true),
			args: args{
				option: HistoricalOption{
					Symbol:    "AAPL",
					Range:     chartrange.Date,
					ExactDate: "20200817",
				},
			},
			want:      d,
			roundTrip: getRoundTripFunc("/stock/AAPL/chart", http.StatusOK, d),
			wantErr:   false,
		},
		{
			name: "Failed to create request",
			o: &Client{
				baseURL: "://",
			},
			args: args{
				option: HistoricalOption{
					Symbol:    "AAPL",
					Range:     chartrange.Date,
					ExactDate: "20200817",
				},
			},
			want:      nil,
			roundTrip: nil,
			wantErr:   true,
		},
		{
			name: "Request Failed",
			o:    NewClient("", true),
			args: args{
				option: HistoricalOption{
					Symbol:    "AAPL",
					Range:     chartrange.Date,
					ExactDate: "20200817",
				},
			},
			want:      nil,
			roundTrip: getRoundTripFunc("/stock/AAPL/chart", http.StatusBadRequest, APIError{StatusCode: http.StatusBadRequest}),
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		if tt.roundTrip != nil {
			tt.o.setTestTransport(tt.roundTrip)
		}
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.o.HistoricalPrice(tt.args.option)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.HistoricalPrice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.HistoricalPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_IntradayPrice(t *testing.T) {
	var d []*IntradayPrice
	getTestData(`[{"date":"2020-08-21","minute":"15:59","label":"3:59 PM","high":522.43,"low":508.56,"open":515.4,"close":514.5,"average":506.195,"volume":13994,"notional":7096203.2,"numberOfTrades":139}]`, &d)
	type args struct {
		option IntradayOption
	}
	tests := []struct {
		name      string
		o         *Client
		args      args
		want      []*IntradayPrice
		roundTrip roundTripFunc
		wantErr   bool
	}{
		{
			name: "Case 1",
			o:    NewClient("", true),
			args: args{
				option: IntradayOption{
					Symbol:           "AAPL",
					ChartIEXOnly:     true,
					ChartReset:       true,
					ChartInterval:    1,
					ChangeFromClose:  true,
					ChartSimplify:    true,
					ChartLast:        1,
					ExactDate:        "20200821",
					ChartIEXWhenNull: true,
				},
			},
			want:      d,
			roundTrip: getRoundTripFunc("/stock/AAPL/intraday-prices", http.StatusOK, d),
			wantErr:   false,
		},
		{
			name: "Failed to create request",
			o: &Client{
				baseURL: "://",
			},
			args: args{
				option: IntradayOption{
					Symbol:    "AAPL",
					ExactDate: "20200817",
				},
			},
			want:      nil,
			roundTrip: nil,
			wantErr:   true,
		},
		{
			name: "Request Failed",
			o:    NewClient("", true),
			args: args{
				option: IntradayOption{
					Symbol:    "AAPL",
					ExactDate: "20200817",
				},
			},
			want:      nil,
			roundTrip: getRoundTripFunc("/stock/AAPL/intraday-prices", http.StatusBadRequest, APIError{StatusCode: http.StatusBadRequest}),
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		if tt.roundTrip != nil {
			tt.o.setTestTransport(tt.roundTrip)
		}
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.o.IntradayPrice(tt.args.option)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.IntradayPrice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.IntradayPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_DelayedQuote(t *testing.T) {
	d := &DelayedQuote{}
	getTestData(`{"symbol": "AAPL","delayedPrice": 143.08,"delayedSize": 200,"delayedPriceTime": 1498762739791,"high": 143.90,"low": 142.26,"totalVolume": 33547893,"processedTime": 1498763640156}`, &d)
	type args struct {
		symbol string
	}
	tests := []struct {
		name      string
		o         *Client
		args      args
		want      *DelayedQuote
		roundTrip roundTripFunc
		wantErr   bool
	}{
		{
			name: "Success",
			o:    NewClient("", true),
			args: args{
				symbol: "AAPL",
			},
			want:      d,
			roundTrip: getRoundTripFunc("/stock/AAPL/delayed-quote", http.StatusOK, *d),
			wantErr:   false,
		},
		{
			name: "Failed to create request",
			o: &Client{
				baseURL: "://",
			},
			args: args{
				symbol: "AAPL",
			},
			want:      nil,
			roundTrip: nil,
			wantErr:   true,
		},
		{
			name: "Request Failed",
			o:    NewClient("", true),
			args: args{
				symbol: "AAPL",
			},
			want:      nil,
			roundTrip: getRoundTripFunc("/stock/AAPL/delayed-quote", http.StatusBadRequest, APIError{StatusCode: http.StatusBadRequest}),
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		if tt.roundTrip != nil {
			tt.o.setTestTransport(tt.roundTrip)
		}
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.o.DelayedQuote(tt.args.symbol)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.DelayedQuote() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.DelayedQuote() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_LargestTrades(t *testing.T) {
	var d []*LargestTrade
	getTestData(`[{"price":186.39,"size":10000,"time":1527090690175,"timeLabel":"11:51:30","venue":"EDGX","venueName":"CboeEDGX"}]`, &d)
	type args struct {
		symbol string
	}
	tests := []struct {
		name      string
		o         *Client
		args      args
		want      []*LargestTrade
		roundTrip roundTripFunc
		wantErr   bool
	}{
		{
			name: "Success",
			o:    NewClient("", true),
			args: args{
				symbol: "AAPL",
			},
			want:      d,
			roundTrip: getRoundTripFunc("/stock/AAPL/largest-trades", http.StatusOK, d),
			wantErr:   false,
		},
		{
			name: "Failed to create request",
			o: &Client{
				baseURL: "://",
			},
			args: args{
				symbol: "AAPL",
			},
			want:      nil,
			roundTrip: nil,
			wantErr:   true,
		},
		{
			name: "Request Failed",
			o:    NewClient("", true),
			args: args{
				symbol: "AAPL",
			},
			want:      nil,
			roundTrip: getRoundTripFunc("/stock/AAPL/largest-trades", http.StatusBadRequest, APIError{StatusCode: http.StatusBadRequest}),
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		if tt.roundTrip != nil {
			tt.o.setTestTransport(tt.roundTrip)
		}
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.o.LargestTrades(tt.args.symbol)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.LargestTrades() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.LargestTrades() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_OHLC(t *testing.T) {
	d := &OHLC{}
	getTestData(`{"open":{"price":154,"time":1506605400394},"close":{"price":153.28,"time":1506605400394},"high":154.80,"low":153.25}`, &d)
	type args struct {
		symbol string
	}
	tests := []struct {
		name      string
		o         *Client
		args      args
		want      *OHLC
		roundTrip roundTripFunc
		wantErr   bool
	}{
		{
			name: "Success",
			o:    NewClient("", true),
			args: args{
				symbol: "AAPL",
			},
			want:      d,
			roundTrip: getRoundTripFunc("/stock/AAPL/ohlc", http.StatusOK, d),
			wantErr:   false,
		},
		{
			name: "Failed to create request",
			o: &Client{
				baseURL: "://",
			},
			args: args{
				symbol: "AAPL",
			},
			want:      nil,
			roundTrip: nil,
			wantErr:   true,
		},
		{
			name: "Request Failed",
			o:    NewClient("", true),
			args: args{
				symbol: "AAPL",
			},
			want:      nil,
			roundTrip: getRoundTripFunc("/stock/AAPL/ohlc", http.StatusBadRequest, APIError{StatusCode: http.StatusBadRequest}),
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		if tt.roundTrip != nil {
			tt.o.setTestTransport(tt.roundTrip)
		}
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.o.OHLC(tt.args.symbol)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.OHLC() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.OHLC() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_PreviousDayPrice(t *testing.T) {
	d := &PreviousDayPrice{}
	getTestData(`{"date":"2019-03-25","open":191.51,"close":188.74,"high":191.98,"low":186.6,"volume":43845293,"uOpen":191.51,"uClose":188.74,"uHigh":191.98,"uLow":186.6,"uVolume":43845293,"change":0,"changePercent":0,"changeOverTime":0,"symbol":"AAPL"}`, &d)
	type args struct {
		symbol string
	}
	tests := []struct {
		name      string
		o         *Client
		args      args
		want      *PreviousDayPrice
		roundTrip roundTripFunc
		wantErr   bool
	}{
		{
			name: "Success",
			o:    NewClient("", true),
			args: args{
				symbol: "AAPL",
			},
			want:      d,
			roundTrip: getRoundTripFunc("/stock/AAPL/previous", http.StatusOK, d),
			wantErr:   false,
		},
		{
			name: "Failed to create request",
			o: &Client{
				baseURL: "://",
			},
			args: args{
				symbol: "AAPL",
			},
			want:      nil,
			roundTrip: nil,
			wantErr:   true,
		},
		{
			name: "Request Failed",
			o:    NewClient("", true),
			args: args{
				symbol: "AAPL",
			},
			want:      nil,
			roundTrip: getRoundTripFunc("/stock/AAPL/previous", http.StatusBadRequest, APIError{StatusCode: http.StatusBadRequest}),
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		if tt.roundTrip != nil {
			tt.o.setTestTransport(tt.roundTrip)
		}
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.o.PreviousDayPrice(tt.args.symbol)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.PreviousDayPrice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.PreviousDayPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_PriceOnly(t *testing.T) {
	d := float64(99)
	type args struct {
		symbol string
	}
	tests := []struct {
		name      string
		o         *Client
		args      args
		want      *float64
		roundTrip roundTripFunc
		wantErr   bool
	}{
		{
			name: "Success",
			o:    NewClient("", true),
			args: args{
				symbol: "AAPL",
			},
			want:      &d,
			roundTrip: getRoundTripFunc("/stock/AAPL/price", http.StatusOK, d),
			wantErr:   false,
		},
		{
			name: "Failed to create request",
			o: &Client{
				baseURL: "://",
			},
			args: args{
				symbol: "AAPL",
			},
			want:      nil,
			roundTrip: nil,
			wantErr:   true,
		},
		{
			name: "Request Failed",
			o:    NewClient("", true),
			args: args{
				symbol: "AAPL",
			},
			want:      nil,
			roundTrip: getRoundTripFunc("/stock/AAPL/price", http.StatusBadRequest, APIError{StatusCode: http.StatusBadRequest}),
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		if tt.roundTrip != nil {
			tt.o.setTestTransport(tt.roundTrip)
		}
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.o.PriceOnly(tt.args.symbol)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.PriceOnly() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.PriceOnly() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Quote(t *testing.T) {
	d := &Quote{}
	getTestData(`{"symbol":"AAPL","companyName":"Apple, Inc.","primaryExchange":"SDQNAA","calculationPrice":"close","open":489.64,"openTime":1657681925689,"openSource":"ifciofla","close":501.31,"closeTime":1653717302906,"closeSource":"iocffial","high":517.662,"highTime":1640950868234,"highSource":"tunpcier l day 1ieemed5","low":487,"lowTime":1673933386320,"lowSource":" re udi l5cinmepdye1aet","latestPrice":504.37,"latestSource":"Close","latestTime":"August 21, 2020","latestUpdate":1609837994673,"latestVolume":84755054,"iexRealtimePrice":null,"iexRealtimeSize":null,"iexLastUpdated":null,"delayedPrice":522.26,"delayedPriceTime":1626341585820,"oddLotDelayedPrice":332.14,"oddLotDelayedPriceTime":1645450107090,"extendedPrice":508.03,"extendedChange":1.51,"extendedChangePercent":0.00299,"extendedPriceTime":1636879893680,"previousClose":496.1,"previousVolume":32909268,"change":25.14,"changePercent":0.05166,"volume":87395983,"iexMarketPercent":null,"iexVolume":null,"avgTotalVolume":43557810,"iexBidPrice":null,"iexBidSize":null,"iexAskPrice":null,"iexAskSize":null,"iexOpen":null,"iexOpenTime":null,"iexClose":503.29,"iexCloseTime":1609280406961,"marketCap":2167167444241,"peRatio":38.21,"week52High":521.511,"week52Low":209,"ytdChange":0.7322,"lastTradeTime":1604632778851,"isUSMarketOpen":false}`, &d)
	type args struct {
		symbol         string
		displayPercent bool
	}
	tests := []struct {
		name      string
		o         *Client
		args      args
		want      *Quote
		roundTrip roundTripFunc
		wantErr   bool
	}{
		{
			name: "Success",
			o:    NewClient("", true),
			args: args{
				symbol:         "AAPL",
				displayPercent: true,
			},
			want:      d,
			roundTrip: getRoundTripFunc("/stock/AAPL/quote", http.StatusOK, d),
			wantErr:   false,
		},
		{
			name: "Failed to create request",
			o: &Client{
				baseURL: "://",
			},
			args: args{
				symbol: "AAPL",
			},
			want:      nil,
			roundTrip: nil,
			wantErr:   true,
		},
		{
			name: "Request Failed",
			o:    NewClient("", true),
			args: args{
				symbol:         "AAPL",
				displayPercent: false,
			},
			want:      nil,
			roundTrip: getRoundTripFunc("/stock/AAPL/quote", http.StatusBadRequest, APIError{StatusCode: http.StatusBadRequest}),
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		if tt.roundTrip != nil {
			tt.o.setTestTransport(tt.roundTrip)
		}
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.o.Quote(tt.args.symbol, tt.args.displayPercent)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Quote() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.Quote() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_VolumeByVenue(t *testing.T) {
	var d []*VolumeByVenue
	getTestData(`[{"volume":128534,"venue":"IXCS","venueName":"an oNEaNSYtli","date":"2020-08-21","marketPercent":0.00158951808219342,"avgMarketPercent":0.0017868032595169216},{"volume":141831,"venue":"AXSE","venueName":"anmSrA icYeEN","date":"2020-08-21","marketPercent":0.0016666903224351562,"avgMarketPercent":0.0009387377400720757},{"volume":292299,"venue":"XOSB","venueName":"d XNBasqa","date":"2020-08-21","marketPercent":0.003458534834747447,"avgMarketPercent":0.0055433584354259964},{"volume":284958,"venue":"TBAY","venueName":"Xe oCbBY","date":"2020-08-21","marketPercent":0.003460473702103633,"avgMarketPercent":0.003530024569310878},{"volume":391067,"venue":"XHPL","venueName":"qNXSP daas","date":"2020-08-21","marketPercent":0.004621416014785481,"avgMarketPercent":0.002193899097686544},{"volume":597489,"venue":"NXSY","venueName":"YNSE","date":"2020-08-21","marketPercent":0.007352907711260765,"avgMarketPercent":0.009256163746696702},{"volume":805126,"venue":"IGXE","venueName":"XIE","date":"2020-08-21","marketPercent":0.009777167434436377,"avgMarketPercent":0.009454111198784783},{"volume":809159,"venue":"XCHI","venueName":"HXC","date":"2020-08-21","marketPercent":0.010442979012030854,"avgMarketPercent":0.0069758385257075},{"volume":1447001,"venue":"AEGD","venueName":"b GoCeEDA","date":"2020-08-21","marketPercent":0.01700497801084891,"avgMarketPercent":0.010939803953861063},{"volume":5036036,"venue":"GDEX","venueName":"EeXDoGbC ","date":"2020-08-21","marketPercent":0.05878454439952958,"avgMarketPercent":0.04871023306885448},{"volume":5884252,"venue":"BATS","venueName":"BbXeCZ o","date":"2020-08-21","marketPercent":0.07173647161506266,"avgMarketPercent":0.05821660989654012},{"volume":6530976,"venue":"RCAX","venueName":"cAN aErSY","date":"2020-08-21","marketPercent":0.07481840061699063,"avgMarketPercent":0.07849704521972577},{"volume":18760292,"venue":"GSXN","venueName":"Nqdasa","date":"2020-08-21","marketPercent":0.22052084929315618,"avgMarketPercent":0.23784064048183048},{"volume":45987894,"venue":"FTR","venueName":"hnOfxcgafeE ","date":"2020-08-21","marketPercent":0.544979824806561,"avgMarketPercent":0.5340732619422156}]`, &d)
	type args struct {
		symbol string
	}
	tests := []struct {
		name      string
		o         *Client
		args      args
		want      []*VolumeByVenue
		roundTrip roundTripFunc
		wantErr   bool
	}{
		{
			name: "Success",
			o:    NewClient("", true),
			args: args{
				symbol: "AAPL",
			},
			want:      d,
			roundTrip: getRoundTripFunc("/stock/AAPL/volume-by-venue", http.StatusOK, d),
			wantErr:   false,
		},
		{
			name: "Failed to create request",
			o: &Client{
				baseURL: "://",
			},
			args: args{
				symbol: "AAPL",
			},
			want:      nil,
			roundTrip: nil,
			wantErr:   true,
		},
		{
			name: "Request Failed",
			o:    NewClient("", true),
			args: args{
				symbol: "AAPL",
			},
			want:      nil,
			roundTrip: getRoundTripFunc("/stock/AAPL/volume-by-venue", http.StatusBadRequest, APIError{StatusCode: http.StatusBadRequest}),
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		if tt.roundTrip != nil {
			tt.o.setTestTransport(tt.roundTrip)
		}
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.o.VolumeByVenue(tt.args.symbol)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.VolumeByVenue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.VolumeByVenue() = %v, want %v", got, tt.want)
			}
		})
	}
}
