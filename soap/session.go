package soap

import (
	"encoding/xml"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type LogonResult string

const (
	LogonResultOk LogonResult = "Ok"

	LogonResultBlocked LogonResult = "Blocked"

	LogonResultUntrusted LogonResult = "Untrusted"

	LogonResultInvalid LogonResult = "Invalid"

	LogonResultDeleted LogonResult = "Deleted"

	LogonResultDisabled LogonResult = "Disabled"

	LogonResultOrganisationInactive LogonResult = "OrganisationInactive"

	LogonResultClientInvalid LogonResult = "ClientInvalid"

	LogonResultFailed LogonResult = "Failed"

	LogonResultTokenInvalid LogonResult = "TokenInvalid"
)

type LogonAction string

const (
	LogonActionNone LogonAction = "None"

	LogonActionSMSLogon LogonAction = "SMSLogon"

	LogonActionChangePassword LogonAction = "ChangePassword"
)

type Logon struct {
	XMLName xml.Name `xml:"http://www.twinfield.com/ Logon"`

	User string `xml:"user,omitempty"`

	Password string `xml:"password,omitempty"`

	Organisation string `xml:"organisation,omitempty"`
}

type LogonResponse struct {
	XMLName  xml.Name `xml:"http://www.twinfield.com/ LogonResponse"`
	Envelope *SOAPEnvelope

	LogonResult *LogonResult `xml:"LogonResult,omitempty"`

	NextAction *LogonAction `xml:"nextAction,omitempty"`

	Cluster string `xml:"cluster,omitempty"`
}

type OAuthLogon struct {
	XMLName xml.Name `xml:"http://www.twinfield.com/ OAuthLogon"`

	ClientToken string `xml:"clientToken,omitempty"`

	ClientSecret string `xml:"clientSecret,omitempty"`

	AccessToken string `xml:"accessToken,omitempty"`

	AccessSecret string `xml:"accessSecret,omitempty"`
}

type OAuthLogonResponse struct {
	XMLName xml.Name `xml:"http://www.twinfield.com/ OAuthLogonResponse"`

	OAuthLogonResult *LogonResult `xml:"OAuthLogonResult,omitempty"`

	NextAction *LogonAction `xml:"nextAction,omitempty"`

	Cluster string `xml:"cluster,omitempty"`
}

type AccessTokenLogon struct {
	XMLName xml.Name `xml:"http://www.twinfield.com/ AccessTokenLogon"`

	AccessToken string `xml:"accessToken,omitempty"`
}

type AccessTokenLogonResponse struct {
	XMLName xml.Name `xml:"http://www.twinfield.com/ AccessTokenLogonResponse"`

	AccessTokenLogonResult *LogonResult `xml:"AccessTokenLogonResult,omitempty"`

	Cluster string `xml:"cluster,omitempty"`
}

// type KeepAliveRequest struct {
// }

// type KeepAliveResponse struct {
// }

type SessionSoap struct {
	client *SOAPClient
}

func NewSessionSoap(url string, tls bool, auth *BasicAuth) *SessionSoap {
	if url == "" {
		url = "https://login.twinfield.com/webservices/session.asmx"
	}
	client := NewSOAPClient(url, tls, auth)

	return &SessionSoap{
		client: client,
	}
}

func (service *SessionSoap) SetHeader(header interface{}) {
	service.client.SetHeader(header)
}

/* Logs on with the user credentials. */
func (service *SessionSoap) Logon(request *Logon) (*LogonResponse, error) {
	response := new(LogonResponse)
	err := service.client.Call("http://www.twinfield.com/Logon", request, response)
	if err != nil {
		return nil, err
	}

	response.Envelope = service.client.respEnvelope
	return response, nil
}

/* Logs on with the OAuth credentials. */
func (service *SessionSoap) OAuthLogon(request *OAuthLogon) (*OAuthLogonResponse, error) {
	response := new(OAuthLogonResponse)
	err := service.client.Call("http://www.twinfield.com/OAuthLogon", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Creates session based on passed access token and returns its ID and external cluster URL */
func (service *SessionSoap) AccessTokenLogon(request *AccessTokenLogon) (*AccessTokenLogonResponse, error) {
	response := new(AccessTokenLogonResponse)
	err := service.client.Call("http://www.twinfield.com/AccessTokenLogon", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// // Keep the session alive, to prevent session time out. A session time out will occur 2 hours after the last web service call for the session.
// func (service *SessionSoap) KeepAlive() (*KeepAliveResponse, error) {
// 	request := new(KeepAliveRequest)
// 	response := new(KeepAliveResponse)
// 	err := service.client.Call("http://www.twinfield.com/KeepAlive", request, response)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return response, nil
// }
