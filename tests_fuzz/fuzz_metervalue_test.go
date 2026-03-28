//go:build fuzz

package testsfuzz

import (
	"testing"

	st "github.com/aasanchez/ocpp16types"
)

func FuzzMeterValue(f *testing.F) {
	f.Add("2024-01-15T10:30:00Z", "42.5")

	f.Fuzz(func(t *testing.T, timestamp string, sampledValueStr string) {
		if len(timestamp) > maxFuzzStringLen || len(sampledValueStr) > maxFuzzStringLen {
			t.Skip("input too large")
		}

		input := st.MeterValueInput{
			Timestamp: timestamp,
			SampledValue: []st.SampledValueInput{
				{
					Value: sampledValueStr,
				},
			},
		}

		mv, err := st.NewMeterValue(input)
		if err != nil {
			if !isExpectedError(err) {
				t.Fatalf("unexpected error: %v", err)
			}
			return
		}

		// If no error, verify Timestamp().String() is non-empty
		if len(mv.Timestamp().String()) == 0 {
			t.Fatal("Timestamp().String() returned empty value")
		}
	})
}
