package vietqr

import "fmt"

func genMCC(mcc string) string {

	if mcc == "" {
		return ""
	}

	return fmt.Sprintf("5204%s", mcc)
}
