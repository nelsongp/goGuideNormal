package currency

import (
	"context"
	"reflect"
	"testing"

	currencyRepo "git.lifemiles.net/LM-Global-Partners/lmgp-rates-svc/repo/currency"
)

func TestMakeGetCurrencyEndpoint(t *testing.T) {
	type args struct {
		s Service
	}

	service := NewCurrencyService(currencyRepo.NewCurrencyStatic())

	tests := []struct {
		name         string
		args         args
		ctx          context.Context
		wantResponse CurrencyResponse
	}{
		{
			name:         "Test Endpoint",
			args:         args{s: service},
			ctx:          context.Background(),
			wantResponse: makeResponse(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn := MakeGetCurrencyEndpoint(tt.args.s)
			got, err := fn(tt.ctx, nil)

			if (err != nil) && !reflect.DeepEqual(got, tt.wantResponse) {
				t.Errorf("CurrencyResponse() = %v, want %v", got, tt.wantResponse)
			}
		})
	}
}

func makeResponse() CurrencyResponse {

	header := Header{
		Code:   "000",
		Result: "Success",
	}

	res := []Currency{
		{"COP", 3277.14},
	}
	return CurrencyResponse{header, res}
}
