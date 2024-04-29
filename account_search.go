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

func (c *Client) NewAccountSearchRequest() AccountSearchRequest {
	r := AccountSearchRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type AccountSearchRequest struct {
	client      *Client
	queryParams *AccountSearchRequestQueryParams
	pathParams  *AccountSearchRequestPathParams
	method      string
	headers     http.Header
	requestBody AccountSearchRequestBody
}

func (r AccountSearchRequest) SOAPAction() string {
	return "search"
}

func (r AccountSearchRequest) NewQueryParams() *AccountSearchRequestQueryParams {
	return &AccountSearchRequestQueryParams{}
}

type AccountSearchRequestQueryParams struct {
}

func (p AccountSearchRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *AccountSearchRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r AccountSearchRequest) NewPathParams() *AccountSearchRequestPathParams {
	return &AccountSearchRequestPathParams{}
}

type AccountSearchRequestPathParams struct {
}

func (p *AccountSearchRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *AccountSearchRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *AccountSearchRequest) SetMethod(method string) {
	r.method = method
}

func (r *AccountSearchRequest) Method() string {
	return r.method
}

func (r AccountSearchRequest) NewRequestBody() AccountSearchRequestBody {
	return AccountSearchRequestBody{
		SearchRecord: SearchRecordBasicAccount{
			Type: "listAcct:AccountSearch",
		},
	}
}

type AccountSearchRequestBody struct {
	XMLName xml.Name `xml:"platformMsgs:search"`

	SearchRecord SearchRecordBasicAccount `xml:"platformMsgs:searchRecord"`
}

func (r *AccountSearchRequest) RequestBody() *AccountSearchRequestBody {
	return &r.requestBody
}

func (r *AccountSearchRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *AccountSearchRequest) SetRequestBody(body AccountSearchRequestBody) {
	r.requestBody = body
}

func (r *AccountSearchRequest) NewResponseBody() *AccountSearchRequestResponseBody {
	return &AccountSearchRequestResponseBody{}
}

type AccountSearchRequestResponseBody struct {
	XMLName xml.Name `xml:"searchResponse"`

	SearchResult struct {
		PlatformCore string `xml:"platformCore,attr"`
		TotalRecords string `xml:"totalRecords"`
		PageSize     string `xml:"pageSize"`
		TotalPages   string `xml:"totalPages"`
		PageIndex    string `xml:"pageIndex"`
		SearchId     string `xml:"searchId"`
		RecordList   struct {
			Record Accounts `xml:"record"`
		} `xml:"recordList"`
	} `xml:"searchResult"`
}

func (r *AccountSearchRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("", r.PathParams())
	return &u, err
}

func (r *AccountSearchRequest) Do() (AccountSearchRequestResponseBody, error) {
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

type AccountSearchBasic struct {
	AccountingContext  SearchMultiSelectField     `xml:"accountingContext,omitempty"`
	Balance            SearchDoubleField          `xml:"balance,omitempty"`
	CashFlowRateType   SearchEnumMultiSelectField `xml:"cashFlowRateType,omitempty"`
	Category1099Misc   SearchMultiSelectField     `xml:"category1099Misc,omitempty"`
	Description        SearchStringField          `xml:"description,omitempty"`
	DisplayName        SearchStringField          `xml:"displayName,omitempty"`
	ExternalID         SearchMultiSelectField     `xml:"externalId,omitempty"`
	ExternalIDString   SearchStringField          `xml:"externalIdString,omitempty"`
	GeneralRateType    SearchEnumMultiSelectField `xml:"generalRateType,omitempty"`
	InternalID         SearchMultiSelectField     `xml:"internalId,omitempty"`
	InternalIDNumber   SearchLongField            `xml:"internalIdNumber,omitempty"`
	IsInactive         SearchBooleanField         `xml:"isInactive,omitempty"`
	LegalName          SearchStringField          `xml:"legalName,omitempty"`
	Locale             SearchEnumMultiSelectField `xml:"locale,omitempty"`
	LocalizedLegalName SearchStringField          `xml:"localizedLegalName,omitempty"`
	LocalizedName      SearchStringField          `xml:"localizedName,omitempty"`
	LocalizedNumber    SearchStringField          `xml:"localizedNumber,omitempty"`
	Name               SearchStringField          `xml:"name,omitempty"`
	Number             SearchStringField          `xml:"number,omitempty"`
	Parent             SearchMultiSelectField     `xml:"parent,omitempty"`
	Subsidiary         SearchMultiSelectField     `xml:"subsidiary,omitempty"`
	Type               SearchEnumMultiSelectField `xml:"type,omitempty"`
	CustomFieldList    SearchCustomFieldList      `xml:"customFieldList,omitempty"`
	TranID             SearchStringField          `xml:"tranId,omitempty"`
}

func (c AccountSearchBasic) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(c, e, start)
}

func (c AccountSearchBasic) IsEmpty() bool {
	return zero.IsZero(c)
}
