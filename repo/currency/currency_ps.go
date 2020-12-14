package currency

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	config "git.lifemiles.net/lm-go-libraries/lifemiles-go/configuration"
	"git.lifemiles.net/lm-go-libraries/lifemiles-go/log"
)

//NewCurrencyPS ...
func NewCurrencyPS(config config.Config, logger log.Logger) *CurrencyPS {
	return &CurrencyPS{config, logger}
}

//CurrencyPS ...
type CurrencyPS struct {
	Config config.Config
	logger log.Logger
}

//GetCurrency ...
func (currency *CurrencyPS) GetCurrency() ([]Currency, error) {
	var result Currencies
	res := []Currency{}
	url := currency.Config.GetString("ps-currency.getCurrencies.endpoint")
	COP := currency.Config.GetString("ps-currency.getCurrencies.currencies.COP")
	resp, err := http.Get(url + COP)

	if resp == nil || err != nil {
		err = errors.New("HttpStatusCode = " + strconv.Itoa(http.StatusInternalServerError))
		currency.logger.Error("Error consulting ps-currency/getCurrencies service", "Error", err.Error)
		return res, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = errors.New("HttpStatusCode = " + strconv.Itoa(resp.StatusCode))
		currency.logger.Error("Error  getting information from  ps-currency/getCurrencies service", "Error", err.Error)
		return res, err
	}

	err = json.NewDecoder(resp.Body).Decode(&result)

	if err != nil {
		currency.logger.Error("Error decoding ps-currency/getCurrencies service response", "error", err.Error())
		return res, err
	}

	for _, list := range result.Items {
		items := Currency{}
		items.CurrencyCode = list.Code
		items.RateToUSD = list.RateToUsd

		res = append(res, items)
	}
	return res, nil

}
