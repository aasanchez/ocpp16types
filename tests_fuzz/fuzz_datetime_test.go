//go:build fuzz

package testsfuzz

import (
	"testing"

	st "github.com/aasanchez/ocpp16types"
)

func FuzzDateTime(f *testing.F) {
	f.Add("2024-01-15T10:30:00Z")

	f.Fuzz(func(t *testing.T, s string) {
		if len(s) > maxFuzzStringLen {
			t.Skip("input too large")
		}

		dt, err := st.NewDateTime(s)
		if err != nil {
			if !isExpectedError(err) {
				t.Fatalf("unexpected error: %v", err)
			}
			return
		}

		// If no error, verify basic invariants
		if len(dt.String()) == 0 {
			t.Fatal("String() returned empty value")
		}
	})
}
