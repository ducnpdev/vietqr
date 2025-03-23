# VietQR - Golang
![](https://res.cloudinary.com/taskmanagereaglob123/image/upload/v1641970995/VietQR.46a78cbb_utwzzh.png)

- thư viện hỗ trợ mã QR để thanh toán thông qua NAPAS với ngôn ngữ lập trình `Golang`

## Table of Contents

- [VietQR - Golang](#vietqr---golang)
  - [Table of Contents](#table-of-contents)
  - [Napas Document](#napas-document)
  - [Chức Năng](#chức-năng)
  - [Cài Đặt](#cài-đặt)
  - [Ví Dụ:](#ví-dụ)
  - [Tách Các Filed Tag](#tách-các-filed-tag)
  - [License](#license)

## Napas Document
- link document của napas: https://vietqr.net/portal-service/download/documents/QR_Format_T&C_v1.0_VN_092021.pdf
## Chức Năng
- hỗ trợ tạo ra qr để thanh toán( accountNo, amount, nội dung, ...)

## Cài Đặt

```bash
go get github.com/ducnpdev/vietqr
```

## Ví Dụ:
```golang
package main

import (
	"fmt"

	"github.com/ducnpdev/vietqr"
)

func main() {
	content := vietqr.GenerateViQR(vietqr.RequestGenerateViQR{
		MerchantAccountInformation: vietqr.MerchantAccountInformation{
			AccountNo: "999990335280715",
		},
		TransactionAmount: "505000",

		AdditionalDataFieldTemplate: vietqr.AdditionalDataFieldTemplate{
			Description: "test noi dungtest noi dungtest noi dungtest noi dung",
		},
		Mcc:          "5139",
		ReceiverName: "Cty ABC",
	})

	fmt.Println("content-main:", content)
}
```

## Tách Các Filed Tag
- golang vietqr detech qr.
```txt
0002
    01
0102
    12
3859
    0010
        A000000727
    0129
        0006
            970437
        0115
            999990335280715
        0208
            QRIBFTTA
5204
    5139
5303
    704
5406
    50500
5802
    VN
6217
    0813
        test noi dung
5907
    Cty ABC
6304
    33C4
```
- detech qr trên mạnh:
```txt
0002010102125406930000
3857
    0010
        A000000727
    0127
        0006
            970437
        0113
            QR50317S75UDT
    0208
        QRIBFTTA
5303
    704
5802
    VN
6292
    0325
        TRUONG TIEU HOC NGHIA TAN
    0719
        02770002virtualNONE
    0836
        BYMII0HL5E AD87PI 03 2025 THT SON 4P
6304614A
```

## License
MIT License