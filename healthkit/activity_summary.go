package healthkit

import (
	"encoding/xml"
	"fmt"
)

type ActivitySummary struct {
	XMLName                xml.Name `xml:"ActivitySummary"`
	DateComponents         string   `xml:"dateComponents,attr"`
	ActiveEnergyBurned     string   `xml:"activeEnergyBurned,attr"`
	ActiveEnergyBurnedGoal string   `xml:"activeEnergyBurnedGoal,attr"`
	ActiveEnergyBurnedUnit string   `xml:"activeEnergyBurnedUnit,attr"`
	AppleMoveMinutes       string   `xml:"appleMoveMinutes,attr"`
	AppleMoveMinutesGoal   string   `xml:"appleMoveMinutesGoal,attr"`
	AppleExerciseTime      string   `xml:"appleExerciseTime,attr"`
	AppleExerciseTimeGoal  string   `xml:"appleExerciseTimeGoal,attr"`
	AppleStandHours        string   `xml:"appleStandHours,attr"`
	AppleStandHoursGoal    string   `xml:"appleStandHoursGoal,attr"`
}

func (as *ActivitySummary) String() string {
	return fmt.Sprint(*as)
}

var _ Data = (*ActivitySummary)(nil)
