package healthkit

import (
	"encoding/xml"
	"fmt"
	"strings"
	"time"
)

type Record struct {
	XMLName xml.Name `xml:"Record"`
	Raw     string   `xml:",innerxml"`

	Type                             string                              `xml:"type,attr"`
	Unit                             string                              `xml:"unit,attr"`
	Value                            string                              `xml:"value,attr"`
	SourceName                       string                              `xml:"sourceName,attr"`
	SourceVersion                    string                              `xml:"sourceVersion,attr"`
	Device                           string                              `xml:"device,attr"`
	CreationDate                     string                              `xml:"creationDate,attr"`
	StartDate                        string                              `xml:"startDate,attr"`
	EndDate                          string                              `xml:"endDate,attr"`
	MetadataEntry                    []*MetadataEntry                    `xml:"MetadataEntry"`
	HeartRateVariabilityMetadataList []*HeartRateVariabilityMetadataList `xml:"HeartRateVariabilityMetadataList"`
}

type HeartRateVariabilityMetadataList struct {
	XMLName                     xml.Name                       `xml:"HeartRateVariabilityMetadataList"`
	InstantaneousBeatsPerMinute []*InstantaneousBeatsPerMinute `xml:"InstantaneousBeatsPerMinute"`
}

type InstantaneousBeatsPerMinute struct {
	XMLName xml.Name `xml:"InstantaneousBeatsPerMinute"`
	Bpm     int      `xml:"bpm,attr"`
	Time    string   `xml:"time,attr"`
}

// String returns the object's string representation useful for logging and debugging.
func (r *Record) String() string {
	rc := *r
	rc.Raw = ""
	return fmt.Sprintf("%+v", rc)
}

func (r *Record) DeviceMap() map[string]string {
	res := make(map[string]string, 5)
	s := strings.TrimPrefix(strings.TrimSuffix(r.Device, ">"), "<")
	for _, part := range strings.Split(s, ", ") {
		pair := strings.Split(part, ":")
		if len(pair) != 2 {
			continue
		}
		if len(pair[0]) == 0 || pair[0][0] == '<' {
			continue
		}
		res[pair[0]] = pair[1]
	}
	return res
}

func (r *Record) CreationDateTime() time.Time {
	t, _ := time.Parse(timeFormat, r.CreationDate)
	return t
}

func (r *Record) StartDateTime() time.Time {
	t, _ := time.Parse(timeFormat, r.StartDate)
	return t
}

func (r *Record) EndDateTime() time.Time {
	t, _ := time.Parse(timeFormat, r.EndDate)
	return t
}

var _ Data = (*Record)(nil)
