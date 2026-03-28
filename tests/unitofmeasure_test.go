package ocpp16types_test

import (
	"testing"

	st "github.com/aasanchez/ocpp16types"
)

const (
	unitOfMeasureWhStr      = "Wh"
	unitOfMeasureKWhStr     = "kWh"
	unitOfMeasureWStr       = "W"
	unitOfMeasureKWStr      = "kW"
	unitOfMeasureVAStr      = "VA"
	unitOfMeasureKVAStr     = "kVA"
	unitOfMeasureVARStr     = "var"
	unitOfMeasureKVARStr    = "kvar"
	unitOfMeasureAStr       = "A"
	unitOfMeasureVStr       = "V"
	unitOfMeasureKStr       = "K"
	unitOfMeasureCelsiusStr = "Celsius"
	unitOfMeasureHzStr      = "Hz"
	unitMethodString        = "UnitOfMeasure.String()"
)

func TestUnitOfMeasure_IsValid_Wh(t *testing.T) {
	t.Parallel()

	if !st.UnitOfMeasureWh.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "UnitOfMeasureWh")
	}
}

func TestUnitOfMeasure_IsValid_KWh(t *testing.T) {
	t.Parallel()

	if !st.UnitOfMeasureKWh.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "UnitOfMeasureKWh")
	}
}

func TestUnitOfMeasure_IsValid_W(t *testing.T) {
	t.Parallel()

	if !st.UnitOfMeasureW.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "UnitOfMeasureW")
	}
}

func TestUnitOfMeasure_IsValid_KW(t *testing.T) {
	t.Parallel()

	if !st.UnitOfMeasureKW.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "UnitOfMeasureKW")
	}
}

func TestUnitOfMeasure_IsValid_VA(t *testing.T) {
	t.Parallel()

	if !st.UnitOfMeasureVA.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "UnitOfMeasureVA")
	}
}

func TestUnitOfMeasure_IsValid_KVA(t *testing.T) {
	t.Parallel()

	if !st.UnitOfMeasureKVA.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "UnitOfMeasureKVA")
	}
}

func TestUnitOfMeasure_IsValid_VAR(t *testing.T) {
	t.Parallel()

	if !st.UnitOfMeasureVAR.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "UnitOfMeasureVAR")
	}
}

func TestUnitOfMeasure_IsValid_KVAR(t *testing.T) {
	t.Parallel()

	if !st.UnitOfMeasureKVAR.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "UnitOfMeasureKVAR")
	}
}

func TestUnitOfMeasure_IsValid_A(t *testing.T) {
	t.Parallel()

	if !st.UnitOfMeasureA.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "UnitOfMeasureA")
	}
}

func TestUnitOfMeasure_IsValid_V(t *testing.T) {
	t.Parallel()

	if !st.UnitOfMeasureV.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "UnitOfMeasureV")
	}
}

func TestUnitOfMeasure_IsValid_K(t *testing.T) {
	t.Parallel()

	if !st.UnitOfMeasureK.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "UnitOfMeasureK")
	}
}

func TestUnitOfMeasure_IsValid_Celsius(t *testing.T) {
	t.Parallel()

	if !st.UnitOfMeasureCelsius.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "UnitOfMeasureCelsius")
	}
}

func TestUnitOfMeasure_IsValid_Hz(t *testing.T) {
	t.Parallel()

	if !st.UnitOfMeasureHz.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "UnitOfMeasureHz")
	}
}

func TestUnitOfMeasure_IsValid_Empty(t *testing.T) {
	t.Parallel()

	measure := st.UnitOfMeasure("")
	if measure.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "UnitOfMeasure(\"\")")
	}
}

func TestUnitOfMeasure_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	measure := st.UnitOfMeasure("Unknown")
	if measure.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "UnitOfMeasure(\"Unknown\")")
	}
}

func TestUnitOfMeasure_IsValid_Lowercase(t *testing.T) {
	t.Parallel()

	measure := st.UnitOfMeasure("wh")
	if measure.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "UnitOfMeasure(\"wh\")")
	}
}

func TestUnitOfMeasure_String_Wh(t *testing.T) {
	t.Parallel()

	got := st.UnitOfMeasureWh.String()
	if got != unitOfMeasureWhStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			unitMethodString,
			got,
			unitOfMeasureWhStr,
		)
	}
}

func TestUnitOfMeasure_String_KWh(t *testing.T) {
	t.Parallel()

	got := st.UnitOfMeasureKWh.String()
	if got != unitOfMeasureKWhStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			unitMethodString,
			got,
			unitOfMeasureKWhStr,
		)
	}
}

func TestUnitOfMeasure_String_W(t *testing.T) {
	t.Parallel()

	got := st.UnitOfMeasureW.String()
	if got != unitOfMeasureWStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			unitMethodString,
			got,
			unitOfMeasureWStr,
		)
	}
}

func TestUnitOfMeasure_String_KW(t *testing.T) {
	t.Parallel()

	got := st.UnitOfMeasureKW.String()
	if got != unitOfMeasureKWStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			unitMethodString,
			got,
			unitOfMeasureKWStr,
		)
	}
}

func TestUnitOfMeasure_String_VA(t *testing.T) {
	t.Parallel()

	got := st.UnitOfMeasureVA.String()
	if got != unitOfMeasureVAStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			unitMethodString,
			got,
			unitOfMeasureVAStr,
		)
	}
}

func TestUnitOfMeasure_String_KVA(t *testing.T) {
	t.Parallel()

	got := st.UnitOfMeasureKVA.String()
	if got != unitOfMeasureKVAStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			unitMethodString,
			got,
			unitOfMeasureKVAStr,
		)
	}
}

func TestUnitOfMeasure_String_VAR(t *testing.T) {
	t.Parallel()

	got := st.UnitOfMeasureVAR.String()
	if got != unitOfMeasureVARStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			unitMethodString,
			got,
			unitOfMeasureVARStr,
		)
	}
}

func TestUnitOfMeasure_String_KVAR(t *testing.T) {
	t.Parallel()

	got := st.UnitOfMeasureKVAR.String()
	if got != unitOfMeasureKVARStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			unitMethodString,
			got,
			unitOfMeasureKVARStr,
		)
	}
}

func TestUnitOfMeasure_String_A(t *testing.T) {
	t.Parallel()

	got := st.UnitOfMeasureA.String()
	if got != unitOfMeasureAStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			unitMethodString,
			got,
			unitOfMeasureAStr,
		)
	}
}

func TestUnitOfMeasure_String_V(t *testing.T) {
	t.Parallel()

	got := st.UnitOfMeasureV.String()
	if got != unitOfMeasureVStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			unitMethodString,
			got,
			unitOfMeasureVStr,
		)
	}
}

func TestUnitOfMeasure_String_K(t *testing.T) {
	t.Parallel()

	got := st.UnitOfMeasureK.String()
	if got != unitOfMeasureKStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			unitMethodString,
			got,
			unitOfMeasureKStr,
		)
	}
}

func TestUnitOfMeasure_String_Celsius(t *testing.T) {
	t.Parallel()

	got := st.UnitOfMeasureCelsius.String()
	if got != unitOfMeasureCelsiusStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			unitMethodString,
			got,
			unitOfMeasureCelsiusStr,
		)
	}
}

func TestUnitOfMeasure_String_Hz(t *testing.T) {
	t.Parallel()

	got := st.UnitOfMeasureHz.String()
	if got != unitOfMeasureHzStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			unitMethodString,
			got,
			unitOfMeasureHzStr,
		)
	}
}
