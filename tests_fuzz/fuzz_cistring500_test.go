//go:build fuzz

package testsfuzz

import (
	"testing"

	st "github.com/aasanchez/ocpp16types"
)

func FuzzCiString500(f *testing.F) {
	f.Add("validstring")

	f.Fuzz(func(t *testing.T, s string) {
		if len(s) > maxFuzzStringLen {
			t.Skip("input too large")
		}

		cs, err := st.NewCiString500Type(s)
		if err != nil {
			if !isExpectedError(err) {
				t.Fatalf("unexpected error: %v", err)
			}
			return
		}

		if cs.Value() != s {
			t.Fatalf("Value() = %q, want %q", cs.Value(), s)
		}
		if cs.String() != s {
			t.Fatalf("String() = %q, want %q", cs.String(), s)
		}
	})
}
