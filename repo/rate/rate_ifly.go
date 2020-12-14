package rate

import (
	"errors"
	"net/http"
	"time"

	"git.lifemiles.net/LM-Global-Partners/lmgp-rates-svc/xmlreq"
	config "git.lifemiles.net/lm-go-libraries/lifemiles-go/configuration"
	"git.lifemiles.net/lm-go-libraries/lifemiles-go/log"
	"github.com/mencosk/soap"
)

func NewRateIfly(config config.Config, logger log.Logger) *RateIfly {
	return &RateIfly{config, logger}
}

type RateIfly struct {
	Config config.Config
	logger log.Logger
}

func (rate *RateIfly) GetRatePerStore(partner string, store string) (RatesPerStore, error) {
	res := RatesPerStore{}
	client := soap.New()
	client.SetTimeOut(rate.Config.GetDuration("ifly.timeout") * time.Second)

	url := rate.Config.GetString("ifly.queryresultservices.endpoint")
	req := makeRequest(rate.Config, partner, store)
	rate.logger.Infof("Request to query result service %v", req)
	rate.logger.Info("URL service", "url", url)

	rep, err := client.R().
		SetUrl(url).
		SetHeader("Content-Type", "text/xml; charset=utf-8").
		SetHeader("SOAPAction", rate.Config.GetString("ifly.queryresultservices.action")).
		SetPayloadRequest(req).
		SetPayloadResponse(&xmlreq.RetrieveQueryResponse{}).
		SetPayloadFault(&xmlreq.RetrieveQueryResponseError{}).
		Call()

	if err != nil {
		rate.logger.Error("Error to call retrieve query result service", "error", err.Error())
		return res, errors.New("Error to call retrieve query result service")
	}

	rate.logger.Info("Response from retrieve query result service",
		"statusCode", rep.StatusCode(),
		"payloadResultError", rep.PayloadResultError(),
		"payloadResult", rep.PayloadResult(),
	)

	if rep.StatusCode() != http.StatusOK {

		fault := rep.PayloadResultError().(*xmlreq.RetrieveQueryResponseError)

		faultErr := fault.Body.Fault.Faultstring
		if faultErr != "" {
			rate.logger.Error(faultErr)
			return res, errors.New(faultErr)
		}

	}

	response := rep.PayloadResult().(*xmlreq.RetrieveQueryResponse)

	if response.Body.QueryResultsResponse.AbsoluteIndex == "0" {
		rate.logger.Error("Query doesn't return data")
		return res, errors.New("Query doesn't return data")
	}
	accRate := response.Body.QueryResultsResponse.QueryResults.QueryResultRow.QueryResultColumn[0].ColumnValue

	rdmRate := response.Body.QueryResultsResponse.QueryResults.QueryResultRow.QueryResultColumn[1].ColumnValue

	res.Sender = response.Body.QueryResultsResponse.QueryResults.QueryResultRow.QueryResultColumn[2].ColumnValue
	res.Partner = partner
	res.Store = store
	res.AccRate = accRate
	res.RdmRate = rdmRate
	return res, nil

}

func makeRequest(config config.Config, partner string, store string) xmlreq.RetrieveQueryRequest {
	now := time.Now()
	timeStamp := now.Format("2006-01-02T15:04:05.000Z")

	request := xmlreq.RetrieveQueryRequest{}
	request.Soapenv = "http://schemas.xmlsoap.org/soap/envelope/"
	request.Type = config.GetString("ifly.queryresultservices.type")
	request.Body.QueryResultsRequest.CompanyCode = config.GetString("ifly.companycode")

	request.Body.QueryResultsRequest.QueryFilter.QueryCode = config.GetString("ifly.queryresultservices.querycode")

	filtro1 := xmlreq.QueryFilterAttributes{
		Text:           "",
		AttributeCode:  config.GetString("ifly.queryresultservices.attributes.store"),
		AttributeValue: store,
	}

	filtro2 := xmlreq.QueryFilterAttributes{
		Text:           "",
		AttributeCode:  config.GetString("ifly.queryresultservices.attributes.partner"),
		AttributeValue: partner,
	}
	request.Body.QueryResultsRequest.QueryFilter.QueryFilterAttributes = append(request.Body.QueryResultsRequest.QueryFilter.QueryFilterAttributes, filtro1, filtro2)

	request.Body.QueryResultsRequest.TxnHeader.UserName = config.GetString("ifly.username")
	request.Body.QueryResultsRequest.TxnHeader.TimeStamp = timeStamp
	request.Body.QueryResultsRequest.PageNumber = "1"
	request.Body.QueryResultsRequest.AbsoluteIndex = "1"
	request.Body.QueryResultsRequest.PageSize = "25"

	return request
}
