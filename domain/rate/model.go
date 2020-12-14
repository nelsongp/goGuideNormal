package rate

//RatePerStore
type RatePerStoreResponse struct {
	Header `json:"header"`
	Body   MilesRate `json:"body"`
}

type Header struct {
	Code   string `json:"code"`
	Result string `json:"result"`
}

type AccRequest struct {
	Sender      string `json:"sender" example:"RBMCO"`
	TerminalID  string `json:"terminalID" example:"0010077923" `
	PartnerCode string `json:"partnerCode" example:"TOTCO"`
}

type MilesRate struct {
	Sender     string `json:"sender"`
	TerminalID string `json:"terminalId"`
	AcrualRate int    `json:"accrualRate"`
}
