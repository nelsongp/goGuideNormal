package rate

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

//MilesRatesEndpoint ...
func MakeGetMilesRatesEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return s.GetMilesRate(request.(AccRequest)), nil
	}
}

//Endpoints ...
type Endpoints struct {
	MilesRate endpoint.Endpoint
}
