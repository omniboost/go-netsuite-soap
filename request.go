package netsuite

import "net/url"

type Request interface {
	Method() string
	// QueryParams() QueryParams
	PathParams() PathParams
	RequestBodyInterface() interface{}
	URL() (*url.URL, error)
	SOAPAction() string
}

// PreferencesRequest can be implemented by a Request that wants its own
// platformMsgs:preferences header. These replace the client-level Preferences
// for that request only, so a client shared between requests is not affected.
type PreferencesRequest interface {
	Preferences() *Preferences
}

type QueryParams interface {
	ToURLValues() (url.Values, error)
}

type PathParams interface {
	Params() map[string]string
}
