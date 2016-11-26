package bills

type ActionDate struct {
	HumanReadable   string `xml:",chardata"`
	EventDate       *Date  `xml:"date,attr"`
	LegislativeDate *Date  `xml:"legis-day,attr"`
}
