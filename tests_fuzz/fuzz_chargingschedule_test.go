//go:build fuzz

package testsfuzz

import (
	"testing"

	st "github.com/aasanchez/ocpp16types"
)

func FuzzChargingSchedule(f *testing.F) {
	f.Add("W", 0, 32.0)

	f.Fuzz(func(t *testing.T, rateUnit string, startPeriod int, limit float64) {
		if len(rateUnit) > maxFuzzStringLen {
			t.Skip("input too large")
		}

		input := st.ChargingScheduleInput{
			ChargingRateUnit: rateUnit,
			ChargingSchedulePeriod: []st.ChargingSchedulePeriodInput{
				{
					StartPeriod: startPeriod,
					Limit:       limit,
				},
			},
		}

		cs, err := st.NewChargingSchedule(input)
		if err != nil {
			if !isExpectedError(err) {
				t.Fatalf("unexpected error: %v", err)
			}
			return
		}

		// If no error, verify ChargingRateUnit().IsValid()
		if !cs.ChargingRateUnit().IsValid() {
			t.Fatalf("ChargingRateUnit().IsValid() = false, want true for unit %q", rateUnit)
		}
	})
}
