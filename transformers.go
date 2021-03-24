// Transform XML Envelope Structs into their Json Counterparts

package vodacomgo

import (
	"encoding/json"
	"reflect"
)

type DataItem struct {
	Text  string `xml:",chardata"`
	Name  string `xml:"name"`
	Type  string `xml:"type"`
	Value string `xml:"value"`
} // `xml:"dataItem"`

func toResponse(lr interface{}, dt interface{}) []byte {
	u := reflect.ValueOf(lr)
	v := reflect.ValueOf(dt)
	typeofU := u.Type()
	dataItemMap := make(map[string]string)
	namevalx := make(map[string]int)
	for i := 0; i < v.NumField(); i++ {
		tag := v.Type().Field(i).Tag //Current tag
		switch tag.Get("xml") {
		case ",chardata":
			continue
		case "name":
			namevalx["name"] = i
			continue
		case "type":
			continue
		case "value":
			namevalx["value"] = i
			continue
		}

	}
	dataItemMap[v.Field(namevalx["name"]).String()] = (v.Field(namevalx["value"]).String())
	lresmap := make(map[string]string)
	for h := 0; h < u.NumField(); h++ {
		// log.Printf("%s\t: %v\n", typeofU.Field(h).Name, u.Field(h).Interface())
		name := typeofU.Field(h).Name
		field := u.Field(h)
		if _, ok := dataItemMap[name]; ok {
			lresmap[string(name)] = string(dataItemMap[name])
		} else {
			lresmap[string(name)] = string(field.String())
		}
	}
	jstr, _ := json.Marshal(lresmap)
	return jstr
}

// LoginEnvelopeToLoginResponse
func (e *LoginEnvelope) ToResponse() LoginResponse {
	var r LoginResponse
	var lrmobj LoginResponse

	r = LoginResponse{
		Code:          e.Body.GetGenericResultResponse.SOAPAPIResult.EventInfo.Code,
		Description:   e.Body.GetGenericResultResponse.SOAPAPIResult.EventInfo.Description,
		Detail:        e.Body.GetGenericResultResponse.SOAPAPIResult.EventInfo.Detail,
		TransactionID: e.Body.GetGenericResultResponse.SOAPAPIResult.EventInfo.TransactionID,
		SessionID:     e.Body.GetGenericResultResponse.SOAPAPIResult.Response.DataItem.Value,
		EventID:       e.Header.Eventid.Text,
	}
	for _, dit := range e.Body.GetGenericResultResponse.SOAPAPIResult.Request.DataItem {
		jstr := toResponse(r, dit)
		json.Unmarshal(jstr, &lrmobj)
		r = lrmobj
	}
	r.EventID = e.Header.Eventid.Text
	// log.Printf("LoginResponse: %v eventid: %v\n", r, r.EventID)
	return r
}

// C2BEnvelopeToC2BResponse
func (e *C2BEnvelope) ToResponse() C2BResponse {
	var r, c2ob C2BResponse
	r = C2BResponse{
		Code:          e.Body.GetGenericResultResponse.SOAPAPIResult.EventInfo.Code,
		Description:   e.Body.GetGenericResultResponse.SOAPAPIResult.EventInfo.Description,
		Detail:        e.Body.GetGenericResultResponse.SOAPAPIResult.EventInfo.Detail,
		EventID:       e.Header.Eventid.Text,
		TransactionID: e.Body.GetGenericResultResponse.SOAPAPIResult.EventInfo.TransactionID,
	}
	for _, dit := range e.Body.GetGenericResultResponse.SOAPAPIResult.Request.DataItem {
		// r = toResponse(r, dit)
		json.Unmarshal(toResponse(r, dit), &c2ob)
		r = c2ob
	}
	for _, dit := range e.Body.GetGenericResultResponse.SOAPAPIResult.Response.DataItem {
		// r = toResponse(r, dit)
		json.Unmarshal(toResponse(r, dit), &c2ob)
		r = c2ob
	}
	r.EventID = e.Header.Eventid.Text
	// log.Printf("C2BResponse: %v eventid: %v\n", r, r.EventID)
	return r
}

// B2CEnvelopeToB2CResponse
func (e *B2CEnvelope) ToResponse() B2CResponse {
	var r, c2ob B2CResponse
	r = B2CResponse{
		Code:          e.Body.GetGenericResultResponse.SOAPAPIResult.EventInfo.Code,
		Description:   e.Body.GetGenericResultResponse.SOAPAPIResult.EventInfo.Description,
		Detail:        e.Body.GetGenericResultResponse.SOAPAPIResult.EventInfo.Detail,
		EventID:       e.Header.Eventid.Text,
		TransactionID: e.Body.GetGenericResultResponse.SOAPAPIResult.EventInfo.TransactionID,
	}
	for _, dit := range e.Body.GetGenericResultResponse.SOAPAPIResult.Request.DataItem {
		// r = toResponse(r, dit)
		json.Unmarshal(toResponse(r, dit), &c2ob)
		r = c2ob
	}
	for _, dit := range e.Body.GetGenericResultResponse.SOAPAPIResult.Response.DataItem {
		// r = toResponse(r, dit)
		json.Unmarshal(toResponse(r, dit), &c2ob)
		r = c2ob
	}
	r.EventID = e.Header.Eventid.Text
	// log.Printf("B2CResponse: %v eventid: %v\n", r, r.EventID)
	return r
}

// C2BCallbackEnvelopeToC2BCallback
func (e *C2BCallbackEnvelope) ToResponse() C2BCallback {
	var r, cobj C2BCallback
	r = C2BCallback{}
	for _, dit := range e.Body.GetGenericResult.Request.DataItem {
		json.Unmarshal(toResponse(r, dit), &cobj)
		r = cobj
	}
	// log.Printf("C2BCallback: %v\n", r)
	return r
}

// B2CCallbackEnvelopeToB2CCallback
func (e *B2CCallbackEnvelope) ToResponse() B2CCallback {
	var r, cobj B2CCallback
	r = B2CCallback{}
	for _, dit := range e.Body.GetGenericResult.Request.DataItem {
		json.Unmarshal(toResponse(r, dit), &cobj)
		r = cobj
	}
	// log.Printf("B2CCallback: %v\n", r)
	return r
}
