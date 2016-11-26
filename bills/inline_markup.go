package bills

import (
	"encoding/xml"
	"strings"
)

// InlineMarkup represents a mixture of raw text and markup elements
// that combine to produce a rich-text string.
type InlineMarkup []Inline

func (m *InlineMarkup) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	*m = make([]Inline, 0, 1)
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}

		switch t := token.(type) {
		case xml.EndElement:
			// all done!
			return nil
		case xml.StartElement:
			obj, err := m.decodeMarkup(d, t)
			if err != nil {
				return err
			}
			*m = append(*m, obj)
		case xml.CharData:
			*m = append(*m, Text(t))
		}
	}
}

func (m *InlineMarkup) decodeMarkup(d *xml.Decoder, start xml.StartElement) (Inline, error) {
	switch start.Name.Local {
	case "added-phrase":
		ret := &AddedPhrase{}
		err := d.DecodeElement(ret, &start)
		return ret, err
	case "act-name":
		ret := &ActName{}
		err := d.DecodeElement(ret, &start)
		return ret, err
	case "bold":
		ret := &Bold{}
		err := d.DecodeElement(ret, &start)
		return ret, err
	case "committee-name":
		ret := &CommitteeName{}
		err := d.DecodeElement(ret, &start)
		return ret, err
	case "cosponsor":
		ret := &CosponsorName{}
		err := d.DecodeElement(ret, &start)
		return ret, err
	case "definition":
		ret := &Definition{}
		err := d.DecodeElement(ret, &start)
		return ret, err
	case "deleted-phrase":
		ret := &DeletedPhrase{}
		err := d.DecodeElement(ret, &start)
		return ret, err
	case "editorial":
		ret := &Editorial{}
		err := d.DecodeElement(ret, &start)
		return ret, err
	case "effective-date":
		ret := &EffectiveDate{}
		err := d.DecodeElement(ret, &start)
		return ret, err
	case "external-xref":
		ret := &ExternalCrossReference{}
		err := d.DecodeElement(ret, &start)
		return ret, err
	case "footnote":
		ret := &Footnote{}
		err := d.DecodeElement(ret, &start)
		return ret, err
	case "footnote-ref":
		ret := &Footnote{}
		err := d.DecodeElement(ret, &start)
		return ret, err
	case "fraction":
		ret := &Fraction{}
		err := d.DecodeElement(ret, &start)
		return ret, err
	case "internal-xref":
		ret := &InternalCrossReference{}
		err := d.DecodeElement(ret, &start)
		return ret, err
	case "italic":
		ret := &Italic{}
		err := d.DecodeElement(ret, &start)
		return ret, err
	case "nonsponsor":
		ret := &NonsponsorName{}
		err := d.DecodeElement(ret, &start)
		return ret, err
	case "linebreak":
		ret := &LineBreak{}
		err := d.DecodeElement(ret, &start)
		return ret, err
	case "nobreak":
		ret := &NoBreak{}
		err := d.DecodeElement(ret, &start)
		return ret, err
	case "omitted-text":
		ret := &OmittedText{}
		err := d.DecodeElement(ret, &start)
		return ret, err
	case "pagebreak":
		ret := &PageBreak{}
		err := d.DecodeElement(ret, &start)
		return ret, err
	case "quote":
		ret := &InlineQuote{}
		err := d.DecodeElement(ret, &start)
		return ret, err
	case "short-title":
		ret := &ShortTitle{}
		err := d.DecodeElement(ret, &start)
		return ret, err
	case "sponsor":
		ret := &SponsorName{}
		err := d.DecodeElement(ret, &start)
		return ret, err
	case "subscript":
		ret := &Subscript{}
		err := d.DecodeElement(ret, &start)
		return ret, err
	case "superscript":
		ret := &Subscript{}
		err := d.DecodeElement(ret, &start)
		return ret, err
	case "term":
		ret := &Subscript{}
		err := d.DecodeElement(ret, &start)
		return ret, err
	default:
		ret := &UnsupportedInlineElement{}
		err := d.DecodeElement(ret, &start)
		return ret, err
	}
}

// Text returns the raw, unformatted text within inline markup.
//
// Using the result for presentation to humans can be dangerous since
// discarding certain markup elements would change the semantics of the text.
func (m InlineMarkup) Text() string {
	texts := make([]string, len(m))
	for i, elem := range m {
		texts[i] = elem.Text()
	}
	return strings.Join(texts, "")
}

func (n InlineMarkup) ChildNodes() InlineMarkup {
	return n
}

// Inline represents any tree node that can appear in an inline markup tree.
type Inline interface {

	// Returns the raw text contained within this node and any child nodes.
	Text() string

	// Returns child nodes, or nil if this node type is a leaf.
	ChildNodes() InlineMarkup
}

// Text represents a raw text string within InlineMarkup
type Text string

func (t Text) Text() string {
	return string(t)
}

func (t Text) ChildNodes() InlineMarkup {
	return nil
}

// UnsupportedInlineElement is a placeholder node type for inline nodes we
// don't yet support.
//
// Callers should ignore nodes of this type except to walk to the children
// when traversing the graph; future versions of this package may start
// to support the given element, which would be a breaking change for any
// caller that specifically depends on recieving UnsupportedInlineElement
// instances.
type UnsupportedInlineElement struct {
	Name  xml.Name
	Attrs map[xml.Name]string
	InlineMarkup
}

func (n *UnsupportedInlineElement) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	n.Name = start.Name
	n.Attrs = make(map[xml.Name]string, len(start.Attr))
	for _, attr := range start.Attr {
		n.Attrs[attr.Name] = attr.Value
	}
	return n.InlineMarkup.UnmarshalXML(d, start)
}

type AddedPhrase struct {
	InlineMarkup
}

type DeletedPhrase struct {
	InlineMarkup
}

type Definition struct {
	InlineMarkup
}

type Editorial struct {
	InlineMarkup
}

type EffectiveDate struct {
	InlineMarkup
}

type Fraction struct {
	InlineMarkup
}

type Footnote struct {
	InlineMarkup
	Id string `xml:"id,attr"`
}

func (n *Footnote) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	err := decodeXMLAttrs(n, start)
	if err != nil {
		return err
	}
	return n.InlineMarkup.UnmarshalXML(d, start)
}

type InternalCrossReference struct {
	InlineMarkup
	IdReference string `xml:"idref,attr"`
}

func (n *InternalCrossReference) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	err := decodeXMLAttrs(n, start)
	if err != nil {
		return err
	}
	return n.InlineMarkup.UnmarshalXML(d, start)
}

type ExternalCrossReference struct {
	InlineMarkup
	TargetTypeCode string `xml:"legal-doc,attr"`
	ParsableCite   string `xml:"parsable-cite,attr"`
}

func (n *ExternalCrossReference) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	err := decodeXMLAttrs(n, start)
	if err != nil {
		return err
	}
	return n.InlineMarkup.UnmarshalXML(d, start)
}

type Superscript struct {
	InlineMarkup
}

type Subscript struct {
	InlineMarkup
}

type Bold struct {
	InlineMarkup
}

type Italic struct {
	InlineMarkup
}

type InlineQuote struct {
	InlineMarkup
}

type ActName struct {
	InlineMarkup
}

type CommitteeName struct {
	InlineMarkup
	CommitteeId string `xml:"committee-id,attr"`
}

func (n *CommitteeName) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	err := decodeXMLAttrs(n, start)
	if err != nil {
		return err
	}
	return n.InlineMarkup.UnmarshalXML(d, start)
}

type SponsorName struct {
	InlineMarkup
	NameId    string `xml:"name-id,attr"`
	ByRequest string `xml:"by-request,attr"`
}

func (n *SponsorName) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	err := decodeXMLAttrs(n, start)
	if err != nil {
		return err
	}
	return n.InlineMarkup.UnmarshalXML(d, start)
}

type CosponsorName struct {
	InlineMarkup
	NameId string `xml:"name-id,attr"`
}

func (n *CosponsorName) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	err := decodeXMLAttrs(n, start)
	if err != nil {
		return err
	}
	return n.InlineMarkup.UnmarshalXML(d, start)
}

type NonsponsorName struct {
	InlineMarkup
	NameId string `xml:"name-id,attr"`
}

func (n *NonsponsorName) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	err := decodeXMLAttrs(n, start)
	if err != nil {
		return err
	}
	return n.InlineMarkup.UnmarshalXML(d, start)
}

type ShortTitle struct {
	InlineMarkup
}

type Term struct {
	InlineMarkup
}

// LeafNode can be embedded in structs representing inline elements that have
// no content.
type LeafNode struct {
}

func (n *LeafNode) Text() string {
	return ""
}

func (n *LeafNode) ChildNodes() InlineMarkup {
	return nil
}

type FootnoteRef struct {
	LeafNode
	IdRef string `xml:"idref,attr"`
}

func (n *FootnoteRef) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return decodeXMLAttrs(n, start)
}

type OmittedText struct {
	LeafNode
}

type LineBreak struct {
	LeafNode
}

type NoBreak struct {
	LeafNode
}

type PageBreak struct {
	LeafNode
}
