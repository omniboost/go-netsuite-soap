package netsuite

import (
	"encoding/xml"
	"fmt"

	"github.com/cydev/zero"
	"github.com/omniboost/go-netsuite-soap/omitempty"
)

type RequestEnvelope struct {
	XMLName xml.Name
	Version string `xml:"-"`

	Header Header `xml:"env:Header"`
	Body   Body   `xml:"env:Body"`
}

func NewRequestEnvelope(version string) RequestEnvelope {
	return RequestEnvelope{
		Version: version,
		Header:  NewHeader(),
	}
}

type ResponseEnvelope struct {
	XMLName xml.Name

	Header Header `xml:"Header"`
	Body   Body   `xml:"Body"`
}

func (env RequestEnvelope) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name = xml.Name{Local: "env:Envelope"}

	v := env.Version
	ns := func(module, domain string) string {
		return fmt.Sprintf("urn:%s_%s.%s.webservices.netsuite.com", module, v, domain)
	}

	namespaces := []xml.Attr{
		{Name: xml.Name{Space: "", Local: "xmlns:xsd"}, Value: "http://www.w3.org/2001/XMLSchema"},
		{Name: xml.Name{Space: "", Local: "xmlns:xsi"}, Value: "http://www.w3.org/2001/XMLSchema-instance"},
		{Name: xml.Name{Space: "", Local: "xmlns:platformMsgs"}, Value: ns("messages", "platform")},
		{Name: xml.Name{Space: "", Local: "xmlns:env"}, Value: "http://schemas.xmlsoap.org/soap/envelope/"},
		{Name: xml.Name{Space: "", Local: "xmlns:platformCore"}, Value: ns("core", "platform")},
		{Name: xml.Name{Space: "", Local: "xmlns:platformCommon"}, Value: ns("common", "platform")},
		{Name: xml.Name{Space: "", Local: "xmlns:listRel"}, Value: ns("relationships", "lists")},
		{Name: xml.Name{Space: "", Local: "xmlns:tranSales"}, Value: ns("sales", "transactions")},
		{Name: xml.Name{Space: "", Local: "xmlns:tranPurch"}, Value: ns("purchases", "transactions")},
		{Name: xml.Name{Space: "", Local: "xmlns:actSched"}, Value: ns("scheduling", "activities")},
		{Name: xml.Name{Space: "", Local: "xmlns:setupCustom"}, Value: ns("customization", "setup")},
		{Name: xml.Name{Space: "", Local: "xmlns:listAcct"}, Value: ns("accounting", "lists")},
		{Name: xml.Name{Space: "", Local: "xmlns:tranBank"}, Value: ns("bank", "transactions")},
		{Name: xml.Name{Space: "", Local: "xmlns:tranCust"}, Value: ns("customers", "transactions")},
		{Name: xml.Name{Space: "", Local: "xmlns:tranEmp"}, Value: ns("employees", "transactions")},
		{Name: xml.Name{Space: "", Local: "xmlns:tranInvt"}, Value: ns("inventory", "transactions")},
		{Name: xml.Name{Space: "", Local: "xmlns:listSupport"}, Value: ns("support", "lists")},
		{Name: xml.Name{Space: "", Local: "xmlns:tranGeneral"}, Value: ns("general", "transactions")},
		{Name: xml.Name{Space: "", Local: "xmlns:commGeneral"}, Value: ns("communication", "general")},
		{Name: xml.Name{Space: "", Local: "xmlns:listMkt"}, Value: ns("marketing", "lists")},
		{Name: xml.Name{Space: "", Local: "xmlns:listWebsite"}, Value: ns("website", "lists")},
		{Name: xml.Name{Space: "", Local: "xmlns:fileCabinet"}, Value: ns("filecabinet", "documents")},
		{Name: xml.Name{Space: "", Local: "xmlns:listEmp"}, Value: ns("employees", "lists")},
		{Name: xml.Name{Space: "", Local: "xmlns:messages"}, Value: ns("messages", "platform")},
	}
	for _, ns := range namespaces {
		start.Attr = append(start.Attr, ns)
	}

	type alias RequestEnvelope
	a := alias(env)
	return e.EncodeElement(a, start)
}

type Body struct {
	ActionBody interface{} `xml:",any"`
}

type Header struct {
	TokenPassport TokenPassport `xml:"platformMsgs:tokenPassport"`
	// DocumentInfo  struct {
	// 	NSInfo string `xml:"platformMsgs:nsId,omitempty"`
	// } `xml:"platformMsgs:documentInfo,omitempty"`
	Preferences       Preferences       `xml:"platformMsgs:preferences"`
	SearchPreferences SearchPreferences `xml:"platformMsgs:searchPreferences"`
}

func (h Header) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(h, e, start)
}

func (h Header) IsEmpty() bool {
	return zero.IsZero(h)
}

func NewHeader() Header {
	return Header{}
}

type ActionBody interface{}

type TokenPassport struct {
	XMLName     xml.Name `xml:"platformMsgs:tokenPassport"`
	Account     string   `xml:"platformMsgs:account"`
	ConsumerKey string   `xml:"platformMsgs:consumerKey"`
	Token       string   `xml:"platformMsgs:token"`
	Nonce       string   `xml:"platformMsgs:nonce"`
	Timestamp   int64    `xml:"platformMsgs:timestamp"`
	Signature   struct {
		Algorithm string `xml:"algorithm,attr"`
		Text      string `xml:",chardata"`
	} `xml:"platformMsgs:signature"`
}

type Preferences struct {
}

type SearchPreferences struct {
	PageSize            int  `xml:"platformMsgs:pageSize,omitempty"`
	BodyFieldsOnly      bool `xml:"platformMsgs:bodyFieldsOnly"`
	ReturnSearchColumns bool `xml:"platformMsgs:returnSearchColumns"`
}

func (s SearchPreferences) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(s, e, start)
}

func (s SearchPreferences) IsEmpty() bool {
	return zero.IsZero(s)
}
