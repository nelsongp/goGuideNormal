package currency

import (
	"errors"
	"reflect"
	"testing"

	currencyRepo "git.lifemiles.net/LM-Global-Partners/lmgp-rates-svc/repo/currency"
)

type dependencies struct {
	cRepo currencyRepo.Repository
}

func TestCurrencyService(t *testing.T) {
	tests := []struct {
		name string
		dependencies
		want CurrencyResponse
	}{
		{
			name: "Test Success response",
			dependencies: dependencies{
				cRepo: currencyRepo.NewCurrencyStatic(),
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
		{
			name: "Test Failed response",
			dependencies: dependencies{
				cRepo: &currencyRepoNil{},
			},
			want: CurrencyResponse{
				Header: Header{
					Code:   "999",
					Result: "Error to get results set",
				},
			},
		},
		{
			name: "Test USD response",
			dependencies: dependencies{
				cRepo: &currencyRepoUSD{},
			},
			want: CurrencyResponse{
				Header: Header{
					Code:   "000",
					Result: "Success",
				},
				Body: []Currency{
					{"USD", 1},
				},
			},
		},
		{
			name: "Test Multiple response",
			dependencies: dependencies{
				cRepo: &currencyRepoMultiple{},
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
			svc := NewCurrencyService(tt.dependencies.cRepo)
			got := svc.GetCurrency()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCurrency() = %v, want %v", got, tt.want)
			}
		})
	}
}

type currencyRepoNil struct{}

func (cc *currencyRepoNil) GetCurrency() ([]currencyRepo.Currency, error) {
	return []currencyRepo.Currency{}, errors.New("this is a custom error")
}

type currencyRepoUSD struct{}

func (cc *currencyRepoUSD) GetCurrency() ([]currencyRepo.Currency, error) {
	res := []currencyRepo.Currency{
		{"USD", 1},
	}

	return res, nil
}

type currencyRepoMultiple struct{}

func (cc *currencyRepoMultiple) GetCurrency() ([]currencyRepo.Currency, error) {
	res := []currencyRepo.Currency{
		{"COP", 3277.14},
		{"USD", 1},
	}

	return res, nil
}
