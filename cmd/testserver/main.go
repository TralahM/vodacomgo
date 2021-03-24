package main

import (
	"bytes"
	"encoding/json"
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
	client  Client
	baseUrl string
}

// ClientFunc is a function tyoe that implements the Client Interface.
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

func (c *APIClient) TestLogin() {
	addr := c.baseUrl + "/login"
	data, _ := json.Marshal(vodacomgo.Elogin)
	resp, _ := DoPost(addr, c.client, bytes.NewReader(data))
	var lresp vodacomgo.LoginResponse
	json.Unmarshal(resp, &lresp)
	log.Printf("Recv Login Response: %v\n", PrettyJson(resp))
	log.Printf("Token: %v\n", lresp.SessionID)
}

func (c *APIClient) TestC2B() {
	addr := c.baseUrl + "/c2b"
	data, _ := json.Marshal(vodacomgo.Ec2b)
	resp, _ := DoPost(addr, c.client, bytes.NewReader(data))
	var lresp vodacomgo.C2BResponse
	json.Unmarshal(resp, &lresp)
	log.Printf("Recv C2B Response: %v\n", PrettyJson(resp))
	log.Printf("ResponseCode: %v\n", lresp.ResponseCode)
}

func (c *APIClient) TestB2C() {
	addr := c.baseUrl + "/b2c"
	data, _ := json.Marshal(vodacomgo.Eb2c)
	resp, _ := DoPost(addr, c.client, bytes.NewReader(data))
	var lresp vodacomgo.B2CResponse
	json.Unmarshal(resp, &lresp)
	log.Printf("Recv B2C Response: %v\n", PrettyJson(resp))
	log.Printf("ResponseCode: %v\n", lresp.ResponseCode)
}

func (c *APIClient) TestB2CCB() {
	addr := c.baseUrl + "/b2c_callback"
	data, _ := json.Marshal(vodacomgo.B2CCallback{
		Amount:              "200",
		ThirdPartyReference: "Somethirdpartypass",
	})
	resp, _ := DoPost(addr, c.client, bytes.NewReader(data))
	var lresp vodacomgo.B2CCallback
	json.Unmarshal(resp, &lresp)
	log.Printf("Recv B2C Callback: %v\n", PrettyJson(resp))
	log.Printf("ThirdPartyReference: %v\n", lresp.ThirdPartyReference)
}
func (c *APIClient) TestC2BCB() {
	addr := c.baseUrl + "/c2b_callback"
	data, _ := json.Marshal(vodacomgo.C2BCallback{
		Amount:              "200",
		ThirdPartyReference: "Somethirdpartypass",
	})
	resp, _ := DoPost(addr, c.client, bytes.NewReader(data))
	var lresp vodacomgo.C2BCallback
	json.Unmarshal(resp, &lresp)
	log.Printf("Recv C2B Callback: %v\n", PrettyJson(resp))
	log.Printf("ThirdPartyReference: %v\n", lresp.ThirdPartyReference)
}

func NewAPIClient(baseUrl string) APIClient {
	client := Decorate(http.DefaultClient,
		Header("Accept", "application/json;charset=utf8"),
	)
	return APIClient{client: client, baseUrl: baseUrl}
}

func PrettyJson(body []byte) string {
	var prettyjson bytes.Buffer
	err := json.Indent(&prettyjson, body, "", "\t")
	if err != nil {
		log.Println("JSON Parse ERror: ", err)
		panic(err)
	}
	return string(prettyjson.Bytes())
}

func main() {
	var baseUrl string = "http://localhost:22080"
	api := NewAPIClient(baseUrl)
	api.TestLogin()
	api.TestC2B()
	api.TestB2C()
	api.TestB2CCB()
	api.TestC2BCB()
}
