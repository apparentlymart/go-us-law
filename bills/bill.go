package bills

import (
	"encoding/xml"
	"io"
)

type Bill struct {
	Form *Form `xml:"form"`
	Body *Body `xml:"legis-body"`
}

func ParseBill(r io.Reader) (*Bill, error) {
	decoder := xml.NewDecoder(r)
	var bill Bill
	err := decoder.Decode(&bill)
	return &bill, err
}

func ParseBillBuffer(buf []byte) (*Bill, error) {
	var bill Bill
	err := xml.Unmarshal(buf, &bill)
	return &bill, err
}
