package bills

type Form struct {
	DistributionCode   string           `xml:"distribution-code"`
	CalendarName       string           `xml:"calendar"`
	CongressName       string           `xml:"congress"`
	SessionName        string           `xml:"session"`
	EnrolledDateline   string           `xml:"enrolled-dateline"`
	LegislationName    string           `xml:"legis-num"`
	AssociatedDocs     []*AssociatedDoc `xml:"associated-doc"`
	CurrentChamberName string           `xml:"current-chamber"`
	Actions            []*Action        `xml:"action"`
	TypeName           string           `xml:"legis-type"`
}

type AssociatedDoc struct {
}

type Action struct {
	StageCode   string         `xml:"stage,attr"`
	Description []InlineMarkup `xml:"action-desc"`
	Instruction []string       `xml:"action-instruction"`
}
