package healthkit

import (
	"encoding/xml"
	"fmt"
	"time"
)

// Correlation is a sample that groups multiple related samples into a single entry.
// https://developer.apple.com/documentation/healthkit/hkcorrelation
type Correlation struct {
	XMLName       xml.Name         `xml:"Correlation"`
	Type          string           `xml:"type,attr"`
	SourceName    string           `xml:"sourceName,attr"`
	SourceVersion string           `xml:"sourceVersion,attr"`
	Device        string           `xml:"device,attr"`
	CreationDate  string           `xml:"creationDate,attr"`
	StartDate     string           `xml:"startDate,attr"`
	EndDate       string           `xml:"endDate,attr"`
	MetadataEntry []*MetadataEntry `xml:"MetadataEntry"`
	Record        []*Record        `xml:"Record"`
}

// String returns the object's string representation useful for logging and debugging.
func (c *Correlation) String() string {
	return fmt.Sprintf("%+v", *c)
}

func (c *Correlation) CreationDateTime() time.Time {
	t, _ := time.Parse(timeFormat, c.CreationDate)
	return t
}

func (c *Correlation) StartDateTime() time.Time {
	t, _ := time.Parse(timeFormat, c.StartDate)
	return t
}

func (c *Correlation) EndDateTime() time.Time {
	t, _ := time.Parse(timeFormat, c.EndDate)
	return t
}

var _ Data = (*Correlation)(nil)
