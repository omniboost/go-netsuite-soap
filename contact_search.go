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

func (c *Client) NewContactSearchRequest() ContactSearchRequest {
	r := ContactSearchRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type ContactSearchRequest struct {
	client      *Client
	queryParams *ContactSearchRequestQueryParams
	pathParams  *ContactSearchRequestPathParams
	method      string
	headers     http.Header
	requestBody ContactSearchRequestBody
}

func (r ContactSearchRequest) SOAPAction() string {
	return "search"
}

func (r ContactSearchRequest) NewQueryParams() *ContactSearchRequestQueryParams {
	return &ContactSearchRequestQueryParams{}
}

type ContactSearchRequestQueryParams struct {
}

func (p ContactSearchRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *ContactSearchRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r ContactSearchRequest) NewPathParams() *ContactSearchRequestPathParams {
	return &ContactSearchRequestPathParams{}
}

type ContactSearchRequestPathParams struct {
}

func (p *ContactSearchRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *ContactSearchRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *ContactSearchRequest) SetMethod(method string) {
	r.method = method
}

func (r *ContactSearchRequest) Method() string {
	return r.method
}

func (r ContactSearchRequest) NewRequestBody() ContactSearchRequestBody {
	body := ContactSearchRequestBody{}
	body.SearchRecord.Type = "listRel:ContactSearch"
	return body
}

type ContactSearchRequestBody struct {
	XMLName xml.Name `xml:"platformMsgs:search"`

	SearchRecord struct {
		XMLName xml.Name `xml:"platformMsgs:searchRecord"`

		Type  string             `xml:"xsi:type,attr"`
		Basic ContactSearchBasic `xml:"listRel:basic"`
	} `xml:"platformMsgs:searchRecord"`
}

func (r *ContactSearchRequest) RequestBody() *ContactSearchRequestBody {
	return &r.requestBody
}

func (r *ContactSearchRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *ContactSearchRequest) SetRequestBody(body ContactSearchRequestBody) {
	r.requestBody = body
}

func (r *ContactSearchRequest) NewResponseBody() *ContactSearchRequestResponseBody {
	return &ContactSearchRequestResponseBody{}
}

type ContactSearchRequestResponseBody struct {
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
			Record Contacts `xml:"record"`
		} `xml:"recordList"`
	} `xml:"searchResult"`
}

func (r *ContactSearchRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("", r.PathParams())
	return &u, err
}

func (r *ContactSearchRequest) Do() (ContactSearchRequestResponseBody, error) {
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

type ContactSearchBasic struct {
	Address                  SearchStringField          `xml:"platformCommon:address,omitempty"`
	Addressee                SearchStringField          `xml:"platformCommon:addressee,omitempty"`
	AddressLabel             SearchStringField          `xml:"platformCommon:addressLabel,omitempty"`
	AddressPhone             SearchStringField          `xml:"platformCommon:addressPhone,omitempty"`
	Attention                SearchStringField          `xml:"platformCommon:attention,omitempty"`
	AvailableOffline         SearchBooleanField         `xml:"platformCommon:availableOffline,omitempty"`
	Category                 SearchMultiSelectField     `xml:"platformCommon:category,omitempty"`
	City                     SearchStringField          `xml:"platformCommon:city,omitempty"`
	Comments                 SearchStringField          `xml:"platformCommon:comments,omitempty"`
	Company                  SearchMultiSelectField     `xml:"platformCommon:company,omitempty"`
	ContactRole              SearchMultiSelectField     `xml:"platformCommon:contactRole,omitempty"`
	ContactSource            SearchMultiSelectField     `xml:"platformCommon:contactSource,omitempty"`
	Country                  SearchEnumMultiSelectField `xml:"platformCommon:country,omitempty"`
	County                   SearchStringField          `xml:"platformCommon:county,omitempty"`
	DateCreated              SearchDateField            `xml:"platformCommon:dateCreated,omitempty"`
	Email                    SearchStringField          `xml:"platformCommon:email,omitempty"`
	Employer                 SearchStringField          `xml:"platformCommon:employer,omitempty"`
	EntityID                 SearchStringField          `xml:"platformCommon:entityId,omitempty"`
	ExternalID               SearchMultiSelectField     `xml:"platformCommon:externalId,omitempty"`
	ExternalIDString         SearchStringField          `xml:"platformCommon:externalIdString,omitempty"`
	Fax                      SearchStringField          `xml:"platformCommon:fax,omitempty"`
	FirstName                SearchStringField          `xml:"platformCommon:firstName,omitempty"`
	GiveAccess               SearchBooleanField         `xml:"platformCommon:giveAccess,omitempty"`
	GlobalSubscriptionStatus SearchEnumMultiSelectField `xml:"platformCommon:globalSubscriptionStatus,omitempty"`
	Group                    SearchMultiSelectField     `xml:"platformCommon:group,omitempty"`
	HasDuplicates            SearchBooleanField         `xml:"platformCommon:hasDuplicates,omitempty"`
	Image                    SearchStringField          `xml:"platformCommon:image,omitempty"`
	InternalID               SearchMultiSelectField     `xml:"platformCommon:internalId,omitempty"`
	InternalIDNumber         SearchLongField            `xml:"platformCommon:internalIdNumber,omitempty"`
	IsDefaultBilling         SearchBooleanField         `xml:"platformCommon:isDefaultBilling,omitempty"`
	IsDefaultShipping        SearchBooleanField         `xml:"platformCommon:isDefaultShipping,omitempty"`
	IsInactive               SearchBooleanField         `xml:"platformCommon:isInactive,omitempty"`
	IsPrivate                SearchBooleanField         `xml:"platformCommon:isPrivate,omitempty"`
	Language                 SearchEnumMultiSelectField `xml:"platformCommon:language,omitempty"`
	LastModifiedDate         SearchDateField            `xml:"platformCommon:lastModifiedDate,omitempty"`
	LastName                 SearchStringField          `xml:"platformCommon:lastName,omitempty"`
	Level                    SearchEnumMultiSelectField `xml:"platformCommon:level,omitempty"`
	MiddleName               SearchStringField          `xml:"platformCommon:middleName,omitempty"`
	Owner                    SearchMultiSelectField     `xml:"platformCommon:owner,omitempty"`
	Permission               SearchEnumMultiSelectField `xml:"platformCommon:permission,omitempty"`
	Phone                    SearchStringField          `xml:"platformCommon:phone,omitempty"`
	PhoneticName             SearchStringField          `xml:"platformCommon:phoneticName,omitempty"`
	Salutation               SearchStringField          `xml:"platformCommon:salutation,omitempty"`
	State                    SearchStringField          `xml:"platformCommon:state,omitempty"`
	Subsidiary               SearchMultiSelectField     `xml:"platformCommon:subsidiary,omitempty"`
	Title                    SearchStringField          `xml:"platformCommon:title,omitempty"`
	Type                     SearchEnumMultiSelectField `xml:"platformCommon:type,omitempty"`
	ZipCode                  SearchStringField          `xml:"platformCommon:zipCode,omitempty"`
	CustomFieldList          SearchCustomFieldList      `xml:"platformCommon:customFieldList,omitempty"`
}

func (c ContactSearchBasic) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(c, e, start)
}

func (c ContactSearchBasic) IsEmpty() bool {
	return zero.IsZero(c)
}
