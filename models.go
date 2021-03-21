package vodacomgo

type Login struct {
	Username string
	Password string
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
