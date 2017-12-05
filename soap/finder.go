package soap

import "encoding/xml"

type MessageType string

type ErrorCodes string

const (
	ErrorCodesNoAccessToOffice                     ErrorCodes = "NoAccessToOffice"
	ErrorCodesOptionNotAllowed                     ErrorCodes = "OptionNotAllowed"
	ErrorCodesInvalidBooleanOptionValue            ErrorCodes = "InvalidBooleanOptionValue"
	ErrorCodesInvalidIntegerOptionValue            ErrorCodes = "InvalidIntegerOptionValue"
	ErrorCodesInvalidDecimalOptionValue            ErrorCodes = "InvalidDecimalOptionValue"
	ErrorCodesInvalidEnumerationOptionValue        ErrorCodes = "InvalidEnumerationOptionValue"
	ErrorCodesOptionValueOutOfRange                ErrorCodes = "OptionValueOutOfRange"
	ErrorCodesParameterOutOfRange                  ErrorCodes = "ParameterOutOfRange"
	ErrorCodesInvalidFinderType                    ErrorCodes = "InvalidFinderType"
	ErrorCodesParameterTooSmall                    ErrorCodes = "ParameterTooSmall"
	ErrorCodesOptionLevelMandatoryForSectionTEQ    ErrorCodes = "OptionLevelMandatoryForSectionTEQ"
	ErrorCodesOptionICIncompatableWithOptionHidden ErrorCodes = "OptionICIncompatableWithOptionHidden"
	ErrorCodesInvalidDateTimeOptionLength          ErrorCodes = "InvalidDateTimeOptionLength"
	ErrorCodesInvalidDateTimeOptionValue           ErrorCodes = "InvalidDateTimeOptionValue"
	ErrorCodesInvalidDateTimeOptionOutOfRange      ErrorCodes = "InvalidDateTimeOptionOutOfRange"
	ErrorCodesOptionMandatory                      ErrorCodes = "OptionMandatory"
	ErrorCodesAccessDenied                         ErrorCodes = "AccessDenied"
	ErrorCodesDisableAccessRulesNotAllowed         ErrorCodes = "DisableAccessRulesNotAllowed"
	ErrorCodesOption1MandatoryIfOption2IsUsed      ErrorCodes = "Option1MandatoryIfOption2IsUsed"
)

type Search struct {
	XMLName xml.Name `xml:"http://www.twinfield.com/ Search"`

	Type_ string `xml:"type,omitempty"`

	Pattern string `xml:"pattern,omitempty"`

	Field int32 `xml:"field,omitempty"`

	FirstRow int32 `xml:"firstRow,omitempty"`

	MaxRows int32 `xml:"maxRows,omitempty"`

	Options *ArrayOfArrayOfString `xml:"options,omitempty"`
}

type SearchResponse struct {
	XMLName xml.Name `xml:"http://www.twinfield.com/ SearchResponse"`

	SearchResult *SearchResult `xml:"SearchResult,omitempty"`

	Data *FinderData `xml:"data,omitempty"`
}

type ArrayOfArrayOfString struct {
	// XMLName xml.Name `xml:"http://www.twinfield.com/ ArrayOfArrayOfString"`

	ArrayOfString []*ArrayOfString `xml:"ArrayOfString,omitempty"`
}

type ArrayOfString struct {
	// XMLName xml.Name `xml:"http://www.twinfield.com/ ArrayOfString"`

	String []string `xml:"string,omitempty"`
}

type SearchResult struct {
	XMLName xml.Name `xml:"http://www.twinfield.com/ SearchResult"`

	MessageOfErrorCodes []*MessageOfErrorCodes `xml:"MessageOfErrorCodes,omitempty"`
}

type MessageOfErrorCodes struct {
	XMLName xml.Name `xml:"http://www.twinfield.com/ MessageOfErrorCodes"`

	Type *MessageType `xml:"Type,omitempty"`

	Text string `xml:"Text,omitempty"`

	Code *ErrorCodes `xml:"Code,omitempty"`

	Parameters *ArrayOfString `xml:"Parameters,omitempty"`
}

type FinderData struct {
	// XMLName xml.Name `xml:"http://www.twinfield.com/ FinderData"`

	TotalRows int32 `xml:"TotalRows,omitempty"`

	Columns *ArrayOfString `xml:"Columns,omitempty"`

	Items *ArrayOfArrayOfString `xml:"Items,omitempty"`
}

type FinderSoap struct {
	client *SOAPClient
}

func NewFinderSoap(url string, tls bool, auth *BasicAuth) *FinderSoap {
	if url == "" {
		url = "https://C4.twinfield.com/webservices/finder.asmx"
	}
	client := NewSOAPClient(url, tls, auth)

	return &FinderSoap{
		client: client,
	}
}

func (service *FinderSoap) SetHeader(header interface{}) {
	service.client.SetHeader(header)
}

/* Searches for different types of data based on the given finder type and search pattern. */
func (service *FinderSoap) Search(request *Search) (*SearchResponse, error) {
	response := new(SearchResponse)
	err := service.client.Call("http://www.twinfield.com/Search", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
