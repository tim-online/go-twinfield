package soap

import "encoding/xml"

type VatDimensionsResponse struct {
	XMLName xml.Name `xml:"general"`
	Result  int      `xml:"result,attr"`
	Vats    []Vat    `xml:"vat"`
}

type Vat struct {
	Office Office `xml:"office" csv:"office"`
	// Type
	Code      string `xml:"code"`
	Name      string `xml:"name"`
	ShortName string `xml:"shortname"`
	Uid       string `xml:"uid"`
	Created   string `xml:"created"`
	Modified  string `xml:"modified"`
	Touched   int    `xml:"touched"`
	// User      User   `xml:"user"`
	Type        string       `xml:"type"`
	Percentages []Percentage `xml:"percentages>percentage"`
}

type Percentage struct {
	Date       string `xml:"date"`
	Status     string `xml:"status,attr"`
	Inuse      bool   `xml:"inuse,attr"`
	Percentage string `xml:"percentage"`
	Created    string `xml:"created"`
	Name       string `xml:"name"`
	ShortName  string `xml:"shortname"`
	// User User `xml:"user"`
	// Discountaccount  string  `xml:"discountaccount"`
	// Writeoffaccount  string  `xml:"writeoffaccount"`
	Accounts []Account `xml:"accounts>account"`
}

type Account struct {
	ID   int       `xml:"id,attr"`
	Dim1 Dimension `xml:"dim1"`
	// GroupCountry GroupCountry `xml:"groupcountry"`
	// Group        Group        `xml:"group"`
	Percentage string `xml:"percentage"`
	Linetype   string `xml:"linetype"`
}

func VatFromXml(data []byte) (Vat, error) {
	vat := Vat{}

	err := xml.Unmarshal(data, &vat)
	return vat, err
}

func VatsFromXml(data []byte) ([]Vat, error) {
	response := VatDimensionsResponse{}
	err := xml.Unmarshal(data, &response)

	return response.Vats, err
}
