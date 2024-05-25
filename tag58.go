package vietqr

import "fmt"

func genCountryCode(code string) string {
	if code == "" {
		code = "VN"
	}

	return fmt.Sprintf("5802%s", code)
}
