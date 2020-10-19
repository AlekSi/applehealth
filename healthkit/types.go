package healthkit

import (
	"encoding/xml"
	"fmt"
	"time"
)

const timeFormat = "2006-01-02 15:04:05 Z0700"

type Data interface {
	fmt.Stringer
	sealed()
}

func (as *ActivitySummary) sealed() {}
func (c *Correlation) sealed()      {}
func (r *Record) sealed()           {}
func (w *Workout) sealed()          {}

// for https://github.com/BurntSushi/go-sumtype
//go-sumtype:decl Data

type Meta struct {
	Locale     string
	ExportDate ExportDate
	Me         Me
}

type ExportDate struct {
	XMLName xml.Name `xml:"ExportDate"`
	Value   string   `xml:"value,attr"`
}

type Me struct {
	XMLName             xml.Name      `xml:"Me"`
	DateOfBirth         string        `xml:"HKCharacteristicTypeIdentifierDateOfBirth,attr"`
	BiologicalSex       BiologicalSex `xml:"HKCharacteristicTypeIdentifierBiologicalSex,attr"`
	BloodType           BloodType     `xml:"HKCharacteristicTypeIdentifierBloodType,attr"`
	FitzpatrickSkinType string        `xml:"HKCharacteristicTypeIdentifierFitzpatrickSkinType,attr"`
}

func (me *Me) DateOfBirthTime() time.Time {
	d, _ := time.Parse("2006-01-02", me.DateOfBirth)
	return d
}

type MetadataEntry struct {
	XMLName xml.Name `xml:"MetadataEntry"`
	Key     string   `xml:"key,attr"`
	Value   string   `xml:"value,attr"`
}

// String returns the object's string representation useful for logging and debugging.
func (me *MetadataEntry) String() string {
	return fmt.Sprintf("%+v", *me)
}
