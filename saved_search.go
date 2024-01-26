package netsuite

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/cydev/zero"
	"github.com/omniboost/go-netsuite-soap/utils"

	"github.com/pkg/errors"
)

func (c *Client) NewSavedSearchRequest() SavedSearchRequest {
	r := SavedSearchRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type SavedSearchRequest struct {
	client      *Client
	queryParams *SearchRequestQueryParams
	pathParams  *SearchRequestPathParams
	method      string
	headers     http.Header
	requestBody SavedSearchRequestBody
}

func (r SavedSearchRequest) SOAPAction() string {
	return "search"
}

func (r SavedSearchRequest) NewQueryParams() *SearchRequestQueryParams {
	return &SearchRequestQueryParams{}
}

type SavedSearchRequestQueryParams struct {
}

func (p SavedSearchRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *SavedSearchRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r SavedSearchRequest) NewPathParams() *SearchRequestPathParams {
	return &SearchRequestPathParams{}
}

type SavedSearchRequestPathParams struct {
}

func (p *SavedSearchRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *SavedSearchRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *SavedSearchRequest) SetMethod(method string) {
	r.method = method
}

func (r *SavedSearchRequest) Method() string {
	return r.method
}

func (r SavedSearchRequest) NewRequestBody() SavedSearchRequestBody {
	return SavedSearchRequestBody{}
}

type SavedSearchRequestBody struct {
	XMLName xml.Name `xml:"platformMsgs:search"`

	SearchRecord struct { // SearchRecordBasic
		Type          string `xml:"xsi:type,attr"`
		SavedSearchId string `xml:"savedSearchId,attr"`
		Criteria      struct {
			XMLName xml.Name `xml:"criteria"`
			Basic   struct {
				XMLName             xml.Name `xml:"basic"`
				SavedSearchCriteria []SavedSearchCriteria
			}
		}
	} `xml:"platformMsgs:searchRecord"`
}

type SavedSearchCriteria struct {
	XMLName     xml.Name
	Operator    string   `xml:"operator,attr"`
	SearchValue []string `xml:"searchValue"`
}

func (r *SavedSearchRequest) RequestBody() *SavedSearchRequestBody {
	return &r.requestBody
}

func (r *SavedSearchRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *SavedSearchRequest) SetRequestBody(body SavedSearchRequestBody) {
	r.requestBody = body
}

func (r *SavedSearchRequest) NewResponseBody() *SearchRequestResponseBody {
	return &SearchRequestResponseBody{}
}

func (r *SavedSearchRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("", r.PathParams())
	return &u, err
}

func (r *SavedSearchRequest) Do() (SearchRequestResponseBody, error) {
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

type SavedSearchStringField struct {
	Operator    SearchStringFieldOperator `xml:"operator,attr"`
	SearchValue string                    `xml:"platformCore:searchValue"`
}

func (s SavedSearchStringField) IsEmpty() bool {
	return zero.IsZero(s)
}

type SavedSearchStringFieldOperator string

type SavedSearchMultiSelectField struct {
	Operator    SearchStringFieldOperator `xml:"platformCore:operator,attr"`
	SearchValue []RecordRef               `xml:"platformCore:searchValue"`
}

func (s SavedSearchMultiSelectField) IsEmpty() bool {
	return zero.IsZero(s)
}

type SavedSearchDoubleField struct{}

func (s SavedSearchDoubleField) IsEmpty() bool {
	return zero.IsZero(s)
}

type SavedSearchEnumMultiSelectField struct {
	Operator    SearchStringFieldOperator `xml:"operator,attr"`
	SearchValue string                    `xml:"platformCore:searchValue"`
}

func (s SavedSearchEnumMultiSelectField) IsEmpty() bool {
	return zero.IsZero(s)
}

type SavedSearchLongField struct {
	Operator     SearchLongFieldOperator `xml:"platformCore:operator,attr"`
	SearchValue  int                     `xml:"platformCore:searchValue"`
	SearchValue2 int                     `xml:"platformCore:searchValue2,omitempty"`
}

func (s SavedSearchLongField) IsEmpty() bool {
	return zero.IsZero(s)
}

type SavedSearchBooleanField struct{}

func (s SavedSearchBooleanField) IsEmpty() bool {
	return zero.IsZero(s)
}

type SavedSearchCustomFieldList []interface{}

func (s SavedSearchCustomFieldList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	alias := struct {
		SearchFields []interface{}
	}{SearchFields: s}

	return e.EncodeElement(alias, start)
}

func (s SavedSearchCustomFieldList) IsEmpty() bool {
	return zero.IsZero(s)
}

type SavedSearchStringCustomField struct {
	XMLName xml.Name `xml:"platformCore:customField"`

	// Type        string                    `xml:"xsi:type,attr"`
	ScriptID    string                    `xml:"platformCore:scriptId,attr,omitempty"`
	InternalID  string                    `xml:"platformCore:internalId,attr,omitempty"`
	Operator    SearchStringFieldOperator `xml:"platformCore:operator,attr"`
	SearchValue string                    `xml:"platformCore:searchValue"`
}

func (s SavedSearchStringCustomField) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type alias SearchStringCustomField

	start.Name = xml.Name{Local: "platformCore:customField"}
	s2 := struct {
		Type string `xml:"xsi:type,attr"`
		alias
	}{Type: "platformCore:SearchStringCustomField", alias: alias(s)}

	return e.EncodeElement(s2, start)
}

type SavedSearchRecordBasic struct { // SearchRecordBasic
	XMLName xml.Name `xml:"platformMsgs:searchRecord"`

	Type  string             `xml:"xsi:type,attr"`
	Basic AccountSearchBasic `xml:"basic"`
}

type SavedSearchDateField struct{}

func (s SavedSearchDateField) IsEmpty() bool {
	return zero.IsZero(s)
}

type SavedSalesTaxItemSearchRecordBasic struct { // SearchRecordBasic
	XMLName xml.Name `xml:"platformMsgs:searchRecord"`

	Type  string                  `xml:"xsi:type,attr"`
	Basic SalesTaxItemSearchBasic `xml:"basic"`
}

type SavedSearchTextNumberField struct {
	// Operator    SearchStringFieldOperator `xml:"operator,attr"`
	// SearchValue string                    `xml:"platformCore:searchValue"`
}

func (s SavedSearchTextNumberField) IsEmpty() bool {
	return zero.IsZero(s)
}

type SavedSearchLongFieldOperator string
