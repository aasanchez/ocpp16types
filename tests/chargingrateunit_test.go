//nolint:dupl // enum test pattern
package ocpp16types_test

import (
	"testing"

	st "github.com/aasanchez/ocpp16types"
)

const (
	unitWattsStr         = "W"
	unitAmperesStr       = "A"
	rateUnitMethodString = "ChargingRateUnit.String()"
)

func TestChargingRateUnit_IsValid_Watts(t *testing.T) {
	t.Parallel()

	if !st.ChargingRateUnitWatts.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "ChargingRateUnitWatts")
	}
}

func TestChargingRateUnit_IsValid_Amperes(t *testing.T) {
	t.Parallel()

	if !st.ChargingRateUnitAmperes.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "ChargingRateUnitAmperes")
	}
}

func TestChargingRateUnit_IsValid_Empty(t *testing.T) {
	t.Parallel()

	unit := st.ChargingRateUnit("")
	if unit.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "ChargingRateUnit(\"\")")
	}
}

func TestChargingRateUnit_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	unit := st.ChargingRateUnit("Unknown")
	if unit.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "ChargingRateUnit(\"Unknown\")")
	}
}

func TestChargingRateUnit_IsValid_Lowercase(t *testing.T) {
	t.Parallel()

	unit := st.ChargingRateUnit("w")
	if unit.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "ChargingRateUnit(\"w\")")
	}
}

func TestChargingRateUnit_String_Watts(t *testing.T) {
	t.Parallel()

	got := st.ChargingRateUnitWatts.String()
	if got != unitWattsStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			rateUnitMethodString,
			got,
			unitWattsStr,
		)
	}
}

func TestChargingRateUnit_String_Amperes(t *testing.T) {
	t.Parallel()

	got := st.ChargingRateUnitAmperes.String()
	if got != unitAmperesStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			rateUnitMethodString,
			got,
			unitAmperesStr,
		)
	}
}
