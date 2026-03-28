//go:build fuzz

package testsfuzz

import (
	"errors"
	"testing"

	st "github.com/aasanchez/ocpp16types"
)

const maxFuzzStringLen = 10000

func FuzzCiString20(f *testing.F) {
	f.Add("validstring")

	f.Fuzz(func(t *testing.T, s string) {
		if len(s) > maxFuzzStringLen {
			t.Skip("input too large")
		}

		cs, err := st.NewCiString20Type(s)
		if err != nil {
			// Verify error is one of the known sentinels
			if !isExpectedError(err) {
				t.Fatalf("unexpected error: %v", err)
			}
			return
		}

		// If no error, verify basic invariants
		if cs.Value() != s {
			t.Fatalf("Value() = %q, want %q", cs.Value(), s)
		}
		if cs.String() != s {
			t.Fatalf("String() = %q, want %q", cs.String(), s)
		}
	})
}

func FuzzCiString25(f *testing.F) {
	f.Add("validstring")

	f.Fuzz(func(t *testing.T, s string) {
		if len(s) > maxFuzzStringLen {
			t.Skip("input too large")
		}

		cs, err := st.NewCiString25Type(s)
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

func FuzzCiString50(f *testing.F) {
	f.Add("validstring")

	f.Fuzz(func(t *testing.T, s string) {
		if len(s) > maxFuzzStringLen {
			t.Skip("input too large")
		}

		cs, err := st.NewCiString50Type(s)
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

func isExpectedError(err error) bool {
	return errors.Is(err, st.ErrEmptyValue) ||
		errors.Is(err, st.ErrInvalidValue)
}
