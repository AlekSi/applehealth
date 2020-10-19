package healthkit

import (
	"encoding/xml"
	"fmt"
)

// ActivitySummary contains the move, exercise, and stand data for a given day.
// https://developer.apple.com/documentation/healthkit/hkactivitysummary
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

// String returns the object's string representation useful for logging and debugging.
func (as *ActivitySummary) String() string {
	return fmt.Sprintf("%+v", *as)
}

var _ Data = (*ActivitySummary)(nil)
