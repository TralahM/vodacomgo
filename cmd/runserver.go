package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/tralahm/vodacomgo/client"
	"github.com/tralahm/vodacomgo/server"
)

var (
	CertFile       = os.Getenv("CERT_FILE")
	KeyFile        = os.Getenv("KEY_FILE")
	ServAddr       = getEnv("SERV_ADDR", "8000")
	c2bCallbackUrl = "https://api.betmondenge.com/en/vodacash_c2b_callback"
	b2cCallbackUrl = "https://api.betmondenge.com/en/vodacash_b2c_callback"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	logger := log.New(os.Stdout, "vodacomgo: ", log.LstdFlags|log.Lshortfile)
	r := mux.NewRouter()
	ipgCli := client.NewIPGClient("thirdpartyc2bw", "thirdpartyc2bw", true)
	h := server.NewHandlers(logger, &ipgCli, c2bCallbackUrl, b2cCallbackUrl)
	r.HandleFunc("api/v1/login", h.LoggingMw(h.Login)).Methods("POST", "GET")
	r.HandleFunc("api/v1/c2b", h.LoggingMw(h.C2B)).Methods("POST", "GET")
	r.HandleFunc("api/v1/b2c", h.LoggingMw(h.B2C)).Methods("POST")
	r.HandleFunc("api/v1/c2b_callback", h.LoggingMw(h.C2BCallback)).Methods("POST")
	r.HandleFunc("api/v1/b2c_callback", h.LoggingMw(h.B2CCallback)).Methods("POST")

	// handlers.SetupRoutes(mux)
	logger.Println("Server starting ...")
	// srv := server.New(r, ServAddr)
	err := http.ListenAndServe("0.0.0.0:"+ServAddr, r)
	if err != nil {
		logger.Fatalf("Server failed to start: %v", err)
	}
}
