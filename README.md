# VietQR - Golang

[![Go Report Card](https://goreportcard.com/badge/github.com/ducnpdev/vietqr)](https://goreportcard.com/report/github.com/ducnpdev/vietqr)
[![GoDoc](https://godoc.org/github.com/ducnpdev/vietqr?status.svg)](https://godoc.org/github.com/ducnpdev/vietqr)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/badge/Go-1.16%2B-blue.svg)](https://golang.org)

<div align="center">
  <img src="https://res.cloudinary.com/taskmanagereaglob123/image/upload/v1641970995/VietQR.46a78cbb_utwzzh.png" alt="VietQR Logo" width="200"/>
  
  A powerful Golang library for generating VietQR codes for NAPAS payment integration
</div>

## 📋 Table of Contents

- [VietQR - Golang](#vietqr---golang)
  - [📋 Table of Contents](#-table-of-contents)
  - [🌟 Overview](#-overview)
  - [✨ Features](#-features)
  - [📥 Installation](#-installation)
  - [�� Quick Start](#-quick-start)
  - [📖 Usage Examples](#-usage-examples)
    - [Basic Payment QR](#basic-payment-qr)
    - [QR with Additional Data](#qr-with-additional-data)
  - [🔍 QR Code Structure](#-qr-code-structure)
  - [�� Documentation](#-documentation)
  - [🤝 Contributing](#-contributing)
  - [📄 License](#-license)

## 🌟 Overview

VietQR is a Golang library that simplifies the generation of QR codes for NAPAS payment integration in Vietnam. It provides a clean and efficient way to create payment QR codes that comply with the NAPAS VietQR standard.

## ✨ Features

- 🚀 Simple and intuitive API
- 💰 Support for all NAPAS VietQR payment fields
- 🔒 Type-safe implementation
- 📝 Comprehensive field validation
- 🎯 QR code generation for various payment scenarios
- 🔍 QR code parsing and validation
- 📚 Well-documented code

## 📥 Installation

```bash
go get github.com/ducnpdev/vietqr
```

## �� Quick Start

```go
package main

import (
    "fmt"
    "github.com/ducnpdev/vietqr"
)

func main() {
    // Create a VietQR payment request
    qrRequest := vietqr.RequestGenerateViQR{
        MerchantAccountInformation: vietqr.MerchantAccountInformation{
            AccountNo: "999990335280715",
        },
        TransactionAmount: "505000",
        AdditionalDataFieldTemplate: vietqr.AdditionalDataFieldTemplate{
            Description: "Payment for services",
        },
        Mcc:          "5139",
        ReceiverName: "Company ABC",
    }

    // Generate the VietQR code
    qrContent := vietqr.GenerateViQR(qrRequest)
    fmt.Println("Generated VietQR Content:", qrContent)
}
```

## 📖 Usage Examples

### Basic Payment QR
```go
qrRequest := vietqr.RequestGenerateViQR{
    MerchantAccountInformation: vietqr.MerchantAccountInformation{
        AccountNo: "999990335280715",
    },
    TransactionAmount: "100000",
    AdditionalDataFieldTemplate: vietqr.AdditionalDataFieldTemplate{
        Description: "Basic payment",
    },
    Mcc:          "5139",
    ReceiverName: "Merchant Name",
}
```

### QR with Additional Data
```go
qrRequest := vietqr.RequestGenerateViQR{
    MerchantAccountInformation: vietqr.MerchantAccountInformation{
        AccountNo: "999990335280715",
    },
    TransactionAmount: "200000",
    AdditionalDataFieldTemplate: vietqr.AdditionalDataFieldTemplate{
        Description: "Detailed payment information",
        // Add more fields as needed
    },
    Mcc:          "5139",
    ReceiverName: "Merchant Name",
}
```

## 🔍 QR Code Structure

The VietQR code follows the NAPAS standard structure. Here's an example of the QR code field breakdown:

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

## �� Documentation

For detailed documentation about the VietQR standard, please refer to the official NAPAS documentation:
- [NAPAS VietQR Documentation](https://vietqr.net/portal-service/download/documents/QR_Format_T&C_v1.0_VN_092021.pdf)

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

<div align="center">
  Made with ❤️ by <a href="https://github.com/ducnpdev">ducnpdev</a>
</div>