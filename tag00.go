package vietqr

// Phiên bản dữ liệu
func genPayloadFormatIndicator(indicator string) string {
	if indicator == "" {
		indicator = "000201"
	}
	return indicator
}
