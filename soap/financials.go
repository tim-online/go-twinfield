package soap

type Financials struct {
	MatchType    string `xml:"matchtype,omitempty"`    // Fixed value customersupplier.
	AccountType  string `xml:"accounttype,omitempty"`  // Fixed value inherit.
	Subanalyse   string `xml:"subanalyse,omitempty"`   // Fixed value false.
	DueDays      int    `xml:"duedays,omitempty"`      // The number of due days.
	PayAvailable string `xml:"payavailable,omitempty"` // Specifies the dimension level. Normally the level of customers is level 2. Read-only attribute.
	// The option none is only allowed in case payavailable is set to false.
	// The option paymentfile is only allowed in case payavailable is set to true.
	MeansOfPayment     string      `xml:"meansofpayment,omitempty"`     // Determines if direct debit is possible.
	Paycode            string      `xml:"paycode,omitempty"`            // The code of the payment type in case direct debit is possible.
	Ebilling           string      `xml:"ebilling,omitempty"`           // Determines if the sales invoices will be sent electronically to the customer.
	EbillMail          string      `xml:"ebillmail,omitempty"`          // The mail address the electronic sales invoice is sent to.
	SubstitutionLevel  int         `xml:"substitutionlevel,omitempty"`  // Level of the balancesheet account. Fixed value 1.
	SubstituteWith     string      `xml:"substitutewith,omitempty"`     // Default customer balancesheet account.
	RelationsReference interface{} `xml:"relationsreference,omitempty"` // Not in use.
	VATType            interface{} `xml:"vattype,omitempty"`            // Not in use.
	VATCode            string      `xml:"vatcode,omitempty"`            // Default VAT code.
	VATObligatory      interface{} `xml:"vatobligatory,omitempty"`      // Not in use.
	PerformanceType    interface{} `xml:"performancetype,omitempty"`    // Not in use.
	// Collect mandate information.
	// Apply this information only when the customer invoices are collected by SEPA direct debit.
	CollectMandate CollectMandate `xml:"collectmandate,omitempty"`
	// Collection schema information.
	// Apply this information only when the customer invoices are collected by SEPA direct debit.
	CollectionSchema CollectionSchema `xml:"collectionschema,omitempty"`
}

type CollectMandate struct {
}

type CollectionSchema struct {
}
