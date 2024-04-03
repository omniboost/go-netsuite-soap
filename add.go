package netsuite

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-netsuite-soap/utils"
	"github.com/pkg/errors"
)

func (c *Client) NewAddRequest() AddRequest {
	r := AddRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type AddRequest struct {
	client      *Client
	queryParams *AddRequestQueryParams
	pathParams  *AddRequestPathParams
	method      string
	headers     http.Header
	requestBody AddRequestBody
}

func (r AddRequest) SOAPAction() string {
	return "add"
}

func (r AddRequest) NewQueryParams() *AddRequestQueryParams {
	return &AddRequestQueryParams{}
}

type AddRequestQueryParams struct {
}

func (p AddRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *AddRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r AddRequest) NewPathParams() *AddRequestPathParams {
	return &AddRequestPathParams{}
}

type AddRequestPathParams struct {
}

func (p *AddRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *AddRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *AddRequest) SetMethod(method string) {
	r.method = method
}

func (r *AddRequest) Method() string {
	return r.method
}

func (r AddRequest) NewRequestBody() AddRequestBody {
	return AddRequestBody{}
}

type AddRequestBody struct {
	XMLName xml.Name `xml:"platformMsgs:add"`

	Record Record `xml:"platformMsgs:record"`
}

func (r *AddRequest) RequestBody() *AddRequestBody {
	return &r.requestBody
}

func (r *AddRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *AddRequest) SetRequestBody(body AddRequestBody) {
	r.requestBody = body
}

func (r *AddRequest) NewResponseBody() *AddRequestResponseBody {
	return &AddRequestResponseBody{}
}

type AddRequestResponseBody struct {
	XMLName xml.Name `xml:"addResponse"`

	AddResult struct {
		Text         string `xml:",chardata"`
		PlatformCore string `xml:"platformCore,attr"`
		Status       struct {
			Text      string `xml:",chardata"`
			IsSuccess string `xml:"isSuccess,attr"`
		} `xml:"status"`
		TotalRecords string `xml:"totalRecords"`
		RecordList   struct {
			Record []interface{} `xml:"recordList>record"`
		} `xml:"recordList"`
	} `xml:"addResult"`

	WriteResponse struct {
		Status       struct {
			Text      string `xml:",chardata"`
			IsSuccess string `xml:"isSuccess,attr"`
		} `xml:"status"`
		BaseRef struct {
			InternalID string `xml:"internalId,attr"`
			ExternalID string `xml:"externalId,attr"`
			Type       string `xml:"type,attr"`
			XSIType    string `xml:"xsi:type,attr"`
		} `xml:"baseRef"`
	} `xml:"writeResponse"`
}

func (r *AddRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("", r.PathParams())
	return &u, err
}

func (r *AddRequest) Do() (AddRequestResponseBody, error) {
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
