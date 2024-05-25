package vietqr

// Phương thức khởi tạo
func genPointInitiationMethod(method string) string {
	if method == "" {
		method = "010212"
	}

	return method
}
