package rate

import (
	"errors"
	"reflect"
	"testing"

	"git.lifemiles.net/LM-Global-Partners/lmgp-rates-svc/repo/rate"
	repo "git.lifemiles.net/LM-Global-Partners/lmgp-rates-svc/repo/rate"
)

func TestGetMilesRateService(t *testing.T) {
	type dependencies struct {
		rateRepo repo.Repository
	}
	request := AccRequest{
		PartnerCode: "TOTCO",
		Sender:      "RBMCO",
		TerminalID:  "0010077923",
	}
	tests := []struct {
		name string
		dependencies
		want RatePerStoreResponse
	}{
		{
			name: "Test Success response",
			dependencies: dependencies{
				rateRepo: repo.NewRateStatic(),
			},
			want: RatePerStoreResponse{
				Header: Header{
					Code:   "000",
					Result: "Success",
				},
				Body: MilesRate{
					Sender:     "RBMCO",
					TerminalID: "0010077923",
					AcrualRate: 2,
				},
			},
		},
		{
			name: "Test Error response",
			dependencies: dependencies{
				rateRepo: &rateRepoNil1{},
			},
			want: RatePerStoreResponse{
				Header: Header{
					Code:   "999",
					Result: "Error - Servicio no disponible",
				},
			},
		},
		{
			name: "Test No Data response",
			dependencies: dependencies{
				rateRepo: &rateRepoNil2{},
			},
			want: RatePerStoreResponse{
				Header: Header{
					Code:   "001",
					Result: "Error - No se encuentra configuración para el comercio solicitado",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := ProcessService{tt.dependencies.rateRepo}
			got := svc.GetMilesRate(request)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMileRate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMilesRateServiceSenderFail(t *testing.T) {
	type dependencies struct {
		rateRepo repo.Repository
	}
	request := AccRequest{
		PartnerCode: "TOTCO",
		Sender:      "RBMCR",
		TerminalID:  "0010077923",
	}
	tests := []struct {
		name string
		dependencies
		want RatePerStoreResponse
	}{
		{
			name: "Test Success response fail",
			dependencies: dependencies{
				rateRepo: repo.NewRateStatic(),
			},
			want: RatePerStoreResponse{
				Header: Header{
					Code:   "001",
					Result: "Error - No se encuentra configuración para el comercio solicitado",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := ProcessService{tt.dependencies.rateRepo}
			got := svc.GetMilesRate(request)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMileRate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMilesRateServiceSenderBadFormat(t *testing.T) {
	type dependencies struct {
		rateRepo repo.Repository
	}
	request := AccRequest{
		PartnerCode: "TOTCO",
		Sender:      "RBMCR",
		TerminalID:  "0010077923",
	}
	tests := []struct {
		name string
		dependencies
		want RatePerStoreResponse
	}{
		{
			name: "Test Success response fail",
			dependencies: dependencies{
				rateRepo: &rateRepoNil3{},
			},
			want: RatePerStoreResponse{
				Header: Header{
					Code:   "001",
					Result: "Error - No se encuentra configuración para el comercio solicitado",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := ProcessService{tt.dependencies.rateRepo}
			got := svc.GetMilesRate(request)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMileRate() = %v, want %v", got, tt.want)
			}
		})
	}
}

type rateRepoNil1 struct{}

func (rt *rateRepoNil1) GetRatePerStore(partner string, store string) (rates rate.RatesPerStore, err error) {
	return rate.RatesPerStore{}, errors.New("Error to call retrieve query result service")
}

type rateRepoNil2 struct{}

func (rt *rateRepoNil2) GetRatePerStore(partner string, store string) (rates rate.RatesPerStore, err error) {
	return rate.RatesPerStore{}, errors.New("Query doesn't return data")
}

type rateRepoNil3 struct{}

func (rt *rateRepoNil3) GetRatePerStore(partner string, store string) (rates rate.RatesPerStore, err error) {
	return rate.RatesPerStore{Sender: "RBMCO",
		AccRate: "valor",
		Partner: "TOTCO",
		RdmRate: "5",
		Store:   "0010077923",
	}, nil
}
