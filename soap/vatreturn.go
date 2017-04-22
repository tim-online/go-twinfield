package soap

import (
	"encoding/xml"
	"time"
)

type AanUBinnenland struct {
	OmzetVerlegd          float64 `xml:" OmzetVerlegd,omitempty" json:"OmzetVerlegd,omitempty"`
	OmzetbelastingVerlegd float64 `xml:" OmzetbelastingVerlegd,omitempty" json:"OmzetbelastingVerlegd,omitempty"`
}

type AanUBuitenland struct {
	OmzetBuitenEUInvoer            float64 `xml:"OmzetBuitenEUInvoer,omitempty" json:"OmzetBuitenEUInvoer,omitempty"`
	OmzetbelastingBuitenEUInvoer   float64 `xml:"OmzetbelastingBuitenEUInvoer,omitempty" json:"OmzetbelastingBuitenEUInvoer,omitempty"`
	OmzetGoederenBinnenEU          float64 `xml:" OmzetGoederenBinnenEU,omitempty" json:"OmzetGoederenBinnenEU,omitempty"`
	OmzetbelastingGoederenBinnenEU float64 `xml:" OmzetbelastingGoederenBinnenEU,omitempty" json:"OmzetbelastingGoederenBinnenEU,omitempty"`
}

type AangifteOmzetbelasting struct {
	AanUBinnenland          *AanUBinnenland          `xml:" AanUBinnenland,omitempty" json:"AanUBinnenland,omitempty"`
	AanUBuitenland          *AanUBuitenland          `xml:"AanUBuitenland,omitempty" json:"AanUBuitenland,omitempty"`
	Berekening              *Berekening              `xml:"Berekening,omitempty" json:"Berekening,omitempty"`
	Communicatiegegevens    []*Communicatiegegevens  `xml:"Communicatiegegevens,omitempty" json:"Communicatiegegevens,omitempty"`
	DoorUBinnenland         *DoorUBinnenland         `xml:"DoorUBinnenland,omitempty" json:"DoorUBinnenland,omitempty"`
	DoorUBuitenland         *DoorUBuitenland         `xml:"DoorUBuitenland,omitempty" json:"DoorUBuitenland,omitempty"`
	IdentificerendeGegevens *IdentificerendeGegevens `xml:"IdentificerendeGegevens,omitempty" json:"IdentificerendeGegevens,omitempty"`
}

type OpgaafICP struct {
	IdentificerendeGegevens *IdentificerendeGegevens `xml:"IdentificerendeGegevens,omitempty" json:"IdentificerendeGegevens,omitempty"`
	Communicatiegegevens    []*Communicatiegegevens  `xml:"Communicatiegegevens,omitempty" json:"Communicatiegegevens,omitempty"`
	// CorrectiesLeveringenDiensten
	// CorrectiesABCLeveringen
	IntracommunautaireLeveringenDiensten []*Levering `xml:"IntracommunautaireLeveringenDiensten>Levering,omitempty" json:"IntracommunautaireLeveringenDiensten,omitempty"`
	// IntracommunautaireABCLeveringen
}

type Levering struct {
	Landcode string  `xml:"Landocde,omitempty" json:"Landcode,omitempty"`
	Nummer   string  `xml:"Nummer,omitempty" json:"Nummer,omitempty"`
	Bedrag   float64 `xml:"Bedrag,omitempty" json:"Bedrag,omitempty"`
}

type Berekening struct {
	TotaalTeBetalenTerugTeVragen float64 `xml:"TotaalTeBetalenTerugTeVragen,omitempty" json:"TotaalTeBetalenTerugTeVragen,omitempty"`
	VerschuldigdeOmzetbelasting  float64 `xml:"VerschuldigdeOmzetbelasting,omitempty" json:"VerschuldigdeOmzetbelasting,omitempty"`
	Voorbelasting                float64 `xml:"Voorbelasting,omitempty" json:"Voorbelasting,omitempty"`
	KleineOndernemersregeling    float64 `xml:"KleineOndernemersregeling,omitempty" json:"KleineOndernemersregeling,omitempty"`
}

type Communicatiegegevens struct {
	ContactpersoonID       string `xml:" ContactpersoonID,omitempty" json:"ContactpersoonID,omitempty"`
	NaamContactpersoon     string `xml:" NaamContactpersoon,omitempty" json:"NaamContactpersoon,omitempty"`
	SoortContactpersoon    string `xml:" SoortContactpersoon,omitempty" json:"SoortContactpersoon,omitempty"`
	TelefoonContactpersoon string `xml:" TelefoonContactpersoon,omitempty" json:"TelefoonContactpersoon,omitempty"`
}

type DoorUBinnenland struct {
	OmzetHoog                  float64 `xml:" OmzetHoog,omitempty" json:"OmzetHoog,omitempty"`
	OmzetLaag                  float64 `xml:" OmzetLaag,omitempty" json:"OmzetLaag,omitempty"`
	OmzetbelastingHoog         float64 `xml:" OmzetbelastingHoog,omitempty" json:"OmzetbelastingHoog,omitempty"`
	OmzetbelastingLaag         float64 `xml:" OmzetbelastingLaag,omitempty" json:"OmzetbelastingLaag,omitempty"`
	OmzetPriveGebruik          float64 `xml:" OmzetPriveGebruik,omitempty" json:"OmzetPriveGebruik,omitempty"`
	OmzetbelastingPriveGebruik float64 `xml:" OmzetbelastingPriveGebruik,omitempty" json:"OmzetbelastingPriveGebruik,omitempty"`
	OmzetBelastMet0OfNiet      float64 `xml:" OmzetBelastMet0OfNiet,omitempty" json:"OmzetBelastMet0OfNiet,omitempty"`
}

type DoorUBuitenland struct {
	OmzetBuitenEUUitvoer float64 `xml:"OmzetBuitenEUUitvoer,omitempty" json:"OmzetBuitenEUUitvoer,omitempty"`
	OmzetBinnenEU        float64 `xml:"OmzetBinnenEU,omitempty" json:"OmzetBinnenEU,omitempty"`
}

type IdentificerendeGegevens struct {
	Aangiftejaar    string          `xml:" Aangiftejaar,omitempty" json:"Aangiftejaar,omitempty"`
	Aangiftetijdvak string          `xml:" Aangiftetijdvak,omitempty" json:"Aangiftetijdvak,omitempty"`
	Mediumsoort     string          `xml:" Mediumsoort,omitempty" json:"Mediumsoort,omitempty"`
	MessageID       string          `xml:" MessageID,omitempty" json:"MessageID,omitempty"`
	OBNummer        string          `xml:" OBNummer,omitempty" json:"OBNummer,omitempty"`
	TijdstipAanmaak TijdstipAanmaak `xml:" TijdstipAanmaak,omitempty" json:"TijdstipAanmaak,omitempty"`
	Valutacode      string          `xml:" Valutacode,omitempty" json:"Valutacode,omitempty"`
}

type TijdstipAanmaak time.Time

func (t *TijdstipAanmaak) UnmarshalText(text []byte) error {
	format := "2006-01-02T15:04:05.999999999"
	t2, err := time.Parse(format, string(text))
	*t = TijdstipAanmaak(t2)
	return err
}

type VatReturnResponse struct {
	AangifteOmzetbelasting *AangifteOmzetbelasting `xml:"vatReturn>AangifteOmzetbelasting,omitempty"`
}

type IctReturnResponse struct {
	OpgaafICP *OpgaafICP `xml:"vatReturn>OpgaafICP,omitempty"`
}

func VatReturnFromXml(data []byte) (*AangifteOmzetbelasting, error) {
	vatReturnResponse := VatReturnResponse{}
	err := xml.Unmarshal(data, &vatReturnResponse)
	if err != nil {
		return nil, err
	}

	return vatReturnResponse.AangifteOmzetbelasting, nil
}

func IctReturnFromXml(data []byte) (*OpgaafICP, error) {
	ictReturnResponse := IctReturnResponse{}
	err := xml.Unmarshal(data, &ictReturnResponse)
	if err != nil {
		return nil, err
	}

	return ictReturnResponse.OpgaafICP, nil
}
