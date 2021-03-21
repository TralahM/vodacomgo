package vodacomgo

var Elogin = Login{
	Username: "thirdpartyc2bw",
	Password: "thirdpartyc2bw",
}

var Ec2b = C2B{
	Token:               "d790ccddc4494976b287d9cfb054ce8d",
	CustomerMSISDN:      "243811222333",
	ServiceProviderCode: "80032",
	Currency:            "CDF",
	Amount:              "328",
	Date:                "20201231",
	ThirdPartyReference: "R20200829123483",
	CommandID:           "InitTrans4allc2b",
	Language:            "EN",
	CallBackChannel:     "2",
	CallBackDestination: "124.39.122.23/api/c2b_callback",
	Surname:             "Surname",
	Initials:            "Initials",
}

var Eb2c = B2C{
	Token:               "d790ccddc4494976b287d9cfb054ce8d",
	CustomerMSISDN:      "243811222333",
	ServiceProviderName: "BMB GRV",
	Shortcode:           "80032",
	Currency:            "CDF",
	Amount:              "328",
	TransactionDateTime: "20201231",
	ThirdPartyReference: "R20200829123483",
	CommandID:           "InitTransForallB2C",
	Language:            "EN",
	CallBackChannel:     "2",
	CallBackDestination: "124.39.122.23/api/b2c_callback",
}
