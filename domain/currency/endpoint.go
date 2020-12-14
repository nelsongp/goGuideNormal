package currency

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

//MakeGetCurrencyEndpoint ...
func MakeGetCurrencyEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return s.GetCurrency(), nil
	}
}

//Endpoints ...
type Endpoints struct {
	Currency endpoint.Endpoint
}
