package vietqr

import "fmt"

func genRecieverName(name string) string {
	if name == "" {
		return ""
	}
	lenName := len(name)

	if lenName >= 10 {
		return fmt.Sprintf("59%d%s", lenName, name)
	}

	return fmt.Sprintf("590%d%s", lenName, name)
}
