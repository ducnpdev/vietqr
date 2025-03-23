package vietqr

type MerchantAccountInformation struct {
	// mặc định là ngân hàng HDBank: 970437
	AcqID     string `json:"acqID"`
	AccountNo string `json:"accountNo"`
}

type AdditionalDataFieldTemplate struct {
	Description string `json:"description"`
}

type (
	RequestGenerateViQR struct {
		// Phiên bản dữ liệu: tag 00
		PayloadFormatIndicator string `json:"payloadFormatIndicator"`

		// Phương thức khởi tạo, phân biệt QR động vào QR tĩnh: tag 01
		// mặc định là QR động: 010212
		PointInitiationMethod string `json:"pointInitiationMethod"`

		// Thông tin định danh ĐVCNTT: tag
		MerchantAccountInformation `json:"merchantAccountInformation"`

		// Merchant Category Code MCC: tag 52
		// identity merchant type
		// apply different fee base on mcc
		Mcc string `json:"mcc"`

		// Mã tiền tệ: tag 53
		// mặc định là VND: 704
		TransactionCurrency string `json:"transactionCurrency"`

		// Số tiền GD: tag 54
		TransactionAmount string `json:"transactionAmount"`

		// Mã quốc gia: tag 58
		// mặc định là VN
		CountryCode string `json:"countryCode"`

		// Merchant Reciever Name: tag 59
		ReceiverName string `json:"recieveName"`

		// Thông tin bổ sung: tag 62
		AdditionalDataFieldTemplate `json:"additionalDataFieldTemplate"`
	}
)
