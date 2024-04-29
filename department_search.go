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

func (c *Client) NewDepartmentSearchRequest() DepartmentSearchRequest {
	r := DepartmentSearchRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type DepartmentSearchRequest struct {
	client      *Client
	queryParams *DepartmentSearchRequestQueryParams
	pathParams  *DepartmentSearchRequestPathParams
	method      string
	headers     http.Header
	requestBody DepartmentSearchRequestBody
}

func (r DepartmentSearchRequest) SOAPAction() string {
	return "search"
}

func (r DepartmentSearchRequest) NewQueryParams() *DepartmentSearchRequestQueryParams {
	return &DepartmentSearchRequestQueryParams{}
}

type DepartmentSearchRequestQueryParams struct {
}

func (p DepartmentSearchRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *DepartmentSearchRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r DepartmentSearchRequest) NewPathParams() *DepartmentSearchRequestPathParams {
	return &DepartmentSearchRequestPathParams{}
}

type DepartmentSearchRequestPathParams struct {
}

func (p *DepartmentSearchRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *DepartmentSearchRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *DepartmentSearchRequest) SetMethod(method string) {
	r.method = method
}

func (r *DepartmentSearchRequest) Method() string {
	return r.method
}

func (r DepartmentSearchRequest) NewRequestBody() DepartmentSearchRequestBody {
	return DepartmentSearchRequestBody{
		SearchRecord: SearchRecordBasicDepartment{
			Type: "listAcct:DepartmentSearch",
		},
	}
}

type DepartmentSearchRequestBody struct {
	XMLName xml.Name `xml:"platformMsgs:search"`

	SearchRecord SearchRecordBasicDepartment `xml:"platformMsgs:searchRecord"`
}

func (r *DepartmentSearchRequest) RequestBody() *DepartmentSearchRequestBody {
	return &r.requestBody
}

func (r *DepartmentSearchRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *DepartmentSearchRequest) SetRequestBody(body DepartmentSearchRequestBody) {
	r.requestBody = body
}

func (r *DepartmentSearchRequest) NewResponseBody() *DepartmentSearchRequestResponseBody {
	return &DepartmentSearchRequestResponseBody{}
}

type DepartmentSearchRequestResponseBody struct {
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
			Record Departments `xml:"record"`
		} `xml:"recordList"`
	} `xml:"searchResult"`
}

func (r *DepartmentSearchRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("", r.PathParams())
	return &u, err
}

func (r *DepartmentSearchRequest) Do() (DepartmentSearchRequestResponseBody, error) {
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

type DepartmentSearchBasic struct {
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

func (c DepartmentSearchBasic) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(c, e, start)
}

func (c DepartmentSearchBasic) IsEmpty() bool {
	return zero.IsZero(c)
}
