package vietqr

import "fmt"

func genMerchantAccountInformation(accInfo MerchantAccountInformation) string {

	valueTag38 := genMerchantAccountInformationTag00() +
		genMerchantAccountInformationTag01(accInfo.AcqID, accInfo.AccountNo) +
		genMerchantAccountInformationTag02()

	tag38 := fmt.Sprintf("38%s%s", fmt.Sprintf("%d", len(valueTag38)), valueTag38)

	return tag38
}

// func tạo ra tag-00 in tag 38
func genMerchantAccountInformationTag00() string {
	return "0010A000000727"
}

// func tạo ra tag-01 in tag 38
func genMerchantAccountInformationTag01(acqID, accountNo string) string {
	var tag01 string

	// sub tag 00 in tag 01
	/*
		00 là sub tab
		06 là độ dài của bank code, mã bank tại VN là 6 số
	*/
	if acqID == "" {
		acqID = "970437"
	}
	subTag00 := fmt.Sprintf("0006%s", acqID)

	// sub tag 01 in tag 01
	subTag01 := fmt.Sprintf("01%s%s", fmt.Sprintf("%d", len(accountNo)), accountNo)

	tag01 = subTag00 + subTag01
	return fmt.Sprintf("01%s%s", fmt.Sprintf("%d", len(tag01)), tag01)
}

// func tạo ra tag-02 in tag 38
func genMerchantAccountInformationTag02() string {
	return "0208QRIBFTTA"
}
