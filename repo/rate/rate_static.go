package rate

import "errors"

type RatesPerStore struct {
	Partner string
	Store   string
	Sender  string
	AccRate string
	RdmRate string
}

type Repository interface {
	GetRatePerStore(partner string, store string) (rates RatesPerStore, err error)
}

type RateStatic struct {
}

//constructor
func NewRateStatic() *RateStatic {
	return &RateStatic{}
}

func (rs *RateStatic) GetRatePerStore(partner string, store string) (rates RatesPerStore, err error) {
	var res RatesPerStore

	if partner == "TOTCO" && store == "0010077923" {
		res.Partner = "TOTCO"
		res.Store = "0010077923"
		res.Sender = "RBMCO"
		res.AccRate = "2"
		res.RdmRate = "22"

	} else {
		return res, errors.New("Query doesn't return data")
	}
	return res, nil
}
