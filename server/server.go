package server

import (
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/tralahm/vodacomgo"
	"github.com/tralahm/vodacomgo/client"
)

type Handlers struct {
	logger         *log.Logger
	ipgClient      *client.APIClient
	c2bCallbackUrl string
	b2cCallbackUrl string
}

func (h *Handlers) C2B(w http.ResponseWriter, r *http.Request) {
	// read request body bytes
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.logger.Printf("Error reading body: %v", err)
		http.Error(w, "cant read body", http.StatusBadRequest)
		return
	}
	var c2b vodacomgo.C2B
	// Unmarshal/deserialize request json body into domain request struct
	err = json.Unmarshal(body, &c2b)
	if err != nil {
		h.logger.Printf("Error reading body: %v", err)
		http.Error(w, "cant read body", http.StatusBadRequest)
		return
	}
	// Here We would call the method to generate a byte slice to be sent as
	// request body to remote IPG which returns a response and
	if h.ipgClient.Sandbox {
		c2bres := vodacomgo.DecodeC2BResponse([]byte(vodacomgo.C2BRx))
		// then send the request to the remote IPG url and receive a responsebodybytes

		// decode response from xml bytes response from IPG
		// transform the decoded response to domain level struct representation
		c2btf := c2bres.ToResponse()
		// Marshall/serialize domain response to json byte string
		jsonstr, _ := json.Marshal(c2btf)
		// write headers
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w = ServerHeader(w)
		w.WriteHeader(http.StatusCreated)
		// write response body
		w.Write(jsonstr)

	} else {
		responseB, err := (h.ipgClient).IpgC2B(c2b)
		if err != nil {
			h.logger.Printf("Error making IPG request : %v", err)
			http.Error(w, "Error making IPG request", http.StatusServiceUnavailable)
			return
		}
		c2bres := vodacomgo.DecodeC2BResponse(responseB)
		// then send the request to the remote IPG url and receive a responsebodybytes

		// decode response from xml bytes response from IPG
		// transform the decoded response to domain level struct representation
		c2btf := c2bres.ToResponse()
		// Marshall/serialize domain response to json byte string
		jsonstr, _ := json.Marshal(c2btf)
		// write headers
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w = ServerHeader(w)
		w.WriteHeader(http.StatusCreated)
		// write response body
		w.Write(jsonstr)

	}

}

func (h *Handlers) B2C(w http.ResponseWriter, r *http.Request) {
	// read request body bytes
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.logger.Printf("Error reading body: %v", err)
		http.Error(w, "cant read body", http.StatusBadRequest)
		return
	}
	var b2c vodacomgo.B2C
	// Unmarshal/deserialize request json body into domain request struct
	err = json.Unmarshal(body, &b2c)
	if err != nil {
		h.logger.Printf("Error reading body: %v", err)
		http.Error(w, "cant read body", http.StatusBadRequest)
		return
	}
	if h.ipgClient.Sandbox {
		b2cres := vodacomgo.DecodeB2CResponse([]byte(vodacomgo.B2CRx))
		b2ctf := b2cres.ToResponse()
		jsonstr, _ := json.Marshal(b2ctf)
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w = ServerHeader(w)
		w.WriteHeader(http.StatusCreated)
		w.Write(jsonstr)

	} else {
		responseB, err := h.ipgClient.IpgB2C(b2c)
		if err != nil {
			h.logger.Printf("Error making IPG request : %v", err)
			http.Error(w, "Error making IPG request", http.StatusServiceUnavailable)
			return
		}
		b2cres := vodacomgo.DecodeB2CResponse(responseB)
		b2ctf := b2cres.ToResponse()
		jsonstr, _ := json.Marshal(b2ctf)
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w = ServerHeader(w)
		w.WriteHeader(http.StatusCreated)
		w.Write(jsonstr)
	}

}

func (h *Handlers) B2CCallback(w http.ResponseWriter, r *http.Request) {
	// read request body bytes
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.logger.Printf("Error reading body: %v", err)
		http.Error(w, "cant read body", http.StatusBadRequest)
		return
	}
	if h.ipgClient.Sandbox {
		callbackb2cres := vodacomgo.DecodeB2CCallback([]byte(vodacomgo.B2CCRx))
		x, _ := xml.MarshalIndent(callbackb2cres, " ", "  ")
		log.Println(string([]byte(x)))
		callbackb2ctf := callbackb2cres.ToResponse()
		jsonstr, _ := json.Marshal(callbackb2ctf)
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w = ServerHeader(w)
		w.WriteHeader(http.StatusCreated)
		w.Write(jsonstr)
	} else {
		var callbackb2c vodacomgo.B2CCallbackEnvelope
		err = xml.Unmarshal(body, &callbackb2c)
		// x, err := xml.MarshalIndent(callbackb2c, "", "\t")
		// log.Println(string(x))
		if err != nil {
			h.logger.Printf("Error reading body: %v", err)
			http.Error(w, "cant read body", http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/xml")
		w = ServerHeader(w)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(vodacomgo.AckB2CT))
	}

}

func (h *Handlers) C2BCallback(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.logger.Printf("Error reading body: %v", err)
		http.Error(w, "cant read body", http.StatusBadRequest)
		return
	}
	if h.ipgClient.Sandbox {
		callbackc2bres := vodacomgo.DecodeC2BCallback([]byte(vodacomgo.C2BCRx))
		x, _ := xml.MarshalIndent(callbackc2bres, " ", "  ")
		log.Println(string([]byte(x)))
		callbackc2btf := callbackc2bres.ToResponse()
		jsonstr, _ := json.Marshal(callbackc2btf)
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w = ServerHeader(w)
		w.WriteHeader(http.StatusCreated)
		w.Write(jsonstr)
	} else {
		var callbackc2b vodacomgo.C2BCallbackEnvelope
		err = xml.Unmarshal(body, &callbackc2b)
		if err != nil {
			h.logger.Printf("Error reading body: %v", err)
			http.Error(w, "cant read body", http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/xml")
		w = ServerHeader(w)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(vodacomgo.AckC2BT))
	}
}

func (h *Handlers) Login(w http.ResponseWriter, r *http.Request) {
	// read request body bytes
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.logger.Printf("Error reading body: %v", err)
		http.Error(w, "cant read body", http.StatusBadRequest)
		return
	}
	var login vodacomgo.Login
	// Unmarshal/deserialize request json body into domain request struct
	err = json.Unmarshal(body, &login)
	if err != nil {
		h.logger.Printf("Error reading body: %v", err)
		http.Error(w, "cant read body", http.StatusBadRequest)
		return
	}
	if h.ipgClient.Sandbox {
		loginres := vodacomgo.DecodeLoginResponse([]byte(vodacomgo.LoginRx))
		logintf := loginres.ToResponse()
		jsonstr, _ := json.Marshal(logintf)
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w = ServerHeader(w)
		w.WriteHeader(http.StatusCreated)
		w.Write(jsonstr)
	} else {
		h.ipgClient.Password = login.Password
		h.ipgClient.Username = login.Username
		responseB, err := h.ipgClient.IpgLogin()
		if err != nil {
			h.logger.Printf("Error making IPG request : %v", err)
			http.Error(w, "Error making IPG request", http.StatusServiceUnavailable)
			return
		}
		loginres := vodacomgo.DecodeLoginResponse(responseB)
		logintf := loginres.ToResponse()
		// Marshall/serialize domain response to json byte string
		jsonstr, _ := json.Marshal(logintf)
		// write headers
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w = ServerHeader(w)
		w.WriteHeader(http.StatusCreated)
		// Write response body
		w.Write(jsonstr)
	}
}

func NewHandlers(logger *log.Logger, ipgClient *client.APIClient, c2bCallbackUrl, b2cCallbackUrl string) *Handlers {
	return &Handlers{
		logger:         logger,
		ipgClient:      ipgClient,
		c2bCallbackUrl: c2bCallbackUrl,
		b2cCallbackUrl: b2cCallbackUrl,
	}
}

func (h *Handlers) LoggingMw(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer h.logger.Printf("Request Processed in %s .", time.Now().Sub(startTime))
		next(w, r)
	}
}

// func (h *Handlers) SetupRoutes(mux *http.ServeMux) {
// 	mux.HandleFunc("/login", h.Login)
// }

func New(mux *http.ServeMux, serverAddress string) *http.Server {
	tlsConfig := &tls.Config{
		// Causes servers to use Go's default ciphersuite preferences,
		// which are tuned to avoid attacks. Does nothing on clients.
		// PreferServerCipherSuites: true,
		// Only use curves which have assembly implementations
		CurvePreferences: []tls.CurveID{
			tls.CurveP256,
			tls.X25519, // Go 1.8 only
		},
		MinVersion: tls.VersionTLS12,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305, // Go 1.8 onl y
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,   // Go 1.8 onl y
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			// Best disabled, as they don't provide Forward Secrecy,
			// but might be necessary for some clients
			// tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			// tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
		},
	}
	srv := &http.Server{
		Addr:         serverAddress,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux,
		TLSConfig:    tlsConfig,
	}
	return srv
}

type Decorator func(http.ResponseWriter) http.ResponseWriter

// Decorate decorates a Client c with all the given Decorators in order.
func Decorate(c http.ResponseWriter, ds ...Decorator) http.ResponseWriter {
	decorated := c

	for _, decorate := range ds {
		decorated = decorate(decorated)
	}
	return decorated
}

func ServerHeader(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Add("Server", "VodacomGo")
	return w
}
