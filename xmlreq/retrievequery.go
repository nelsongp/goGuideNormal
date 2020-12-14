package xmlreq

import "encoding/xml"

//Retrieve Query Result Request
type RetrieveQueryRequest struct {
	XMLName xml.Name `xml:"soapenv:Envelope"`
	Text    string   `xml:",chardata"`
	Soapenv string   `xml:"xmlns:soapenv,attr"`
	Type    string   `xml:"xmlns:type,attr"`
	Header  string   `xml:"soapenv:Header"`
	Body    struct {
		Text                string `xml:",chardata"`
		QueryResultsRequest struct {
			Text        string `xml:",chardata"`
			CompanyCode string `xml:"companyCode"`
			QueryFilter struct {
				Text                  string `xml:",chardata"`
				QueryCode             string `xml:"queryCode"`
				QueryFilterAttributes []struct {
					Text           string `xml:",chardata"`
					AttributeCode  string `xml:"attributeCode"`
					AttributeValue string `xml:"attributeValue"`
				} `xml:"queryFilterAttributes"`
			} `xml:"queryFilter"`
			PageNumber    string `xml:"pageNumber"`
			AbsoluteIndex string `xml:"absoluteIndex"`
			PageSize      string `xml:"pageSize"`
			TxnHeader     struct {
				Text                    string `xml:",chardata"`
				TransactionID           string `xml:"transactionID"`
				UserName                string `xml:"userName"`
				ChannelUserCode         string `xml:"channelUserCode"`
				TransactionToken        string `xml:"transactionToken"`
				TimeStamp               string `xml:"timeStamp"`
				DeviceId                string `xml:"deviceId"`
				DeviceIP                string `xml:"deviceIP"`
				DeviceOperatingSystem   string `xml:"deviceOperatingSystem"`
				DeviceLocationLatitude  string `xml:"deviceLocationLatitude"`
				DeviceLocationLongitude string `xml:"deviceLocationLongitude"`
				DeviceCountryCode       string `xml:"deviceCountryCode"`
				AdditionalInfo          string `xml:"additionalInfo"`
				Remarks                 string `xml:"remarks"`
			} `xml:"txnHeader"`
		} `xml:"type:QueryResultsRequest"`
	} `xml:"soapenv:Body"`
}

//Retrieve Query Result Response
type RetrieveQueryResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Soapenv string   `xml:"soapenv,attr"`
	Header  string   `xml:"Header"`
	Body    struct {
		Text                 string `xml:",chardata"`
		QueryResultsResponse struct {
			Text         string `xml:",chardata"`
			Ns1          string `xml:"ns1,attr"`
			CompanyCode  string `xml:"companyCode"`
			QueryResults struct {
				Text             string `xml:",chardata"`
				QueryCode        string `xml:"queryCode"`
				QueryName        string `xml:"queryName"`
				QueryDescription struct {
					Text string `xml:",chardata"`
					Nil  string `xml:"nil,attr"`
					Xsi  string `xml:"xsi,attr"`
				} `xml:"queryDescription"`
				QueryResultColumnHeader []struct {
					Text           string `xml:",chardata"`
					ColumnIndex    string `xml:"columnIndex"`
					ColumnName     string `xml:"columnName"`
					ColumnAlias    string `xml:"columnAlias"`
					ColumnDataType string `xml:"columnDataType"`
				} `xml:"queryResultColumnHeader"`
				QueryResultRow struct {
					Text              string `xml:",chardata"`
					RowIndex          string `xml:"rowIndex"`
					QueryResultColumn []struct {
						Text        string `xml:",chardata"`
						ColumnIndex string `xml:"columnIndex"`
						ColumnValue string `xml:"columnValue"`
					} `xml:"queryResultColumn"`
				} `xml:"queryResultRow"`
			} `xml:"queryResults"`
			HasNextPage   string `xml:"hasNextPage"`
			AbsoluteIndex string `xml:"absoluteIndex"`
			TxnHeader     struct {
				Text          string `xml:",chardata"`
				TransactionID string `xml:"transactionID"`
				UserName      string `xml:"userName"`
				TimeStamp     string `xml:"timeStamp"`
			} `xml:"txnHeader"`
		} `xml:"QueryResultsResponse"`
	} `xml:"Body"`
}

type RetrieveQueryResponseError struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Soapenv string   `xml:"soapenv,attr"`
	Header  string   `xml:"Header"`
	Body    struct {
		Text  string `xml:",chardata"`
		Fault struct {
			Text        string `xml:",chardata"`
			Faultcode   string `xml:"faultcode"`
			Faultstring string `xml:"faultstring"`
			Detail      struct {
				Text                       string `xml:",chardata"`
				CRMCoreWebServiceException struct {
					Text        string `xml:",chardata"`
					Ns1         string `xml:"ns1,attr"`
					Faultcode   string `xml:"faultcode"`
					Faultstring string `xml:"faultstring"`
					TxnHeader   struct {
						Text string `xml:",chardata"`
						Nil  string `xml:"nil,attr"`
						Xsi  string `xml:"xsi,attr"`
					} `xml:"txnHeader"`
				} `xml:"ns1:CRMCoreWebServiceException"`
			} `xml:"detail"`
		} `xml:"Fault"`
	} `xml:"Body"`
}

type QueryFilterAttributes struct {
	Text           string `xml:",chardata"`
	AttributeCode  string `xml:"attributeCode"`
	AttributeValue string `xml:"attributeValue"`
}
