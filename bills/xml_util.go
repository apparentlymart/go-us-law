package bills

import (
	"encoding/xml"
	"fmt"
	"reflect"
	"strings"
)

// decodeXMLAttrs is a bit like xml.Unmarshal, but it considers only
// struct members that are tagged as XML attributes and can thus do
// its work with only an xml.StartElement, and no need for reading
// any other tokens from the token stream.
//
// This is useful when the attributes need to be decoded separately
// from an element's contents, such as for the structs that embed
// TextWrapper but add their own attributes that need to be decoded
// before delegating to TextWrapper's decoding function.
//
// Only non-namespaced attributes are supported.
func decodeXMLAttrs(target interface{}, start xml.StartElement) error {
	val := reflect.ValueOf(target)

	if val.Kind() == reflect.Interface {
		if val.IsNil() {
			return fmt.Errorf("can't apply decodeXMLAttrs to nil interface")
		}

		e := val.Elem()
		if e.Kind() == reflect.Ptr && !e.IsNil() {
			val = e
		}
	}

	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			val.Set(reflect.New(val.Type().Elem()))
		}
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return fmt.Errorf("decodeXMLAttrs requires struct target")
	}

	attrs := map[string]string{}
	for _, attr := range start.Attr {
		if attr.Name.Space != "" {
			continue
		}
		attrs[attr.Name.Local] = attr.Value
	}

	typ := val.Type()
	n := typ.NumField()
	for i := 0; i < n; i++ {
		f := typ.Field(i)
		tag := f.Tag.Get("xml")
		if f.PkgPath != "" && !f.Anonymous || tag == "-" {
			continue // Private field
		}

		if i := strings.Index(tag, " "); i >= 0 {
			// namespaced attributes not supported
			continue
		}

		tokens := strings.Split(tag, ",")
		if len(tokens) == 1 {
			// If there's only one token then it can't be an attribute
			continue
		}

		isAttr := false
		for _, flag := range tokens[1:] {
			if flag == "attr" {
				isAttr = true
				break
			}
		}
		if !isAttr {
			continue
		}

		attrName := tokens[0]
		if attrName == "" {
			continue
		}

		val.Field(i).Set(reflect.ValueOf(attrs[attrName]))
	}

	return nil
}
