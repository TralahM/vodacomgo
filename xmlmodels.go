package vodacomgo

import "encoding/xml"

type LoginEnvelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	S       string   `xml:"S,attr"`
	SOAPENV string   `xml:"SOAP-ENV,attr"`
	Header  struct {
		Text    string `xml:",chardata"`
		Eventid struct {
			Text string `xml:",chardata"`
			Ns2  string `xml:"ns2,attr"`
			Ns3  string `xml:"ns3,attr"`
		} `xml:"eventid"`
	} `xml:"Header"`
	Body struct {
		Text                     string `xml:",chardata"`
		GetGenericResultResponse struct {
			Text          string `xml:",chardata"`
			Ns2           string `xml:"ns2,attr"`
			Ns3           string `xml:"ns3,attr"`
			SOAPAPIResult struct {
				Text      string `xml:",chardata"`
				EventInfo struct {
					Text          string `xml:",chardata"`
					Code          string `xml:"code"`
					Description   string `xml:"description"`
					Detail        string `xml:"detail"`
					TransactionID string `xml:"transactionID"`
				} `xml:"eventInfo"`
				Request struct {
					Text     string `xml:",chardata"`
					DataItem []struct {
						Text  string `xml:",chardata"`
						Name  string `xml:"name"`
						Type  string `xml:"type"`
						Value string `xml:"value"`
					} `xml:"dataItem"`
				} `xml:"request"`
				Response struct {
					Text     string `xml:",chardata"`
					DataItem struct {
						Text  string `xml:",chardata"`
						Name  string `xml:"name"`
						Type  string `xml:"type"`
						Value string `xml:"value"`
					} `xml:"dataItem"`
				} `xml:"response"`
			} `xml:"SOAPAPIResult"`
		} `xml:"getGenericResultResponse"`
	} `xml:"Body"`
}

type C2BEnvelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	S       string   `xml:"S,attr"`
	SOAPENV string   `xml:"SOAP-ENV,attr"`
	Header  struct {
		Text    string `xml:",chardata"`
		Eventid struct {
			Text string `xml:",chardata"`
			Ns2  string `xml:"ns2,attr"`
			Ns3  string `xml:"ns3,attr"`
		} `xml:"eventid"`
	} `xml:"Header"`
	Body struct {
		Text                     string `xml:",chardata"`
		GetGenericResultResponse struct {
			Text          string `xml:",chardata"`
			Ns2           string `xml:"ns2,attr"`
			Ns3           string `xml:"ns3,attr"`
			SOAPAPIResult struct {
				Text      string `xml:",chardata"`
				EventInfo struct {
					Text          string `xml:",chardata"`
					Code          string `xml:"code"`
					Description   string `xml:"description"`
					Detail        string `xml:"detail"`
					TransactionID string `xml:"transactionID"`
				} `xml:"eventInfo"`
				Request struct {
					Text     string `xml:",chardata"`
					DataItem []struct {
						Text  string `xml:",chardata"`
						Name  string `xml:"name"`
						Type  string `xml:"type"`
						Value string `xml:"value"`
					} `xml:"dataItem"`
				} `xml:"request"`
				Response struct {
					Text     string `xml:",chardata"`
					DataItem []struct {
						Text  string `xml:",chardata"`
						Name  string `xml:"name"`
						Type  string `xml:"type"`
						Value string `xml:"value"`
					} `xml:"dataItem"`
				} `xml:"response"`
			} `xml:"SOAPAPIResult"`
		} `xml:"getGenericResultResponse"`
	} `xml:"Body"`
}

type B2CEnvelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	S       string   `xml:"S,attr"`
	SOAPENV string   `xml:"SOAP-ENV,attr"`
	Header  struct {
		Text    string `xml:",chardata"`
		Eventid struct {
			Text string `xml:",chardata"`
			Ns2  string `xml:"ns2,attr"`
			Ns3  string `xml:"ns3,attr"`
		} `xml:"eventid"`
	} `xml:"Header"`
	Body struct {
		Text                     string `xml:",chardata"`
		GetGenericResultResponse struct {
			Text          string `xml:",chardata"`
			Ns2           string `xml:"ns2,attr"`
			Ns3           string `xml:"ns3,attr"`
			SOAPAPIResult struct {
				Text      string `xml:",chardata"`
				EventInfo struct {
					Text          string `xml:",chardata"`
					Code          string `xml:"code"`
					Description   string `xml:"description"`
					Detail        string `xml:"detail"`
					TransactionID string `xml:"transactionID"`
				} `xml:"eventInfo"`
				Request struct {
					Text     string `xml:",chardata"`
					DataItem []struct {
						Text  string `xml:",chardata"`
						Name  string `xml:"name"`
						Type  string `xml:"type"`
						Value string `xml:"value"`
					} `xml:"dataItem"`
				} `xml:"request"`
				Response struct {
					Text     string `xml:",chardata"`
					DataItem []struct {
						Text  string `xml:",chardata"`
						Name  string `xml:"name"`
						Type  string `xml:"type"`
						Value string `xml:"value"`
					} `xml:"dataItem"`
				} `xml:"response"`
			} `xml:"SOAPAPIResult"`
		} `xml:"getGenericResultResponse"`
	} `xml:"Body"`
}

type B2CCallbackEnvelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Soapenv string   `xml:"soapenv,attr"`
	Soap    string   `xml:"soap,attr"`
	Gen     string   `xml:"gen,attr"`
	Header  struct {
		Text    string `xml:",chardata"`
		EventID string `xml:"EventID"`
	} `xml:"Header"`
	Body struct {
		Text             string `xml:",chardata"`
		GetGenericResult struct {
			Text    string `xml:",chardata"`
			Request struct {
				Text     string `xml:",chardata"`
				DataItem []struct {
					Text  string `xml:",chardata"`
					Name  string `xml:"name"`
					Value string `xml:"value"`
					Type  string `xml:"type"`
				} `xml:"dataItem"`
			} `xml:"Request"`
		} `xml:"getGenericResult"`
	} `xml:"Body"`
}

type C2BCallbackEnvelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Soapenv string   `xml:"soapenv,attr"`
	Soap    string   `xml:"soap,attr"`
	Gen     string   `xml:"gen,attr"`
	Body    struct {
		Text             string `xml:",chardata"`
		GetGenericResult struct {
			Text    string `xml:",chardata"`
			Request struct {
				Text     string `xml:",chardata"`
				DataItem []struct {
					Text  string `xml:",chardata"`
					Name  string `xml:"name"`
					Value string `xml:"value"`
					Type  string `xml:"type"`
				} `xml:"dataItem"`
			} `xml:"Request"`
		} `xml:"getGenericResult"`
	} `xml:"Body"`
}
