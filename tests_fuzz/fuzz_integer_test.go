//go:build fuzz

package testsfuzz

import (
	"testing"

	st "github.com/aasanchez/ocpp16types"
)

func FuzzInteger(f *testing.F) {
	f.Add(100)

	f.Fuzz(func(t *testing.T, i int) {
		integer, err := st.NewInteger(i)
		if err != nil {
			if !isExpectedError(err) {
				t.Fatalf("unexpected error: %v", err)
			}
			return
		}

		// If no error, verify Value() matches input (cast to uint16)
		if int(integer.Value()) != i {
			t.Fatalf("Value() = %d, want %d", integer.Value(), uint16(i))
		}
	})
}
