package ocpp16types

// Phase represents the phase of a measurement as defined in OCPP 1.6.
type Phase string

// Alias for shorter constant declarations.
type ph = Phase

const (
	// PhaseSinglePhase indicates single-phase measurement.
	PhaseSinglePhase ph = "Single"
	// PhaseL1 indicates phase L1.
	PhaseL1 ph = "L1"
	// PhaseL2 indicates phase L2.
	PhaseL2 ph = "L2"
	// PhaseL3 indicates phase L3.
	PhaseL3 ph = "L3"
	// PhaseN indicates neutral phase.
	PhaseN ph = "N"
	// PhaseL1L2 indicates phases L1 and L2.
	PhaseL1L2 ph = "L1-L2"
	// PhaseL2L3 indicates phases L2 and L3.
	PhaseL2L3 ph = "L2-L3"
	// PhaseL3L1 indicates phases L3 and L1.
	PhaseL3L1 ph = "L3-L1"
)

// IsValid checks if the Phase value is valid per OCPP 1.6.
func (t Phase) IsValid() bool {
	switch t {
	case PhaseSinglePhase, PhaseL1, PhaseL2, PhaseL3, PhaseN,
		PhaseL1L2, PhaseL2L3, PhaseL3L1:
		return true
	default:
		return false
	}
}

// String returns the string representation of Phase.
func (t Phase) String() string {
	return string(t)
}
