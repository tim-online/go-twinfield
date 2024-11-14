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

const (
	libraryVersion = "0.0.1"
	// loginEndpoint  = "https://login.twinfield.com/webservices/session.asmx?wsdl"
	userAgent      = "go-twinfield/" + libraryVersion
	mediaType      = "text/xml"
	sessionTimeout = time.Minute * 30
)

func NewClient(httpClient *http.Client, user string, password string, organisation string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	// baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{client: httpClient, UserAgent: userAgent, User: user, Password: password, Organisation: organisation}
	// c.Products = &ProductsService{client: c}
	return c
}

func NewOauthClient(httpClient *http.Client, clientID, clientSecret, refreshToken string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	// baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{client: httpClient, UserAgent: userAgent, ClientID: clientID, ClientSecret: clientSecret, RefreshToken: refreshToken}
	// c.Products = &ProductsService{client: c}
	return c
}

type Client struct {
	// HTTP client used to communicate with the DO API.
	client *http.Client

	// User agent for client
	UserAgent string

	// Services used for communicating with the API
	LoginService *LoginService

	// Optional function called after every successful request made to the DO APIs
	onRequestCompleted RequestCompletionCallback

	// Credentials
	sessionID    string
	sessionStart time.Time
	cluster      string
	User         string
	Password     string
	Organisation string
	company      string

	// Oauth
	token        *oauth2.Token
	ClientID     string
	ClientSecret string
	RefreshToken string
}

// RequestCompletionCallback defines the type of the request callback function
type RequestCompletionCallback func(*http.Request, *http.Response)

type Oauth2Config struct {
	oauth2.Config
}

func (c *Client) Login() error {
	if c.IsSessionClient() {
		return c.LoginCredentials()
	}

	if c.IsOauthClient() {
		return c.LoginOauth()
	}

	return errors.New("neither user and oauth credentials are set")
}

func (c *Client) IsSessionClient() bool {
	return c.User != "" && c.Password != "" && c.Organisation != ""
}

func (c *Client) IsOauthClient() bool {
	return c.ClientID != "" && c.ClientSecret != "" && c.RefreshToken != ""
}

func (c *Client) LoginCredentials() error {
	var err error
	c.sessionID, c.cluster, err = c.loginCredentials()
	return err
}

func (c *Client) LoginOauth() error {
	// @TODO: implement retry strategy?
	var err error
	c.sessionID, err = c.loginOauth()
	return err
}

func (c *Client) loginOauth() (string, error) {
	var err error
	if c.ClientID == "" {
		return "", errors.New("ClientID is required")
	}

	if c.ClientSecret == "" {
		return "", errors.New("ClientSecret is required")
	}

	if c.RefreshToken == "" {
		return "", errors.New("RefreshToken is required")
	}

	c.token, err = c.RequestAccessTokenWithRefreshToken()
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://login.twinfield.com/auth/authentication/connect/accesstokenvalidation?token=%s", c.token.AccessToken)
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

	c.cluster = val.ClusterURL // set global variable
	return "", nil
}

func (c *Client) RequestAccessTokenWithRefreshToken() (*oauth2.Token, error) {
	// build client
	oauthConfig := c.NewOauth2Config()
	oauthConfig.ClientID = c.ClientID
	oauthConfig.ClientSecret = c.ClientSecret

	token := &oauth2.Token{
		RefreshToken: c.RefreshToken,
	}

	// get http client with automatic oauth logic
	// httpClient := oauthConfig.Client(oauth2.NoContext, refreshToken)

	tokenSource := oauthConfig.TokenSource(oauth2.NoContext, token)
	return tokenSource.Token()
}

func (c *Client) NewOauth2Config() *Oauth2Config {
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
func (c *Client) loginCredentials() (string, string, error) {
	session := soap.NewSessionSoap("", true, nil)
	logon := &soap.Logon{
		User:         c.User,
		Password:     c.Password,
		Organisation: c.Organisation,
	}

	logonResp, err := session.Logon(logon)
	if err != nil {
		return "", "", err
	}

	if *logonResp.LogonResult != soap.LogonResultOk {
		return "", "", fmt.Errorf("logonResult: %s", *logonResp.LogonResult)
	}
	cluster := logonResp.Cluster
	sessionID := logonResp.Envelope.Header.Header.(*soap.Header).SessionID
	return sessionID, cluster, nil
}

func (c *Client) GetSessionID() (string, error) {
	if c.sessionID == "" {
		err := c.Login()
		if err == nil {
			// logon succesful: update session start
			c.sessionStart = time.Now().UTC()
		}
		return c.sessionID, err
	}

	sessionAge := time.Now().UTC().Sub(c.sessionStart)
	if sessionAge > sessionTimeout {
		err := c.Login()
		if err == nil {
			// logon succesful: update session start
			c.sessionStart = time.Now().UTC()
		}
		return c.sessionID, err
	}

	return c.sessionID, nil
}

func (c *Client) GetAccessToken() (string, error) {
	if c.token.Valid() {
		return c.token.AccessToken, nil
	}

	err := c.Login()
	if err != nil {
		return "", err
	}

	// start new session with accesstoken
	return c.token.AccessToken, nil
}

func (c *Client) GetOffices() ([]soap.Office, error) {
	xml := &soap.ProcessXmlString{
		XmlRequest: "<list><type>offices</type></list>",
	}

	process, err := c.GetProcessXmlSoap()
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

func (c *Client) GetProcessXmlSoap() (*soap.ProcessXmlSoap, error) {
	header, err := c.GetSoapHeader()
	if err != nil {
		return nil, err
	}

	url := c.cluster + "/webservices/processxml.asmx"
	process := soap.NewProcessXmlSoap(url, true, nil)
	process.SetHeader(header)
	return process, nil
}

func (c *Client) GetFinderXmlSoap() (*soap.FinderSoap, error) {
	header, err := c.GetSoapHeader()
	if err != nil {
		return nil, err
	}

	url := c.cluster + "/webservices/finder.asmx"
	finder := soap.NewFinderSoap(url, true, nil)
	finder.SetHeader(header)
	return finder, nil
}

func (c *Client) GetDeclarationsXmlSoap() (*soap.DeclarationsSoap, error) {
	header, err := c.GetSoapHeader()
	if err != nil {
		return nil, err
	}

	url := c.cluster + "/webservices/declarations.asmx"
	req := soap.NewDeclarationsSoap(url, true, nil)
	req.SetHeader(header)
	return req, nil
}

func (c *Client) GetSessionXmlSoap() (*soap.SessionSoap, error) {
	header, err := c.GetSoapHeader()
	if err != nil {
		return nil, err
	}

	url := c.cluster + "/webservices/session.asmx"
	req := soap.NewSessionSoap(url, true, nil)
	req.SetHeader(header)
	return req, nil
}

func (c *Client) GetOfficeByID(officeID string) (*soap.Office, error) {
	xml := &soap.ProcessXmlString{
		XmlRequest: fmt.Sprintf(`<read>
			<type>office</type>
			<code>%s</code>
		</read>`, officeID),
	}

	process, err := c.GetProcessXmlSoap()
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

func (c *Client) GetSoapHeader() (*soap.Header, error) {
	header := &soap.Header{}

	if c.IsOauthClient() {
		token, err := c.GetAccessToken()
		if err != nil {
			return header, err
		}
		header.AccessToken = token
		header.CompanyCode = c.company
		return header, nil
	}

	sessionID, err := c.GetSessionID()
	if err != nil {
		return header, err
	}
	header.SessionID = sessionID
	return header, nil
}

func (c *Client) SetCompany(company string) {
	c.company = company
}

type Validation struct {
	ClusterURL string `json:"twf.clusterUrl"`
}
