//nolint:dupl // enum test pattern
package ocpp16types_test

import (
	"testing"

	st "github.com/aasanchez/ocpp16types"
)

const (
	locationBodyStr      = "Body"
	locationCableStr     = "Cable"
	locationEVStr        = "EV"
	locationInletStr     = "Inlet"
	locationOutletStr    = "Outlet"
	locationMethodString = "Location.String()"
)

func TestLocation_IsValid_Body(t *testing.T) {
	t.Parallel()

	if !st.LocationBody.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "LocationBody")
	}
}

func TestLocation_IsValid_Cable(t *testing.T) {
	t.Parallel()

	if !st.LocationCable.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "LocationCable")
	}
}

func TestLocation_IsValid_EV(t *testing.T) {
	t.Parallel()

	if !st.LocationEV.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "LocationEV")
	}
}

func TestLocation_IsValid_Inlet(t *testing.T) {
	t.Parallel()

	if !st.LocationInlet.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "LocationInlet")
	}
}

func TestLocation_IsValid_Outlet(t *testing.T) {
	t.Parallel()

	if !st.LocationOutlet.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "LocationOutlet")
	}
}

func TestLocation_IsValid_Empty(t *testing.T) {
	t.Parallel()

	location := st.Location("")
	if location.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "Location(\"\")")
	}
}

func TestLocation_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	location := st.Location("Unknown")
	if location.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "Location(\"Unknown\")")
	}
}

func TestLocation_IsValid_Lowercase(t *testing.T) {
	t.Parallel()

	location := st.Location("body")
	if location.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "Location(\"body\")")
	}
}

func TestLocation_String_Body(t *testing.T) {
	t.Parallel()

	got := st.LocationBody.String()
	if got != locationBodyStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			locationMethodString,
			got,
			locationBodyStr,
		)
	}
}

func TestLocation_String_Cable(t *testing.T) {
	t.Parallel()

	got := st.LocationCable.String()
	if got != locationCableStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			locationMethodString,
			got,
			locationCableStr,
		)
	}
}

func TestLocation_String_EV(t *testing.T) {
	t.Parallel()

	got := st.LocationEV.String()
	if got != locationEVStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			locationMethodString,
			got,
			locationEVStr,
		)
	}
}

func TestLocation_String_Inlet(t *testing.T) {
	t.Parallel()

	got := st.LocationInlet.String()
	if got != locationInletStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			locationMethodString,
			got,
			locationInletStr,
		)
	}
}

func TestLocation_String_Outlet(t *testing.T) {
	t.Parallel()

	got := st.LocationOutlet.String()
	if got != locationOutletStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			locationMethodString,
			got,
			locationOutletStr,
		)
	}
}
