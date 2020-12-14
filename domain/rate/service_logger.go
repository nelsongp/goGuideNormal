package rate

import (
	"git.lifemiles.net/lm-go-libraries/lifemiles-go/log"
)

//ServiceLogger
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

//GetMilesRate ...
func (s *ServiceLogger) GetMilesRate(accRate AccRequest) RatePerStoreResponse {

	s.logger.Info("start request miles rate")
	response := s.next.GetMilesRate(accRate)
	s.logger.Info("end request to miles rate")
	return response
}
