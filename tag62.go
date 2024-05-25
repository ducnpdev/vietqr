package vietqr

import "fmt"

func genAdditionalDataFieldTemplate(req AdditionalDataFieldTemplate) string {

	valueTag62 := genAdditionalDataFieldTemplateTag08(req.Description)
	if valueTag62 == "" {
		return ""
	}

	tag62 := fmt.Sprintf("62%s%s", fmt.Sprintf("%d", len(valueTag62)), valueTag62)
	return tag62
}

// func tạo ra tag-08 in tag 62
// nội dung giao dịch
func genAdditionalDataFieldTemplateTag08(description string) string {
	if description == "" {
		return ""
	}
	tag08 := fmt.Sprintf("08%d%s", len(description), description)
	fmt.Println("tag08:", tag08)
	return tag08
}
