//go:build fuzz

package testsfuzz

import (
	"testing"

	st "github.com/aasanchez/ocpp16types"
)

func FuzzChargingSchedulePeriod(f *testing.F) {
	f.Add(0, 32.0, 3)

	f.Fuzz(func(t *testing.T, startPeriod int, limit float64, numberPhases int) {
		// Create input with optional NumberPhases
		input := st.ChargingSchedulePeriodInput{
			StartPeriod: startPeriod,
			Limit:       limit,
		}

		// Only include NumberPhases if it's in valid range
		if numberPhases >= 1 && numberPhases <= 3 {
			input.NumberPhases = &numberPhases
		}

		period, err := st.NewChargingSchedulePeriod(input)
		if err != nil {
			if !isExpectedError(err) {
				t.Fatalf("unexpected error: %v", err)
			}
			return
		}

		// If no error, verify StartPeriod().Value() matches
		if int(period.StartPeriod().Value()) != startPeriod {
			t.Fatalf(
				"StartPeriod().Value() = %d, want %d",
				period.StartPeriod().Value(),
				startPeriod,
			)
		}
	})
}
