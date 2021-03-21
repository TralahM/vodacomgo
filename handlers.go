package vodacomgo

import (
	"encoding/xml"
	"io"
	"text/template"
)

func GenLogin(wr io.Writer, data Login) {
	t, err := template.New("login").Parse(LoginT)
	if err != nil {
		panic(err)
	}
	err = t.Execute(wr, data)
	if err != nil {
		panic(err)
	}
}

func GenB2C(wr io.Writer, data B2C) {
	t, err := template.New("b2c").Parse(B2CT)
	if err != nil {
		panic(err)
	}
	err = t.Execute(wr, data)
	if err != nil {
		panic(err)
	}
}

func GenC2B(wr io.Writer, data C2B) {
	t, err := template.New("c2b").Parse(C2BT)
	if err != nil {
		panic(err)
	}
	err = t.Execute(wr, data)
	if err != nil {
		panic(err)
	}
}

func DecodeLoginResponse(content []byte) LoginEnvelope {
	var lxml LoginEnvelope
	_ = xml.Unmarshal([]byte(content), &lxml)
	return lxml
}
func DecodeC2BResponse(content []byte) C2BEnvelope {
	var lxml C2BEnvelope
	_ = xml.Unmarshal([]byte(content), &lxml)
	return lxml
}
func DecodeB2CResponse(content []byte) B2CEnvelope {
	var lxml B2CEnvelope
	_ = xml.Unmarshal([]byte(content), &lxml)
	return lxml
}
func DecodeC2BCallback(content []byte) C2BCallbackEnvelope {
	var lxml C2BCallbackEnvelope
	_ = xml.Unmarshal([]byte(content), &lxml)
	return lxml
}
func DecodeB2CCallback(content []byte) B2CCallbackEnvelope {
	var lxml B2CCallbackEnvelope
	_ = xml.Unmarshal([]byte(content), &lxml)
	return lxml
}
