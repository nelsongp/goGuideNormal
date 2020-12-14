package currency

//CurrencyResponse ...
type CurrencyResponse struct {
	Header `json:"header"`
	Body   []Currency `json:"body"`
}

//Header ...
type Header struct {
	Code   string `json:"code"`
	Result string `json:"result"`
}

//Currency ...
type Currency struct {
	CurrencyCode string  `json:"currencyCode"`
	RateToUSD    float32 `json:"rateToUSD"`
}
