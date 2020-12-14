package rate

import (
	"context"
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

//NewHTTPHandler ...
// @Summary Return RatePerStoreResponse
// @Description Retrieve Rate Per Store Response
// @Tags AccRequest
// @Param accRequest body AccRequest true "AccRequest"
// @Accept  json
// @Produce  json
// @Success 200 {object} RatePerStoreResponse
// @Failure 401 "Unauthorized"
// @Failure 403 "Forbidden"
// @Failure 404 "Not Found"
// @Router /milesRate [post]
func NewHTTPHandler(endpoints Endpoints) http.Handler {

	r := mux.NewRouter()

	r.Handle("/milesRate",
		httptransport.NewServer(
			endpoints.MilesRate,
			DecodeGetRequest,
			EncodeGetResponse,
		)).Methods("POST")

	return r
}

//DecodeGetRequest
func DecodeGetRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var accRequest AccRequest
	if err := json.NewDecoder(r.Body).Decode(&accRequest); err != nil {
		return nil, err
	}
	return accRequest, nil
}

//EncodeGetResponse ...
func EncodeGetResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
