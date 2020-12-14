package currency

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	//SUCCESS ...
	SUCCESS = `{
					"responseCode": "200",
					"items": [
						{
							"code": "COP",
							"decimals": 0,
							"rateToUsd": 3277.14,
							"installments": 24,
							"quoteCountry": "CO",
							"stationId": "MDE",
							"redAgentId": "777774",
							"milAgentId": "777774",
							"product": "AVL",
							"iata": "76992613",
							"default": true
						}
					],
					"message": null
				}`
	//SUCCESS2 ...
	SUCCESS2 = `{
		"responseCode": "200",
		"items": [
			{
				"code": "COP",
				"decimals": 0,
				"rateToUsd": 3277.14,
				"installments": 24,
				"quoteCountry": "CO",
				"stationId": "MDE",
				"redAgentId": "777774",
				"milAgentId": "777774",
				"product": "AVL",
				"iata": "76992613",
				"default": true
			},
			{
				"code": "USD",
				"decimals": 2,
				"rateToUsd": 1.0,
				"installments": 1,
				"quoteCountry": "CO",
				"stationId": "BOG",
				"redAgentId": "777774",
				"milAgentId": "777774",
				"product": "AVL",
				"iata": "76992613",
				"default": false
			}
		],
		"message": null
	}`
	//FAULTERROR ...
	FAULTERROR = `{
					"response": "200",
					"messageText": null
				}`

	//ERROR ...
	ERROR = `{
						"responseCode": 404,
						"messageError": "Service not available"
					}`
)

func mockService(t *testing.T) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path == "/success" {

			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(200)
			w.Write([]byte(SUCCESS))

		}

		if r.URL.Path == "/success2" {

			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(200)
			w.Write([]byte(SUCCESS2))

		}

		/* 	if r.URL.Path == "/fault" {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(404)
			w.Write([]byte(FAULTERROR))
		} */

		if r.URL.Path == "/app-err" {

			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(404)
			w.Write([]byte(ERROR))

		}

		/* if r.URL.Path == "/" {

			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(400)
			//w.Write([]byte(APPERROR))

		} */
	}))
	return server
}
