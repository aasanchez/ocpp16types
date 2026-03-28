//nolint:dupl // enum test pattern
package ocpp16types_test

import (
	"testing"

	st "github.com/aasanchez/ocpp16types"
)

const (
	valueFormatRawStr        = "Raw"
	valueFormatSignedDataStr = "SignedData"
	valueFormatMethodString  = "ValueFormat.String()"
)

func TestValueFormat_IsValid_Raw(t *testing.T) {
	t.Parallel()

	if !st.ValueFormatRaw.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "ValueFormatRaw")
	}
}

func TestValueFormat_IsValid_SignedData(t *testing.T) {
	t.Parallel()

	if !st.ValueFormatSignedData.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "ValueFormatSignedData")
	}
}

func TestValueFormat_IsValid_Empty(t *testing.T) {
	t.Parallel()

	format := st.ValueFormat("")
	if format.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "ValueFormat(\"\")")
	}
}

func TestValueFormat_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	format := st.ValueFormat("Unknown")
	if format.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "ValueFormat(\"Unknown\")")
	}
}

func TestValueFormat_IsValid_Lowercase(t *testing.T) {
	t.Parallel()

	format := st.ValueFormat("raw")
	if format.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "ValueFormat(\"raw\")")
	}
}

func TestValueFormat_String_Raw(t *testing.T) {
	t.Parallel()

	got := st.ValueFormatRaw.String()
	if got != valueFormatRawStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			valueFormatMethodString,
			got,
			valueFormatRawStr,
		)
	}
}

func TestValueFormat_String_SignedData(t *testing.T) {
	t.Parallel()

	got := st.ValueFormatSignedData.String()
	if got != valueFormatSignedDataStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			valueFormatMethodString,
			got,
			valueFormatSignedDataStr,
		)
	}
}
