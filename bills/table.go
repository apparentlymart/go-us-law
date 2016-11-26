package bills

type TableGroup struct {
	Columns []*TableColumn `xml:"colspec"`
	Head    *TableRowSeq   `xml:"thead"`
	Bodies  []*TableRowSeq `xml:"tbody"`
}

type TableColumn struct {
	// TODO: column model
}

type TableRowSeq struct {
	Rows []TableRow `xml:"row"`
}

type TableRow struct {
	Entries []InlineMarkup `xml:"entry"`
}
