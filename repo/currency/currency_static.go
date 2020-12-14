package currency

//Currencies ...
type Currencies struct {
	ResponseCode string `json:"responseCode"`
	Items        []struct {
		Code         string  `json:"code"`
		Decimals     int     `json:"decimals"`
		RateToUsd    float32 `json:"rateToUsd"`
		Installments int     `json:"installments"`
		QuoteCountry string  `json:"quoteCountry"`
		StationID    string  `json:"stationId"`
		RedAgentID   string  `json:"redAgentId"`
		MilAgentID   string  `json:"milAgentId"`
		Product      string  `json:"product"`
		Iata         string  `json:"iata"`
		Default      bool    `json:"default"`
	} `json:"items"`
	//Body         Items  `json:"items"`
}

//Items ...
type Items struct {
	Code         string  `json:"code"`
	Decimals     int     `json:"decimals"`
	RateToUsd    float32 `json:"rateToUsd"`
	Installments int     `json:"installments"`
	QuoteCountry string  `json:"quoteCountry"`
	StationID    string  `json:"stationId"`
	RedAgentID   string  `json:"redAgentId"`
	MilAgentID   string  `json:"milAgentId"`
	Product      string  `json:"product"`
	Iata         string  `json:"iata"`
	Default      bool    `json:"default"`
}

//Currency ...
type Currency struct {
	CurrencyCode string
	RateToUSD    float32
}

//Repository ...
type Repository interface {
	GetCurrency() (currencies []Currency, err error)
}

//CurrencyStatic ...
type CurrencyStatic struct {
}

//NewCurrencyStatic ...
//constructor
func NewCurrencyStatic() *CurrencyStatic {
	return &CurrencyStatic{}
}

//GetCurrency ...
func (cs *CurrencyStatic) GetCurrency() (currencies []Currency, err error) {

	res := []Currency{
		{"COP", 3277.14},
	}

	return res, nil
}
