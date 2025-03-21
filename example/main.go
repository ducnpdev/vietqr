package main

import (
	"fmt"

	"github.com/ducnpdev/vietqr"
)

func main() {
	content := vietqr.GenerateViQR(vietqr.RequestGenerateViQR{
		MerchantAccountInformation: vietqr.MerchantAccountInformation{
			AccountNo: "999990335280715",
		},
		TransactionAmount: "505000",

		AdditionalDataFieldTemplate: vietqr.AdditionalDataFieldTemplate{
			Description: "test noi dungtest noi dungtest noi dungtest noi dung",
		},
		// Mcc: "5139",
	})

	fmt.Println("content-main:", content)
}
