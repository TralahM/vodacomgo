[![License: GPLv3](https://img.shields.io/badge/License-GPLV2-green.svg)](https://opensource.org/licenses/GPLV2)
[![Organization](https://img.shields.io/badge/Org-TralahTek-blue.svg)](https://github.com/TralahTek)
[![Views](http://hits.dwyl.io/TralahM/vodacomgo.svg)](http://dwyl.io/TralahM/vodacomgo)
[![PRs Welcome](https://img.shields.io/badge/PRs-Welcome-brightgreen.svg?style=flat-square)](https://github.com/TralahM/vodacomgo/pull/)
[![GitHub pull-requests](https://img.shields.io/badge/Issues-pr-red.svg?style=flat-square)](https://github.com/TralahM/vodacomgo/pull/)
[![Language](https://img.shields.io/badge/Language-go-00ADD8.svg)](https://github.com/TralahM)
<img title="Watching" src="https://img.shields.io/github/watchers/TralahM/vodacomgo?label=Watchers&color=blue&style=flat-square">
<img title="Stars" src="https://img.shields.io/github/stars/TralahM/vodacomgo?color=red&style=flat-square">
<img title="Forks" src="https://img.shields.io/github/forks/TralahM/vodacomgo?color=green&style=flat-square">
<noscript><a href="https://liberapay.com/TralahM/donate"><img alt="Donate using Liberapay" src="https://liberapay.com/assets/widgets/donate.svg"></a></noscript>

# vodacomgo.
A simple package to parse and interact with vodacom mpesa integrated payment
gateway.


# Installation
```console
# In terminal do:
$ go get -u github.com/tralahm/vodacomgo
```


# Documentation
## Generate XML Request Bodies for Login,C2B,B2C
```go
package main

import "ioutil"
import "bytes"
import "github.com/tralahm/vodacomgo"

login:=vodacomgo.Login{Username:"thirdpartyc2bw",Password:"thirdpartyc2bw"}
loginpayload:=new(bytes.Buffer)
vodacomgo.GenLogin(loginpayload,login)

c2b:=vodacomgo.C2B{
    Token: "sessionID",
    CustomerMSISDN: "243812345678",
    Amount: "4000",
    Currency: "CDF",
    ThirdPartyReference: "Your thirdparty Reference.",
    Language: "EN",
    Surname: "Surname",
    Initials: "Initials"
    ServiceProviderCode: "39209392",
    CommandID: "CommandID",
    CallBackChannel: "2",
    CallBackDestination: "url",
    Date: "date",
}
c2bpayload:=new(bytes.Buffer)
vodacomgo.GenC2B(c2bpayload,c2b)
req,err:=http.NewRequest(http.MethodPost,addr,c2bpayload)
response,_:=http.DefaultClient().Do(req)

b2c:=vodacomgo.B2C{
    Token              :"string",
    ServiceProviderName:"string",
    CustomerMSISDN     :"string",
    Currency           :"string",
    Amount             :"string",
    TransactionDateTime:"string",
    Shortcode          :"string",
    Language           :"string",
    ThirdPartyReference:"string",
    CallBackChannel    :"string",
    CallBackDestination:"string",
    CommandID          :"string",
}
b2cpayload:=new(bytes.Buffer)
vodacomgo.GenB2C(b2cpayload,b2c)

```
## Parse and Decode XML Responses to Structs and Generate Simpler Structs from Parsed XML Struct Responses

```go
responseBytes,_:=ioutil.ReadAll(response.Body)

loginresult:=vodacomgo.DecodeLoginResponse(responseBytes)
loginresponse:=loginresult.ToResponse()
sessionID:=loginresponse.SessionID

c2bresult:=vodacomgo.DecodeC2BResponse(responseBytes)
c2bresponse:=c2bresult.ToResponse()

b2cresult:=vodacomgo.DecodeC2BResponse(responseBytes)
b2cresponse:=b2cresult.ToResponse()
```

## Use the Defined XML strings to send callback acknowledgement responses to IPG
```go
b2c_acknowledgment:=[]byte(vodacomgo.AckB2CT)
http.NewRequest(http.MethodPost,addr,b2c_acknowledgment)
c2b_acknowledgment:=[]byte(vodacomgo.AckC2BT)
```


# LICENCE

[Read the license here](LICENSE)


# Self-Promotion

[![](https://img.shields.io/badge/Github-TralahM-green?style=for-the-badge&logo=github)](https://github.com/TralahM)
[![](https://img.shields.io/badge/X-%40tralahtek-blue?style=for-the-badge&logo=x)](https://x.com/TralahM)
[![TralahM](https://img.shields.io/badge/Kaggle-TralahM-purple.svg?style=for-the-badge&logo=kaggle)](https://kaggle.com/TralahM)
[![TralahM](https://img.shields.io/badge/LinkedIn-TralahM-white.svg?style=for-the-badge&logo=linkedin)](https://linkedin.com/in/TralahM)

[![Blog](https://img.shields.io/badge/Blog-tralahm.github.io-blue.svg?style=for-the-badge&logo=rss)](https://tralahm.github.io)

[![TralahTek](https://img.shields.io/badge/Organization-TralahTek-cyan.svg?style=for-the-badge)](https://org.tralahtek.com)

