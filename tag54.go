package vietqr

import "fmt"

func genTransactionAmount(amount string) string {

	lenAmount := len(amount)

	if lenAmount >= 10 {
		return fmt.Sprintf("54%d%s", lenAmount, amount)
	}

	return fmt.Sprintf("540%d%s", lenAmount, amount)
}
