//nolint:dupl // enum test pattern
package ocpp16types_test

import (
	"testing"

	st "github.com/aasanchez/ocpp16types"
)

const (
	purposeChargePointMaxProfileStr = "ChargePointMaxProfile"
	purposeTxDefaultProfileStr      = "TxDefaultProfile"
	purposeTxProfileStr             = "TxProfile"
	purposeTypeMethodString         = "ChargingProfilePurposeType.String()"
)

//nolint:revive // test function name
func TestChargingProfilePurposeType_IsValid_ChargePointMaxProfile(t *testing.T) {
	t.Parallel()

	if !st.ChargePointMaxProfile.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "ChargePointMaxProfile")
	}
}

func TestChargingProfilePurposeType_IsValid_TxDefaultProfile(t *testing.T) {
	t.Parallel()

	if !st.TxDefaultProfile.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "TxDefaultProfile")
	}
}

func TestChargingProfilePurposeType_IsValid_TxProfile(t *testing.T) {
	t.Parallel()

	if !st.TxProfile.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "TxProfile")
	}
}

func TestChargingProfilePurposeType_IsValid_Empty(t *testing.T) {
	t.Parallel()

	purpose := st.ChargingProfilePurposeType("")
	if purpose.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "ChargingProfilePurposeType(\"\")")
	}
}

func TestChargingProfilePurposeType_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	purpose := st.ChargingProfilePurposeType("Unknown")
	if purpose.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "ChargingProfilePurposeType(\"Unknown\")")
	}
}

func TestChargingProfilePurposeType_IsValid_Lowercase(t *testing.T) {
	t.Parallel()

	purpose := st.ChargingProfilePurposeType("chargepointa")
	if purpose.IsValid() {
		t.Errorf(
			st.ErrorIsValidTrue,
			"ChargingProfilePurposeType(\"chargepointa\")",
		)
	}
}

func TestChargingProfilePurposeType_String_ChargePointMaxProfile(t *testing.T) {
	t.Parallel()

	got := st.ChargePointMaxProfile.String()
	if got != purposeChargePointMaxProfileStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			purposeTypeMethodString,
			got,
			purposeChargePointMaxProfileStr,
		)
	}
}

func TestChargingProfilePurposeType_String_TxDefaultProfile(t *testing.T) {
	t.Parallel()

	got := st.TxDefaultProfile.String()
	if got != purposeTxDefaultProfileStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			purposeTypeMethodString,
			got,
			purposeTxDefaultProfileStr,
		)
	}
}

func TestChargingProfilePurposeType_String_TxProfile(t *testing.T) {
	t.Parallel()

	got := st.TxProfile.String()
	if got != purposeTxProfileStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			purposeTypeMethodString,
			got,
			purposeTxProfileStr,
		)
	}
}
