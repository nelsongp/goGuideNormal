package currency

import (
	currencyRepo "git.lifemiles.net/LM-Global-Partners/lmgp-rates-svc/repo/currency"
)

//Service ...
type Service interface {
	GetCurrency() CurrencyResponse
}

//ProcessService ...
type ProcessService struct {
	cRepo currencyRepo.Repository
}

//NewCurrencyService ...
func NewCurrencyService(cRepo currencyRepo.Repository) *ProcessService {
	return &ProcessService{cRepo}
}

//GetCurrency ...
func (s *ProcessService) GetCurrency() CurrencyResponse {

	var header Header
	var response []Currency
	currencies, err := s.cRepo.GetCurrency()
	if err != nil {
		header = Header{
			Code:   "999",
			Result: "Error to get results set",
		}
		return CurrencyResponse{
			Header: header,
		}
	}

	for _, currency := range currencies {
		if currency.CurrencyCode != "USD" && len(currencies) > 1 {
			response = append(response, Currency{
				currency.CurrencyCode,
				currency.RateToUSD,
			})
		} else if len(currencies) == 1 {
			response = append(response, Currency{
				currency.CurrencyCode,
				currency.RateToUSD,
			})
		}
	}
	header = Header{
		Code:   "000",
		Result: "Success",
	}
	return CurrencyResponse{
		Header: header,
		Body:   response,
	}

}
