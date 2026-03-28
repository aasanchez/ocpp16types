package ocpp16types_test

import (
	"fmt"

	st "github.com/aasanchez/ocpp16types"
)

func ExampleNewDateTime() {
	dateTime, err := st.NewDateTime("2024-01-15T10:30:00Z")
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(dateTime.String())
	// Output:
	// 2024-01-15T10:30:00Z
}

func ExampleNewDateTime_invalid() {
	_, err := st.NewDateTime("not-a-date")
	if err != nil {
		fmt.Println(err)

		return
	}
	//nolint:lll,revive // Output comment must match exact error string
	// Output:
	// NewDateTime: value: invalid value: parsing time "not-a-date" as "2006-01-02T15:04:05Z07:00": cannot parse "not-a-date" as "2006"
}
