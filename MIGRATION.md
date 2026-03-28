# Migration Plan: ocpp16messages → ocpp16types

> Goal: Replace `ocpp16messages/types/` with a dependency on
> `github.com/aasanchez/ocpp16types`, eliminating duplicate type definitions
> and establishing `ocpp16types` as the single source of truth.

---

## Prerequisites

Before starting this migration:

1. **ocpp16types README.md** — Create the missing README (PLAN.md phase 13).
2. **ocpp16types CODE_OF_CONDUCT.md** — Create the missing community doc (PLAN.md phase 14).
3. **Verify full CI green** on `ocpp16types` main branch.
4. **Tag a release** on `ocpp16types` (e.g., `v0.1.0`) so `ocpp16messages` can pin a version.

---

## Scope

### In scope (core shared types — 20 files in `ocpp16messages/types/`)

All 20 type files in `ocpp16messages/types/` have 1:1 equivalents in `ocpp16types`:

| Category | Types |
|----------|-------|
| Value types | CiString20, CiString25, CiString50, CiString255, CiString500, DateTime, Integer |
| Enums (shared) | Location, Measurand, Phase, ReadingContext, UnitOfMeasure, ValueFormat |
| Composites | IdToken, IdTagInfo, SampledValue, MeterValue, ChargingSchedule, ChargingSchedulePeriod |
| Infrastructure | errors.go, doc.go |

### Out of scope (message-specific types)

These types live in per-message `types/` sub-packages and are **not** part of this
migration. They remain owned by `ocpp16messages`:

- `bootnotification/types/` — RegistrationStatus (note: also in ocpp16types)
- `statusnotification/types/` — ChargePointErrorCode, ChargePointStatus (also in ocpp16types)
- `stoptransaction/types/` — Reason / StopReason (also in ocpp16types)
- `remotestarttransaction/types/` — RemoteStartTransactionStatus
- `sendlocallist/types/` — AuthorizationData, UpdateStatus, UpdateType
- `setchargingprofile/types/` — ChargingProfile, ChargingProfileKindType, RecurrencyKindType
- ~17 additional message-specific type packages

**Future consideration:** Some message-specific types (RegistrationStatus,
ChargePointErrorCode, ChargePointStatus, StopReason) already exist in `ocpp16types`.
A follow-up migration can collapse those message-specific packages into
`ocpp16types` imports as well, but that is a separate effort.

---

## Migration Strategy

### Approach: Big-bang replacement per message package

Since the type APIs are identical (same names, same constructors, same signatures),
the migration is a mechanical import-path rewrite. No behavioral changes.

### Phase 1: Add ocpp16types dependency

**File:** `ocpp16messages/go.mod`

```
require github.com/aasanchez/ocpp16types v0.1.0
```

If working locally with the Go workspace (`go.work`), a `replace` directive
works during development:

```
replace github.com/aasanchez/ocpp16types => ../ocpp16types
```

### Phase 2: Rewrite imports in all message packages

For every message file that currently imports `ocpp16messages/types`:

**Before:**
```go
import (
    st "github.com/aasanchez/ocpp16messages/types"
)
```

**After:**
```go
import (
    st "github.com/aasanchez/ocpp16types"
)
```

The `st` alias is already used consistently across message files, so all
downstream references (`st.CiString20Type`, `st.NewDateTime(...)`, etc.)
remain unchanged.

Files that import without alias (`types.XYZ`) should adopt the `st` alias
to match the established convention and avoid collision with message-specific
`types` sub-packages:

**Before:**
```go
import (
    "github.com/aasanchez/ocpp16messages/types"
)
// types.IdToken
```

**After:**
```go
import (
    st "github.com/aasanchez/ocpp16types"
)
// st.IdToken
```

### Phase 3: Handle dual-import files

Some message files import **both** shared types and message-specific types:

```go
import (
    mbt "github.com/aasanchez/ocpp16messages/bootnotification/types"
    st  "github.com/aasanchez/ocpp16messages/types"
)
```

These become:

```go
import (
    mbt "github.com/aasanchez/ocpp16messages/bootnotification/types"
    st  "github.com/aasanchez/ocpp16types"
)
```

The message-specific alias (`mbt`, `slt`, etc.) remains untouched.

### Phase 4: Update tests

Message-level tests that reference shared types via `st` or `types` need
the same import rewrite. The test logic itself does not change since the
API surface is identical.

### Phase 5: Delete `ocpp16messages/types/`

Once all imports are rewritten and CI is green:

1. Delete `ocpp16messages/types/` entirely (20 source files + `tests/` + `doc.go` + `errors.go`).
2. Delete `ocpp16messages/types/tests/` (18 test files + shared helper).
3. Verify no remaining internal references: `grep -r "ocpp16messages/types" .`

### Phase 6: Update go.work (if applicable)

If the workspace file lists both modules, ensure the dependency graph is correct:

```
use (
    ./ocpp16types
    ./ocpp16messages
    ./ocpp16store
)
```

### Phase 7: Validate

1. `go build ./...` — all message packages compile.
2. `go vet ./...` — no issues.
3. `make test` (or equivalent) — all message-level tests pass.
4. `make lint` — golangci-lint clean.
5. Verify `go mod tidy` does not reintroduce removed internal imports.

---

## Affected Files (estimated)

| Area | Count | Change |
|------|-------|--------|
| Message source files (`*/request.go`, `*/confirmation.go`) | ~23 | Import rewrite |
| Message test files | ~23 | Import rewrite |
| `go.mod` | 1 | Add `require ocpp16types` |
| `types/` directory | ~38 files | **Delete** |
| Makefile / CI | ~2 | Update coverpkg exclusions |
| **Total** | ~87 | |

---

## Risk Assessment

| Risk | Mitigation |
|------|-----------|
| API drift between ocpp16types and ocpp16messages/types | Verified 1:1 parity — identical names, constructors, signatures |
| Breaking downstream consumers of ocpp16messages | ocpp16messages exports message types, not shared types directly; consumers use message structs |
| Message-specific types confused with shared types | Out of scope — clearly separated by per-message sub-packages |
| Go module version resolution | Pin ocpp16types to a tagged release; use `go.work` replace for local dev |

---

## Execution Order

| Step | Action | Validates |
|------|--------|-----------|
| 0 | Tag `ocpp16types` v0.1.0 | Stable dependency target |
| 1 | Add `require` to `ocpp16messages/go.mod` | Dependency resolves |
| 2 | Rewrite imports in one message package (e.g., `authorize/`) | Pattern works end-to-end |
| 3 | Run tests for that package | No regressions |
| 4 | Rewrite remaining message packages | Mechanical repetition |
| 5 | Run full test suite | Global validation |
| 6 | Delete `ocpp16messages/types/` | Clean break |
| 7 | Final `go mod tidy` + `go build ./...` + `make test-all` | Full green |

---

## Post-Migration Follow-ups

1. **Collapse message-specific duplicates:** RegistrationStatus, ChargePointErrorCode,
   ChargePointStatus, and StopReason exist in both `ocpp16types` and their respective
   message-specific `types/` packages. A subsequent migration can eliminate those
   message-level duplicates too.

2. **Update depguard allowlist:** `golangci.yml` in `ocpp16messages` should allow
   `github.com/aasanchez/ocpp16types` and disallow `github.com/aasanchez/ocpp16messages/types`.

3. **Update ocpp16messages README/docs** to document the dependency on `ocpp16types`.

4. **Consider automation:** A `sed`/`gofmt` script can perform the mechanical
   import rewrite across all files in one pass.
