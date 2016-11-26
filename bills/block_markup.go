package bills

import (
	"encoding/xml"
)

// BlockMarkup represents a sequence of blocks contained within a
// structural element in a bill.
type BlockMarkup []Block

func (m *BlockMarkup) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	*m = make([]Block, 0, 1)
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
			obj, err := decodeBlockElement(d, t)
			if err != nil {
				return err
			}
			*m = append(*m, obj)
		}
	}
}

// Block represents an element rendered in a block-formatting layout
// (on a line of its own).
//
// Blocks are the elements that have the most variation between types.
// There is no common interface between blocks, so code using Blocks
// must always use a type switch or similar construct to find the
// specific type of block and then implement handling for each block
// type separately.
type Block interface {
	// Do-nothing method that just declares that this interface is implemented.
	Block()
}

// isBlockElement returns true if the given XML element name corresponds
// to a supported block element type. Used by the structural element parser
// to recognize block elements that are directly nested inside structural
// elements.
func isBlockElement(name xml.Name) bool {
	switch name.Local {
	case "quoted-block", "graphic", "formula", "toc", "table", "list":
		return true
	default:
		return false
	}
}

func decodeBlockElement(d *xml.Decoder, start xml.StartElement) (Block, error) {
	switch start.Name.Local {
	case "formula":
		ret := &Formula{}
		err := d.DecodeElement(ret, &start)
		return ret, err
	case "graphic":
		ret := &Graphic{}
		err := d.DecodeElement(ret, &start)
		return ret, err
	case "list":
		ret := &List{}
		err := d.DecodeElement(ret, &start)
		return ret, err
	case "quoted-block":
		ret := &QuotedBlock{}
		err := d.DecodeElement(ret, &start)
		return ret, err
	case "table":
		ret := &Table{}
		err := d.DecodeElement(ret, &start)
		return ret, err
	case "toc":
		ret := &TableOfContents{}
		err := d.DecodeElement(ret, &start)
		return ret, err
	default:
		ret := &UnsupportedBlockElement{}
		err := d.DecodeElement(ret, &start)
		return ret, err
	}
}

type QuotedBlock struct {
	ActName      string `xml:"act-name,attr"`
	Id           string `xml:"id,attr"`
	ParsableCite string `xml:"parsable-cite,attr"`
	StyleCode    string `xml:"style,attr"`

	// Quoted blocks have a mixed content model. Content elements can either be
	// implementations of Block or Structural or they can be InlineMarkup
	// values representing directly-quoted paragraphs of text.
	Content []interface{}

	AfterText string
}

func (n *QuotedBlock) Block() {
}

func (n *QuotedBlock) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	*n = QuotedBlock{}
	err := decodeXMLAttrs(n, start)
	if err != nil {
		return err
	}

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

			if t.Name.Local == "after-quoted-block" {
				err := d.DecodeElement(&n.AfterText, &t)
				if err != nil {
					return err
				}
				continue
			}

			// Quote blocks can contain both block elements and structural
			// elements, so we need to recognize which one we're working
			// with and decode as appropriate.
			switch {
			case isBlockElement(t.Name):
				obj, err := decodeBlockElement(d, t)
				if err != nil {
					return err
				}
				n.Content = append(n.Content, obj)
			case t.Name.Local == "text":
				var obj InlineMarkup
				err := d.DecodeElement(&obj, &t)
				if err != nil {
					return err
				}
				n.Content = append(n.Content, obj)
			default:
				obj, err := decodeStructuralElement(d, t)
				if err != nil {
					return err
				}
				n.Content = append(n.Content, obj)
			}
		}
	}

	return nil
}

type Graphic struct {
	Depth          string `xml:"depth,attr"`
	File           string `xml:"file,attr"`
	Description    string `xml:"graphic-desc,attr"`
	Indent         string `xml:"graphic-indent,attr"`
	HorizAlignCode string `xml:"halign,attr"`
	RotationCode   string `xml:"rotation,attr`
	Span           string `xml:"span,attr"`
}

func (n *Graphic) Block() {
}

type Formula struct {
	Id      string   `xml:"id,attr"`
	Graphic *Graphic `xml:"graphic"`
}

func (n *Formula) Block() {
}

type TableOfContents struct {
	ContainerLevelCode    string `xml:"container-level,attr"`
	IdRef                 string `xml:"idref,attr"`
	LowestBoldedLevelCode string `xml:"lowest-bolded-level,attr"`
	LowestLevelCode       string `xml:"lowest-level,attr"`
	QuotedBlockCode       string `xml:"quoted-block,attr"`
	RegenerationCode      string `xml:"regeneration,attr"`

	Header               InlineMarkup `xml:"header"`
	InstructiveParagraph InlineMarkup `xml:"instructive-para"`

	Entries TOCList
}

func (n *TableOfContents) Block() {
}

func (n *TableOfContents) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	err := decodeXMLAttrs(n, start)
	if err != nil {
		return err
	}

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
			switch t.Name.Local {
			case "header":
				err := d.DecodeElement(&n.Header, &t)
				if err != nil {
					return err
				}
			case "instructive-para":
				err := d.DecodeElement(&n.InstructiveParagraph, &t)
				if err != nil {
					return err
				}
			default:
				obj, err := decodeTOCEntry(d, t)
				if err != nil {
					return err
				}
				n.Entries = append(n.Entries, obj)
			}
		}
	}
	return nil
}

type Table struct {
	Titles       []string      `xml:"ttitle"`
	Descriptions []string      `xml:"tdesc"`
	Groups       []*TableGroup `xml:"tgroup"`
}

func (n *Table) Block() {
}

type List struct {
	Items []InlineMarkup `xml:"list-item"`
}

func (n *List) Block() {
}

type UnsupportedBlockElement struct {
	Name    xml.Name
	Attrs   map[xml.Name]string
	Content []byte
}

func (n *UnsupportedBlockElement) Block() {
}

func (n *UnsupportedBlockElement) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	n.Name = start.Name
	n.Attrs = make(map[xml.Name]string, len(start.Attr))
	for _, attr := range start.Attr {
		n.Attrs[attr.Name] = attr.Value
	}
	d.DecodeElement(&n.Content, &start)
	return nil
}
