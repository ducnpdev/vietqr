package vietqr

func GenerateViQR(req RequestGenerateViQR) string {
	return genContent(req)
}

func genContent(req RequestGenerateViQR) string {

	content :=
		genPayloadFormatIndicator(req.PayloadFormatIndicator) +
			genPointInitiationMethod(req.PointInitiationMethod) +
			genMerchantAccountInformation(MerchantAccountInformation{
				AcqID:     req.AcqID,
				AccountNo: req.AccountNo,
			}) +
			genMCC(req.Mcc) +
			genTransactionCurrency(req.TransactionCurrency) +
			genTransactionAmount(req.TransactionAmount) +
			genCountryCode(req.CountryCode) +
			genRecieverName(req.ReceiverName) +
			genAdditionalDataFieldTemplate(req.AdditionalDataFieldTemplate)

	crc := genCRC(content)

	contentCRC := content + crc
	return contentCRC
}
