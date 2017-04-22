package soap

import "encoding/xml"

type DebDimensionsResponse struct {
	XMLName    xml.Name  `xml:"general"`
	Result     int       `xml:"result,attr"`
	Dimensions []Debitor `xml:"dimension"`
}

type Debitor struct {
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
	// Creditmanagement        Creditmanagement
	// Remittanceadvice  Remittanceadvice
	Addresses []Address `xml:"addresses>address"`
	Banks     []Bank    `xml:"banks>bank"`
	// Postingrules
}

func DebitorFromXml(data []byte) (Debitor, error) {
	debitor := Debitor{}

	err := xml.Unmarshal(data, &debitor)
	return debitor, err
}

func DebitorsFromXml(data []byte) ([]Debitor, error) {
	response := DebDimensionsResponse{}
	err := xml.Unmarshal(data, &response)

	return response.Dimensions, err
}
