package rate

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	//SUCCESS....
	SUCCESS = `<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/">
			<soapenv:Header></soapenv:Header>
			<soapenv:Body>
			<ns1:QueryResultsResponse xmlns:ns1="http://www.ibsplc.com/iloyal/crmcore/querybuilder/retrievequeryresults/type/">
				<companyCode>LM</companyCode>
				<queryResults>
					<queryCode>QRY1025</queryCode>
					<queryName>UberRatesSQL</queryName>
					<queryDescription xsi:nil="1" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"/>
					<queryResultColumnHeader>
						<columnIndex>2</columnIndex>
						<columnName>strtxn.rdmrat</columnName>
						<columnAlias>red_rate</columnAlias>
						<columnDataType>NUMBER</columnDataType>
					</queryResultColumnHeader>
					<queryResultColumnHeader>
						<columnIndex>1</columnIndex>
						<columnName>strtxn.acrrat</columnName>
						<columnAlias>acc_rate</columnAlias>
						<columnDataType>NUMBER</columnDataType>
					</queryResultColumnHeader>
					<queryResultColumnHeader>
						<columnIndex>3</columnIndex>
						<columnName>strdtl.sndcod</columnName>
						<columnAlias>sender_code</columnAlias>
						<columnDataType>VARCHAR2</columnDataType>
					</queryResultColumnHeader>
					<queryResultRow>
						<rowIndex>1</rowIndex>
						<queryResultColumn>
						<columnIndex>1</columnIndex>
						<columnValue>2</columnValue>
						</queryResultColumn>
						<queryResultColumn>
						<columnIndex>2</columnIndex>
						<columnValue>22</columnValue>
						</queryResultColumn>
						<queryResultColumn>
						<columnIndex>3</columnIndex>
						<columnValue>RBMCO</columnValue>
						</queryResultColumn>
					</queryResultRow>
				</queryResults>
				<hasNextPage>false</hasNextPage>
				<absoluteIndex>1</absoluteIndex>
				<txnHeader>
					<transactionID>1</transactionID>
					<userName>m-portal</userName>
					<timeStamp>2020-08-12T22:44:25.000Z</timeStamp>
				</txnHeader>
			</ns1:QueryResultsResponse>
			</soapenv:Body>
		</soapenv:Envelope>`
	NODATA = `<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/">
			   <soapenv:Header></soapenv:Header>
			   <soapenv:Body>
			      <ns1:QueryResultsResponse xmlns:ns1="http://www.ibsplc.com/iloyal/crmcore/querybuilder/retrievequeryresults/type/">
			         <companyCode>LM</companyCode>
			         <queryResults>
			            <queryCode>QRY1025</queryCode>
			            <queryName>UberRatesSQL</queryName>
			            <queryDescription xsi:nil="1" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"/>
			         </queryResults>
			         <hasNextPage>false</hasNextPage>
			         <absoluteIndex>0</absoluteIndex>
			         <txnHeader>
			            <transactionID>1</transactionID>
			            <userName>m-portal</userName>
			            <timeStamp>2020-08-12T22:44:25.000Z</timeStamp>
			         </txnHeader>
			      </ns1:QueryResultsResponse>
			   </soapenv:Body>
			</soapenv:Envelope>`
	//ERROR
	ERROR = `<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/">
			   <soapenv:Header></soapenv:Header>
			   <soapenv:Body>
			      <soapenv:Fault>
			         <faultcode>soapenv:Server</faultcode>
			         <faultstring>CRMCoreWebServiceException</faultstring>
			         <detail>
			            <ns1:CRMCoreWebServiceException xmlns:ns1="http://www.ibsplc.com/iloyal/crmcore/querybuilder/retrievequeryresults/type/">
			               <faultcode>crmcore.querybuilder.queryNotFound</faultcode>
			               <faultstring>crmcore.querybuilder.queryNotFound</faultstring>
			               <txnHeader xsi:nil="1" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"/>
			            </ns1:CRMCoreWebServiceException>
			         </detail>
			      </soapenv:Fault>
			   </soapenv:Body>
			</soapenv:Envelope>`
)

func mockService(t *testing.T) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path == "/success" {

			w.Header().Set("Content-Type", "text/xml; charset=utf-8")
			w.WriteHeader(200)
			w.Write([]byte(SUCCESS))

		}
		if r.URL.Path == "/nodata" {
			w.Header().Set("Content-Type", "text/xml; charset=utf-8")
			w.WriteHeader(200)
			w.Write([]byte(NODATA))
		}

		if r.URL.Path == "/app-err" {

			w.Header().Set("Content-Type", "text/xml; charset=utf-8")
			w.WriteHeader(500)
			w.Write([]byte(ERROR))

		}

	}))
	return server
}
