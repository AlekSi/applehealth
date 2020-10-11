package applehealth

import (
	"encoding/xml"
	"fmt"
	"sync/atomic"

	"github.com/AlekSi/applehealth/healthkit"
)

// Parser is responsible for parsing Apple Health XML stream.
//
// Most users should use Unmarshaler instead.
type Parser struct {
	DisallowUnhandledElements bool // if true, Next() will return UnhandledElementError for unhandled data

	d      *xml.Decoder
	offset int64
	meta   *healthkit.Meta
}

// UnhandledElementError is returned by the Parser from the Next() method
// if DisallowUnhandledElements is true and unhandled data is encountered.
type UnhandledElementError struct {
	Name string
}

func (u *UnhandledElementError) Error() string {
	return fmt.Sprintf("unhandled element %s", u.Name)
}

// NewParser creates a new parser with given XML decoder.
//
// Most users should use NewUnmarshaler instead.
func NewParser(d *xml.Decoder) *Parser {
	return &Parser{
		d:    d,
		meta: new(healthkit.Meta),
	}
}

// Meta returns parsed metadata after the first call to Next().
func (p *Parser) Meta() *healthkit.Meta {
	return p.meta
}

// InputOffset returns XML stream input offset. Unlike calling xml.Decoder.InputOffset directly,
// that method can be called concurrently with Next().
func (p *Parser) InputOffset() int64 {
	return atomic.LoadInt64(&p.offset)
}

func (p *Parser) updateOffset() {
	atomic.StoreInt64(&p.offset, p.d.InputOffset())
}

var newData = map[string]func() healthkit.Data{
	"Record":          func() healthkit.Data { return new(healthkit.Record) },
	"Correlation":     func() healthkit.Data { return new(healthkit.Correlation) },
	"Workout":         func() healthkit.Data { return new(healthkit.Workout) },
	"ActivitySummary": func() healthkit.Data { return new(healthkit.ActivitySummary) },
}

// Next returns the next health data object, or error
// (io.EOF, *UnhandledElementError, or XML parsing error).
func (p *Parser) Next() (healthkit.Data, error) {
	defer p.updateOffset()

	for {
		t, err := p.d.Token()
		if err != nil {
			return nil, err
		}

		se, ok := t.(xml.StartElement)
		if !ok {
			continue
		}

		name := se.Name.Local
		switch name {
		case "HealthData":
			// not using DecodeElement to avoid reading the whole huge element
			for _, attr := range t.Attr {
				expected := xml.Name{Local: "locale"}
				if attr.Name == expected {
					p.meta.Locale = attr.Value
				}
			}

		case "ExportDate":
			var ed healthkit.ExportDate
			if err = p.d.DecodeElement(&ed, &t); err != nil {
				return nil, err
			}
			p.meta.ExportDate = ed

		case "Me":
			var m healthkit.Me
			if err = p.d.DecodeElement(&m, &t); err != nil {
				return nil, err
			}
			p.meta.Me = m

		default:
			if ndf := newData[name]; ndf != nil {
				d := ndf()
				if err = p.d.DecodeElement(d, &t); err != nil {
					return nil, err
				}
				return d, nil
			}

			if p.DisallowUnhandledElements {
				return nil, &UnhandledElementError{Name: name}
			}
		}
	}
}
