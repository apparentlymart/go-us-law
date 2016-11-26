package bills

import (
	"encoding/xml"
	"reflect"
	"strconv"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestUnmarshalInlineMarkup(t *testing.T) {
	tests := []struct {
		Input    string
		Expected InlineMarkup
	}{
		{
			"<t></t>",
			InlineMarkup{},
		},
		{
			"<t> </t>",
			InlineMarkup{
				Text(" "),
			},
		},
		{
			"<t>hello</t>",
			InlineMarkup{
				Text("hello"),
			},
		},
		{
			`<t><not-a-real-element foo="baz">...</not-a-real-element></t>`,
			InlineMarkup{
				&UnsupportedInlineElement{
					Name: xml.Name{"", "not-a-real-element"},
					Attrs: map[xml.Name]string{
						xml.Name{"", "foo"}: "baz",
					},
					InlineMarkup: InlineMarkup{
						Text("..."),
					},
				},
			},
		},
		{
			"<t>hello <italic>world</italic></t>",
			InlineMarkup{
				Text("hello "),
				&Italic{
					InlineMarkup{
						Text("world"),
					},
				},
			},
		},
		{
			`<t><sponsor name-id="S000033">Bernie Sanders</sponsor></t>`,
			InlineMarkup{
				&SponsorName{
					InlineMarkup: InlineMarkup{
						Text("Bernie Sanders"),
					},
					NameId: "S000033",
				},
			},
		},
		{
			`<t><cosponsor name-id="S000033">Bernie Sanders</cosponsor></t>`,
			InlineMarkup{
				&CosponsorName{
					InlineMarkup: InlineMarkup{
						Text("Bernie Sanders"),
					},
					NameId: "S000033",
				},
			},
		},
		{
			`<t><committee-name committee-id="HWM00">Committee on Ways and Means</committee-name></t>`,
			InlineMarkup{
				&CommitteeName{
					InlineMarkup: InlineMarkup{
						Text("Committee on Ways and Means"),
					},
					CommitteeId: "HWM00",
				},
			},
		},
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var got InlineMarkup
			err := xml.Unmarshal([]byte(test.Input), &got)
			if err != nil {
				t.Fatalf("error: %s", err)
			}

			if !reflect.DeepEqual(got, test.Expected) {
				t.Errorf(
					"incorrect result\ngot:  %s\nwant: %s",
					spew.Sdump(got),
					spew.Sdump(test.Expected),
				)
			}
		})
	}
}
