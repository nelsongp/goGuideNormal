package rate

import (
	"context"
	"reflect"
	"testing"

	rateRepo "git.lifemiles.net/LM-Global-Partners/lmgp-rates-svc/repo/rate"
)

func TestMakeGetMilesRatesEndpoint(t *testing.T) {
	type args struct {
		s Service
	}
	request := AccRequest{
		PartnerCode: "TOTCO",
		Sender:      "RBMCO",
		TerminalID:  "0010077923",
	}
	service := NewMilesRateService(rateRepo.NewRateStatic())
	tests := []struct {
		name         string
		args         args
		ctx          context.Context
		wantResponse RatePerStoreResponse
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
			fn := MakeGetMilesRatesEndpoint(tt.args.s)
			got, err := fn(tt.ctx, request)

			if (err != nil) && !reflect.DeepEqual(got, tt.wantResponse) {
				t.Errorf("MilesRateEndpoint() = %v, want %v", got, tt.wantResponse)
			}
		})
	}
}
func makeResponse() RatePerStoreResponse {

	header := Header{
		Code:   "000",
		Result: "Success",
	}

	res := MilesRate{

		Sender:     "RBMCO",
		AcrualRate: 2,
		TerminalID: "0010077923",
	}
	return RatePerStoreResponse{header, res}
}
