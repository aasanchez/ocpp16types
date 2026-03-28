package ocpp16types_test

import (
	"testing"

	st "github.com/aasanchez/ocpp16types"
)

const (
	reasonDeAuthorizedStr   = "DeAuthorized"
	reasonEmergencyStopStr  = "EmergencyStop"
	reasonEVDisconnectedStr = "EVDisconnected"
	reasonHardResetStr      = "HardReset"
	reasonLocalStr          = "Local"
	reasonOtherStr          = "Other"
	reasonPowerLossStr      = "PowerLoss"
	reasonRebootStr         = "Reboot"
	reasonRemoteStr         = "Remote"
	reasonSoftResetStr      = "SoftReset"
	reasonUnlockCommandStr  = "UnlockCommand"
	stopReasonMethodString  = "StopReason.String()"
)

func TestStopReason_IsValid_DeAuthorized(t *testing.T) {
	t.Parallel()

	if !st.StopReasonDeAuthorized.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "StopReasonDeAuthorized")
	}
}

func TestStopReason_IsValid_EmergencyStop(t *testing.T) {
	t.Parallel()

	if !st.StopReasonEmergencyStop.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "StopReasonEmergencyStop")
	}
}

func TestStopReason_IsValid_EVDisconnected(t *testing.T) {
	t.Parallel()

	if !st.StopReasonEVDisconnected.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "StopReasonEVDisconnected")
	}
}

func TestStopReason_IsValid_HardReset(t *testing.T) {
	t.Parallel()

	if !st.StopReasonHardReset.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "StopReasonHardReset")
	}
}

func TestStopReason_IsValid_Local(t *testing.T) {
	t.Parallel()

	if !st.StopReasonLocal.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "StopReasonLocal")
	}
}

func TestStopReason_IsValid_Other(t *testing.T) {
	t.Parallel()

	if !st.StopReasonOther.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "StopReasonOther")
	}
}

func TestStopReason_IsValid_PowerLoss(t *testing.T) {
	t.Parallel()

	if !st.StopReasonPowerLoss.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "StopReasonPowerLoss")
	}
}

func TestStopReason_IsValid_Reboot(t *testing.T) {
	t.Parallel()

	if !st.StopReasonReboot.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "StopReasonReboot")
	}
}

func TestStopReason_IsValid_Remote(t *testing.T) {
	t.Parallel()

	if !st.StopReasonRemote.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "StopReasonRemote")
	}
}

func TestStopReason_IsValid_SoftReset(t *testing.T) {
	t.Parallel()

	if !st.StopReasonSoftReset.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "StopReasonSoftReset")
	}
}

func TestStopReason_IsValid_UnlockCommand(t *testing.T) {
	t.Parallel()

	if !st.StopReasonUnlockCommand.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "StopReasonUnlockCommand")
	}
}

func TestStopReason_IsValid_Empty(t *testing.T) {
	t.Parallel()

	reason := st.StopReason("")
	if reason.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "StopReason(\"\")")
	}
}

func TestStopReason_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	reason := st.StopReason("Unknown")
	if reason.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "StopReason(\"Unknown\")")
	}
}

func TestStopReason_IsValid_Lowercase(t *testing.T) {
	t.Parallel()

	reason := st.StopReason("deauthorized")
	if reason.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "StopReason(\"deauthorized\")")
	}
}

func TestStopReason_String_DeAuthorized(t *testing.T) {
	t.Parallel()

	got := st.StopReasonDeAuthorized.String()
	if got != reasonDeAuthorizedStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			stopReasonMethodString,
			got,
			reasonDeAuthorizedStr,
		)
	}
}

func TestStopReason_String_EmergencyStop(t *testing.T) {
	t.Parallel()

	got := st.StopReasonEmergencyStop.String()
	if got != reasonEmergencyStopStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			stopReasonMethodString,
			got,
			reasonEmergencyStopStr,
		)
	}
}

func TestStopReason_String_EVDisconnected(t *testing.T) {
	t.Parallel()

	got := st.StopReasonEVDisconnected.String()
	if got != reasonEVDisconnectedStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			stopReasonMethodString,
			got,
			reasonEVDisconnectedStr,
		)
	}
}

func TestStopReason_String_HardReset(t *testing.T) {
	t.Parallel()

	got := st.StopReasonHardReset.String()
	if got != reasonHardResetStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			stopReasonMethodString,
			got,
			reasonHardResetStr,
		)
	}
}

func TestStopReason_String_Local(t *testing.T) {
	t.Parallel()

	got := st.StopReasonLocal.String()
	if got != reasonLocalStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			stopReasonMethodString,
			got,
			reasonLocalStr,
		)
	}
}

func TestStopReason_String_Other(t *testing.T) {
	t.Parallel()

	got := st.StopReasonOther.String()
	if got != reasonOtherStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			stopReasonMethodString,
			got,
			reasonOtherStr,
		)
	}
}

func TestStopReason_String_PowerLoss(t *testing.T) {
	t.Parallel()

	got := st.StopReasonPowerLoss.String()
	if got != reasonPowerLossStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			stopReasonMethodString,
			got,
			reasonPowerLossStr,
		)
	}
}

func TestStopReason_String_Reboot(t *testing.T) {
	t.Parallel()

	got := st.StopReasonReboot.String()
	if got != reasonRebootStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			stopReasonMethodString,
			got,
			reasonRebootStr,
		)
	}
}

func TestStopReason_String_Remote(t *testing.T) {
	t.Parallel()

	got := st.StopReasonRemote.String()
	if got != reasonRemoteStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			stopReasonMethodString,
			got,
			reasonRemoteStr,
		)
	}
}

func TestStopReason_String_SoftReset(t *testing.T) {
	t.Parallel()

	got := st.StopReasonSoftReset.String()
	if got != reasonSoftResetStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			stopReasonMethodString,
			got,
			reasonSoftResetStr,
		)
	}
}

func TestStopReason_String_UnlockCommand(t *testing.T) {
	t.Parallel()

	got := st.StopReasonUnlockCommand.String()
	if got != reasonUnlockCommandStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			stopReasonMethodString,
			got,
			reasonUnlockCommandStr,
		)
	}
}
