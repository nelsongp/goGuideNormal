package rate

import (
	"reflect"
	"testing"
	"time"

	"git.lifemiles.net/LM-Global-Partners/lmgp-rates-svc/xmlreq"
	"git.lifemiles.net/lm-go-libraries/lifemiles-go/configuration"
	"git.lifemiles.net/lm-go-libraries/lifemiles-go/log"
)

func getMockConfig() configuration.Config {

	config := configuration.GetInstance(configuration.NewSetting("../../", "application", "yaml", false))

	config.Set("ifly.queryresultservices.endpoint", "localhost/test")
	config.Set("ifly.queryresultservices.action", "localhost/test")
	config.Set("ifly.queryresultservices.type", "localhost/test")
	config.Set("ifly.queryresultservices.querycode", "QRY1025")
	config.Set("ifly.queryresultservices.attributes.partner", "PARTNER")
	config.Set("ifly.queryresultservices.attributes.store", "STORE")
	config.Set("ifly.timeout", 40)
	config.Set("ifly.companycode", "LM")
	config.Set("ifly.programcode", "LMS")
	config.Set("ifly.username", "m-portal")
	return config
}

func TestNewIflyRepo(t *testing.T) {
	type args struct {
		config configuration.Config
		logger log.Logger
	}

	config := getMockConfig()
	logger := log.NewLogger()

	tests := []struct {
		name string
		args args
		want *RateIfly
	}{
		{
			name: "Testing NewRateIfly",
			args: args{
				config: config,
				logger: logger,
			},
			want: &RateIfly{config, logger},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRateIfly(tt.args.config, tt.args.logger); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRateIfly() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakeRequest(t *testing.T) {
	type args struct {
		config configuration.Config
	}

	config := getMockConfig()
	now := time.Now()
	timeStamp := now.Format("2006-01-02T15:04:05.000Z")

	request := xmlreq.RetrieveQueryRequest{}

	request.Soapenv = "http://schemas.xmlsoap.org/soap/envelope/"
	request.Type = "localhost/test"
	request.Body.QueryResultsRequest.CompanyCode = "LM"

	request.Body.QueryResultsRequest.QueryFilter.QueryCode = "QRY1025"

	filtro1 := xmlreq.QueryFilterAttributes{
		Text:           "",
		AttributeCode:  "STORE",
		AttributeValue: "0010077923",
	}

	filtro2 := xmlreq.QueryFilterAttributes{
		Text:           "",
		AttributeCode:  "PARTNER",
		AttributeValue: "TOTCO",
	}
	request.Body.QueryResultsRequest.QueryFilter.QueryFilterAttributes = append(request.Body.QueryResultsRequest.QueryFilter.QueryFilterAttributes, filtro1, filtro2)
	request.Body.QueryResultsRequest.TxnHeader.UserName = "m-portal"
	request.Body.QueryResultsRequest.TxnHeader.TimeStamp = timeStamp
	request.Body.QueryResultsRequest.AbsoluteIndex = "1"
	request.Body.QueryResultsRequest.PageNumber = "1"
	request.Body.QueryResultsRequest.PageSize = "25"

	tests := []struct {
		name      string
		args      args
		timestamp string
		want      xmlreq.RetrieveQueryRequest
	}{
		{
			name: "Test makeRequest",
			args: args{
				config: config,
			},
			timestamp: timeStamp,
			want:      request,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := makeRequest(tt.args.config, "TOTCO", "0010077923")
			got.Body.QueryResultsRequest.TxnHeader.TimeStamp = tt.timestamp
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRatesPerStore(t *testing.T) {
	type fields struct {
		Config configuration.Config
		logger log.Logger
	}
	svc := mockService(t)
	logger := log.NewLogger()
	config := getMockConfig()

	tests := []struct {
		name     string
		endpoint string
		action   string
		fields   fields
		want     RatesPerStore
		wantErr  bool
		typeErr  string
	}{
		{
			name:     "Test QueryResultRequest Success",
			endpoint: svc.URL + "/success",
			action:   svc.URL + "/success",
			fields:   fields{config, logger},
			want: RatesPerStore{
				Partner: "TOTCO",
				Store:   "0010077923",
				Sender:  "RBMCO",
				AccRate: "2",
				RdmRate: "22",
			},
			wantErr: false,
		},
		{
			name:     "Test QueryResultRequest Error",
			endpoint: svc.URL + "/app-err",
			action:   svc.URL + "/app-err",
			fields:   fields{config, logger},
			want:     RatesPerStore{},
			wantErr:  true,
			typeErr:  "CRMCoreWebServiceException",
		},
		{
			name:     "Test QueryResultRequest Bad Request",
			endpoint: svc.URL + "/",
			action:   svc.URL + "/",
			fields:   fields{config, logger},
			want:     RatesPerStore{},
			wantErr:  true,
			typeErr:  "Error to call retrieve query result service",
		},
		{
			name:     "Test QueryResultRequest No Data",
			endpoint: svc.URL + "/nodata",
			action:   svc.URL + "/nodata",
			fields:   fields{config, logger},
			want:     RatesPerStore{},
			wantErr:  true,
			typeErr:  "Query doesn't return data",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.Config.Set("ifly.queryresultservices.endpoint", tt.endpoint)
			tt.fields.Config.Set("ifly.queryresultservices.action", tt.action)
			mr := &RateIfly{
				Config: tt.fields.Config,
				logger: tt.fields.logger,
			}
			got, err := mr.GetRatePerStore("TOTCO", "0010077923")
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRatePerStore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if (err != nil) && tt.wantErr && !reflect.DeepEqual(err.Error(), tt.typeErr) {
				t.Errorf("GetRatePerStore() typeErr = %v, wantTypeErr %v", err, tt.typeErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRatePerStore() got = %v, want %v", got, tt.want)
			}
		})
	}
}
