package soap

import "encoding/xml"

type PeriodRequest struct {
	XMLName xml.Name `xml:"period"`

	Office     string `xml:"office"`
	YearPeriod string `xml:"yearperiod"`
}

type PeriodResponse struct {
	XMLName xml.Name `xml:"period"`

	Result     int                     `xml:"result,attr,omitempty"`
	MsgType    string                  `xml:"msgtype,attr"`
	Msg        string                  `xml:"msg,attr"`
	Office     PeriodResponseOffice    `xml:"office"`
	YearPeriod PeriodReponseYearperiod `xml:"yearperiod"`
	Year       int                     `xml:"year"`
}

type PeriodResponseOffice struct {
	ID        string `xml:",chardata"`
	Name      string `xml:"name,attr"`
	ShortName string `xml:"shortname,attr"`
}

type PeriodReponseYearperiod struct {
	Result  int    `xml:"result,attr,omitempty"`
	MsgType string `xml:"msgtype,attr"`
	Msg     string `xml:"msg,attr"`
	Period  string `xml:",chardata"`
}
