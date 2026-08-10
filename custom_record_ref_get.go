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

func (c *Client) NewCustomRecordRefGetRequest() CustomRecordRefGetRequest {
	r := CustomRecordRefGetRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CustomRecordRefGetRequest struct {
	client      *Client
	queryParams *CustomRecordRefGetRequestQueryParams
	pathParams  *CustomRecordRefGetRequestPathParams
	method      string
	headers     http.Header
	requestBody CustomRecordRefGetRequestBody
}

func (r CustomRecordRefGetRequest) SOAPAction() string {
	return "get"
}

func (r CustomRecordRefGetRequest) NewQueryParams() *CustomRecordRefGetRequestQueryParams {
	return &CustomRecordRefGetRequestQueryParams{}
}

type CustomRecordRefGetRequestQueryParams struct {
}

func (p CustomRecordRefGetRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *CustomRecordRefGetRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r CustomRecordRefGetRequest) NewPathParams() *CustomRecordRefGetRequestPathParams {
	return &CustomRecordRefGetRequestPathParams{}
}

type CustomRecordRefGetRequestPathParams struct {
}

func (p *CustomRecordRefGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *CustomRecordRefGetRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *CustomRecordRefGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *CustomRecordRefGetRequest) Method() string {
	return r.method
}

func (r CustomRecordRefGetRequest) NewRequestBody() CustomRecordRefGetRequestBody {
	return CustomRecordRefGetRequestBody{}
}

type CustomRecordRefGetRequestBody struct {
	XMLName xml.Name `xml:"get"`

	Xmlns   string `xml:"xmlns,attr"`
	BaseRef struct {
		ScriptID   string `xml:"scriptId,attr,omitempty"`
		ExternalID string `xml:"externalId,attr,omitempty"`
		InternalID string `xml:"internalId,attr,omitempty"`
		XSIType    string `xml:"xsi:type,attr"`
		Type       string `xml:"type,attr,omitempty"`
		TypeID     string `xml:"typeId,attr,omitempty"`
		Xmlns      string `xml:"xmlns:q1,attr"`
	} `xml:"baseRef"`
}

func (r *CustomRecordRefGetRequest) RequestBody() *CustomRecordRefGetRequestBody {
	return &r.requestBody
}

func (r *CustomRecordRefGetRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *CustomRecordRefGetRequest) SetRequestBody(body CustomRecordRefGetRequestBody) {
	r.requestBody = body
}

func (r *CustomRecordRefGetRequest) NewResponseBody() *CustomRecordRefGetRequestResponseBody {
	return &CustomRecordRefGetRequestResponseBody{}
}

type CustomRecordRefGetRequestResponseBody struct {
	XMLName xml.Name `xml:"getResponse"`

	// Text         string `xml:",chardata"`
	// Xmlns        string `xml:"xmlns,attr"`
	ReadResponse struct {
		// Text         string `xml:",chardata"`
		PlatformMsgs string `xml:"platformMsgs,attr"`
		Status       struct {
			// Text         string `xml:",chardata"`
			IsSuccess    string `xml:"isSuccess,attr"`
			PlatformCore string `xml:"platformCore,attr"`
		} `xml:"status"`
		Record struct {
			// Text           string `xml:",chardata"`
			InternalID     string     `xml:"internalId,attr"`
			Type           string     `xml:"type,attr"`
			SetupCustom    string     `xml:"setupCustom,attr"`
			Label          string     `xml:"label"`
			ScriptID       string     `xml:"scriptId"`
			RecordScriptID string     `xml:"recordScriptId"`
			RecordType     RecordType `xml:"recordType"`
			FieldType      string     `xml:"fieldType"`
			IsInactive     string     `xml:"isInactive"`
			ShowInList     string     `xml:"showInList"`
			HasGLImpact    string     `xml:"hasGLImpact"`
			IsMandatory    string     `xml:"isMandatory"`
		} `xml:"record"`
	} `xml:"readResponse"`
}

func (r *CustomRecordRefGetRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("", r.PathParams())
	return &u, err
}

func (r *CustomRecordRefGetRequest) Do() (CustomRecordRefGetRequestResponseBody, error) {
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

type CustomRecordRefGetBasic struct {
	ExternalID       SearchMultiSelectField `xml:"externalId,omitempty"`
	ExternalIDString SearchStringField      `xml:"externalIdString,omitempty"`
	InternalID       SearchMultiSelectField `xml:"internalId,omitempty"`
	InternalIDNumber SearchLongField        `xml:"internalIdNumber,omitempty"`
	IsInactive       SearchBooleanField     `xml:"isInactive,omitempty"`
	Name             SearchStringField      `xml:"name,omitempty"`
	NameNoHierarchy  SearchStringField      `xml:"nameNoHierarchy,omitempty"`
	Subsidiary       SearchMultiSelectField `xml:"subsidiary,omitempty"`
	CustomFieldList  SearchCustomFieldList  `xml:"platformCommon:customFieldList,omitempty"`
}

func (c CustomRecordRefGetBasic) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(c, e, start)
}

func (c CustomRecordRefGetBasic) IsEmpty() bool {
	return zero.IsZero(c)
}
