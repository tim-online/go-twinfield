package soap

type Address struct {
	ID      int    `xml:"id,attr"`
	Type    string `xml:"type,attr"`
	Default bool   `xml:"default,attr"`
	Name    string `xml:"name"`
	// Country   string `xml:"country,innerxml"`
	City      string `xml:"city"`
	Postcode  string `xml:"postcode"`
	Telephone string `xml:"telephone"`
	Telefax   string `xml:"telefax"`
	Email     string `xml:"email"`
	Contact   string `xml:"contact"`
	Field1    string `xml:"field1"`
	Field2    string `xml:"field2"`
	Field3    string `xml:"field3"`
	Field4    string `xml:"field4"`
	Field5    string `xml:"field5"`
	Field6    string `xml:"field6"`
}
