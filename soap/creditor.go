package soap

import "encoding/xml"

type obsolete string
type notinuse string

type CrdDimensionsResponse struct {
	XMLName    xml.Name   `xml:"general"`
	Result     int        `xml:"result,attr"`
	Dimensions []Creditor `xml:"dimension"`
}

// type Dimension Creditor

type Creditor struct {
	XMLName xml.Name `xml:"dimension"`

	Office Office `xml:"office" csv:"office"`
	// Type
	Code              string `xml:"code" csv:"code"`
	Uid               string `xml:"uid"`
	Name              string `xml:"name"`
	ShortName         string `xml:"shortname"`
	Inuse             bool
	Behaviour         string
	Touched           int
	Begionperiod      int
	Beginyear         int
	Endperiod         int
	Endyear           int
	Website           string
	Cocnumber         obsolete
	Vatnumber         obsolete
	Editdimensionname notinuse
	// Financials        Financial
	// Remittanceadvice  Remittanceadvice
	Addresses []Address `xml:"addresses>address"`
	Banks     []Bank    `xml:"banks>bank"`
}

func CreditorFromXml(data []byte) (Creditor, error) {
	creditor := Creditor{}

	err := xml.Unmarshal(data, &creditor)
	return creditor, err
}

func CreditorsFromXml(data []byte) ([]Creditor, error) {
	response := CrdDimensionsResponse{}
	err := xml.Unmarshal(data, &response)

	return response.Dimensions, err
}
