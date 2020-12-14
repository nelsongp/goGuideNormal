package currency

import (
	"reflect"
	"testing"

	currencyRepo "git.lifemiles.net/LM-Global-Partners/lmgp-rates-svc/repo/currency"
	"git.lifemiles.net/lm-go-libraries/lifemiles-go/log"
)

func TestNewServiceLogger(t *testing.T) {
	type args struct {
		logger log.Logger
		n      Service
	}

	cRepo := currencyRepo.NewCurrencyStatic()
	CurrencyService := NewCurrencyService(cRepo)

	logger := log.NewLogger()
	next := CurrencyService
	//service := Service.GetCurrency()
	tests := []struct {
		name string
		args args
		want *ServiceLogger
	}{
		{
			name: "Testing NewServiceLogger",
			args: args{
				logger: logger,
				n:      next,
			},
			want: &ServiceLogger{logger, next},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewServiceLogger(tt.args.logger, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewServiceLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_serviceLogger_GetCurrency(t *testing.T) {

	type args struct {
		logger log.Logger
		n      Service
	}

	cRepo := currencyRepo.NewCurrencyStatic()
	CurrencyService := NewCurrencyService(cRepo)

	logger := log.NewLogger()
	next := CurrencyService

	mw := &ServiceLogger{
		logger: logger,
		next:   next,
	}

	tests := []struct {
		name string
		args args
		want CurrencyResponse
	}{
		{
			name: "Testing ServiceLogger_GetCurrency",
			args: args{
				logger: logger,
				n:      next,
			},
			want: CurrencyResponse{
				Header: Header{
					Code:   "000",
					Result: "Success",
				},
				Body: []Currency{
					{"COP", 3277.14},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mw.GetCurrency(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ServiceLogger_GetCurrency() = %v, want %v", got, tt.want)
			}
		})
	}

}

/* func makeCurrency() CurrencyResponse {
	header := Header{
		Code: "000",
		Result: "Success",
	}

	body := Currency{
		CurrencyCode : "COP",
		RateToUSD: 3277.14,
	}

	response := CurrencyResponse{
		Header: header,
		Body: []body,
	}

	return response

} */
