//go:build fuzz

package testsfuzz

import (
	"testing"

	st "github.com/aasanchez/ocpp16types"
)

func FuzzSampledValue(f *testing.F) {
	f.Add("42.5")

	f.Fuzz(func(t *testing.T, value string) {
		if len(value) > maxFuzzStringLen {
			t.Skip("input too large")
		}

		input := st.SampledValueInput{
			Value: value,
		}

		sv, err := st.NewSampledValue(input)
		if err != nil {
			if !isExpectedError(err) {
				t.Fatalf("unexpected error: %v", err)
			}
			return
		}

		// If no error, verify Value.Value() == input
		if sv.Value.Value() != value {
			t.Fatalf("Value.Value() = %q, want %q", sv.Value.Value(), value)
		}
	})
}
