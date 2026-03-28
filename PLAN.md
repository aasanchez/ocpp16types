# ocpp16types — Implementation Plan

> Goal: Bring `ocpp16types` to full parity with the design, structure, testing
> pyramid, and tooling conventions established in `ocpp16messages/types/`.

---

## 1. errors.go — Complete the Error Constants

The current `errors.go` has 3 constants and 2 sentinel errors.
`ocpp16messages/types/errors.go` has 10 constants and 2 sentinel errors.

**Add these missing constants:**

| Constant               | Format                                          | Purpose                                      |
|------------------------|-------------------------------------------------|----------------------------------------------|
| `ErrorExpectedError`   | `"Expected Error: %v"`                          | Test: expected an error, format its value     |
| `ErrorMismatchValue`   | `"Expected %v, got %v"`                         | Test: `%v` variant of `ErrorMismatch`         |
| `ErrorUnexpectedError` | `"Unexpected Error: %v"`                        | Test: did not expect an error                 |
| `ErrorWantContains`    | `"error = %v, want error containing '%s'"`      | Test: assert error contains substring         |
| `ErrorWantNonNil`      | `"%s = nil, want non-nil"`                      | Test: expected non-nil value                  |
| `ErrorWrapping`        | `"%v, wants wrapping %v"`                       | Test: assert `errors.Is` wrapping             |
| `ErrorMethodMismatch`  | `"%s = %v, want %v"`                            | Test: method return value mismatch            |
| `ErrorIsValidTrue`     | `"%s.IsValid() = true, want false"`             | Test: enum should have been invalid           |
| `ErrorIsValidFalse`    | `"%s.IsValid() = false, want true"`             | Test: enum should have been valid             |

These constants are consumed by the test files (black-box package `ocpp16types_test`)
and establish a single source of truth for assertion formatting.

---

## 2. doc.go — Enrich Package Documentation

Replace the current minimal doc.go with a comprehensive version matching the
`ocpp16messages/types/doc.go` pattern:

- List the three core validated types (CiString, DateTime, Integer)
- List all enum types provided
- List composite types (IdTagInfo, IdToken, MeterValue, SampledValue, ChargingSchedule, ChargingSchedulePeriod)
- State the zero-dependency guarantee
- State the construction-time validation contract
- State thread-safety via immutable fields and value receivers

---

## 3. Test Directory Structure

Create the `tests/` subdirectory for black-box unit tests:

```
ocpp16types/
  tests/                            ← NEW directory
    cistring_test.go
    datetime_test.go
    integer_test.go
    errors_test.go
    authorizationstatus_test.go
    chargepointstatus_test.go
    chargepointerrorcode_test.go
    chargingprofilepurposetype_test.go
    chargingrateunit_test.go
    registrationstatus_test.go
    stopreason_test.go
    location_test.go
    measurand_test.go
    phase_test.go
    readingcontext_test.go
    unitofmeasure_test.go
    valueformat_test.go
    idtoken_test.go
    idtaginfo_test.go
    sampledvalue_test.go
    metervalue_test.go
    chargingschedule_test.go
    chargingscheduleperiod_test.go
    enum_assert_test.go             ← shared helper (enumValidator + assertEnumValid/assertEnumInvalid)
```

All test files use:

- **Package**: `ocpp16types_test` (black-box testing)
- **Import alias**: `st "github.com/aasanchez/ocpp16types"` (matches `ocpp16messages` convention)
- **`t.Parallel()`** on every test function
- **Shared error format constants** from `errors.go` (e.g., `st.ErrorIsValidFalse`)
- **No external test dependencies** — stdlib only (`testing`, `errors`, `strings`, `fmt`)

### 3.1 Enum Test Pattern (per enum type)

For every enum (AuthorizationStatus, ChargePointStatus, ChargePointErrorCode,
ChargingProfilePurposeType, ChargingRateUnit, RegistrationStatus, StopReason,
Location, Measurand, Phase, ReadingContext, UnitOfMeasure, ValueFormat):

1. One `TestXxx_IsValid_<Value>` per valid constant → asserts `IsValid() == true`
2. One `TestXxx_IsValid_Empty` → asserts `IsValid() == false` for `""`
3. One `TestXxx_IsValid_Unknown` → asserts `IsValid() == false` for `"Unknown"`
4. One `TestXxx_IsValid_Lowercase` → asserts `IsValid() == false` for lowercase variant
5. One `TestXxx_String_<Value>` per valid constant → asserts `String()` output

Uses `st.ErrorIsValidTrue`, `st.ErrorIsValidFalse`, `st.ErrorMethodMismatch`.

### 3.2 Value Type Test Pattern (CiString, DateTime, Integer)

**CiString (per variant: 20, 25, 50, 255, 500):**
1. `TestNewCiStringXXType` — valid max-length string succeeds
2. `TestNewCiStringXX_Empty` — empty string returns `ErrEmptyValue`
3. `TestNewCiStringXX_TooLong` — exceeding max length returns error
4. `TestNewCiStringXX_TestValue` — `.Value()` returns original input
5. `TestCiStringXXType_String` — `.String()` returns original input

**DateTime:**
1. `TestNewDateTime` — valid RFC3339 UTC string succeeds
2. `TestNewDateTime_Empty` — empty string returns error
3. `TestNewDateTime_InvalidFormat` — non-RFC3339 returns error
4. `TestNewDateTime_NonUTC` — non-UTC offset returns error
5. `TestDateTime_String` — `.String()` returns RFC3339Nano

**Integer:**
1. `TestNewInteger` — valid value succeeds
2. `TestNewInteger_Zero` — 0 succeeds
3. `TestNewInteger_Max` — 65535 succeeds
4. `TestNewInteger_Negative` — negative returns error
5. `TestNewInteger_Overflow` — >65535 returns error
6. `TestInteger_Value` — `.Value()` returns original

### 3.3 Composite Type Test Pattern

**IdToken:**
1. `TestNewIdToken` — valid CiString20 succeeds
2. `TestNewIdToken_Empty` — empty returns `ErrEmptyValue`
3. `TestNewIdToken_TooLong` — exceeds 20 returns error
4. `TestIdToken_String` — `.String()` returns original

**IdTagInfo:**
1. `TestNewIdTagInfo_Accepted` — valid construction
2. `TestNewIdTagInfo_InvalidStatus` — invalid status returns error
3. `TestIdTagInfo_WithExpiryDate` — builder method works
4. `TestIdTagInfo_WithParentIdTag` — builder method works
5. `TestIdTagInfo_Getters` — all getters return correct values (including nil optionals)

**SampledValue:**
1. `TestNewSampledValue_MinimalValid` — only required field
2. `TestNewSampledValue_AllFields` — all optional fields set
3. `TestNewSampledValue_InvalidValue` — empty required value returns error
4. `TestSampledValue_Getters` — optional getters return correct values or nil

**MeterValue:**
1. `TestNewMeterValue_Valid` — valid timestamp + sampled values
2. `TestNewMeterValue_InvalidTimestamp` — bad timestamp returns error
3. `TestNewMeterValue_EmptySampledValues` — empty slice returns error
4. `TestMeterValue_Getters` — verify all getters

**ChargingSchedulePeriod:**
1. `TestNewChargingSchedulePeriod_Valid`
2. `TestNewChargingSchedulePeriod_InvalidStartPeriod`
3. `TestNewChargingSchedulePeriod_InvalidLimit`
4. `TestChargingSchedulePeriod_Getters`

**ChargingSchedule:**
1. `TestNewChargingSchedule_Valid`
2. `TestNewChargingSchedule_InvalidRateUnit`
3. `TestNewChargingSchedule_EmptyPeriods`
4. `TestNewChargingSchedule_MultipleErrors` — accumulated errors
5. `TestChargingSchedule_OptionalFields` — Duration, MinChargingRate, StartSchedule

### 3.4 Shared Test Helper: `enum_assert_test.go`

Replicate the `enumValidator` pattern from `ocpp16messages`:

```go
type enumValidator struct {
    value     fmt.Stringer
    isValidFn func() bool
}

func assertEnumValid(t *testing.T, validator enumValidator, wantString string)
func assertEnumInvalid(t *testing.T, validator enumValidator)
```

---

## 4. Example Tests (godoc)

Create example test files in the root package directory (not in `tests/`):

```
ocpp16types/
  example_cistring_test.go
  example_datetime_test.go
  example_integer_test.go
  example_idtoken_test.go
  example_idtaginfo_test.go
  example_authorizationstatus_test.go
```

Each file:

- Package: `ocpp16types_test`
- Contains `ExampleNewXxx()` and `ExampleNewXxx_invalid()` functions
- Uses `// Output:` comments for `go test` verification
- Matches the pattern from `ocpp16messages/types/example_cistring_test.go`

---

## 5. Fuzz Tests

Create a `tests_fuzz/` directory at the project root:

```
ocpp16types/
  tests_fuzz/                       ← NEW directory
    doc.go                          ← build tag documentation
    fuzz_cistring_test.go           ← CiString20/25/50 fuzz
    fuzz_cistring255_test.go        ← CiString255 fuzz (separate for length)
    fuzz_cistring500_test.go        ← CiString500 fuzz
    fuzz_datetime_test.go
    fuzz_integer_test.go
    fuzz_idtaginfo_test.go
    fuzz_chargingschedule_test.go
    fuzz_chargingscheduleperiod_test.go
    fuzz_metervalue_test.go
    fuzz_sampledvalue_test.go
```

Each file:

- Build tag: `//go:build fuzz`
- Package: `tests_fuzz`
- Seeds valid corpus values via `f.Add(...)`
- Verifies: no panics, errors are always one of the known sentinels (`ErrEmptyValue`, `ErrInvalidValue`)
- Guards against overly-large inputs with early `t.Skip`

---

## 6. Makefile

Create a `Makefile` mirroring `ocpp16messages`'s structure, adapted to the
flat package layout (no message sub-packages):

### Targets:

| Target          | Description                                              |
|-----------------|----------------------------------------------------------|
| `help`          | Display available targets                                |
| `test`          | Unit tests + coverage (`reports/coverage.out`)           |
| `test-coverage` | HTML coverage report                                     |
| `test-example`  | Run `Example*` tests                                     |
| `test-fuzz`     | Run fuzz tests from `tests_fuzz/`                        |
| `test-race`     | Full suite with `-race`                                  |
| `test-all`      | `lint test test-example test-fuzz test-race`             |
| `lint`          | `golangci-lint` + `go vet` + `staticcheck`               |
| `format`        | `gci` + `gofumpt` + `golines` + `gofmt`                 |
| `pkgsite`       | Local `pkgsite` server at `:8080`                        |

### Adaptations from `ocpp16messages`:

- `coverpkg` filter: exclude `tests_fuzz` directory
- `pkgsite` URL: `github.com/aasanchez/ocpp16types`
- `depguard` allowlist: `github.com/aasanchez/ocpp16types`

---

## 7. golangci.yml — Full v2 Configuration

Replace the current minimal `.golangci.yml` with the full v2 config from
`ocpp16messages`, with these adaptations:

- `depguard.rules.main.allow`: `github.com/aasanchez/ocpp16types` (+ `$gostd`)
- `revive.rules.package-naming`: no exclusion needed (no `types/` sub-packages)
- Keep: `version: "2"`, all formatters, `wsl_v5`, `revive` with `enable-all-rules`
- Keep: output to `reports/golangci-lint.txt`

---

## 8. codecov.yml

Create `codecov.yml` identical to `ocpp16messages`:

```yaml
coverage:
  status:
    project:
      default:
        target: 10%
        threshold: 1%
    patch:
      default:
        target: 10%
        threshold: 1%
```

---

## 9. GitHub Actions CI Enhancement

The current `.github/workflows/ci.yml` is functional but minimal. Align it
with the Makefile-driven approach:

**Recommended changes:**

1. Add `make test` step (replaces bare `go test`)
2. Add `make test-example` step
3. Add codecov upload step after coverage generation
4. Add `make test-race` step
5. Keep the existing `go vet`, `staticcheck`, and `golangci-lint` steps

---

## 10. README.md

Create a README following the `ocpp16messages` pattern:

- Module name and purpose (shared OCPP 1.6 domain types)
- Installation: `go get github.com/aasanchez/ocpp16types`
- Type inventory table (value types, enums, composites)
- Usage examples (constructor, validation, builder)
- Test commands (`make test`, `make test-fuzz`, etc.)
- Zero-dependency guarantee
- License

---

## 11. Supporting Documentation Files

Create if not present (check `ocpp16messages` for exact wording):

- `CONTRIBUTING.md`
- `CLA.md`
- `CODE_OF_CONDUCT.md`

These are typically identical across the ecosystem.

---

## Implementation Order

The recommended execution sequence, respecting dependencies:

| Phase | Files                              | Rationale                                     |
|-------|------------------------------------|-----------------------------------------------|
| 1     | `errors.go`                        | Test constants needed by all test files        |
| 2     | `doc.go`                           | Package documentation before test scaffolding  |
| 3     | `tests/enum_assert_test.go`        | Shared helper needed by enum tests             |
| 4     | `tests/*_test.go` (enums)          | 13 enum types, highest file count              |
| 5     | `tests/*_test.go` (value types)    | CiString, DateTime, Integer                    |
| 6     | `tests/*_test.go` (composites)     | IdToken, IdTagInfo, SampledValue, MeterValue, ChargingSchedule* |
| 7     | `example_*_test.go`                | Godoc examples                                 |
| 8     | `tests_fuzz/`                      | Fuzz tests                                     |
| 9     | `Makefile`                         | Build automation                               |
| 10    | `golangci.yml`                     | Linter config                                  |
| 11    | `codecov.yml`                      | Coverage config                                |
| 12    | `.github/workflows/ci.yml`         | CI pipeline enhancement                        |
| 13    | `README.md`                        | Project documentation                          |
| 14    | `CONTRIBUTING.md`, `CLA.md`, `CODE_OF_CONDUCT.md` | Community docs               |

---

## File Count Summary

| Category         | Files to Create | Notes                               |
|------------------|-----------------|-------------------------------------|
| errors.go        | 1 (modify)      | Add 7 missing constants             |
| doc.go           | 1 (modify)      | Enrich documentation                |
| Unit tests       | 23              | In `tests/` directory               |
| Example tests    | 6               | In root directory                   |
| Fuzz tests       | ~11             | In `tests_fuzz/` directory          |
| Makefile         | 1               | New                                 |
| golangci.yml     | 1 (modify)      | Replace minimal config              |
| codecov.yml      | 1               | New                                 |
| CI workflow       | 1 (modify)      | Enhance existing                    |
| README.md        | 1               | New                                 |
| Community docs   | 3               | New (CONTRIBUTING, CLA, COC)        |
| **Total**        | **~48 files**   | 2 modified + ~46 new                |
