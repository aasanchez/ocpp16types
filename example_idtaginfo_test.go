package ocpp16types_test

import (
	"fmt"

	st "github.com/aasanchez/ocpp16types"
)

func ExampleNewIdTagInfo() {
	info, err := st.NewIdTagInfo(st.AuthorizationStatusAccepted)
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(info.Status().String())
	// Output:
	// Accepted
}

func ExampleNewIdTagInfo_invalid() {
	_, err := st.NewIdTagInfo(st.AuthorizationStatus("Bogus"))
	if err != nil {
		fmt.Println(err)

		return
	}
	// Output:
	// invalid value
}
