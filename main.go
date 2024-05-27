package vietqr

import (
	"fmt"
)

func main() {
	content := GenerateViQR(RequestGenerateViQR{
		MerchantAccountInformation: MerchantAccountInformation{
			AccountNo: "999990335280715",
		},
		TransactionAmount: "505000",

		AdditionalDataFieldTemplate: AdditionalDataFieldTemplate{
			Description: "test noi dungtest noi dungtest noi dungtest noi dung",
		},
	})

	fmt.Println("content-main:", content)
}
