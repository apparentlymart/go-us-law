package bills

import (
	"encoding/xml"
)

type TOCList []TOCEntry

func (m *TOCList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	*m = make([]TOCEntry, 0, 1)
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
			obj, err := decodeTOCEntry(d, t)
			if err != nil {
				return err
			}
			*m = append(*m, obj)
		}
	}
}

type TOCEntry interface {
	// Placeholder method to indicate implementation of this interface
	TOCEntry() TOCEntry
}

func decodeTOCEntry(d *xml.Decoder, start xml.StartElement) (TOCEntry, error) {
	switch start.Name.Local {
	case "toc-entry":
		ret := &SimpleTOCEntry{}
		err := d.DecodeElement(ret, &start)
		return ret, err
	case "multi-column-toc-entry":
		ret := &MultiColumnTOCEntry{}
		err := d.DecodeElement(ret, &start)
		return ret, err
	case "toc-quoted-entry":
		ret := &QuotedSimpleTOCEntry{}
		err := d.DecodeElement(ret, &start)
		return ret, err
	case "toc-multi-column-quoted-entry":
		ret := &QuotedMultiColumnTOCEntry{}
		err := d.DecodeElement(ret, &start)
		return ret, err
	default:
		ret := &UnsupportedTOCEntry{}
		err := d.DecodeElement(ret, &start)
		return ret, err
	}
}

type SimpleTOCEntry struct {
	BoldCode  string `xml:"bold,attr"`
	IdRef     string `xml:"idref,attr"`
	LevelCode string `xml:"level,attr"`
	Header    InlineMarkup
}

func (e *SimpleTOCEntry) TOCEntry() TOCEntry {
	return e
}

func (n *SimpleTOCEntry) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	err := decodeXMLAttrs(n, start)
	if err != nil {
		return err
	}
	return n.Header.UnmarshalXML(d, start)
}

type MultiColumnTOCEntry struct {
	SimpleTOCEntry
	Target     InlineMarkup `xml:"target"`
	PageNumber string       `xml:"page-num"`
}

func (e *MultiColumnTOCEntry) TOCEntry() TOCEntry {
	return e
}

type QuotedSimpleTOCEntry struct {
	StyleCode string `xml:"style,attr"`
	Entry     *SimpleTOCEntry
}

func (e *QuotedSimpleTOCEntry) TOCEntry() TOCEntry {
	return e
}

type QuotedMultiColumnTOCEntry struct {
	StyleCode string `xml:"style,attr"`
	Entry     *MultiColumnTOCEntry
}

func (e *QuotedMultiColumnTOCEntry) TOCEntry() TOCEntry {
	return e
}

type UnsupportedTOCEntry struct {
	Name    xml.Name
	Attrs   map[xml.Name]string
	Content []byte
}

func (e *UnsupportedTOCEntry) TOCEntry() TOCEntry {
	return e
}
