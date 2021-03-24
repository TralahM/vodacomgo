package client

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/tralahm/vodacomgo"
)

// A Client sends http.Requests and returns http.Responses or errors in case of
// failure.

type Client interface {
	Do(*http.Request) (*http.Response, error)
}

type APIClient struct {
	client   Client
	BaseUrl  string
	Username string
	Password string
	Token    string
	Sandbox  bool
}

// ClientFunc is a function type that implements the Client Interface.
type ClientFunc func(*http.Request) (*http.Response, error)

func (f ClientFunc) Do(r *http.Request) (*http.Response, error) {
	return f(r)
}

// A Decorator wraps a Client with extra behaviour

type Decorator func(Client) Client

// Decorate decorates a Client c with all the given Decorators in order.
func Decorate(c Client, ds ...Decorator) Client {
	decorated := c

	for _, decorate := range ds {
		decorated = decorate(decorated)
	}
	return decorated
}

func Logging(l *log.Logger) Decorator {
	return func(c Client) Client {
		return ClientFunc(func(r *http.Request) (*http.Response, error) {
			l.Printf("%s: %s %s", r.UserAgent(), r.Method, r.URL)
			return c.Do(r)
		})
	}
}

// Authorization returns a Decorator that authorizes every Client request with
// the given token.
func Authorization(token string) Decorator {
	return Header("Authorization", token)
}

// Header returns a Decorator that adds the given HTTP header to every request
// done by a Client
func Header(name, value string) Decorator {
	return func(c Client) Client {
		return ClientFunc(func(r *http.Request) (*http.Response, error) {
			r.Header.Add(name, value)
			return c.Do(r)
		})
	}
}

func DoPost(addr string, client Client, data io.Reader) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPost, addr, data)

	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer res.Body.Close()
	log.Println("statusCode: ", res.Status)
	responseBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error reading HTTP response: %s", err)
	}
	return responseBytes, err
}

func (api *APIClient) IpgLogin() ([]byte, error) {
	endpoint := ":8091/insight/SOAPIn"
	addr := api.BaseUrl + endpoint
	rwr := new(bytes.Buffer)
	if api.Username == "" || api.Password == "" {
		panic("username and password must be provided in NewAPIClient.")
	}
	login := vodacomgo.Login{Username: api.Username, Password: api.Password}
	vodacomgo.GenLogin(rwr, login)
	responseBytes, err := DoPost(addr, api.client, rwr)
	lrep := vodacomgo.DecodeLoginResponse(responseBytes)
	psd := lrep.ToResponse()
	if err != nil {
		log.Fatalf("Got %v\n", err)
	}
	if err == nil {
		api.Token = psd.SessionID
	}
	return responseBytes, err
}

func (api *APIClient) IpgC2B(c2b vodacomgo.C2B) ([]byte, error) {
	endpoint := ":8091/insight/SOAPIn"
	if api.Token == "" {
		api.IpgLogin()
	}
	addr := api.BaseUrl + endpoint
	rwr := new(bytes.Buffer)
	if c2b.Token == "" {
		c2b.Token = api.Token
	}
	vodacomgo.GenC2B(rwr, c2b)
	responseBytes, err := DoPost(addr, api.client, rwr)
	return responseBytes, err
}
func (api *APIClient) IpgB2C(b2c vodacomgo.B2C) ([]byte, error) {
	endpoint := ":8094/iPG/B2C"
	if api.Token == "" {
		api.IpgLogin()
	}
	addr := api.BaseUrl + endpoint
	rwr := new(bytes.Buffer)
	if b2c.Token == "" {
		b2c.Token = api.Token
	}
	vodacomgo.GenB2C(rwr, b2c)
	responseBytes, err := DoPost(addr, api.client, rwr)
	return responseBytes, err
}

func NewIPGClient(username, password string, sandbox bool) APIClient {
	var baseUrl string
	sandboxBase := "https://uatipg.m-pesa.vodacom.cd"
	productionBase := "https://ipg.m-pesa.vodacom.cd"
	client := Decorate(http.DefaultClient,
		Header("Accept", "application/xml,text/xml"),
	)
	if sandbox {
		baseUrl = sandboxBase
	} else {
		baseUrl = productionBase
	}
	return APIClient{client: client, BaseUrl: baseUrl, Username: username, Password: password, Sandbox: sandbox}
}
