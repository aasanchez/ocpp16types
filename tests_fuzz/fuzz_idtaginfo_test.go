//go:build fuzz

package testsfuzz

import (
	"testing"

	st "github.com/aasanchez/ocpp16types"
)

func FuzzIdTagInfo(f *testing.F) {
	f.Add("Accepted")

	f.Fuzz(func(t *testing.T, statusStr string) {
		if len(statusStr) > maxFuzzStringLen {
			t.Skip("input too large")
		}

		// Cast to AuthorizationStatus
		status := st.AuthorizationStatus(statusStr)

		info, err := st.NewIdTagInfo(status)
		if err != nil {
			if !isExpectedError(err) {
				t.Fatalf("unexpected error: %v", err)
			}
			return
		}

		// If no error, verify Status().IsValid()
		if !info.Status().IsValid() {
			t.Fatalf("Status().IsValid() = false, want true for status %q", status)
		}
	})
}
