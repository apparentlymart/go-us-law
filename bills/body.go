package bills

import (
	"encoding/xml"
)

type Body struct {
	StyleCode string `xml:"style,attr"`
	StructuralMarkup
}

func (b *Body) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	err := decodeXMLAttrs(b, start)
	if err != nil {
		return err
	}
	return b.StructuralMarkup.UnmarshalXML(d, start)
}
