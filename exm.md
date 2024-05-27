# VIETQR
![](https://res.cloudinary.com/taskmanagereaglob123/image/upload/v1641970995/VietQR.46a78cbb_utwzzh.png)

- Support draw QR code from data bank ( accountName, amount, memo,....) with many templates 
- Support create link URL from QR code

## Table of Contents

- [VIETQR](#vietqr)
  - [Table of Contents](#table-of-contents)
  - [Features](#features)
  - [Choose what you need](#choose-what-you-need)
  - [Try It!](#try-it)
  - [Demo preview](#demo-preview)
  - [Installation](#installation)
  - [Example](#example)
  - [VietQR API](#vietqr-api)
    - [getTemplate()](#gettemplate)
      - [Response successfully](#response-successfully)
    - [getBanks()](#getbanks)
      - [Response successfully](#response-successfully-1)
    - [genQuickLink()](#genquicklink)
    - [genQRCodeBase64()](#genqrcodebase64)
      - [Response successfully](#response-successfully-2)
  - [License](#license)

## Features
- Support draw QR code from data bank ( accountName, amount, memo,....) with many templates 
- Support create link URL from QR code

## Choose what you need

| Project | Support |
| --- | --- |
| [VietQR](https://github.com/vietqr/vietqr-node) | - Running with DOM on CLIENT-SIDE: Browser(IE6+, Chrome, Firefox, Safari, Opera, Mobile Safari, Android, Windows Mobile, etc.), Electron, NW.js, etc. <br/>- Running without DOM on SERVER-SIDE: Save image to file(PNG/JPEG/Base64) or get data url text.  NodeJS, Electron, NW.js, etc. <br/>- A QRCode generator for React Native: Generate QRCode image or get base64 data url text.|
## Try It!

[Try It!](http://my.vietqr.io/ "Easy To Try It!")

## Demo preview

![Demo preview](https://res.cloudinary.com/taskmanagereaglob123/image/upload/v1641971750/photo_2022-01-12_14.15.13_ozxy6g.jpg)

## Installation
- Download install

    [https://github.com/vietqr/vietqr-node](https://github.com/vietqr/vietqr-node)

- Npm install

	```BASH
	npm install vietqr
	```
## Example

```javascript
import {VietQR} from 'vietqr';


let vietQR = new VietQR({
    clientID: 'de8a0804-a76d-41e5-8ad6-31503ce7d5f4',
    apiKey: '17c29f09-4ea2-4417-b9c2-7f020d35de42',
});

// list banks are supported create QR code by Vietqr
vietQR.getBanks().then((banks)=>{console.log(banks)})
.catch((err)=>{});

// list templates are supported by Vietqr
vietQR.getTemplate().then((data)=>{console.log(data)})
.catch((err)=>{});


// create QR code from data
vietQR.genQRCodeBase64({
    bank: '970415',
    accountName: 'QUY VAC XIN PHONG CHONG COVID',
    accountNumber: '113366668888',
    amount: '79000',
    memo: 'Ung Ho Quy Vac Xin',
    template: 'compact'
}).then((data)=>{console.log(data)})
.catch((err)=>{});

```


## VietQR API

| Option | Description | 
| --- | --- |
| **bank** | Get the bin field corresponding to the data returned in the banks API (VietQR.getTemplate() or [API Bank](https://www.vietqr.io/danh-sach-api/api-tao-ma-qr/api-danh-sach-ma-ngan-hang)) |
| **accountName** | Bank account name |
| **accountNumber** | Bank account number |
| **amount** | Amount of money |
| **memo** | Money transfer content |
| **template** | Type of template returned to the user |
| **clientID** | api_key is provided when registering an account at http://my.vietqr.io/ 
| **apiKey** | client_key is provided when registering an account at http://my.vietqr.io/ 

### getTemplate()
```javascript
vietQR.getTemplate().then((data)=>{console.log(data)})
.catch((err)=>{});
```
#### Response successfully
```javascript
 {
  code: '00',
  desc: 'success',
  data: [
    {
      name: 'QR Only',
      template: 'qr_only',
      demo: 'https://api.vietqr.io/Vietinbank/113366668888/790000/Gop%20Quy/qr_only.jpg?accountName=Quy%20Vacxin%20Covid'
    },
    {
      name: 'Compact',
      template: 'compact',
      demo: 'https://api.vietqr.io/Vietinbank/113366668888/790000/Gop%20Quy/compact.jpg?accountName=Quy%20Vacxin%20Covid'
    },
    {
      name: 'Compact 2',
      template: 'compact2',
      demo: 'https://api.vietqr.io/Vietinbank/113366668888/790000/Gop%20Quy/compact2.jpg?accountName=Quy%20Vacxin%20Covid'
    }
  ]
}
```

### getBanks()
```javascript
vietQR.getBanks().then((data)=>{console.log(data)})
.catch((err)=>{});
```
#### Response successfully
```javascript
{
  code: '00',
  desc: 'Get Bank list successful! Total 52 banks',
  data: [
    {
      id: 17,
      name: 'Ngân hàng TMCP Công thương Việt Nam',
      code: 'ICB',
      bin: '970415',
      isTransfer: 1,
      short_name: 'VietinBank',
      logo: 'https://api.vietqr.io/img/ICB.3d4d6760.png',
      support: 3
    },
    ....
  ]
}
```


### genQuickLink()
```javascript
let link = vietQR.genQuickLink({
        bank: '970415',
        accountName: 'QUY VAC XIN PHONG CHONG COVID',
        accountNumber: '113366668888',
        amount: '79000',
        memo: 'Ung Ho Quy Vac Xin',
        template: 'compact', 
        media: '.jpg' 
    });
console.log(link);
```

### genQRCodeBase64()
```javascript
vietQR.genQRCodeBase64({
    bank: '970415',
    accountName: 'QUY VAC XIN PHONG CHONG COVID',
    accountNumber: '113366668888',
    amount: '79000',
    memo: 'Ung Ho Quy Vac Xin',
    template : 'qr_only'
}).then((data)=>{console.log(data)})
.catch((err)=>{});
```
#### Response successfully
```javascript
{
    "code": "00",
    "desc": "Gen VietQR successful!",
    "data": {
        "acqId": "970415",
        "accountName": "QUY VAC XIN PHONG CHONG COVID",
        "qrDataURL": "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAOwAAADeCAYAAAA..."
    }
}
```


## License
MIT License