package soap

import (
	"encoding/xml"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

const (
	MessageTypeError MessageType = "Error"

	MessageTypeWarning MessageType = "Warning"

	MessageTypeInformational MessageType = "Informational"
)

type LoadMessage string

const (
	LoadMessageNoAccess LoadMessage = "NoAccess"

	LoadMessageInvalidCode LoadMessage = "InvalidCode"

	LoadMessageInvalidParameter LoadMessage = "InvalidParameter"

	LoadMessageNotSupported LoadMessage = "NotSupported"

	LoadMessageInvalidDocumentType LoadMessage = "InvalidDocumentType"
)

type SetStatusMessages string

const (
	SetStatusMessagesNoAccess SetStatusMessages = "NoAccess"

	SetStatusMessagesDocumentDoesNotExist SetStatusMessages = "DocumentDoesNotExist"

	SetStatusMessagesStartStatusInCorrectForTransition SetStatusMessages = "StartStatusInCorrectForTransition"

	SetStatusMessagesUnableToPerformTransition SetStatusMessages = "UnableToPerformTransition"
)

type GetAllSummaries struct {
	XMLName xml.Name `xml:"http://www.twinfield.com/ GetAllSummaries"`

	CompanyCode string `xml:"companyCode,omitempty"`
}

type GetAllSummariesResponse struct {
	XMLName xml.Name `xml:"http://www.twinfield.com/ GetAllSummariesResponse"`

	GetAllSummariesResult *ArrayOfMessageOfLoadMessage `xml:"GetAllSummariesResult,omitempty"`

	VatReturn *ArrayOfDeclarationSummary `xml:"vatReturn,omitempty"`
}

type GetRangeOfSummaries struct {
	XMLName xml.Name `xml:"http://www.twinfield.com/ GetRangeOfSummaries"`

	CompanyCode string `xml:"companyCode,omitempty"`

	StartIndex int32 `xml:"startIndex,omitempty"`

	Count int32 `xml:"count,omitempty"`
}

type GetRangeOfSummariesResponse struct {
	XMLName xml.Name `xml:"http://www.twinfield.com/ GetRangeOfSummariesResponse"`

	GetRangeOfSummariesResult *ArrayOfMessageOfLoadMessage `xml:"GetRangeOfSummariesResult,omitempty"`

	VatReturn *ArrayOfDeclarationSummary `xml:"vatReturn,omitempty"`
}

type GetNumberOfDeclarations struct {
	XMLName xml.Name `xml:"http://www.twinfield.com/ GetNumberOfDeclarations"`

	CompanyCode string `xml:"companyCode,omitempty"`
}

type GetNumberOfDeclarationsResponse struct {
	XMLName xml.Name `xml:"http://www.twinfield.com/ GetNumberOfDeclarationsResponse"`

	GetNumberOfDeclarationsResult *ArrayOfMessageOfLoadMessage `xml:"GetNumberOfDeclarationsResult,omitempty"`

	NumberOfVatReturns int32 `xml:"numberOfVatReturns,omitempty"`
}

type GetVatReturnAsXbrl struct {
	XMLName xml.Name `xml:"http://www.twinfield.com/ GetVatReturnAsXbrl"`

	DocumentId int32 `xml:"documentId,omitempty"`

	IsMessageIdRequired bool `xml:"isMessageIdRequired,omitempty"`
}

type GetVatReturnAsXbrlResponse struct {
	XMLName xml.Name `xml:"http://www.twinfield.com/ GetVatReturnAsXbrlResponse"`

	GetVatReturnAsXbrlResult *ArrayOfMessageOfLoadMessage `xml:"GetVatReturnAsXbrlResult,omitempty"`

	// VatReturn struct {} `xml:"vatReturn,omitempty"`
	VatReturn []byte `xml:",innerxml"`
}

type GetVatReturnAsXml struct {
	XMLName xml.Name `xml:"http://www.twinfield.com/ GetVatReturnAsXml"`

	DocumentId int32 `xml:"documentId,omitempty"`
}

type GetVatReturnAsXmlResponse struct {
	XMLName xml.Name `xml:"http://www.twinfield.com/ GetVatReturnAsXmlResponse"`

	GetVatReturnAsXmlResult *ArrayOfMessageOfLoadMessage `xml:"GetVatReturnAsXmlResult,omitempty"`

	// VatReturn struct {} `xml:"vatReturn,omitempty"`
	VatReturn []byte `xml:",innerxml"`
}

type GetIctReturnAsXbrl struct {
	XMLName xml.Name `xml:"http://www.twinfield.com/ GetIctReturnAsXbrl"`

	DocumentId int32 `xml:"documentId,omitempty"`

	IsMessageIdRequired bool `xml:"isMessageIdRequired,omitempty"`
}

type GetIctReturnAsXbrlResponse struct {
	XMLName xml.Name `xml:"http://www.twinfield.com/ GetIctReturnAsXbrlResponse"`

	GetIctReturnAsXbrlResult *ArrayOfMessageOfLoadMessage `xml:"GetIctReturnAsXbrlResult,omitempty"`

	// VatReturn struct {} `xml:"vatReturn,omitempty"`
	VatReturn []byte `xml:",innerxml"`
}

type GetIctReturnAsXml struct {
	XMLName xml.Name `xml:"http://www.twinfield.com/ GetIctReturnAsXml"`

	DocumentId int32 `xml:"documentId,omitempty"`
}

type GetIctReturnAsXmlResponse struct {
	XMLName xml.Name `xml:"http://www.twinfield.com/ GetIctReturnAsXmlResponse"`

	GetIctReturnAsXmlResult *ArrayOfMessageOfLoadMessage `xml:"GetIctReturnAsXmlResult,omitempty"`

	OpgaafICP []byte `xml:",innerxml"`
}

type GetTaxGroupVatReturnAsXbrl struct {
	XMLName xml.Name `xml:"http://www.twinfield.com/ GetTaxGroupVatReturnAsXbrl"`

	DocumentId int32 `xml:"documentId,omitempty"`

	IsMessageIdRequired bool `xml:"isMessageIdRequired,omitempty"`
}

type GetTaxGroupVatReturnAsXbrlResponse struct {
	XMLName xml.Name `xml:"http://www.twinfield.com/ GetTaxGroupVatReturnAsXbrlResponse"`

	GetTaxGroupVatReturnAsXbrlResult *ArrayOfMessageOfLoadMessage `xml:"GetTaxGroupVatReturnAsXbrlResult,omitempty"`

	VatReturn struct {
	} `xml:"vatReturn,omitempty"`
}

type GetTaxGroupVatReturnAsXml struct {
	XMLName xml.Name `xml:"http://www.twinfield.com/ GetTaxGroupVatReturnAsXml"`

	DocumentId int32 `xml:"documentId,omitempty"`
}

type GetTaxGroupVatReturnAsXmlResponse struct {
	XMLName xml.Name `xml:"http://www.twinfield.com/ GetTaxGroupVatReturnAsXmlResponse"`

	GetTaxGroupVatReturnAsXmlResult *ArrayOfMessageOfLoadMessage `xml:"GetTaxGroupVatReturnAsXmlResult,omitempty"`

	VatReturn struct {
	} `xml:"vatReturn,omitempty"`
}

type SetToSent struct {
	XMLName xml.Name `xml:"http://www.twinfield.com/ SetToSent"`

	DocumentId int32 `xml:"documentId,omitempty"`
}

type SetToSentResponse struct {
	XMLName xml.Name `xml:"http://www.twinfield.com/ SetToSentResponse"`

	SetToSentResult *ArrayOfMessageOfSetStatusMessages `xml:"SetToSentResult,omitempty"`
}

type SetToApproved struct {
	XMLName xml.Name `xml:"http://www.twinfield.com/ SetToApproved"`

	DocumentId int32 `xml:"documentId,omitempty"`
}

type SetToApprovedResponse struct {
	XMLName xml.Name `xml:"http://www.twinfield.com/ SetToApprovedResponse"`

	SetToApprovedResult *ArrayOfMessageOfSetStatusMessages `xml:"SetToApprovedResult,omitempty"`
}

type SetToRejected struct {
	XMLName xml.Name `xml:"http://www.twinfield.com/ SetToRejected"`

	DocumentId int32 `xml:"documentId,omitempty"`

	Reason string `xml:"reason,omitempty"`
}

type SetToRejectedResponse struct {
	XMLName xml.Name `xml:"http://www.twinfield.com/ SetToRejectedResponse"`

	SetToRejectedResult *ArrayOfMessageOfSetStatusMessages `xml:"SetToRejectedResult,omitempty"`
}

type ArrayOfMessageOfLoadMessage struct {
	MessageOfLoadMessage []*MessageOfLoadMessage `xml:"MessageOfLoadMessage,omitempty"`
}

type MessageOfLoadMessage struct {
	XMLName xml.Name `xml:"http://www.twinfield.com/ MessageOfLoadMessage"`

	Type *MessageType `xml:"Type,omitempty"`

	Text string `xml:"Text,omitempty"`

	Code *LoadMessage `xml:"Code,omitempty"`

	Parameters *ArrayOfString `xml:"Parameters,omitempty"`
}

type ArrayOfDeclarationSummary struct {
	DeclarationSummary []DeclarationSummary `xml:"DeclarationSummary,omitempty"`
}

type DeclarationSummary struct {
	XMLName xml.Name `xml:"http://www.twinfield.com/ DeclarationSummary"`

	Id int32 `xml:"Id,omitempty"`

	DocumentCode string `xml:"DocumentCode,omitempty"`

	Name string `xml:"Name,omitempty"`

	DocumentTimeFrame *DocumentTimeFrame `xml:"DocumentTimeFrame,omitempty"`

	Status *DocumentStatus `xml:"Status,omitempty"`

	Assignee *CodeName `xml:"Assignee,omitempty"`

	Company *CodeName `xml:"Company,omitempty"`
}

type DocumentTimeFrame struct {
	XMLName xml.Name `xml:"http://www.twinfield.com/ DocumentTimeFrame"`

	Year int `xml:"Year,omitempty"`

	Period string `xml:"Period,omitempty"`
}

type DocumentStatus struct {
	Description string `xml:"Description,omitempty"`

	StepIndex int32 `xml:"StepIndex,omitempty"`

	ExtraInformation string `xml:"ExtraInformation,omitempty"`
}

type CodeName struct {
	Code string `xml:"Code,omitempty"`

	Name string `xml:"Name,omitempty"`
}

type ArrayOfMessageOfSetStatusMessages struct {
	XMLName xml.Name `xml:"http://www.twinfield.com/ ArrayOfMessageOfSetStatusMessages"`

	MessageOfSetStatusMessages []*MessageOfSetStatusMessages `xml:"MessageOfSetStatusMessages,omitempty"`
}

type MessageOfSetStatusMessages struct {
	XMLName xml.Name `xml:"http://www.twinfield.com/ MessageOfSetStatusMessages"`

	Type *MessageType `xml:"Type,omitempty"`

	Text string `xml:"Text,omitempty"`

	Code *SetStatusMessages `xml:"Code,omitempty"`

	Parameters *ArrayOfString `xml:"Parameters,omitempty"`
}

type DeclarationsSoap struct {
	client *SOAPClient
}

func NewDeclarationsSoap(url string, tls bool, auth *BasicAuth) *DeclarationsSoap {
	if url == "" {
		url = "https://C4.twinfield.com/webservices/declarations.asmx"
	}
	client := NewSOAPClient(url, tls, auth)

	return &DeclarationsSoap{
		client: client,
	}
}

func (service *DeclarationsSoap) SetHeader(header interface{}) {
	service.client.SetHeader(header)
}

/* Returns all the declaration summaries for a given company. */
func (service *DeclarationsSoap) GetAllSummaries(request *GetAllSummaries) (*GetAllSummariesResponse, error) {
	response := new(GetAllSummariesResponse)
	err := service.client.Call("http://www.twinfield.com/GetAllSummaries", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Returns a range of declaration summaries for a given company. */
func (service *DeclarationsSoap) GetRangeOfSummaries(request *GetRangeOfSummaries) (*GetRangeOfSummariesResponse, error) {
	response := new(GetRangeOfSummariesResponse)
	err := service.client.Call("http://www.twinfield.com/GetRangeOfSummaries", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Returns number of declarations for the given company. */
func (service *DeclarationsSoap) GetNumberOfDeclarations(request *GetNumberOfDeclarations) (*GetNumberOfDeclarationsResponse, error) {
	response := new(GetNumberOfDeclarationsResponse)
	err := service.client.Call("http://www.twinfield.com/GetNumberOfDeclarations", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Returns Vat return declaration details in Xbrl format. */
func (service *DeclarationsSoap) GetVatReturnAsXbrl(request *GetVatReturnAsXbrl) (*GetVatReturnAsXbrlResponse, error) {
	response := new(GetVatReturnAsXbrlResponse)
	err := service.client.Call("http://www.twinfield.com/GetVatReturnAsXbrl", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Returns Vat return declaration details in Xml format. */
func (service *DeclarationsSoap) GetVatReturnAsXml(request *GetVatReturnAsXml) (*GetVatReturnAsXmlResponse, error) {
	response := new(GetVatReturnAsXmlResponse)
	err := service.client.Call("http://www.twinfield.com/GetVatReturnAsXml", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Returns Ict return declaration details in Xbrl format. */
func (service *DeclarationsSoap) GetIctReturnAsXbrl(request *GetIctReturnAsXbrl) (*GetIctReturnAsXbrlResponse, error) {
	response := new(GetIctReturnAsXbrlResponse)
	err := service.client.Call("http://www.twinfield.com/GetIctReturnAsXbrl", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Returns Ict return declaration details in Xml format. */
func (service *DeclarationsSoap) GetIctReturnAsXml(request *GetIctReturnAsXml) (*GetIctReturnAsXmlResponse, error) {
	response := new(GetIctReturnAsXmlResponse)
	err := service.client.Call("http://www.twinfield.com/GetIctReturnAsXml", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Returns tax group vat return declaration details in Xbrl format. */
func (service *DeclarationsSoap) GetTaxGroupVatReturnAsXbrl(request *GetTaxGroupVatReturnAsXbrl) (*GetTaxGroupVatReturnAsXbrlResponse, error) {
	response := new(GetTaxGroupVatReturnAsXbrlResponse)
	err := service.client.Call("http://www.twinfield.com/GetTaxGroupVatReturnAsXbrl", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Returns tax group vat return declaration details in Xml format. */
func (service *DeclarationsSoap) GetTaxGroupVatReturnAsXml(request *GetTaxGroupVatReturnAsXml) (*GetTaxGroupVatReturnAsXmlResponse, error) {
	response := new(GetTaxGroupVatReturnAsXmlResponse)
	err := service.client.Call("http://www.twinfield.com/GetTaxGroupVatReturnAsXml", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Sets the declaration status to Sent. */
func (service *DeclarationsSoap) SetToSent(request *SetToSent) (*SetToSentResponse, error) {
	response := new(SetToSentResponse)
	err := service.client.Call("http://www.twinfield.com/SetToSent", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Sets the declaration status to Approved. */
func (service *DeclarationsSoap) SetToApproved(request *SetToApproved) (*SetToApprovedResponse, error) {
	response := new(SetToApprovedResponse)
	err := service.client.Call("http://www.twinfield.com/SetToApproved", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Sets the declaration status to Rejected. */
func (service *DeclarationsSoap) SetToRejected(request *SetToRejected) (*SetToRejectedResponse, error) {
	response := new(SetToRejectedResponse)
	err := service.client.Call("http://www.twinfield.com/SetToRejected", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
