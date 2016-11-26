package bills

// StructuralVisitorImpl provides all of the methods of StructuralVisitor with no-op
// implementations. Embedding this into another visitor struct avoids the
// need to implement all of the methods.
//
// Note that you should almost always implement EnterStructuralElement,
// because the default implementation will not traverse any child elements
// and thus the walk will visit nothing.
type StructuralVisitorImpl struct {
}

func (i *StructuralVisitorImpl) EnterCaption(Structural) {
}

func (i *StructuralVisitorImpl) ExitCaption(Structural) {
}

func (i *StructuralVisitorImpl) EnterEnum(InlineMarkup) InlineVisitor {
	return nil
}

func (i *StructuralVisitorImpl) ExitEnum(InlineMarkup, InlineVisitor) {
}

func (i *StructuralVisitorImpl) EnterHeader(InlineMarkup) InlineVisitor {
	return nil
}

func (i *StructuralVisitorImpl) ExitHeader(InlineMarkup, InlineVisitor) {
}

func (i *StructuralVisitorImpl) EnterText(InlineMarkup) InlineVisitor {
	return nil
}

func (i *StructuralVisitorImpl) ExitText(InlineMarkup, InlineVisitor) {
}

func (i *StructuralVisitorImpl) EnterQuotedBlock(*QuotedBlock) {
}

func (i *StructuralVisitorImpl) ExitQuotedBlock(*QuotedBlock) {
}

func (i *StructuralVisitorImpl) VisitGraphic(*Graphic) {
}

func (i *StructuralVisitorImpl) VisitFormula(*Formula) {
}

func (i *StructuralVisitorImpl) EnterTOC(*TableOfContents) TOCVisitor {
	return nil
}

func (i *StructuralVisitorImpl) ExitTOC(*TableOfContents, TOCVisitor) {
}

func (i *StructuralVisitorImpl) EnterTable(*Table) TableVisitor {
	return nil
}

func (i *StructuralVisitorImpl) ExitTable(*Table, TableVisitor) {
}

func (i *StructuralVisitorImpl) EnterList(*List) ListVisitor {
	return nil
}

func (i *StructuralVisitorImpl) ExitList(*List, ListVisitor) {
}

func (i *StructuralVisitorImpl) EnterStructuralElement(Structural) StructuralVisitor {
	return nil
}

func (i *StructuralVisitorImpl) ExitStructuralElement(Structural, StructuralVisitor) {
}

// InlineVisitorImpl provides all of the methods of InlineVisitor with no-op
// implementations. Embedding this into another visitor struct avoids the
// need to implement all of the methods.
type InlineVisitorImpl struct {
}

func (i *InlineVisitorImpl) EnterInlineElement(Inline) InlineVisitor {
	return nil
}

func (i *InlineVisitorImpl) ExitInlineElement(Inline, InlineVisitor) {
}

func (i *InlineVisitorImpl) VisitInlineElement(Inline) {
}

func (i *InlineVisitorImpl) VisitInlineText(Text) {
}

// TableVisitorImpl provides all of the methods of TableVisitor with no-op
// implementations. Embedding this into another visitor struct avoids the
// need to implement all of the methods.
type TableVisitorImpl struct {
}

func (i *StructuralVisitorImpl) EnterTableGroup(*TableGroup) {
}

func (i *StructuralVisitorImpl) ExitTableGroup(*TableGroup) {
}

func (i *StructuralVisitorImpl) EnterTableBody(*TableRowSeq) {
}

func (i *StructuralVisitorImpl) ExitTableHead(*TableRowSeq) {
}

func (i *StructuralVisitorImpl) EnterTableHead(*TableRowSeq) {
}

func (i *StructuralVisitorImpl) ExitTableBody(*TableRowSeq) {
}

func (i *StructuralVisitorImpl) EnterTableRow(*TableRow) {
}

func (i *StructuralVisitorImpl) ExitTableRow(*TableRow) {
}

func (i *StructuralVisitorImpl) EnterTableCell(InlineMarkup) InlineVisitor {
	return nil
}

func (i *StructuralVisitorImpl) ExitTableCell(InlineMarkup) {
}

// TOCVisitorImpl provides all of the methods of TOCVisitor with no-op
// implementations. Embedding this into another visitor struct avoids
// the need to implement all of the methods.
type TOCVisitorImpl struct {
}

func (i *TOCVisitorImpl) EnterTOCEntry(TOCEntry) {
}

func (i *TOCVisitorImpl) ExitTOCEntry(TOCEntry) {
}

func (i *TOCVisitorImpl) EnterTOCEnum(InlineMarkup) InlineVisitor {
	return nil
}

func (i *TOCVisitorImpl) ExitTOCEnum(InlineMarkup, InlineVisitor) {
}

func (i *TOCVisitorImpl) EnterTOCHeading(InlineMarkup) InlineVisitor {
	return nil
}

func (i *TOCVisitorImpl) ExitTOCHeading(InlineMarkup, InlineVisitor) {
}

func (i *TOCVisitorImpl) EnterTOCQuoted() TOCVisitor {
	return nil
}

func (i *TOCVisitorImpl) ExitTOCQuoted(TOCEntry, TOCVisitor) {
}

// ListVisitorImpl provides all of the methods of ListVisitor with no-op
// implementations. Embedding this into another visitor struct avoids
// the need to implement all of the methods.
type ListVisitorImpl struct {
}

func (i *TOCVisitorImpl) EnterListItem(InlineMarkup) InlineVisitor {
	return nil
}

func (i *TOCVisitorImpl) ExitListItem(InlineMarkup, InlineVisitor) {
}
