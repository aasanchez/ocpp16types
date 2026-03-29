package ocpp16types_test

import (
	"fmt"

	st "github.com/aasanchez/ocpp16types"
)

func ExampleNewAuthorizationData() {
	expiryDate := "2027-12-31T23:59:59Z"

	authData, err := st.NewAuthorizationData(st.AuthorizationDataInput{
		IdTag: "TAG001",
		IdTagInfo: &st.IdTagInfoInput{
			Status:     "Accepted",
			ExpiryDate: &expiryDate,
		},
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(authData.IdTag().Value())
	fmt.Println(authData.IdTagInfo().Status().String())
	// Output:
	// TAG001
	// Accepted
}

func ExampleNewAuthorizationData_deleteEntry() {
	authData, err := st.NewAuthorizationData(st.AuthorizationDataInput{
		IdTag:     "TAG002",
		IdTagInfo: nil,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(authData.IdTag().Value())
	fmt.Println(authData.IdTagInfo() == nil)
	// Output:
	// TAG002
	// true
}
