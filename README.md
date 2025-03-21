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
package vietqr

import (
	"fmt"
)

func main() {
	content := GenerateViQR(RequestGenerateViQR{
		MerchantAccountInformation: MerchantAccountInformation{
			AccountNo: "999990335280715",
		},
		TransactionAmount: "505000",
		AdditionalDataFieldTemplate: AdditionalDataFieldTemplate{
			Description: "test noi dung",
		},
	})

	fmt.Println("content-main:", content)
    // => 00020101021238590010A0000007270129000697043701159999903352807150208QRIBFTTA53037045406505005802VN62170813test noi dung630433C4
}
```

## Tách Các Filed Tag
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
6304
    33C4
```

## License
MIT License