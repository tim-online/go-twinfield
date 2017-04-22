package soap

import (
	"encoding/xml"
	"strings"
)

type Dimension struct {
	ID            string `xml:",innerxml"`
	Name          string `xml:"name,attr"`
	ShortName     string `xml:"shortname,attr"`
	DimensionType string `xml:"dimensiontype,attr"`
	Code          string `xml:"innerxml"`
}

// <dimensions office="50000295" type="CRD" result="1">
// 	<dimension name="Diversen" shortname="">
// 		2000
// 	</dimension>
// 	<dimension name="FJ Stuyts" shortname="">
// 		2001
// 	</dimension>
// 	<dimension name="Gekoelde Transporten Zeeland BV" shortname="">
// 		2002
// 	</dimension>
// 	<dimension name="EasyTrip" shortname="">
// 		2003
// 	</dimension>
// </dimensions>

type DimensionsResponse struct {
	XMLName    xml.Name    `xml:"dimensions"`
	Office     string      `xml:"office,attr"`
	Type       string      `xml:"type,attr"`
	Result     int         `xml:"result,attr"`
	Dimensions []Dimension `xml:"dimension"`
}

func DimensionsFromXml(data []byte) ([]Dimension, error) {
	dimensionsResp := DimensionsResponse{}

	err := xml.Unmarshal([]byte(data), &dimensionsResp)
	if err != nil {
		return nil, err
	}

	dimensions := dimensionsResp.Dimensions
	for i, dimension := range dimensions {
		dimensions[i].ID = strings.TrimSpace(dimension.ID)
	}

	return dimensions, err
}
