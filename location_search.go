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

func (c *Client) NewLocationSearchRequest() LocationSearchRequest {
	r := LocationSearchRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type LocationSearchRequest struct {
	client      *Client
	queryParams *LocationSearchRequestQueryParams
	pathParams  *LocationSearchRequestPathParams
	method      string
	headers     http.Header
	requestBody LocationSearchRequestBody
}

func (r LocationSearchRequest) SOAPAction() string {
	return "search"
}

func (r LocationSearchRequest) NewQueryParams() *LocationSearchRequestQueryParams {
	return &LocationSearchRequestQueryParams{}
}

type LocationSearchRequestQueryParams struct {
}

func (p LocationSearchRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *LocationSearchRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r LocationSearchRequest) NewPathParams() *LocationSearchRequestPathParams {
	return &LocationSearchRequestPathParams{}
}

type LocationSearchRequestPathParams struct {
}

func (p *LocationSearchRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *LocationSearchRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *LocationSearchRequest) SetMethod(method string) {
	r.method = method
}

func (r *LocationSearchRequest) Method() string {
	return r.method
}

func (r LocationSearchRequest) NewRequestBody() LocationSearchRequestBody {
	return LocationSearchRequestBody{
		SearchRecord: SearchRecordBasicLocation{
			Type: "listAcct:LocationSearch",
		},
	}
}

type LocationSearchRequestBody struct {
	XMLName xml.Name `xml:"platformMsgs:search"`

	SearchRecord SearchRecordBasicLocation `xml:"platformMsgs:searchRecord"`
}

func (r *LocationSearchRequest) RequestBody() *LocationSearchRequestBody {
	return &r.requestBody
}

func (r *LocationSearchRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *LocationSearchRequest) SetRequestBody(body LocationSearchRequestBody) {
	r.requestBody = body
}

func (r *LocationSearchRequest) NewResponseBody() *LocationSearchRequestResponseBody {
	return &LocationSearchRequestResponseBody{}
}

type LocationSearchRequestResponseBody struct {
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
		SearchId     string `xml:"searchId"`
		RecordList   struct {
			Record Locations `xml:"record"`
		} `xml:"recordList"`
	} `xml:"searchResult"`
}

func (r *LocationSearchRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("", r.PathParams())
	return &u, err
}

func (r *LocationSearchRequest) Do() (LocationSearchRequestResponseBody, error) {
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

type LocationSearchBasic struct {
	ExternalId       SearchMultiSelectField `xml:"externalId,omitempty"`
	ExternalIdString SearchStringField      `xml:"externalIdString,omitempty"`
	InternalId       SearchMultiSelectField `xml:"internalId,omitempty"`
	InternalIdNumber SearchLongField        `xml:"internalIdNumber,omitempty"`
	IsInactive       SearchBooleanField     `xml:"isInactive,omitempty"`
	Name             SearchStringField      `xml:"name,omitempty"`
	NameNoHierarchy  SearchStringField      `xml:"nameNoHierarchy,omitempty"`
	Subsidiary       SearchMultiSelectField `xml:"subsidiary,omitempty"`
	CustomFieldList  SearchCustomFieldList  `xml:"platformCommon:customFieldList,omitempty"`
}

func (c LocationSearchBasic) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(c, e, start)
}

func (c LocationSearchBasic) IsEmpty() bool {
	return zero.IsZero(c)
}

