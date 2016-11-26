package bills

// StructuralVisitor is the argument to StructuralMarkup's Walk method, and
// its methods are called as the descendent elements are traversed.
type StructuralVisitor interface {

	// EnterCaption and ExitCaption delimit the calls to EnterEnum, ExitEnum,
	// EnterHeader and ExitHeader to mark the (virtual) boundaries of the
	// "caption" of this item, for the benefit of visitors that e.g. want
	// to combine the caption elements together into a single string.
	//
	// Not called if both the enumerator and the header are nil.
	EnterCaption(Structural)
	ExitCaption(Structural)

	EnterEnum(InlineMarkup) InlineVisitor
	ExitEnum(InlineMarkup, InlineVisitor)

	EnterHeader(InlineMarkup) InlineVisitor
	ExitHeader(InlineMarkup, InlineVisitor)

	EnterText(InlineMarkup) InlineVisitor
	ExitText(InlineMarkup, InlineVisitor)

	EnterQuotedBlock(*QuotedBlock)
	ExitQuotedBlock(*QuotedBlock)

	VisitGraphic(*Graphic)

	VisitFormula(*Formula)

	EnterTOC(*TableOfContents) TOCVisitor
	ExitTOC(*TableOfContents, TOCVisitor)

	EnterTable(*Table) TableVisitor
	ExitTable(*Table, TableVisitor)

	EnterList(*List) ListVisitor
	ExitList(*List, ListVisitor)

	EnterStructuralElement(Structural) StructuralVisitor
	ExitStructuralElement(Structural, StructuralVisitor)
}

type InlineVisitor interface {

	// EnterInlineElement and ExitInlineElement are used for element types
	// that have more inline markup inside them.
	EnterInlineElement(Inline) InlineVisitor
	ExitInlineElement(Inline, InlineVisitor)

	// VisitInlineElement is used for element types that are leaf nodes.
	VisitInlineElement(Inline)

	// VisitInlineText is used for raw text strings around and within the
	// markup.
	VisitInlineText(Text)
}

type TableVisitor interface {
	EnterTableGroup(*TableGroup)
	ExitTableGroup(*TableGroup)

	EnterTableHead(*TableRowSeq)
	ExitTableHead(*TableRowSeq)

	EnterTableBody(*TableRowSeq)
	ExitTableBody(*TableRowSeq)

	EnterTableRow(*TableRow)
	ExitTableRow(*TableRow)

	EnterTableCell(InlineMarkup)
	ExitTableCell(InlineMarkup)
}

type TOCVisitor interface {
	EnterTOCEntry(TOCEntry)
	ExitTOCEntry(TOCEntry)

	EnterTOCEnum(InlineMarkup) InlineVisitor
	ExitTOCEnum(InlineMarkup, InlineVisitor)

	EnterTOCHeading(InlineMarkup) InlineVisitor
	ExitTOCHeading(InlineMarkup, InlineVisitor)

	EnterTOCQuoted(TOCEntry) TOCVisitor
	ExitTOCQuoted(TOCEntry, TOCVisitor)
}

type ListVisitor interface {
	EnterListItem(InlineMarkup) InlineVisitor
	ExitListItem(InlineMarkup) InlineVisitor
}

func (m StructuralMarkup) Walk(v StructuralVisitor) {
	for _, node := range m {
		structuralWalk(v, node)
	}
}

func (m InlineMarkup) Walk(v InlineVisitor) {
	for _, n := range m {
		if text, ok := n.(Text); ok {
			v.VisitInlineText(text)
			continue
		}

		cn := n.ChildNodes()
		if cn == nil {
			v.VisitInlineElement(n)
		} else {
			cv := v.EnterInlineElement(cn)
			if cv != nil {
				cn.Walk(v)
				v.ExitInlineElement(cn, cv)
			}
		}
	}
}

func structuralWalk(v StructuralVisitor, n Structural) {
	childNodes := n.ChildElements()

	cv := v.EnterStructuralElement(n)
	if cv == nil {
		return
	}

	enum := n.Enumerator()
	header := n.Header()

	if enum != nil || header != nil {
		v.EnterCaption(n)
		if enum != nil {
			cv := v.EnterEnum(enum)
			if cv != nil {
				enum.Walk(cv)
				v.ExitEnum(enum, cv)
			}
		}
		if header != nil {
			cv := v.EnterHeader(header)
			if cv != nil {
				header.Walk(cv)
				v.ExitHeader(header, cv)
			}
		}
		v.ExitCaption(n)
	}

	childNodes.Walk(cv)

	v.ExitStructuralElement(n, cv)
}
