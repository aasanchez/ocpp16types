package ocpp16types

// Measurand represents measurable quantities as defined in OCPP 1.6.
type Measurand string

// Alias for shorter constant declarations.
type m = Measurand

const (
	// MeasurandCurrentExport measures exported current.
	MeasurandCurrentExport m = "Current.Export"
	// MeasurandCurrentImport measures imported current.
	MeasurandCurrentImport m = "Current.Import"
	// MeasurandCurrentOffered measures offered current.
	MeasurandCurrentOffered m = "Current.Offered"
	// MeasurandEnergyActiveExportRegister measures active export
	// energy register.
	MeasurandEnergyActiveExportRegister m = "Energy.Active.Export.Register"
	// MeasurandEnergyActiveImportRegister measures active import
	// energy register.
	MeasurandEnergyActiveImportRegister m = "Energy.Active.Import.Register"
	// MeasurandEnergyReactiveExportRegister measures reactive export
	// energy register.
	MeasurandEnergyReactiveExportRegister m = "Energy.Reactive.Export.Register"
	// MeasurandEnergyReactiveImportRegister measures reactive import
	// energy register.
	MeasurandEnergyReactiveImportRegister m = "Energy.Reactive.Import.Register"
	// MeasurandEnergyActiveExportInterval measures active export
	// energy interval.
	MeasurandEnergyActiveExportInterval m = "Energy.Active.Export.Interval"
	// MeasurandEnergyActiveImportInterval measures active import
	// energy interval.
	MeasurandEnergyActiveImportInterval m = "Energy.Active.Import.Interval"
	// MeasurandEnergyReactiveExportInterval measures reactive export
	// energy interval.
	MeasurandEnergyReactiveExportInterval m = "Energy.Reactive.Export.Interval"
	// MeasurandEnergyReactiveImportInterval measures reactive import
	// energy interval.
	MeasurandEnergyReactiveImportInterval m = "Energy.Reactive.Import.Interval"
	// MeasurandFrequency measures frequency.
	MeasurandFrequency m = "Frequency"
	// MeasurandPowerActiveExport measures active exported power.
	MeasurandPowerActiveExport m = "Power.Active.Export"
	// MeasurandPowerActiveImport measures active imported power.
	MeasurandPowerActiveImport m = "Power.Active.Import"
	// MeasurandPowerFactor measures power factor.
	MeasurandPowerFactor m = "Power.Factor"
	// MeasurandPowerOffered measures offered power.
	MeasurandPowerOffered m = "Power.Offered"
	// MeasurandPowerReactiveExport measures reactive exported power.
	MeasurandPowerReactiveExport m = "Power.Reactive.Export"
	// MeasurandPowerReactiveImport measures reactive imported power.
	MeasurandPowerReactiveImport m = "Power.Reactive.Import"
	// MeasurandRPM measures RPM.
	MeasurandRPM m = "RPM"
	// MeasurandSoC measures state of charge.
	MeasurandSoC m = "SoC"
	// MeasurandTemperature measures temperature.
	MeasurandTemperature m = "Temperature"
	// MeasurandVoltage measures voltage.
	MeasurandVoltage m = "Voltage"
)

// IsValid checks if the Measurand value is valid per OCPP 1.6.
func (t Measurand) IsValid() bool {
	switch t {
	case MeasurandCurrentExport, MeasurandCurrentImport,
		MeasurandCurrentOffered,
		MeasurandEnergyActiveExportRegister,
		MeasurandEnergyActiveImportRegister,
		MeasurandEnergyReactiveExportRegister,
		MeasurandEnergyReactiveImportRegister,
		MeasurandEnergyActiveExportInterval,
		MeasurandEnergyActiveImportInterval,
		MeasurandEnergyReactiveExportInterval,
		MeasurandEnergyReactiveImportInterval,
		MeasurandFrequency, MeasurandPowerActiveExport,
		MeasurandPowerActiveImport, MeasurandPowerFactor,
		MeasurandPowerOffered, MeasurandPowerReactiveExport,
		MeasurandPowerReactiveImport, MeasurandRPM,
		MeasurandSoC, MeasurandTemperature, MeasurandVoltage:
		return true
	default:
		return false
	}
}

// String returns the string representation of Measurand.
func (t Measurand) String() string {
	return string(t)
}
