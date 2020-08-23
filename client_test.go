//Package iex provides a http client to access IEX data

package iex

import (
	"net/http"
	"reflect"
	"testing"
)

func TestNewClient(t *testing.T) {
	type args struct {
		pk      string
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
				pk:      "",
				sk:      "",
				sandbox: true,
			},
		},
		{
			name: "Live",
			args: args{
				pk:      "",
				sk:      "",
				sandbox: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClient(tt.args.pk, tt.args.sk, tt.args.sandbox); got == nil {
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
			o:    NewClient("", "", true),
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
// 			o:    NewClient("", "", true),
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
	tests := []struct {
		name      string
		o         *Client
		want      *Metadata
		roundTrip roundTripFunc
		wantErr   bool
	}{
		{
			name: "Success",
			o:    NewClient("", "", true),
			want: &Metadata{
				PayAsYouGoEnabled:    false,
				EffectiveDate:        1,
				SubscriptionTermType: "unaanl",
				TierName:             "slodtrta_",
				MessageLimit:         501660,
				MessagesUsed:         0,
				CircuitBreaker:       nil,
			},
			roundTrip: getRoundTripFunc("/account/metadata", http.StatusOK, Metadata{
				PayAsYouGoEnabled:    false,
				EffectiveDate:        1,
				SubscriptionTermType: "unaanl",
				TierName:             "slodtrta_",
				MessageLimit:         501660,
				MessagesUsed:         0,
				CircuitBreaker:       nil,
			}),
			wantErr: false,
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
			name: "Request Failed",
			o: &Client{
				baseURL: "https://example.com",
			},
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
