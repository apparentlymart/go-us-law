package bills

import (
	"encoding/xml"
)

type StructuralMarkup []Structural

func (m StructuralMarkup) ChildElements() StructuralMarkup {
	return m
}

func (m *StructuralMarkup) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	*m = make([]Structural, 0, 1)
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
			obj, err := decodeStructuralElement(d, t)
			if err != nil {
				return err
			}
			*m = append(*m, obj)
		}
	}
}

func decodeStructuralElement(d *xml.Decoder, start xml.StartElement) (Structural, error) {
	switch start.Name.Local {
	case "chapter":
		ret := &Chapter{}
		err := d.DecodeElement(ret, &start)
		return ret, err
	case "clause":
		ret := &Clause{}
		err := d.DecodeElement(ret, &start)
		return ret, err
	case "division":
		ret := &Division{}
		err := d.DecodeElement(ret, &start)
		return ret, err
	case "item":
		ret := &Item{}
		err := d.DecodeElement(ret, &start)
		return ret, err
	case "paragraph":
		ret := &Paragraph{}
		err := d.DecodeElement(ret, &start)
		return ret, err
	case "part":
		ret := &Part{}
		err := d.DecodeElement(ret, &start)
		return ret, err
	case "section":
		ret := &Section{}
		err := d.DecodeElement(ret, &start)
		return ret, err
	case "title":
		ret := &Title{}
		err := d.DecodeElement(ret, &start)
		return ret, err
	case "subchapter":
		ret := &Chapter{}
		err := d.DecodeElement(ret, &start)
		return ret, err
	case "subclause":
		ret := &Subclause{}
		err := d.DecodeElement(ret, &start)
		return ret, err
	case "subdivision":
		ret := &Subdivision{}
		err := d.DecodeElement(ret, &start)
		return ret, err
	case "subitem":
		ret := &Subitem{}
		err := d.DecodeElement(ret, &start)
		return ret, err
	case "subparagraph":
		ret := &Subparagraph{}
		err := d.DecodeElement(ret, &start)
		return ret, err
	case "subpart":
		ret := &Part{}
		err := d.DecodeElement(ret, &start)
		return ret, err
	case "subsection":
		ret := &Subsection{}
		err := d.DecodeElement(ret, &start)
		return ret, err
	case "subtitle":
		ret := &Subtitle{}
		err := d.DecodeElement(ret, &start)
		return ret, err
	default:
		ret := &UnsupportedStructuralElement{}
		err := d.DecodeElement(ret, &start)
		return ret, err
	}
}

type Structural interface {
	Enumerator() InlineMarkup
	Header() InlineMarkup
	Text() InlineMarkup
	Blocks() BlockMarkup
	ChildElements() StructuralMarkup
	ContinuationText() InlineMarkup
}

type StructuralElement struct {
	enumerator       InlineMarkup
	header           InlineMarkup
	text             InlineMarkup
	blocks           BlockMarkup
	childElements    StructuralMarkup
	continuationText InlineMarkup
}

func (m *StructuralElement) Enumerator() InlineMarkup {
	return m.enumerator
}

func (m *StructuralElement) Header() InlineMarkup {
	return m.header
}

func (m *StructuralElement) Text() InlineMarkup {
	return m.text
}

func (m *StructuralElement) Blocks() BlockMarkup {
	return m.blocks
}

func (m *StructuralElement) ChildElements() StructuralMarkup {
	return m.childElements
}

func (m *StructuralElement) ContinuationText() InlineMarkup {
	return m.text
}

func (m *StructuralElement) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	*m = StructuralElement{}
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

			// Block elements can be directly nested in structural elements.
			if isBlockElement(t.Name) {
				obj, err := decodeBlockElement(d, t)
				if err != nil {
					return err
				}
				m.blocks = append(m.blocks, obj)
				continue
			}

			switch t.Name.Local {
			case "continuation-text":
				err := d.DecodeElement(&m.continuationText, &t)
				if err != nil {
					return err
				}
			case "enum":
				err := d.DecodeElement(&m.enumerator, &t)
				if err != nil {
					return err
				}
			case "header":
				err := d.DecodeElement(&m.header, &t)
				if err != nil {
					return err
				}
			case "text":
				err := d.DecodeElement(&m.text, &t)
				if err != nil {
					return err
				}
			default:
				obj, err := decodeStructuralElement(d, t)
				if err != nil {
					return err
				}
				m.childElements = append(m.childElements, obj)
			}
		}
	}
}

type Chapter struct {
	StructuralElement
}

type SubChapter struct {
	StructuralElement
}

type Clause struct {
	StructuralElement
}

type Subclause struct {
	StructuralElement
}

type Division struct {
	StructuralElement
}

type Subdivision struct {
	StructuralElement
}

type Item struct {
	StructuralElement
}

type Subitem struct {
	StructuralElement
}

type Paragraph struct {
	StructuralElement
}

type Subparagraph struct {
	StructuralElement
}

type Part struct {
	StructuralElement
}

type Subpart struct {
	StructuralElement
}

type Section struct {
	StructuralElement
}

type Subsection struct {
	StructuralElement
}

type Title struct {
	StructuralElement
}

type Subtitle struct {
	StructuralElement
}

// UnsupportedStructuralElement is a placeholder node type for structural nodes
// we don't yet support.
//
// Callers should ignore nodes of this type except to walk to the children
// when traversing the graph; future versions of this package may start
// to support the given element, which would be a breaking change for any
// caller that specifically depends on recieving UnsupportedStructuralElement
// instances.
type UnsupportedStructuralElement struct {
	Name  xml.Name
	Attrs map[xml.Name]string
	StructuralElement
}

func (n *UnsupportedStructuralElement) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	n.Name = start.Name
	n.Attrs = make(map[xml.Name]string, len(start.Attr))
	for _, attr := range start.Attr {
		n.Attrs[attr.Name] = attr.Value
	}
	return n.StructuralElement.UnmarshalXML(d, start)
}
