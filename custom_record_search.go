package netsuite

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/cydev/zero"
	"github.com/omniboost/go-netsuite-soap/omitempty"
	"github.com/omniboost/go-netsuite-soap/utils"
	"github.com/pkg/errors"
)

func (c *Client) NewCustomRecordSearchRequest() CustomRecordSearchRequest {
	r := CustomRecordSearchRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CustomRecordSearchRequest struct {
	client      *Client
	queryParams *CustomRecordSearchRequestQueryParams
	pathParams  *CustomRecordSearchRequestPathParams
	method      string
	headers     http.Header
	requestBody CustomRecordSearchRequestBody
}

func (r CustomRecordSearchRequest) SOAPAction() string {
	return "search"
}

func (r CustomRecordSearchRequest) NewQueryParams() *CustomRecordSearchRequestQueryParams {
	return &CustomRecordSearchRequestQueryParams{}
}

type CustomRecordSearchRequestQueryParams struct {
}

func (p CustomRecordSearchRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *CustomRecordSearchRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r CustomRecordSearchRequest) NewPathParams() *CustomRecordSearchRequestPathParams {
	return &CustomRecordSearchRequestPathParams{}
}

type CustomRecordSearchRequestPathParams struct {
}

func (p *CustomRecordSearchRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *CustomRecordSearchRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *CustomRecordSearchRequest) SetMethod(method string) {
	r.method = method
}

func (r *CustomRecordSearchRequest) Method() string {
	return r.method
}

func (r CustomRecordSearchRequest) NewRequestBody() CustomRecordSearchRequestBody {
	body := CustomRecordSearchRequestBody{}
	body.SearchRecord.Type = "setupCustom:CustomRecordSearch"
	return body
}

type CustomRecordSearchRequestBody struct {
	XMLName xml.Name `xml:"platformMsgs:search"`

	SearchRecord struct {
		XMLName xml.Name `xml:"platformMsgs:searchRecord"`

		Type  string                  `xml:"xsi:type,attr"`
		Basic CustomRecordSearchBasic `xml:"setupCustom:basic"`
	} `xml:"platformMsgs:searchRecord"`
}

func (r *CustomRecordSearchRequest) RequestBody() *CustomRecordSearchRequestBody {
	return &r.requestBody
}

func (r *CustomRecordSearchRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *CustomRecordSearchRequest) SetRequestBody(body CustomRecordSearchRequestBody) {
	r.requestBody = body
}

func (r *CustomRecordSearchRequest) NewResponseBody() *CustomRecordSearchRequestResponseBody {
	return &CustomRecordSearchRequestResponseBody{}
}

type CustomRecordSearchRequestResponseBody struct {
	XMLName xml.Name `xml:"searchResponse"`

	SearchResult struct {
		PlatformCore string `xml:"platformCore,attr"`
		Status       struct {
			IsSuccess string `xml:"isSuccess,attr"`
		} `xml:"status"`
		TotalRecords string `xml:"totalRecords"`
		PageSize     string `xml:"pageSize"`
		TotalPages   string `xml:"totalPages"`
		PageIndex    string `xml:"pageIndex"`
		SearchID     string `xml:"searchId"`
		RecordList   struct {
			Record CustomRecords `xml:"record"`
		} `xml:"recordList"`
	} `xml:"searchResult"`
}

func (r *CustomRecordSearchRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("", r.PathParams())
	return &u, err
}

func (r *CustomRecordSearchRequest) Do() (CustomRecordSearchRequestResponseBody, error) {
	var err error

	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	if err != nil {
		return *responseBody, errors.WithStack(err)
	}

	return *responseBody, nil
}

type CustomRecordSearchBasic struct {
	RecType RecordRef `xml:"platformCommon:recType,omitempty"`
}

func (c CustomRecordSearchBasic) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(c, e, start)
}

func (c CustomRecordSearchBasic) IsEmpty() bool {
	return zero.IsZero(c)
}
