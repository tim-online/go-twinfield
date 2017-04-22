package twinfield

import "net/http"

const (
	libraryVersion = "0.0.1"
	// loginEndpoint  = "https://login.twinfield.com/webservices/session.asmx?wsdl"
	userAgent = "go-twinfield/" + libraryVersion
	mediaType = "text/xml"
)

func NewClient(httpClient *http.Client, user string, password string, organisation string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	// set package globals :(
	User = user
	Password = password
	Organisation = organisation

	// baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{client: httpClient, UserAgent: userAgent}
	// c.Products = &ProductsService{client: c}
	return c
}

type Client struct {
	// HTTP client used to communicate with the DO API.
	client *http.Client

	// User agent for client
	UserAgent string

	// Services used for communicating with the API
	Login *LoginService

	// Optional function called after every successful request made to the DO APIs
	onRequestCompleted RequestCompletionCallback
}

// RequestCompletionCallback defines the type of the request callback function
type RequestCompletionCallback func(*http.Request, *http.Response)
