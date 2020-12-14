package currency

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDecodeGetCurrencyRequest(t *testing.T) {
	type args struct {
		ctx     context.Context
		r       *http.Request
		request interface{}
	}

	ctx := context.Background()

	svc := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tests := []struct {
			name    string
			args    args
			wantErr bool
		}{
			{
				name: "Test CurrencyRequest Success",
				args: args{
					ctx: ctx,
					r:   r,
				},
				wantErr: false,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if request, err := DecodeGetRequest(tt.args.ctx, tt.args.r); (err != nil) != tt.wantErr && request != nil {
					t.Errorf("DecodeGetRequest() error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		}
	}))

	http.Get(svc.URL)
}

func TestEncodeGetCurrencyResponse(t *testing.T) {
	type args struct {
		ctx      context.Context
		w        http.ResponseWriter
		response interface{}
	}

	resp := `{
		"header": {
			"code": "000",
			"result": "Success"
		},
		"body": [
			{
				"currencyCode": "COP",
				"rateToUSD": 3277.14
			}
		]
	}`

	ctx := context.Background()
	//var wrt http.ResponseWriter

	svc := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		w.Write([]byte(resp))

		tests := []struct {
			name    string
			args    args
			wantErr bool
		}{
			{
				name: "Test CurrencyResponse Success",
				args: args{
					ctx:      ctx,
					w:        w,
					response: CurrencyResponse{},
				},
				wantErr: false,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if err := EncodeGetResponse(tt.args.ctx, tt.args.w, tt.args.response); (err != nil) != tt.wantErr {
					t.Errorf("EncodeGetResponse() error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		}
	}))

	http.Get(svc.URL)
}

func TestNewHTTPHandler(t *testing.T) {
	type args struct {
		endpoints Endpoints
	}

	endpoint := Endpoints{
		Currency: MakeGetCurrencyEndpoint(TestImpl{}),
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test Handler Success",
			args: args{endpoints: endpoint},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHTTPCurrencyHandler(tt.args.endpoints); got == nil {
				t.Errorf("NewHTTPCurrencyHandler() = %v, want not nil", got)
			}
		})
	}
}

type TestImpl struct{}

func (TestImpl) GetCurrency() CurrencyResponse {
	return CurrencyResponse{}
}
