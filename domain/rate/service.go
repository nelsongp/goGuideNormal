package rate

import (
	rateRepo "git.lifemiles.net/LM-Global-Partners/lmgp-rates-svc/repo/rate"
	"strconv"
)

//Service (analogo a la clase donde se encuentra el metodo a ejecutar)
//Aca se define el metodo que le vas a regresar osea el service regresa la interfax GetMilesRate - Nelson
type Service interface {
	GetMilesRate(accRate AccRequest) RatePerStoreResponse
}

//ProcessService (analogo de la llamada a constructor de clase en java)
//Aca se manda a llamar el repo de donde se saca la info (en este caso ifly) - Nelson
type ProcessService struct {
	rateRepo rateRepo.Repository
}

//inicializacion de constuctor (o asi lo veo yo)
//Entenderia yo que es como la inicializadon del repo al parecer - Nelson
func NewMilesRateService(rateRepo rateRepo.Repository) *ProcessService {
	return &ProcessService{rateRepo}
}

func (s *ProcessService) GetMilesRate(accRate AccRequest) RatePerStoreResponse {
	var header Header
	var response MilesRate
	ratesPerStore, err := s.rateRepo.GetRatePerStore(accRate.PartnerCode, accRate.TerminalID)
	if err != nil {
		//problemas de conexion con Ifly
		if err.Error() == "Error to call retrieve query result service" {
			header = Header{
				Code:   "999",
				Result: "Error - Servicio no disponible",
			}
			return RatePerStoreResponse{
				Header: header,
			}
		} else { // se realizo la busqueda en Ifly pero no encontro nada
			header = Header{
				Code:   "001",
				Result: "Error - No se encuentra configuración para el comercio solicitado",
			}
			return RatePerStoreResponse{
				Header: header,
			}
		}
	}
	//revisamos si el sender obtenido es igual al recibido
	if ratesPerStore.Sender != accRate.Sender {
		header = Header{
			Code:   "001",
			Result: "Error - No se encuentra configuración para el comercio solicitado",
		}
		return RatePerStoreResponse{
			Header: header,
		}
	}

	accrualRate, err := strconv.Atoi(ratesPerStore.AccRate)

	if err != nil {
		//el valor recibido en el servicio de Ifly no devolvio un entero
		header = Header{
			Code:   "001",
			Result: "Error - No se encuentra configuración para el comercio solicitado",
		}
		return RatePerStoreResponse{
			Header: header,
		}
	}
	response = MilesRate{
		AcrualRate: accrualRate,
		Sender:     ratesPerStore.Sender,
		TerminalID: ratesPerStore.Store,
	}

	header = Header{
		Code:   "000",
		Result: "Success",
	}
	return RatePerStoreResponse{
		Header: header,
		Body:   response,
	}
}
