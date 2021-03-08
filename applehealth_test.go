package applehealth

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/AlekSi/applehealth/healthkit"
)

func TestUnmarshaler(t *testing.T) {
	for _, ext := range []string{"xml", "zip"} {
		t.Run(ext, func(t *testing.T) {
			t.Parallel()

			f := filepath.Join("testdata", "testdata."+ext)
			if _, err := os.Stat(f); ext == "zip" && os.IsNotExist(err) {
				t.Skipf("Generate ZIP archive by running `make %s`.", f)
			}

			u, err := NewUnmarshaler(f)
			require.NoError(t, err)
			defer func() {
				assert.NoError(t, u.Close())
			}()
			u.DisallowUnhandledElements = true

			expectedData := []healthkit.Data{
				&healthkit.Record{
					XMLName:       xml.Name{Local: "Record"},
					Type:          "HKQuantityTypeIdentifierHeight",
					Unit:          "cm",
					Value:         "168",
					SourceName:    "My Processing Tool",
					SourceVersion: "80",
					CreationDate:  "2017-10-29 21:29:16 +0300",
					StartDate:     "2015-12-21 13:46:53 +0300",
					EndDate:       "2015-12-21 13:46:53 +0300",
					MetadataEntry: []*healthkit.MetadataEntry{{
						XMLName: xml.Name{Local: "MetadataEntry"},
						Key:     "Original Creation Date",
						Value:   "2015-12-21 13:46:53 +0300",
					}, {
						XMLName: xml.Name{Local: "MetadataEntry"},
						Key:     "Original Source",
						Value:   "MyWatch",
					}, {
						XMLName: xml.Name{Local: "MetadataEntry"},
						Key:     "Source File Export Date",
						Value:   "2017-10-29 21:15:43 +0300",
					}},
				},
				&healthkit.Record{
					XMLName:       xml.Name{Local: "Record"},
					Type:          "HKQuantityTypeIdentifierBodyMass",
					Unit:          "kg",
					Value:         "65",
					SourceName:    "Здоровье",
					SourceVersion: "10.3.3",
					CreationDate:  "2017-07-25 08:18:55 +0300",
					StartDate:     "2017-07-25 08:18:00 +0300",
					EndDate:       "2017-07-25 08:18:00 +0300",
					MetadataEntry: []*healthkit.MetadataEntry{{
						XMLName: xml.Name{Local: "MetadataEntry"},
						Key:     "HKWasUserEntered",
						Value:   "1",
					}},
				},
			}

			for i, expected := range expectedData {
				t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
					actual, err := u.Next()
					assert.NoError(t, err)
					assert.Equal(t, expected, actual)
				})
			}

			actual, err := u.Next()
			assert.Equal(t, io.EOF, err)
			assert.Nil(t, actual)

			expectedMeta := &healthkit.Meta{
				Locale: "ru_RU",
				ExportDate: healthkit.ExportDate{
					XMLName: xml.Name{Local: "ExportDate"},
					Value:   "2019-12-26 08:20:53 +0300",
				},
				Me: healthkit.Me{
					XMLName:             xml.Name{Local: "Me"},
					DateOfBirth:         "1987-10-02",
					BiologicalSex:       healthkit.BiologicalSexFemale,
					BloodType:           healthkit.BloodTypeABPositive,
					FitzpatrickSkinType: "HKFitzpatrickSkinTypeVI",
				},
			}
			assert.Equal(t, expectedMeta, u.Meta())
			assert.Equal(t, time.Date(1987, 10, 2, 0, 0, 0, 0, time.UTC), expectedMeta.Me.DateOfBirthTime())
		})
	}
}

func Example() {
	file := filepath.Join("testdata", "testdata.zip")
	u, err := NewUnmarshaler(file)
	if err != nil {
		log.Fatal(err)
	}
	defer u.Close()

	// call Next() once to be able to read metadata
	var data healthkit.Data
	if data, err = u.Next(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Meta: %s\n", u.Meta())

	for {
		fmt.Printf("Got %T:\n", data)
		switch data := data.(type) {
		case *healthkit.Record:
			fmt.Printf("\t%s\n", data.Type)
		}

		if data, err = u.Next(); err != nil {
			break
		}
	}
	if err != io.EOF {
		log.Fatal(err)
	}

	// Output:
	// Meta: {Locale:ru_RU ExportDate:{XMLName:{Space: Local:ExportDate} Value:2019-12-26 08:20:53 +0300} Me:{XMLName:{Space: Local:Me} DateOfBirth:1987-10-02 BiologicalSex:HKBiologicalSexFemale BloodType:HKBloodTypeABPositive FitzpatrickSkinType:HKFitzpatrickSkinTypeVI}}
	// Got *healthkit.Record:
	// 	HKQuantityTypeIdentifierHeight
	// Got *healthkit.Record:
	// 	HKQuantityTypeIdentifierBodyMass
}

func BenchmarkUnmarshaler(b *testing.B) {
	xml, err := filepath.Glob(filepath.FromSlash("benchdata/*.xml"))
	if err != nil {
		b.Fatal(err)
	}
	zip, err := filepath.Glob(filepath.FromSlash("benchdata/*.zip"))
	if err != nil {
		b.Fatal(err)
	}
	matches := make([]string, 0, len(xml)+len(zip))
	matches = append(matches, xml...)
	matches = append(matches, zip...)

	for _, file := range matches {
		file := file
		b.Run(file, func(b *testing.B) {
			if _, err := os.Stat(file); os.IsNotExist(err) {
				b.Skipf("%s is absent.", file)
			}

			b.ReportAllocs()

			u, err := NewUnmarshaler(file)
			require.NoError(b, err)
			defer u.Close()
			u.DisallowUnhandledElements = true

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, err = u.Next()
				if err != nil {
					if err == io.EOF {
						b.Skipf("%s is too small for N = %d.", file, b.N)
					}
					b.Fatal(err)
				}
			}
			b.StopTimer()

			offset := u.d.InputOffset()
			b.SetBytes(offset / int64(b.N))
			b.ReportMetric(float64(offset)/1024/1024, "MB")
		})
	}
}
