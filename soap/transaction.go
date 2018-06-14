package soap

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strconv"
	"time"
)

var (
	DestinyTemporary Destiny = "temporary"
	DestinyFinal     Destiny = "final"

	RegimeFiscal     Regime = "Fiscal"
	RegimeCommercial Regime = "Commercial"
	RegimeEconomic   Regime = "Economic"
	RegimeGeneric    Regime = "Generic"

	LineTypeDetail LineType = "detail"
	LineTypeVat    LineType = "vat"

	Debit  DebitCredit = "debit"
	Credit DebitCredit = "credit"

	MatchStatusAvailable    MatchStatus = "available"
	MatchStatusMatched      MatchStatus = "matched"
	MatchStatusProposed     MatchStatus = "proposed"
	MatchStatusNotMatchable MatchStatus = "notmatchable"

	PerformanceTypeServices PerformanceType = "services"
	PerformanceTypeGoods    PerformanceType = "goods"

	MsgTypeError MsgType = "error"
)

type Destiny string
type LineType string
type DebitCredit string
type Money float64
type MatchStatus string
type Regime string
type PerformanceType string
type MsgType string

type Transaction struct {
	XMLName        xml.Name `xml:"transaction"`
	Destiny        Destiny  `xml:"destiny,attr"`
	Location       Destiny  `xml:"location,attr,omitempty"`
	AutobalanceVat bool     `xml:"autobalancevat,attr"`
	RaiseWarning   bool     `xml:"raisewarning,attr"`
	MsgType        MsgType  `xml:"msgtype,attr,omitempty"`
	Msg            string   `xml:"msg,attr,omitempty"`
	Result         int      `xml:"result,attr,omitempty"`

	Header TransactionHeader `xml:"header"`
	Lines  []TransactionLine `xml:"lines>line"`
}

type TransactionHeader struct {
	Office           string  `xml:"office"` // Office code
	Code             string  `xml:"code"`   // Transaction type code
	Number           int     `xml:"number"` // Transaction number
	Period           string  `xml:"period"` // Period in YYYY/PP format
	Currency         string  `xml:"currency"`
	Regime           Regime  `xml:"regime,omitempty"`
	Date             *Date   `xml:"date,omitempty"` // date in YYYYmmdd format
	Origin           Destiny `xml:"origin"`
	ModificationDate string  `xml:"modificationdate,omitempty"`
	User             string  `xml:"user,omitempty"`
	Inputdate        string  `xml:"inputdate,omitempty"`
	Freetext1        string  `xml:"freetext1,omitempty"`
	Freetext2        string  `xml:"freetext2,omitempty"`
	Freetext3        string  `xml:"freetext3,omitempty"`
	InvoiceNumber    string  `xml:"invoicenumber"`
}

type TransactionLine struct {
	Type                 LineType        `xml:"type,attr"`
	ID                   int             `xml:"id,attr,omitempty"`
	Dim1                 string          `xml:"dim1,omitempty"`
	Dim2                 string          `xml:"dim2,omitempty"`
	Dim3                 string          `xml:"dim3,omitempty,omitempty"`
	Dim4                 string          `xml:"dim4,omitempty,omitempty"`
	Dim5                 string          `xml:"dim5,omitempty,omitempty"`
	Dim6                 string          `xml:"dim6,omitempty,omitempty"`
	DebitCredit          DebitCredit     `xml:"debitcredit"`
	Value                Money           `xml:"value"`
	BaseValue            Money           `xml:"basevalue,omitempty"`
	Rate                 float64         `xml:"rate,omitempty"`
	RepValue             Money           `xml:"repvalue,omitempty"`
	RepRate              float64         `xml:"reprate,omitempty"`
	Description          string          `xml:"description"`
	MatchStatus          MatchStatus     `xml:"matchstatus,omitempty"`
	MatchLevel           int             `xml:"matchlevel,omitempty"`
	Relation             int             `xml:"relation,omitempty"`
	BaseValueOpen        Money           `xml:"basevalueopen,omitempty"`
	RepValueOpen         Money           `xml:"repvalueopen,omitempty"`
	VatCode              string          `xml:"vatcode,omitempty"`
	VatValue             Money           `xml:"vatvalue,omitempty"`
	VatBaseValue         Money           `xml:"vatbasevalue,omitempty"`
	VatRepValue          Money           `xml:"vatrepvalue,omitempty"`
	VatTurnover          Money           `xml:"vatturnover,omitempty"`
	VatBaseTurnover      Money           `xml:"vatbaseturnover,omitempty"`
	VatRepTurnover       Money           `xml:"vatrepturnover,omitempty"`
	Baseline             int             `xml:"baseline,omitempty"`
	PerformanceType      PerformanceType `xml:"performancetype,omitempty"`
	PerformanceCountry   string          `xml:"performancecountry,omitempty"`
	PerformanceVatNumber string          `xml:"performancevatnumber,omitempty"`
	PerformanceDate      *Date           `xml:"performancedate,omitempty"`
	DestOffice           string          `xml:"destoffice,omitempty"`
	CurrencyDate         *Date           `xml:"currencydate,omitempty"` // date in YYYYmmdd format
	FreeChar             string          `xml:"freechar,omitempty"`
	Comment              string          `xml:"comment,omitempty"`
	Matches              []Match         `xml:"matches>match,omitempty"`

	MsgType MsgType `xml:"msgtype,attr,omitempty"`
	Msg     string  `xml:"msg,attr,omitempty"`
}

type Match struct {
}

func (m *Money) UnmarshalJSON(b []byte) error {
	value := ""
	err := json.Unmarshal(b, &value)
	if err == nil {
		if value == "" {
			return nil
		}

		f, err := strconv.ParseFloat(value, 64)

		m2 := Money(f)
		*m = m2
		return err
	}

	f := 0.0
	err = json.Unmarshal(b, &f)
	if err == nil {
		m2 := Money(f)
		*m = m2
		return err
	}

	return err
}

func (m Money) MarshalText() (text []byte, err error) {
	f := float64(m)
	return []byte(fmt.Sprintf("%.2f", f)), nil
}

type Date time.Time

func (d *Date) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	t := time.Time(*d)
	return e.EncodeElement(t.Format("20060102"), start)
}

func (d *Date) UnmarshalText(text []byte) (err error) {
	value := string(text)
	if value == "" {
		return nil
	}

	// 20170119
	// Mon Jan 2 15:04:05 -0700 MST 2006
	layout := "20060102"
	time, err := time.Parse(layout, string(text))
	date := Date(time)
	*d = date
	return err
}

func (d *Date) MarshalJSON() ([]byte, error) {
	t := time.Time(*d)
	return json.Marshal(t.Format("20060102"))
}
