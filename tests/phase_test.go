package ocpp16types_test

import (
	"testing"

	st "github.com/aasanchez/ocpp16types"
)

const (
	phaseL1Str        = "L1"
	phaseL2Str        = "L2"
	phaseL3Str        = "L3"
	phaseNStr         = "N"
	phaseL1L2Str      = "L1-L2"
	phaseL2L3Str      = "L2-L3"
	phaseL3L1Str      = "L3-L1"
	phaseMethodString = "Phase.String()"
)

func TestPhase_IsValid_L1(t *testing.T) {
	t.Parallel()

	if !st.PhaseL1.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "PhaseL1")
	}
}

func TestPhase_IsValid_L2(t *testing.T) {
	t.Parallel()

	if !st.PhaseL2.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "PhaseL2")
	}
}

func TestPhase_IsValid_L3(t *testing.T) {
	t.Parallel()

	if !st.PhaseL3.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "PhaseL3")
	}
}

func TestPhase_IsValid_N(t *testing.T) {
	t.Parallel()

	if !st.PhaseN.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "PhaseN")
	}
}

func TestPhase_IsValid_L1L2(t *testing.T) {
	t.Parallel()

	if !st.PhaseL1L2.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "PhaseL1L2")
	}
}

func TestPhase_IsValid_L2L3(t *testing.T) {
	t.Parallel()

	if !st.PhaseL2L3.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "PhaseL2L3")
	}
}

func TestPhase_IsValid_L3L1(t *testing.T) {
	t.Parallel()

	if !st.PhaseL3L1.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "PhaseL3L1")
	}
}

func TestPhase_IsValid_Empty(t *testing.T) {
	t.Parallel()

	phase := st.Phase("")
	if phase.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "Phase(\"\")")
	}
}

func TestPhase_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	phase := st.Phase("Unknown")
	if phase.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "Phase(\"Unknown\")")
	}
}

func TestPhase_IsValid_Lowercase(t *testing.T) {
	t.Parallel()

	phase := st.Phase("single")
	if phase.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "Phase(\"single\")")
	}
}

func TestPhase_String_L1(t *testing.T) {
	t.Parallel()

	got := st.PhaseL1.String()
	if got != phaseL1Str {
		t.Errorf(
			st.ErrorMethodMismatch,
			phaseMethodString,
			got,
			phaseL1Str,
		)
	}
}

func TestPhase_String_L2(t *testing.T) {
	t.Parallel()

	got := st.PhaseL2.String()
	if got != phaseL2Str {
		t.Errorf(
			st.ErrorMethodMismatch,
			phaseMethodString,
			got,
			phaseL2Str,
		)
	}
}

func TestPhase_String_L3(t *testing.T) {
	t.Parallel()

	got := st.PhaseL3.String()
	if got != phaseL3Str {
		t.Errorf(
			st.ErrorMethodMismatch,
			phaseMethodString,
			got,
			phaseL3Str,
		)
	}
}

func TestPhase_String_N(t *testing.T) {
	t.Parallel()

	got := st.PhaseN.String()
	if got != phaseNStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			phaseMethodString,
			got,
			phaseNStr,
		)
	}
}

func TestPhase_String_L1L2(t *testing.T) {
	t.Parallel()

	got := st.PhaseL1L2.String()
	if got != phaseL1L2Str {
		t.Errorf(
			st.ErrorMethodMismatch,
			phaseMethodString,
			got,
			phaseL1L2Str,
		)
	}
}

func TestPhase_String_L2L3(t *testing.T) {
	t.Parallel()

	got := st.PhaseL2L3.String()
	if got != phaseL2L3Str {
		t.Errorf(
			st.ErrorMethodMismatch,
			phaseMethodString,
			got,
			phaseL2L3Str,
		)
	}
}

func TestPhase_String_L3L1(t *testing.T) {
	t.Parallel()

	got := st.PhaseL3L1.String()
	if got != phaseL3L1Str {
		t.Errorf(
			st.ErrorMethodMismatch,
			phaseMethodString,
			got,
			phaseL3L1Str,
		)
	}
}
