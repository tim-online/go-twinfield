package soap

import "encoding/xml"

type DebDimensionsResponse struct {
	XMLName    xml.Name  `xml:"general"`
	Result     int       `xml:"result,attr"`
	Dimensions []Debitor `xml:"dimension"`
}

type Debitor struct {
	XMLName xml.Name `xml:"dimension"`

	Office            Office   `xml:"office,omitempty" csv:"office"`
	Type              string   `xml:"type"`
	Code              string   `xml:"code" csv:"code"`
	Uid               string   `xml:"uid,omitempty"`
	Name              string   `xml:"name"`
	ShortName         string   `xml:"shortname"`
	Inuse             bool     `xml:"inuse,omitempty"`
	Behaviour         string   `xml:"behaviour,omitempty"`
	Touched           int      `xml:"touched,omitempty"`
	Beginperiod       int      `xml:"begionperiod,omitempty"`
	Beginyear         int      `xml:"beginyear,omitempty"`
	Endperiod         int      `xml:"endperiod,omitempty"`
	Endyear           int      `xml:"endyear,omitempty"`
	Website           string   `xml:"website"`
	Cocnumber         obsolete `xml:"cocnumber,omitempty"`
	Vatnumber         obsolete `xml:"vatnumber,omitempty"`
	Editdimensionname notinuse `xml:"editdimensionname,omitempty"`
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
