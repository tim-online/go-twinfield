package soap

import "encoding/xml"

type GeneralRequest struct {
	XMLName xml.Name `xml:"general"`

	Transactions []Transaction   `xml:"transaction"`
	Periods      []PeriodRequest `xml:"periods>period"`
}

// <general result="1">
// 	<periods result="1">
// 		<period result="1">
// 			<office name="Test Bergvliet" shortname="ZZZ LEEG_4CFG">00005659</office>
// 			<yearperiod>2017/01</yearperiod>
// 			<year>2017</year>
// 			<period>1</period>
// 		</period>
// 	</periods>
// </general>
type GeneralResponse struct {
	XMLName xml.Name `xml:"general"`

	Result  int              `xml:"result,attr,omitempty"`
	Periods []PeriodResponse `xml:"periods>period"`
}
