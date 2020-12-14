package currency

import (
	"reflect"
	"strings"
	"testing"

	"git.lifemiles.net/lm-go-libraries/lifemiles-go/configuration"
	"git.lifemiles.net/lm-go-libraries/lifemiles-go/log"
)

func getMockConfig() configuration.Config {

	config := configuration.GetInstance(configuration.NewSetting("../../", "application", "yaml", false))

	config.Set("ps-currency.getCurrencies.currencies.COP", "CO")
	return config
}

func TestNewPSRepo(t *testing.T) {
	type args struct {
		config configuration.Config
		logger log.Logger
	}

	config := getMockConfig()
	logger := log.NewLogger()

	tests := []struct {
		name string
		args args
		want *CurrencyPS
	}{
		{
			name: "Testing NewPSRepo",
			args: args{
				config: config,
				logger: logger,
			},
			want: &CurrencyPS{config, logger},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCurrencyPS(tt.args.config, tt.args.logger); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCurrencyPS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCurrency(t *testing.T) {
	type fields struct {
		Config configuration.Config
		logger log.Logger
	}

	svc := mockService(t)
	logger := log.NewLogger()
	config := getMockConfig()

	tests := []struct {
		name    string
		url     string
		COP     string
		fields  fields
		want    []Currency
		wantErr bool
		typeErr string
	}{
		{
			name:   "Test ps-currency/getCurrencies Success",
			url:    svc.URL + "/success",
			COP:    "",
			fields: fields{config, logger},
			want: []Currency{
				{"COP", 3277.14},
			},
			wantErr: false,
		},
		{
			name:    "Test ps-currency/getCurrencies Application Error",
			url:     svc.URL + "/app-err",
			COP:     "",
			fields:  fields{config, logger},
			want:    []Currency{},
			wantErr: true,
			typeErr: "HttpStatusCode = 404",
		},
		{
			name:    "Test ps-currency/getCurrencies Service Unavailable",
			url:     strings.Replace(svc.URL, "127.0.0.1", "", 1),
			COP:     "COL",
			fields:  fields{config, logger},
			want:    []Currency{},
			wantErr: true,
			typeErr: "HttpStatusCode = 500",
		},
		/* 		{
			name:    "Test ps-currency/getCurrencies Service Unavailable",
			url:     svc.URL + "/fault",
			COP:     "COL",
			fields:  fields{config, logger},
			want:    []Currency{},
			wantErr: true,
			typeErr: "EOF",
		}, */
		{
			name:   "Test ps-currency/getCurrencies Success",
			url:    svc.URL + "/success2",
			COP:    "",
			fields: fields{config, logger},
			want: []Currency{
				{"COP", 3277.14},
				{"USD", 1},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.Config.Set("ps-currency.getCurrencies.endpoint", tt.url)
			tt.fields.Config.Set("ps-currency.getCurrencies.currencies.COP", tt.COP)
			mr := &CurrencyPS{
				Config: tt.fields.Config,
				logger: tt.fields.logger,
			}
			got, err := mr.GetCurrency()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCurrency() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if (err != nil) && tt.wantErr && !reflect.DeepEqual(err.Error(), tt.typeErr) {
				t.Errorf("GetCurrency() typeErr = %v, wantTypeErr %v", err, tt.typeErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCurrency() got = %v, want %v", got, tt.want)
			}
		})
	}
}
