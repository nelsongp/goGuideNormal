package currency

import (
	"reflect"
	"testing"

	"git.lifemiles.net/lm-go-libraries/lifemiles-go/configuration"
	"git.lifemiles.net/lm-go-libraries/lifemiles-go/log"
)

func TestNewStaticRepo(t *testing.T) {
	type args struct {
		config configuration.Config
		logger log.Logger
	}

	tests := []struct {
		name string
		want *CurrencyStatic
	}{
		{
			name: "Testing NewStaticRepo",
			want: &CurrencyStatic{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCurrencyStatic(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCurrencyStatic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCurrencyStatic(t *testing.T) {

	svc := mockService(t)
	tests := []struct {
		name    string
		url     string
		COP     string
		want    []Currency
		wantErr bool
		typeErr string
	}{
		{
			name: "Test static-currency/getCurrencies Success",
			url:  svc.URL + "/success",
			COP:  "",
			want: []Currency{
				{"COP", 3277.14},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mr := &CurrencyStatic{}
			got, err := mr.GetCurrency()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCurrencyStatic() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if (err != nil) && tt.wantErr && !reflect.DeepEqual(err.Error(), tt.typeErr) {
				t.Errorf("GetCurrencyStatic() typeErr = %v, wantTypeErr %v", err, tt.typeErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCurrencyStatic() got = %v, want %v", got, tt.want)
			}
		})
	}
}
