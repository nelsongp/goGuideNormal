package rate

import (
	"git.lifemiles.net/lm-go-libraries/lifemiles-go/log"
	"reflect"
	"testing"
)

func TestServiceLogger(t *testing.T) {

	logger := log.NewLogger()
	n := ServiceImp{}

	type args struct {
		logger log.Logger
		n      Service
	}

	tests := []struct {
		name string
		args
		want RatePerStoreResponse
	}{
		{
			name: "Test Service Logger",
			args: args{logger, n},
			want: RatePerStoreResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := ServiceLogger{tt.args.logger, tt.args.n}
			if got := svc.GetMilesRate(AccRequest{}); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMilesRate() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestNewServiceLogger(t *testing.T) {
	type args struct {
		logger log.Logger
		next   Service
	}

	logger := log.NewLogger()
	n := ServiceImp{}

	tests := []struct {
		name string
		args args
		want *ServiceLogger
	}{
		{
			name: "Testing NewServiceLogger",
			args: args{
				logger: logger,
				next:   n,
			},
			want: &ServiceLogger{logger, n},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewServiceLogger(tt.args.logger, tt.args.next); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewServiceLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}

type ServiceImp struct {
}

func (ServiceImp) GetMilesRate(accRequest AccRequest) RatePerStoreResponse {
	return RatePerStoreResponse{}
}
