package soap

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// <office name="Logivisi 2.0 B.V." shortname="ZUN 50000043">50000043</office>
type Office struct {
	XMLName   xml.Name `xml:"office"`
	ID        string   `xml:",innerxml" csv:"office_id"`
	Name      string   `xml:"name,attr" csv:"office_name"`
	ShortName string   `xml:"shortname,attr" csv:"office_shortname"`
}

func (o *Office) String() string {
	return o.ID
}

type OfficesResponse struct {
	XMLName xml.Name `xml:"offices"`
	Offices []Office `xml:"office"`
}

func OfficesFromXml(data []byte) ([]Office, error) {
	officesResp := OfficesResponse{}

	err := xml.Unmarshal([]byte(data), &officesResp)
	if err != nil {
		return nil, err
	}

	offices := officesResp.Offices
	for i, office := range offices {
		offices[i].ID = strings.TrimSpace(office.ID)
	}

	return offices, err
}

type OfficeResponse struct {
	Code      string `xml:"code"`
	Name      string `xml:"name"`
	Created   string `xml:"created"`
	Modified  string `xml:"modified"`
	ShortName string `xml:"shortname"`
	Touched   int    `xml:"touched"`
	User      struct {
		ID        string `xml:",innerxml"`
		Name      string `xml:"name,attr"`
		ShortName string `xml:"shortname,attr"`
	} `xml:"user"`
	General struct {
	} `xml:"general"`
	SystemDimensions struct {
	} `xml:"systemdimensions"`
	SystemDimensionTypes struct {
	} `xml:"systemdimensiontypes"`
	Creditmanagement struct {
	} `xml:"creditmanagement"`
	Vatmanagement struct {
	} `xml:"vatmanagement"`
	FixedAssets struct {
	} `xml:"fixedassets"`
	InterCompany struct {
	} `xml:"intercompany"`
	Regional struct {
	} `xml:"regional"`

	// error shit
	MsgType string `xml:"msgtype,attr"`
	Msg     string `xml:"msg,attr"`
}

func OfficeFromXml(data []byte) (*Office, error) {
	officeResp := OfficeResponse{}

	err := xml.Unmarshal([]byte(data), &officeResp)
	if err != nil {
		return nil, err
	}

	if officeResp.Msg != "" {
		return nil, fmt.Errorf(officeResp.Msg)
	}

	office := &Office{
		ID:        officeResp.Code,
		Name:      officeResp.Name,
		ShortName: officeResp.ShortName,
	}

	return office, err
}
