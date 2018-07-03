package twinfield

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/tim-online/go-twinfield/soap"
	"golang.org/x/oauth2"
)

const sessionTimeout = time.Minute * 30

var (
	sessionID    string
	sessionStart time.Time
	User         = ""
	Password     = ""
	Organisation = ""
	cluster      = ""

	ClientID     = ""
	ClientSecret = ""
	RefreshToken = ""
	token        *oauth2.Token
	company      string // allow this to be set
)

type Oauth2Config struct {
	oauth2.Config
}

func Login() error {
	if IsSessionClient() {
		return LoginCredentials()
	}

	if IsOauthClient() {
		return LoginOauth()
	}

	return errors.New("neither user and oauth credentials are set")
}

func IsSessionClient() bool {
	return User != "" && Password != "" && Organisation != ""
}

func IsOauthClient() bool {
	return ClientID != "" && ClientSecret != "" && RefreshToken != ""
}

func LoginCredentials() error {
	var err error
	sessionID, err = loginCredentials()
	return err
}

func LoginOauth() error {
	// @TODO: implement retry strategy?
	var err error
	sessionID, err = loginOauth()
	return err
}

func loginOauth() (string, error) {
	var err error
	if ClientID == "" {
		return "", errors.New("ClientID is required")
	}

	if ClientSecret == "" {
		return "", errors.New("ClientSecret is required")
	}

	if RefreshToken == "" {
		return "", errors.New("RefreshToken is required")
	}

	token, err = RequestAccessTokenWithRefreshToken()
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://login.twinfield.com/auth/authentication/connect/accesstokenvalidation?token=%s", token.AccessToken)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	d := json.NewDecoder(resp.Body)
	val := Validation{}
	err = d.Decode(&val)
	if err != nil {
		return "", err
	}

	cluster = val.ClusterURL // set global variable
	return "", nil
}

func RequestAccessTokenWithRefreshToken() (*oauth2.Token, error) {
	// build client
	oauthConfig := NewOauth2Config()
	oauthConfig.ClientID = ClientID
	oauthConfig.ClientSecret = ClientSecret

	token := &oauth2.Token{
		RefreshToken: RefreshToken,
	}

	// get http client with automatic oauth logic
	// httpClient := oauthConfig.Client(oauth2.NoContext, refreshToken)

	tokenSource := oauthConfig.TokenSource(oauth2.NoContext, token)
	return tokenSource.Token()
}

func NewOauth2Config() *Oauth2Config {
	authURL, _ := url.Parse("https://login.twinfield.com/auth/authentication/connect/authorize")
	tokenURL, _ := url.Parse("https://login.twinfield.com/auth/authentication/connect/token")

	// These are not registered in the oauth library by default
	oauth2.RegisterBrokenAuthHeaderProvider("https://login.twinfield.com")

	return &Oauth2Config{
		Config: oauth2.Config{
			RedirectURL:  "",
			ClientID:     "",
			ClientSecret: "",
			Scopes:       []string{},
			Endpoint: oauth2.Endpoint{
				AuthURL:  authURL.String(),
				TokenURL: tokenURL.String(),
			},
		},
	}
}

// @TODO: add KeepAlive() in separate thread or shutdown session after
// # of connections = 0
func loginCredentials() (string, error) {
	session := soap.NewSessionSoap("", true, nil)
	logon := &soap.Logon{
		User:         User,
		Password:     Password,
		Organisation: Organisation,
	}

	logonResp, err := session.Logon(logon)
	if err != nil {
		return "", err
	}

	if *logonResp.LogonResult != soap.LogonResultOk {
		return "", fmt.Errorf("logonResult: %s", *logonResp.LogonResult)
	}
	cluster = logonResp.Cluster
	sessionID := logonResp.Envelope.Header.Header.(*soap.Header).SessionID
	return sessionID, nil
}

func GetSessionID() (string, error) {
	if sessionID == "" {
		err := Login()
		if err == nil {
			// logon succesful: update session start
			sessionStart = time.Now().UTC()
		}
		return sessionID, err
	}

	sessionAge := time.Now().UTC().Sub(sessionStart)
	if sessionAge > sessionTimeout {
		err := Login()
		if err == nil {
			// logon succesful: update session start
			sessionStart = time.Now().UTC()
		}
		return sessionID, err
	}

	return sessionID, nil
}

func GetAccessToken() (string, error) {
	if token.Valid() {
		return token.AccessToken, nil
	}

	err := Login()
	if err != nil {
		return "", err
	}

	// start new session with accesstoken
	return token.AccessToken, nil
}

func GetOffices() ([]soap.Office, error) {
	xml := &soap.ProcessXmlString{
		XmlRequest: "<list><type>offices</type></list>",
	}

	process, err := GetProcessXmlSoap()
	if err != nil {
		return nil, err
	}

	processResp, err := process.ProcessXmlString(xml)
	if err != nil {
		return nil, err
	}

	data := []byte(processResp.ProcessXmlStringResult)
	offices, err := soap.OfficesFromXml(data)
	return offices, err
}

func GetProcessXmlSoap() (*soap.ProcessXmlSoap, error) {
	header, err := GetSoapHeader()
	if err != nil {
		return nil, err
	}

	url := cluster + "/webservices/processxml.asmx"
	process := soap.NewProcessXmlSoap(url, true, nil)
	process.SetHeader(header)
	return process, nil
}

func GetFinderXmlSoap() (*soap.FinderSoap, error) {
	header, err := GetSoapHeader()
	if err != nil {
		return nil, err
	}

	url := cluster + "/webservices/finder.asmx"
	finder := soap.NewFinderSoap(url, true, nil)
	finder.SetHeader(header)
	return finder, nil
}

func GetDeclarationsXmlSoap() (*soap.DeclarationsSoap, error) {
	header, err := GetSoapHeader()
	if err != nil {
		return nil, err
	}

	url := cluster + "/webservices/declarations.asmx"
	req := soap.NewDeclarationsSoap(url, true, nil)
	req.SetHeader(header)
	return req, nil
}

func GetSessionXmlSoap() (*soap.SessionSoap, error) {
	header, err := GetSoapHeader()
	if err != nil {
		return nil, err
	}

	url := cluster + "/webservices/session.asmx"
	req := soap.NewSessionSoap(url, true, nil)
	req.SetHeader(header)
	return req, nil
}

func GetOfficeByID(officeID string) (*soap.Office, error) {
	xml := &soap.ProcessXmlString{
		XmlRequest: fmt.Sprintf(`<read>
			<type>office</type>
			<code>%s</code>
		</read>`, officeID),
	}

	process, err := GetProcessXmlSoap()
	if err != nil {
		return nil, err
	}

	processResp, err := process.ProcessXmlString(xml)
	if err != nil {
		return nil, err
	}

	data := []byte(processResp.ProcessXmlStringResult)
	office, err := soap.OfficeFromXml(data)
	return office, err
}

func Cluster() string {
	return cluster
}

func GetSoapHeader() (*soap.Header, error) {
	header := &soap.Header{}

	if IsOauthClient() {
		token, err := GetAccessToken()
		if err != nil {
			return header, err
		}
		header.AccessToken = token
		header.CompanyCode = company
		return header, nil
	}

	sessionID, err := GetSessionID()
	if err != nil {
		return header, err
	}
	header.SessionID = sessionID
	return header, nil
}

type Validation struct {
	ClusterURL string `json:"twf.clusterUrl"`
}

func SetCompany(company string) {
	company = company
}
