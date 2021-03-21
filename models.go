package vodacomgo

type Login struct {
	Username string
	Password string
}

type LoginResponse struct {
	Code          string `json:"code"`
	Description   string `json:"description"`
	Detail        string `json:"detail"`
	TransactionID string `json:"transactionID"`
	EventID       string `json:"event_id"`
	Username      string `json:"Username"`
	Password      string `json:"Password"`
	SessionID     string `json:"SessionID"`
}

type C2B struct {
	Token               string
	CustomerMSISDN      string
	ServiceProviderCode string
	Currency            string
	Amount              string
	Date                string
	ThirdPartyReference string
	CommandID           string
	Language            string
	CallBackChannel     string
	CallBackDestination string
	Surname             string
	Initials            string
}
type C2BResponse struct {
	Amount              string `json:"Amount"`
	CallBackChannel     string `json:"CallBackChannel"`
	CallBackDestination string `json:"CallBackDestination"`
	Code                string `json:"code"`
	CommandID           string `json:"CommandId"`
	Currency            string `json:"Currency"`
	CustomerMSISDN      string `json:"CustomerMSISDN"`
	Date                string `json:"Date"`
	Description         string `json:"description"`
	Detail              string `json:"detail"`
	EventID             string `json:"event_id"`
	Initials            string `json:"Initials"`
	InsightReference    string `json:"InsightReference"`
	Language            string `json:"Language"`
	ResponseCode        string `json:"ResponseCode"`
	ServiceProviderCode string `json:"ServiceProviderCode"`
	Surname             string `json:"Surname"`
	ThirdPartyReference string `json:"ThirdPartyReference"`
	TransactionID       string `json:"transactionID"`
}


type C2BCallback struct {
	Amount                   string `json:"Amount"`
	ConversationID           string `json:"ConversationID"`
	InsightReference         string `json:"InsightReference"`
	OriginatorConversationID string `json:"OriginatorConversationID"`
	ResultCode               string `json:"ResultCode"`
	ResultDesc               string `json:"ResultDesc"`
	ResultType               string `json:"ResultType"`
	ThirdPartyReference      string `json:"ThirdPartyReference"`
	TransactionID            string `json:"TransactionID"`
	TransactionTime          string `json:"TransactionTime"`
}

type B2C struct {
	Token               string
	ServiceProviderName string
	CustomerMSISDN      string
	Currency            string
	Amount              string
	TransactionDateTime string
	Shortcode           string
	Language            string
	ThirdPartyReference string
	CallBackChannel     string
	CallBackDestination string
	CommandID           string
}

type B2CResponse struct {
	Amount              string `json:"Amount"`
	CallBackChannel     string `json:"CallBackChannel"`
	CallBackDestination string `json:"CallBackDestination"`
	Code                string `json:"code"`
	CommandID           string `json:"CommandID"`
	Currency            string `json:"Currency"`
	CustomerMSISDN      string `json:"CustomerMSISDN"`
	Description         string `json:"description"`
	Detail              string `json:"detail"`
	EventID             string `json:"event_id"`
	InsightReference    string `json:"InsightReference"`
	Language            string `json:"Language"`
	ResponseCode        string `json:"ResponseCode"`
	ServiceProviderName string `json:"ServiceProviderName"`
	Shortcode           string `json:"Shortcode"`
	ThirdPartyReference string `json:"ThirdPartyReference"`
	TransactionDateTime string `json:"TransactionDateTime"`
	TransactionID       string `json:"transactionID"`
}

type B2CCallback struct {
	Amount                   string `json:"Amount"`
	ConversationID           string `json:"ConversationId"`
	InsightReference         string `json:"InsightReference"`
	OriginatorConversationID string `json:"OriginatorConversationId"`
	ResultCode               string `json:"ResultCode"`
	ResultDesc               string `json:"ResultDesc"`
	ResultType               string `json:"ResultType"`
	ThirdPartyReference      string `json:"ThirdPartyReference"`
	TransactionID            string `json:"TransactionID"`
	TransactionTime          string `json:"TransactionTime"`
}
