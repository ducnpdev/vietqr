package vietqr

import (
	"fmt"
)

func main() {
	content := GenerateViQR(RequestGenerateViQR{
		// PayloadFormatIndicator: "000201",
		// PointInitiationMethod:  "010212",

		MerchantAccountInformation: MerchantAccountInformation{
			// AcqID:     "970437",
			AccountNo: "999990335280715",
		},
		// TransactionCurrency: "704",

		TransactionAmount: "505000",

		AdditionalDataFieldTemplate: AdditionalDataFieldTemplate{
			Description: "test noi dungtest noi dungtest noi dungtest noi dung",
		},
	})

	fmt.Println("content-main:", content)
}
