package currency

import (
	"context"
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

//NewHTTPCurrencyHandler ...
// @Summary Return CurrencyResponse
// @Description Retrieve a list of currencies and rates to USD
// @Tags Currency
// @Accept  json
// @Produce  json
// @Success 200 {object} CurrencyResponse
// @Failure 401 "Unauthorized"
// @Failure 403 "Forbidden"
// @Failure 404 "Not Found"
// @Router /getCurrencies [get]
func NewHTTPCurrencyHandler(endpoints Endpoints) http.Handler {

	r := mux.NewRouter()

	r.Handle("/getCurrencies",
		httptransport.NewServer(
			endpoints.Currency,
			DecodeGetRequest,
			EncodeGetResponse,
		)).Methods("GET")

	return r
}

//DecodeGetRequest ...
func DecodeGetRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

//EncodeGetResponse ...
func EncodeGetResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
