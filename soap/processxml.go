package soap

import (
	"encoding/xml"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type ProcessXmlString struct {
	XMLName xml.Name `xml:"http://www.twinfield.com/ ProcessXmlString"`

	XmlRequest string `xml:"xmlRequest,omitempty"`
}

type ProcessXmlStringResponse struct {
	XMLName xml.Name `xml:"http://www.twinfield.com/ ProcessXmlStringResponse"`

	ProcessXmlStringResult string `xml:"ProcessXmlStringResult,omitempty"`
}

type ProcessXmlDocument struct {
	XMLName xml.Name `xml:"http://www.twinfield.com/ ProcessXmlDocument"`

	XmlRequest struct {
	} `xml:"xmlRequest,omitempty"`
}

type ProcessXmlDocumentResponse struct {
	XMLName xml.Name `xml:"http://www.twinfield.com/ ProcessXmlDocumentResponse"`

	ProcessXmlDocumentResult struct {
	} `xml:"ProcessXmlDocumentResult,omitempty"`
}

type ProcessXmlCompressed struct {
	XMLName xml.Name `xml:"http://www.twinfield.com/ ProcessXmlCompressed"`

	XmlRequest []byte `xml:"xmlRequest,omitempty"`
}

type ProcessXmlCompressedResponse struct {
	XMLName xml.Name `xml:"http://www.twinfield.com/ ProcessXmlCompressedResponse"`

	ProcessXmlCompressedResult []byte `xml:"ProcessXmlCompressedResult,omitempty"`
}

type ProcessXmlSoap struct {
	client *SOAPClient
}

func NewProcessXmlSoap(url string, tls bool, auth *BasicAuth) *ProcessXmlSoap {
	if url == "" {
		url = "https://C3.twinfield.com/webservices/processxml.asmx"
	}
	client := NewSOAPClient(url, tls, auth)

	return &ProcessXmlSoap{
		client: client,
	}
}

func (service *ProcessXmlSoap) SetHeader(header interface{}) {
	service.client.SetHeader(header)
}

/* Processes xml in string format. */
func (service *ProcessXmlSoap) ProcessXmlString(request *ProcessXmlString) (*ProcessXmlStringResponse, error) {
	response := new(ProcessXmlStringResponse)
	err := service.client.Call("http://www.twinfield.com/ProcessXmlString", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Processes xml in document object model format. */
func (service *ProcessXmlSoap) ProcessXmlDocument(request *ProcessXmlDocument) (*ProcessXmlDocumentResponse, error) {
	response := new(ProcessXmlDocumentResponse)
	err := service.client.Call("http://www.twinfield.com/ProcessXmlDocument", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Processes a G-zipped UTF-8 XML string and returns a G-zipped UTF-8 result. */
func (service *ProcessXmlSoap) ProcessXmlCompressed(request *ProcessXmlCompressed) (*ProcessXmlCompressedResponse, error) {
	response := new(ProcessXmlCompressedResponse)
	err := service.client.Call("http://www.twinfield.com/ProcessXmlCompressed", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
