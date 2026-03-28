package ocpp16types_test

import (
	"testing"

	st "github.com/aasanchez/ocpp16types"
)

const (
	measurandCurrentExportStr                = "Current.Export"
	measurandCurrentImportStr                = "Current.Import"
	measurandCurrentOfferedStr               = "Current.Offered"
	measurandEnergyActiveExportRegisterStr   = "Energy.Active.Export.Register"
	measurandEnergyActiveImportRegisterStr   = "Energy.Active.Import.Register"
	measurandEnergyReactiveExportRegisterStr = "Energy.Reactive.Export.Register"
	measurandEnergyReactiveImportRegisterStr = "Energy.Reactive.Import.Register"
	measurandEnergyActiveExportIntervalStr   = "Energy.Active.Export.Interval"
	measurandEnergyActiveImportIntervalStr   = "Energy.Active.Import.Interval"
	measurandEnergyReactiveExportIntervalStr = "Energy.Reactive.Export.Interval"
	measurandEnergyReactiveImportIntervalStr = "Energy.Reactive.Import.Interval"
	measurandFrequencyStr                    = "Frequency"
	measurandPowerActiveExportStr            = "Power.Active.Export"
	measurandPowerActiveImportStr            = "Power.Active.Import"
	measurandPowerFactorStr                  = "Power.Factor"
	measurandPowerOfferedStr                 = "Power.Offered"
	measurandPowerReactiveExportStr          = "Power.Reactive.Export"
	measurandPowerReactiveImportStr          = "Power.Reactive.Import"
	measurandRPMStr                          = "RPM"
	measurandSoCStr                          = "SoC"
	measurandTemperatureStr                  = "Temperature"
	measurandVoltageStr                      = "Voltage"
	measurandMethodString                    = "Measurand.String()"
)

func TestMeasurand_IsValid_CurrentExport(t *testing.T) {
	t.Parallel()

	if !st.MeasurandCurrentExport.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "MeasurandCurrentExport")
	}
}

func TestMeasurand_IsValid_CurrentImport(t *testing.T) {
	t.Parallel()

	if !st.MeasurandCurrentImport.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "MeasurandCurrentImport")
	}
}

func TestMeasurand_IsValid_CurrentOffered(t *testing.T) {
	t.Parallel()

	if !st.MeasurandCurrentOffered.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "MeasurandCurrentOffered")
	}
}

func TestMeasurand_IsValid_EnergyActiveExportRegister(t *testing.T) {
	t.Parallel()

	if !st.MeasurandEnergyActiveExportRegister.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "MeasurandEnergyActiveExportRegister")
	}
}

func TestMeasurand_IsValid_EnergyActiveImportRegister(t *testing.T) {
	t.Parallel()

	if !st.MeasurandEnergyActiveImportRegister.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "MeasurandEnergyActiveImportRegister")
	}
}

func TestMeasurand_IsValid_EnergyReactiveExportRegister(t *testing.T) {
	t.Parallel()

	if !st.MeasurandEnergyReactiveExportRegister.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "MeasurandEnergyReactiveExportRegister")
	}
}

func TestMeasurand_IsValid_EnergyReactiveImportRegister(t *testing.T) {
	t.Parallel()

	if !st.MeasurandEnergyReactiveImportRegister.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "MeasurandEnergyReactiveImportRegister")
	}
}

func TestMeasurand_IsValid_EnergyActiveExportInterval(t *testing.T) {
	t.Parallel()

	if !st.MeasurandEnergyActiveExportInterval.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "MeasurandEnergyActiveExportInterval")
	}
}

func TestMeasurand_IsValid_EnergyActiveImportInterval(t *testing.T) {
	t.Parallel()

	if !st.MeasurandEnergyActiveImportInterval.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "MeasurandEnergyActiveImportInterval")
	}
}

func TestMeasurand_IsValid_EnergyReactiveExportInterval(t *testing.T) {
	t.Parallel()

	if !st.MeasurandEnergyReactiveExportInterval.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "MeasurandEnergyReactiveExportInterval")
	}
}

func TestMeasurand_IsValid_EnergyReactiveImportInterval(t *testing.T) {
	t.Parallel()

	if !st.MeasurandEnergyReactiveImportInterval.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "MeasurandEnergyReactiveImportInterval")
	}
}

func TestMeasurand_IsValid_Frequency(t *testing.T) {
	t.Parallel()

	if !st.MeasurandFrequency.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "MeasurandFrequency")
	}
}

func TestMeasurand_IsValid_PowerActiveExport(t *testing.T) {
	t.Parallel()

	if !st.MeasurandPowerActiveExport.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "MeasurandPowerActiveExport")
	}
}

func TestMeasurand_IsValid_PowerActiveImport(t *testing.T) {
	t.Parallel()

	if !st.MeasurandPowerActiveImport.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "MeasurandPowerActiveImport")
	}
}

func TestMeasurand_IsValid_PowerFactor(t *testing.T) {
	t.Parallel()

	if !st.MeasurandPowerFactor.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "MeasurandPowerFactor")
	}
}

func TestMeasurand_IsValid_PowerOffered(t *testing.T) {
	t.Parallel()

	if !st.MeasurandPowerOffered.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "MeasurandPowerOffered")
	}
}

func TestMeasurand_IsValid_PowerReactiveExport(t *testing.T) {
	t.Parallel()

	if !st.MeasurandPowerReactiveExport.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "MeasurandPowerReactiveExport")
	}
}

func TestMeasurand_IsValid_PowerReactiveImport(t *testing.T) {
	t.Parallel()

	if !st.MeasurandPowerReactiveImport.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "MeasurandPowerReactiveImport")
	}
}

func TestMeasurand_IsValid_RPM(t *testing.T) {
	t.Parallel()

	if !st.MeasurandRPM.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "MeasurandRPM")
	}
}

func TestMeasurand_IsValid_SoC(t *testing.T) {
	t.Parallel()

	if !st.MeasurandSoC.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "MeasurandSoC")
	}
}

func TestMeasurand_IsValid_Temperature(t *testing.T) {
	t.Parallel()

	if !st.MeasurandTemperature.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "MeasurandTemperature")
	}
}

func TestMeasurand_IsValid_Voltage(t *testing.T) {
	t.Parallel()

	if !st.MeasurandVoltage.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "MeasurandVoltage")
	}
}

func TestMeasurand_IsValid_Empty(t *testing.T) {
	t.Parallel()

	measurand := st.Measurand("")
	if measurand.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "Measurand(\"\")")
	}
}

func TestMeasurand_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	measurand := st.Measurand("Unknown")
	if measurand.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "Measurand(\"Unknown\")")
	}
}

func TestMeasurand_IsValid_Lowercase(t *testing.T) {
	t.Parallel()

	measurand := st.Measurand("current.export")
	if measurand.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "Measurand(\"current.export\")")
	}
}

func TestMeasurand_String_CurrentExport(t *testing.T) {
	t.Parallel()

	got := st.MeasurandCurrentExport.String()
	if got != measurandCurrentExportStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			measurandMethodString,
			got,
			measurandCurrentExportStr,
		)
	}
}

func TestMeasurand_String_CurrentImport(t *testing.T) {
	t.Parallel()

	got := st.MeasurandCurrentImport.String()
	if got != measurandCurrentImportStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			measurandMethodString,
			got,
			measurandCurrentImportStr,
		)
	}
}

func TestMeasurand_String_CurrentOffered(t *testing.T) {
	t.Parallel()

	got := st.MeasurandCurrentOffered.String()
	if got != measurandCurrentOfferedStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			measurandMethodString,
			got,
			measurandCurrentOfferedStr,
		)
	}
}

func TestMeasurand_String_EnergyActiveExportRegister(t *testing.T) {
	t.Parallel()

	got := st.MeasurandEnergyActiveExportRegister.String()
	if got != measurandEnergyActiveExportRegisterStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			measurandMethodString,
			got,
			measurandEnergyActiveExportRegisterStr,
		)
	}
}

func TestMeasurand_String_EnergyActiveImportRegister(t *testing.T) {
	t.Parallel()

	got := st.MeasurandEnergyActiveImportRegister.String()
	if got != measurandEnergyActiveImportRegisterStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			measurandMethodString,
			got,
			measurandEnergyActiveImportRegisterStr,
		)
	}
}

func TestMeasurand_String_EnergyReactiveExportRegister(t *testing.T) {
	t.Parallel()

	got := st.MeasurandEnergyReactiveExportRegister.String()
	if got != measurandEnergyReactiveExportRegisterStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			measurandMethodString,
			got,
			measurandEnergyReactiveExportRegisterStr,
		)
	}
}

func TestMeasurand_String_EnergyReactiveImportRegister(t *testing.T) {
	t.Parallel()

	got := st.MeasurandEnergyReactiveImportRegister.String()
	if got != measurandEnergyReactiveImportRegisterStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			measurandMethodString,
			got,
			measurandEnergyReactiveImportRegisterStr,
		)
	}
}

func TestMeasurand_String_EnergyActiveExportInterval(t *testing.T) {
	t.Parallel()

	got := st.MeasurandEnergyActiveExportInterval.String()
	if got != measurandEnergyActiveExportIntervalStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			measurandMethodString,
			got,
			measurandEnergyActiveExportIntervalStr,
		)
	}
}

func TestMeasurand_String_EnergyActiveImportInterval(t *testing.T) {
	t.Parallel()

	got := st.MeasurandEnergyActiveImportInterval.String()
	if got != measurandEnergyActiveImportIntervalStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			measurandMethodString,
			got,
			measurandEnergyActiveImportIntervalStr,
		)
	}
}

func TestMeasurand_String_EnergyReactiveExportInterval(t *testing.T) {
	t.Parallel()

	got := st.MeasurandEnergyReactiveExportInterval.String()
	if got != measurandEnergyReactiveExportIntervalStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			measurandMethodString,
			got,
			measurandEnergyReactiveExportIntervalStr,
		)
	}
}

func TestMeasurand_String_EnergyReactiveImportInterval(t *testing.T) {
	t.Parallel()

	got := st.MeasurandEnergyReactiveImportInterval.String()
	if got != measurandEnergyReactiveImportIntervalStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			measurandMethodString,
			got,
			measurandEnergyReactiveImportIntervalStr,
		)
	}
}

func TestMeasurand_String_Frequency(t *testing.T) {
	t.Parallel()

	got := st.MeasurandFrequency.String()
	if got != measurandFrequencyStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			measurandMethodString,
			got,
			measurandFrequencyStr,
		)
	}
}

func TestMeasurand_String_PowerActiveExport(t *testing.T) {
	t.Parallel()

	got := st.MeasurandPowerActiveExport.String()
	if got != measurandPowerActiveExportStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			measurandMethodString,
			got,
			measurandPowerActiveExportStr,
		)
	}
}

func TestMeasurand_String_PowerActiveImport(t *testing.T) {
	t.Parallel()

	got := st.MeasurandPowerActiveImport.String()
	if got != measurandPowerActiveImportStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			measurandMethodString,
			got,
			measurandPowerActiveImportStr,
		)
	}
}

func TestMeasurand_String_PowerFactor(t *testing.T) {
	t.Parallel()

	got := st.MeasurandPowerFactor.String()
	if got != measurandPowerFactorStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			measurandMethodString,
			got,
			measurandPowerFactorStr,
		)
	}
}

func TestMeasurand_String_PowerOffered(t *testing.T) {
	t.Parallel()

	got := st.MeasurandPowerOffered.String()
	if got != measurandPowerOfferedStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			measurandMethodString,
			got,
			measurandPowerOfferedStr,
		)
	}
}

func TestMeasurand_String_PowerReactiveExport(t *testing.T) {
	t.Parallel()

	got := st.MeasurandPowerReactiveExport.String()
	if got != measurandPowerReactiveExportStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			measurandMethodString,
			got,
			measurandPowerReactiveExportStr,
		)
	}
}

func TestMeasurand_String_PowerReactiveImport(t *testing.T) {
	t.Parallel()

	got := st.MeasurandPowerReactiveImport.String()
	if got != measurandPowerReactiveImportStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			measurandMethodString,
			got,
			measurandPowerReactiveImportStr,
		)
	}
}

func TestMeasurand_String_RPM(t *testing.T) {
	t.Parallel()

	got := st.MeasurandRPM.String()
	if got != measurandRPMStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			measurandMethodString,
			got,
			measurandRPMStr,
		)
	}
}

func TestMeasurand_String_SoC(t *testing.T) {
	t.Parallel()

	got := st.MeasurandSoC.String()
	if got != measurandSoCStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			measurandMethodString,
			got,
			measurandSoCStr,
		)
	}
}

func TestMeasurand_String_Temperature(t *testing.T) {
	t.Parallel()

	got := st.MeasurandTemperature.String()
	if got != measurandTemperatureStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			measurandMethodString,
			got,
			measurandTemperatureStr,
		)
	}
}

func TestMeasurand_String_Voltage(t *testing.T) {
	t.Parallel()

	got := st.MeasurandVoltage.String()
	if got != measurandVoltageStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			measurandMethodString,
			got,
			measurandVoltageStr,
		)
	}
}
