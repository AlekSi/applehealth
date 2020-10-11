// Package healthkit contains Apple's HealthKit types and constans.
package healthkit // import "github.com/AlekSi/applehealth/healthkit"

type BiologicalSex string

const (
	BiologicalSexFemale = BiologicalSex("HKBiologicalSexFemale")
	BiologicalSexMale   = BiologicalSex("HKBiologicalSexMale")
	BiologicalSexOther  = BiologicalSex("HKBiologicalSexOther")
)

type BloodType string

const (
	BloodTypeAPositive  = BloodType("HKBloodTypeAPositive")
	BloodTypeANegative  = BloodType("HKBloodTypeANegative")
	BloodTypeBPositive  = BloodType("HKBloodTypeBPositive")
	BloodTypeBNegative  = BloodType("HKBloodTypeBNegative")
	BloodTypeABPositive = BloodType("HKBloodTypeABPositive")
	BloodTypeABNegative = BloodType("HKBloodTypeABNegative")
	BloodTypeOPositive  = BloodType("HKBloodTypeOPositive")
	BloodTypeONegative  = BloodType("HKBloodTypeONegative")
)

type CategoryValueSleepAnalysis string

const (
	CategoryValueSleepAnalysisInBed  = CategoryValueSleepAnalysis("HKCategoryValueSleepAnalysisInBed")
	CategoryValueSleepAnalysisAsleep = CategoryValueSleepAnalysis("HKCategoryValueSleepAnalysisAsleep")
	CategoryValueSleepAnalysisAwake  = CategoryValueSleepAnalysis("HKCategoryValueSleepAnalysisAwake")
)
