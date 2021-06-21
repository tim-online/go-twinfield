package soap

type Creditmanagement struct {
	ResponsibleUser string  `xml:"responsibleuser,omitempty"` // The credit manager.
	BaseCreditLimit float64 `xml:"basecreditlimit,omitempty"` // The credit limit amount.
	SendReminder    string  `xml:"sendreminder,omitempty"`    // Determines if and how a customer will be reminded.
	ReminderEmail   string  `xml:"reminderemail,omitempty"`   // Mandatory if sendreminder is email
	Blocked         string  `xml:"blocked,omitempty"`         // Indicates if related projects for this customer are blocked in time & expenses.
	Freetext1       string  `xml:"freetext1,omitempty"`       // Right of use.
	Freetext2       string  `xml:"freetext2,omitempty"`       // Segment code.
	Freetext3       string  `xml:"freetext3,omitempty"`       //  Not in use.
	Comment         string  `xml:"comment,omitempty"`         // Comment.
}
