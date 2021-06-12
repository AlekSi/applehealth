package healthkit

import (
	"encoding/xml"
	"fmt"
	"time"
)

type Workout struct {
	XMLName xml.Name `xml:"Workout"`
	Raw     string   `xml:",innerxml"`

	WorkoutActivityType   string           `xml:"workoutActivityType,attr"`
	Duration              string           `xml:"duration,attr"`
	DurationUnit          string           `xml:"durationUnit,attr"`
	TotalDistance         string           `xml:"totalDistance,attr"`
	TotalDistanceUnit     string           `xml:"totalDistanceUnit,attr"`
	TotalEnergyBurned     string           `xml:"totalEnergyBurned,attr"`
	TotalEnergyBurnedUnit string           `xml:"totalEnergyBurnedUnit,attr"`
	SourceName            string           `xml:"sourceName,attr"`
	SourceVersion         string           `xml:"sourceVersion,attr"`
	CreationDate          string           `xml:"creationDate,attr"`
	StartDate             string           `xml:"startDate,attr"`
	EndDate               string           `xml:"endDate,attr"`
	MetadataEntry         []*MetadataEntry `xml:"MetadataEntry"`
}

// String returns the object's string representation useful for logging and debugging.
func (w *Workout) String() string {
	wc := *w
	wc.Raw = ""
	return fmt.Sprintf("%+v", wc)
}

func (w *Workout) CreationDateTime() time.Time {
	t, _ := time.Parse(timeFormat, w.CreationDate)
	return t
}

func (w *Workout) StartDateTime() time.Time {
	t, _ := time.Parse(timeFormat, w.StartDate)
	return t
}

func (w *Workout) EndDateTime() time.Time {
	t, _ := time.Parse(timeFormat, w.EndDate)
	return t
}

var _ Data = (*Workout)(nil)
