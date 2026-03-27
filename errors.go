package ocpp16types

import "errors"

// ErrorFieldFormat is the format string for field-level errors.
const ErrorFieldFormat = "%s: %w"

// ErrorMismatch is the format for value comparison errors.
const ErrorMismatch = "Expected %q, got %q"

// ErrorWantNil is for asserting nil errors in tests.
const ErrorWantNil = "error = nil, want error for %s"

// ErrEmptyValue indicates a required value was empty.
var ErrEmptyValue = errors.New("value cannot be empty")

// ErrInvalidValue indicates a value failed validation.
var ErrInvalidValue = errors.New("invalid value")
