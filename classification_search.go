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

func (c *Client) NewClassificationSearchRequest() ClassificationSearchRequest {
	r := ClassificationSearchRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type ClassificationSearchRequest struct {
	client      *Client
	queryParams *ClassificationSearchRequestQueryParams
	pathParams  *ClassificationSearchRequestPathParams
	method      string
	headers     http.Header
	requestBody ClassificationSearchRequestBody
}

func (r ClassificationSearchRequest) SOAPAction() string {
	return "search"
}

func (r ClassificationSearchRequest) NewQueryParams() *ClassificationSearchRequestQueryParams {
	return &ClassificationSearchRequestQueryParams{}
}

type ClassificationSearchRequestQueryParams struct {
}

func (p ClassificationSearchRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *ClassificationSearchRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r ClassificationSearchRequest) NewPathParams() *ClassificationSearchRequestPathParams {
	return &ClassificationSearchRequestPathParams{}
}

type ClassificationSearchRequestPathParams struct {
}

func (p *ClassificationSearchRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *ClassificationSearchRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *ClassificationSearchRequest) SetMethod(method string) {
	r.method = method
}

func (r *ClassificationSearchRequest) Method() string {
	return r.method
}

func (r ClassificationSearchRequest) NewRequestBody() ClassificationSearchRequestBody {
	return ClassificationSearchRequestBody{
		SearchRecord: SearchRecordBasicClassification{
			Type: "listAcct:ClassificationSearch",
		},
	}
}

type ClassificationSearchRequestBody struct {
	XMLName xml.Name `xml:"platformMsgs:search"`

	SearchRecord SearchRecordBasicClassification `xml:"platformMsgs:searchRecord"`
}

func (r *ClassificationSearchRequest) RequestBody() *ClassificationSearchRequestBody {
	return &r.requestBody
}

func (r *ClassificationSearchRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *ClassificationSearchRequest) SetRequestBody(body ClassificationSearchRequestBody) {
	r.requestBody = body
}

func (r *ClassificationSearchRequest) NewResponseBody() *ClassificationSearchRequestResponseBody {
	return &ClassificationSearchRequestResponseBody{}
}

type ClassificationSearchRequestResponseBody struct {
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
			Record Classifications `xml:"record"`
		} `xml:"recordList"`
	} `xml:"searchResult"`
}

func (r *ClassificationSearchRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("", r.PathParams())
	return &u, err
}

func (r *ClassificationSearchRequest) Do() (ClassificationSearchRequestResponseBody, error) {
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

type ClassificationSearchBasic struct {
	ExternalId       SearchMultiSelectField `xml:"externalId,omitempty"`
	ExternalIdString SearchStringField      `xml:"externalIdString,omitempty"`
	InternalId       SearchMultiSelectField `xml:"internalId,omitempty"`
	InternalIdNumber SearchLongField        `xml:"internalIdNumber,omitempty"`
	IsInactive       SearchBooleanField     `xml:"isInactive,omitempty"`
	Name             SearchStringField      `xml:"name,omitempty"`
	NameNoHierarchy  SearchStringField      `xml:"nameNoHierarchy,omitempty"`
	Subsidiary       SearchMultiSelectField `xml:"subsidiary,omitempty"`
	CustomFieldList  SearchCustomFieldList  `xml:"customFieldList,omitempty"`
}

func (c ClassificationSearchBasic) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(c, e, start)
}

func (c ClassificationSearchBasic) IsEmpty() bool {
	return zero.IsZero(c)
}
