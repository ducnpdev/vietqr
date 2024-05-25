package vietqr

import "fmt"

func genTransactionCurrency(currency string) string {

	if currency == "" {
		currency = "704"
	}

	return fmt.Sprintf("5303%s", currency)

}
