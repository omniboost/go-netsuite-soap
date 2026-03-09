package netsuite

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path"
	"strings"
	"text/template"
	"time"

	"github.com/omniboost/go-httperr"
	"github.com/pkg/errors"
)

const (
	libraryVersion = "0.0.1"
	userAgent      = "go-netsuite-soap/" + libraryVersion
	mediaType      = "text/xml"
	charset        = "utf-8"
)

var (
	BaseURL string = "https://webservices.netsuite.com/services/NetSuitePort_2022_2"
)

// NewClient returns a new Exact Globe Client client
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	client := &Client{}

	client.SetHTTPClient(httpClient)
	client.SetBaseURL(BaseURL)
	client.SetDebug(false)
	client.SetUserAgent(userAgent)
	client.SetMediaType(mediaType)
	client.SetCharset(charset)
	client.SearchPreferences = &SearchPreferences{}

	return client
}

// Client manages communication with Exact Globe Client
type Client struct {
	// HTTP client used to communicate with the Client.
	http *http.Client

	SearchPreferences *SearchPreferences

	debug   bool
	baseURL string

	// credentials
	// applicationID string
	clientID     string
	clientSecret string
	tokenID      string
	tokenSecret  string
	accountID    string

	// User agent for client
	userAgent string

	mediaType string
	charset   string

	// Optional function called after every successful request made to the DO Clients
	beforeRequestDo    BeforeRequestDoCallback
	onRequestCompleted RequestCompletionCallback
}

type BeforeRequestDoCallback func(*http.Client, *http.Request, interface{})

// RequestCompletionCallback defines the type of the request callback function
type RequestCompletionCallback func(*http.Request, *http.Response)

func (c *Client) SetHTTPClient(client *http.Client) {
	c.http = client
}

func (c Client) Debug() bool {
	return c.debug
}

func (c *Client) SetDebug(debug bool) {
	c.debug = debug
}

// func (c Client) ApplicationID() string {
// 	return c.applicationID
// }

// func (c *Client) SetApplicationID(applicationID string) {
// 	c.applicationID = applicationID
// }

func (c Client) ClientID() string {
	return c.clientID
}

func (c *Client) SetClientID(clientID string) {
	c.clientID = clientID
}

func (c Client) ClientSecret() string {
	return c.clientSecret
}

func (c *Client) SetClientSecret(clientSecret string) {
	c.clientSecret = clientSecret
}

func (c Client) TokenID() string {
	return c.tokenID
}

func (c *Client) SetTokenID(tokenID string) {
	c.tokenID = tokenID
}

func (c Client) TokenSecret() string {
	return c.tokenSecret
}

func (c *Client) SetTokenSecret(tokenSecret string) {
	c.tokenSecret = tokenSecret
}

func (c Client) AccountID() string {
	return c.accountID
}

func (c *Client) SetAccountID(accountID string) {
	c.accountID = accountID
}

func (c Client) BaseURL() (*url.URL, error) {
	tmpl, err := template.New("host").Parse(c.baseURL)
	if err != nil {
		return &url.URL{}, err
	}
	buf := new(bytes.Buffer)
	// err = tmpl.Execute(buf, map[string]interface{}{"account_id": c.companyID})
	err = tmpl.Execute(buf, map[string]interface{}{})
	if err != nil {
		return &url.URL{}, err
	}
	return url.Parse(buf.String())
}

func (c *Client) SetBaseURL(baseURL string) {
	c.baseURL = baseURL
}

func (c *Client) SetMediaType(mediaType string) {
	c.mediaType = mediaType
}

func (c Client) MediaType() string {
	return mediaType
}

func (c *Client) SetCharset(charset string) {
	c.charset = charset
}

func (c Client) Charset() string {
	return charset
}

func (c *Client) SetUserAgent(userAgent string) {
	c.userAgent = userAgent
}

func (c Client) UserAgent() string {
	return userAgent
}

func (c *Client) SetBeforeRequestDo(fun BeforeRequestDoCallback) {
	c.beforeRequestDo = fun
}

func (c *Client) GetEndpointURL(p string, pathParams PathParams) (url.URL, error) {
	clientURL, err := c.BaseURL()
	if err != nil {
		return url.URL{}, err
	}

	parsed, err := url.Parse(p)
	if err != nil {
		log.Fatal(err)
	}
	q := clientURL.Query()
	for k, vv := range parsed.Query() {
		for _, v := range vv {
			q.Add(k, v)
		}
	}
	clientURL.RawQuery = q.Encode()

	clientURL.Path = path.Join(clientURL.Path, parsed.Path)

	tmpl, err := template.New("path").Parse(clientURL.Path)
	if err != nil {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	params := pathParams.Params()
	// params["administration_id"] = c.Administration()
	err = tmpl.Execute(buf, params)
	if err != nil {
		return url.URL{}, err
	}

	clientURL.Path = buf.String()
	return *clientURL, nil
}

func (c *Client) NewRequest(ctx context.Context, req Request) (*http.Request, error) {
	// convert body struct to xml
	buf := new(bytes.Buffer)
	if req.RequestBodyInterface() != nil {
		soapRequest := NewRequestEnvelope()
		soapRequest.Body.ActionBody = req.RequestBodyInterface()
		soapRequest.Header.SearchPreferences = *c.SearchPreferences

		// Add passport header
		g := c.NewSignatureGenerator()
		signature, err := g.Generate()
		if err != nil {
			return nil, err
		}

		soapRequest.Header.TokenPassport.Account = c.AccountID()
		soapRequest.Header.TokenPassport.ConsumerKey = c.ClientID()
		soapRequest.Header.TokenPassport.Token = c.TokenID()
		soapRequest.Header.TokenPassport.Nonce = g.Nonce
		soapRequest.Header.TokenPassport.Signature.Algorithm = g.SignatureMethod.String()
		soapRequest.Header.TokenPassport.Signature.Text = signature
		soapRequest.Header.TokenPassport.Timestamp = g.Timestamp

		enc := xml.NewEncoder(buf)
		enc.Indent("", "  ")
		err = enc.Encode(soapRequest)
		if err != nil {
			return nil, err
		}

		err = enc.Flush()
		if err != nil {
			return nil, err
		}
	}

	// create new http request
	u, err := req.URL()
	if err != nil {
		return nil, err
	}

	r, err := http.NewRequest(req.Method(), u.String(), buf)
	if err != nil {
		return nil, err
	}

	// values := url.Values{}
	// err = utils.AddURLValuesToRequest(values, req, true)
	// if err != nil {
	// 	return nil, err
	// }

	// optionally pass along context
	if ctx != nil {
		r = r.WithContext(ctx)
	}

	// set other headers
	r.Header.Add("Content-Type", fmt.Sprintf("%s; charset=%s", c.MediaType(), c.Charset()))
	r.Header.Add("Accept", c.MediaType())
	r.Header.Add("User-Agent", c.UserAgent())
	r.Header.Add("SOAPAction", req.SOAPAction())

	return r, nil
}

func (c *Client) NewSignatureGenerator() *SignatureGenerator {
	return &SignatureGenerator{
		SignatureMethod: HMACSHA256,
		ClientID:        c.ClientID(),
		ClientSecret:    c.ClientSecret(),
		TokenID:         c.TokenID(),
		TokenSecret:     c.TokenSecret(),
		AccountID:       c.AccountID(),
		Nonce:           GenerateNonce(),
		Timestamp:       time.Now().Unix(),
	}
	// return &SignatureGenerator{
	// 	SignatureMethod: HMACSHA256,
	// 	ClientID:        "71cc02b731f05895561ef0862d71553a3ac99498a947c3b7beaf4a1e4a29f7c4",
	// 	ClientSecret:    "7278da58caf07f5c336301a601203d10a58e948efa280f0618e25fcee1ef2abd",
	// 	TokenID:         "89e08d9767c5ac85b374415725567d05b54ecf0960ad2470894a52f741020d82",
	// 	TokenSecret:     "060cd9ab3ffbbe1e3d3918e90165ffd37ab12acc76b4691046e2d29c7d7674c2",
	// 	AccountID:       "1234567",
	// 	Nonce:           "6obMKq0tmY8ylVOdEkA1",
	// 	Timestamp:       1439829974,
	// }
}

// Do sends an Client request and returns the Client response. The Client response is xml decoded and stored in the value
// pointed to by v, or returned as an error if an Client error has occurred. If v implements the io.Writer interface,
// the raw response will be written to v, without attempting to decode it.
func (c *Client) Do(req *http.Request, body interface{}) (*http.Response, error) {
	if c.beforeRequestDo != nil {
		c.beforeRequestDo(c.http, req, body)
	}

	if c.debug == true {
		dump, _ := httputil.DumpRequestOut(req, true)
		log.Println(string(dump))
	}

	httpResp, err := c.http.Do(req)
	if err != nil {
		// wrap error in http error so we can handle it properly
		return nil, &httperr.Error{StatusCode: httpResp.StatusCode, Err: err}
	}

	if c.onRequestCompleted != nil {
		c.onRequestCompleted(req, httpResp)
	}

	// close body io.Reader
	defer func() {
		if rerr := httpResp.Body.Close(); err == nil {
			err = rerr
		}
	}()

	if c.debug == true {
		dump, _ := httputil.DumpResponse(httpResp, true)
		log.Println(string(dump))
	}

	// check if the response isn't an error
	err = CheckResponse(httpResp)
	if err != nil {
		return httpResp, err
	}

	// check the provided interface parameter
	if httpResp == nil {
		return httpResp, nil
	}

	if body == nil {
		return httpResp, err
	}

	if httpResp.ContentLength == 0 {
		return httpResp, nil
	}

	soapResponse := &ResponseEnvelope{
		Header: Header{},
		Body: Body{
			ActionBody: body,
		},
	}

	soapError := SOAPError{Response: httpResp}
	errResp := &ResponseEnvelope{
		Header: Header{},
		Body: Body{
			ActionBody: &soapError,
		},
	}

	soapFault := SOAPFault{Response: httpResp}
	faultResp := &ResponseEnvelope{
		Header: Header{},
		Body: Body{
			ActionBody: &soapFault,
		},
	}

	statusResponseBody := StatusResponseBody{Response: httpResp}
	statusResp := &ResponseEnvelope{
		Header: Header{},
		Body: Body{
			ActionBody: &statusResponseBody,
		},
	}

	err = c.Unmarshal(httpResp.Body, soapResponse, errResp, faultResp, statusResp)
	if err != nil {
		return httpResp, err
	}

	if statusResponseBody.Node.Status.Error() != "" {
		return httpResp, statusResponseBody.Node.Status
	}

	if soapError.Error() != "" {
		return httpResp, soapError
	}

	if soapFault.Error() != "" {
		return httpResp, soapFault
	}

	return httpResp, nil
}

func (c *Client) Unmarshal(r io.Reader, vv ...interface{}) error {
	if len(vv) == 0 {
		return nil
	}

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	errs := []error{}
	for _, v := range vv {
		r := bytes.NewReader(b)
		dec := xml.NewDecoder(r)

		err := dec.Decode(v)
		if err != nil && err != io.EOF {
			errs = append(errs, err)
		}
	}

	if len(errs) == len(vv) {
		// Everything errored
		msgs := make([]string, len(errs))
		for i, e := range errs {
			log.Println(e)
			msgs[i] = fmt.Sprint(e)
		}
		return errors.New(strings.Join(msgs, ", "))
	}

	return nil
}

// CheckResponse checks the Client response for errors, and returns them if
// present. A response is considered an error if it has a status code outside
// the 200 range. Client error responses are expected to have either no response
// body, or a xml response body that maps to ErrorResponse. Any other response
// body will be silently ignored.
func CheckResponse(r *http.Response) error {
	errorResponse := &ErrorResponse{Response: r}

	// Don't check content-lenght: a created response, for example, has no body
	// if r.Header.Get("Content-Length") == "0" {
	// 	errorResponse.Errors.Message = r.Status
	// 	return errorResponse
	// }

	if c := r.StatusCode; c >= 200 && c <= 299 {
		return nil
	}

	// read data and copy it back
	data, err := ioutil.ReadAll(r.Body)
	r.Body = ioutil.NopCloser(bytes.NewReader(data))
	if err != nil {
		return errorResponse
	}

	err = checkContentType(r)
	if err != nil {
		return errors.WithStack(err)
	}

	if r.ContentLength == 0 {
		return errors.New("response body is empty")
	}

	// convert xml to struct
	if len(data) != 0 {
		err = xml.Unmarshal(data, &errorResponse)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	if errorResponse.Error() != "" {
		return errorResponse
	}

	return nil
}

func checkContentType(response *http.Response) error {
	header := response.Header.Get("Content-Type")
	contentType := strings.Split(header, ";")[0]
	if contentType != mediaType {
		return fmt.Errorf("Expected Content-Type \"%s\", got \"%s\"", mediaType, contentType)
	}

	return nil
}

// {
//   "type": "https://www.w3.org/Protocols/rfc2616/rfc2616-sec10.html#sec10.4.4",
//   "title": "Forbidden",
//   "status": 403,
//   "o:errorDetails": [
//     {
//       "detail": "The account record is only available as a beta record. Enable the REST Record Service (Beta) feature in Setup > Company > Enable Features to work with this record.",
//       "o:errorCode": "INSUFFICIENT_PERMISSION"
//     }
//   ]
// }

type ErrorResponse struct {
	// HTTP response that caused this error
	Response *http.Response
	Err      string
}

func (r *ErrorResponse) Error() string {
	return r.Err
}

type SOAPError struct {
	// HTTP response that caused this error
	Response *http.Response
}

func (e SOAPError) Error() string {
	return ""
}

type SOAPFault struct {
	// HTTP response that caused this error
	Response *http.Response

	XMLName     xml.Name `xml:"Fault"`
	Faultcode   string   `xml:"faultcode"`
	Faultstring string   `xml:"faultstring"`
	Detail      struct {
		Fault struct {
			PlatformFaults string `xml:"platformFaults,attr"`
			Code           string `xml:"code"`
			Message        string `xml:"message"`
		} `xml:"any"`
		Hostname struct {
			Ns1 string `xml:"ns1,attr"`
		} `xml:"hostname"`
	} `xml:"detail"`
}

func (f SOAPFault) Error() string {
	l := []string{f.Faultcode, f.Faultstring, f.Detail.Fault.Code, f.Detail.Fault.Message}
	ll := []string{}
	for _, v := range l {
		if v != "" {
			ll = append(ll, v)
		}
	}

	return strings.Join(ll, ", ")
}

type StatusResponseBody struct {
	// HTTP response that caused this error
	Response *http.Response

	Node struct {
		Status Status `xml:"status"`
	} `xml:",any"`
}
