package ocpp16types_test

import (
	"testing"

	st "github.com/aasanchez/ocpp16types"
)

const (
	statusAvailableStr     = "Available"
	statusPreparingStr     = "Preparing"
	statusChargingStr      = "Charging"
	statusSuspendedEVStr   = "SuspendedEV"
	statusSuspendedEVSEStr = "SuspendedEVSE"
	statusFinishingStr     = "Finishing"
	statusReservedStr      = "Reserved"
	statusUnavailableStr   = "Unavailable"
	statusFaultedStr       = "Faulted"
	cpStatusMethodString   = "ChargePointStatus.String()"
)

func TestChargePointStatus_IsValid_Available(t *testing.T) {
	t.Parallel()

	if !st.ChargePointStatusAvailable.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "ChargePointStatusAvailable")
	}
}

func TestChargePointStatus_IsValid_Preparing(t *testing.T) {
	t.Parallel()

	if !st.ChargePointStatusPreparing.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "ChargePointStatusPreparing")
	}
}

func TestChargePointStatus_IsValid_Charging(t *testing.T) {
	t.Parallel()

	if !st.ChargePointStatusCharging.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "ChargePointStatusCharging")
	}
}

func TestChargePointStatus_IsValid_SuspendedEV(t *testing.T) {
	t.Parallel()

	if !st.ChargePointStatusSuspendedEV.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "ChargePointStatusSuspendedEV")
	}
}

func TestChargePointStatus_IsValid_SuspendedEVSE(t *testing.T) {
	t.Parallel()

	if !st.ChargePointStatusSuspendedEVSE.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "ChargePointStatusSuspendedEVSE")
	}
}

func TestChargePointStatus_IsValid_Finishing(t *testing.T) {
	t.Parallel()

	if !st.ChargePointStatusFinishing.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "ChargePointStatusFinishing")
	}
}

func TestChargePointStatus_IsValid_Reserved(t *testing.T) {
	t.Parallel()

	if !st.ChargePointStatusReserved.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "ChargePointStatusReserved")
	}
}

func TestChargePointStatus_IsValid_Unavailable(t *testing.T) {
	t.Parallel()

	if !st.ChargePointStatusUnavailable.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "ChargePointStatusUnavailable")
	}
}

func TestChargePointStatus_IsValid_Faulted(t *testing.T) {
	t.Parallel()

	if !st.ChargePointStatusFaulted.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "ChargePointStatusFaulted")
	}
}

func TestChargePointStatus_IsValid_Empty(t *testing.T) {
	t.Parallel()

	status := st.ChargePointStatus("")
	if status.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "ChargePointStatus(\"\")")
	}
}

func TestChargePointStatus_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	status := st.ChargePointStatus("Unknown")
	if status.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "ChargePointStatus(\"Unknown\")")
	}
}

func TestChargePointStatus_IsValid_Lowercase(t *testing.T) {
	t.Parallel()

	status := st.ChargePointStatus("available")
	if status.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "ChargePointStatus(\"available\")")
	}
}

func TestChargePointStatus_String_Available(t *testing.T) {
	t.Parallel()

	got := st.ChargePointStatusAvailable.String()
	if got != statusAvailableStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			cpStatusMethodString,
			got,
			statusAvailableStr,
		)
	}
}

func TestChargePointStatus_String_Preparing(t *testing.T) {
	t.Parallel()

	got := st.ChargePointStatusPreparing.String()
	if got != statusPreparingStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			cpStatusMethodString,
			got,
			statusPreparingStr,
		)
	}
}

func TestChargePointStatus_String_Charging(t *testing.T) {
	t.Parallel()

	got := st.ChargePointStatusCharging.String()
	if got != statusChargingStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			cpStatusMethodString,
			got,
			statusChargingStr,
		)
	}
}

func TestChargePointStatus_String_SuspendedEV(t *testing.T) {
	t.Parallel()

	got := st.ChargePointStatusSuspendedEV.String()
	if got != statusSuspendedEVStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			cpStatusMethodString,
			got,
			statusSuspendedEVStr,
		)
	}
}

func TestChargePointStatus_String_SuspendedEVSE(t *testing.T) {
	t.Parallel()

	got := st.ChargePointStatusSuspendedEVSE.String()
	if got != statusSuspendedEVSEStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			cpStatusMethodString,
			got,
			statusSuspendedEVSEStr,
		)
	}
}

func TestChargePointStatus_String_Finishing(t *testing.T) {
	t.Parallel()

	got := st.ChargePointStatusFinishing.String()
	if got != statusFinishingStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			cpStatusMethodString,
			got,
			statusFinishingStr,
		)
	}
}

func TestChargePointStatus_String_Reserved(t *testing.T) {
	t.Parallel()

	got := st.ChargePointStatusReserved.String()
	if got != statusReservedStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			cpStatusMethodString,
			got,
			statusReservedStr,
		)
	}
}

func TestChargePointStatus_String_Unavailable(t *testing.T) {
	t.Parallel()

	got := st.ChargePointStatusUnavailable.String()
	if got != statusUnavailableStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			cpStatusMethodString,
			got,
			statusUnavailableStr,
		)
	}
}

func TestChargePointStatus_String_Faulted(t *testing.T) {
	t.Parallel()

	got := st.ChargePointStatusFaulted.String()
	if got != statusFaultedStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			cpStatusMethodString,
			got,
			statusFaultedStr,
		)
	}
}
