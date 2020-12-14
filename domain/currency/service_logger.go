package currency

import (
	"git.lifemiles.net/lm-go-libraries/lifemiles-go/log"
)

//ServiceLogger ...
type ServiceLogger struct {
	logger log.Logger
	next   Service
}

//NewServiceLogger ...
func NewServiceLogger(logger log.Logger, n Service) *ServiceLogger {
	return &ServiceLogger{
		logger: logger,
		next:   n,
	}
}

//GetCurrency ...
func (s *ServiceLogger) GetCurrency() CurrencyResponse {

	s.logger.Info("start request to getCurrency")

	response := s.next.GetCurrency()

	s.logger.Info("end request to getCurrency")

	return response
}
