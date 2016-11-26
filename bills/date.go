package bills

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"time"
)

type Date struct {
	Year  int
	Month time.Month
	Day   int
}

func (d *Date) UnmarshalXMLAttr(attr xml.Attr) error {
	value := attr.Value
	if len(value) != 8 {
		return fmt.Errorf("invalid %s value %q", attr.Name, value)
	}

	year, err := strconv.Atoi(value[0:4])
	if err != nil {
		return fmt.Errorf("invalid year in %s", attr.Name)
	}

	day, err := strconv.Atoi(value[6:8])
	if err != nil {
		return fmt.Errorf("invalid day in %s", attr.Name)
	}

	monthNum, err := strconv.Atoi(value[4:6])
	if err != nil {
		return fmt.Errorf("invalid month in %s", attr.Name)
	}

	d.Year = year
	d.Month = time.Month(monthNum)
	d.Day = day

	return nil
}
