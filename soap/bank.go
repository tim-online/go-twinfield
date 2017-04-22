package soap

type Bank struct {
	ID            int     `xml:"id,attr"`
	Type          string  `xml:"type,attr"`
	Default       bool    `xml:"default,attr"`
	Ascription    string  `xml:"ascription"`
	Accountnumber string  `xml:"accountnumber"`
	Address       Address `xml:"address"`
	Name          string  `xml:"bankname"`
	Biccode       string  `xml:"biccode"`
	City          string  `xml:"city"`
	// Country   string `xml:"country,innerxml"`
	Iban       string `xml:"iban"`
	Natbiccode string `xml:"natbiccode"`
	Postcode   string `xml:"postcode"`
	State      string `xml:"state"`
}
