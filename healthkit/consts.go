// Package healthkit contains Apple's HealthKit types and constans.
package healthkit // import "github.com/AlekSi/applehealth/healthkit"

// BiologicalSex indicates the user’s sex.
// https://developer.apple.com/documentation/healthkit/hkbiologicalsex
type BiologicalSex string

// BiologicalSex constants.
const (
	BiologicalSexFemale = BiologicalSex("HKBiologicalSexFemale")
	BiologicalSexMale   = BiologicalSex("HKBiologicalSexMale")
	BiologicalSexOther  = BiologicalSex("HKBiologicalSexOther")
)

// BloodType indicates the user’s blood type.
// https://developer.apple.com/documentation/healthkit/hkbloodtype
type BloodType string

// BloodType constants.
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

// CategoryValueSleepAnalysis represents the result of a sleep analysis.
// https://developer.apple.com/documentation/healthkit/hkcategoryvaluesleepanalysis
type CategoryValueSleepAnalysis string

// CategoryValueSleepAnalysis constants.
const (
	CategoryValueSleepAnalysisInBed  = CategoryValueSleepAnalysis("HKCategoryValueSleepAnalysisInBed")
	CategoryValueSleepAnalysisAsleep = CategoryValueSleepAnalysis("HKCategoryValueSleepAnalysisAsleep")
	CategoryValueSleepAnalysisAwake  = CategoryValueSleepAnalysis("HKCategoryValueSleepAnalysisAwake")
)

// MetadataKey is used to add metadata to objects stored in HealthKit.
// https://developer.apple.com/documentation/healthkit/samples/metadata_keys
type MetadataKey string

// MetadataKey constants.
const (
	MetadataKeyExternalUUID   = MetadataKey("HKExternalUUID")
	MetadataKeyTimeZone       = MetadataKey("HKTimeZone")
	MetadataKeyWasUserEntered = MetadataKey("HKWasUserEntered")
)
