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
	CertFile = os.Getenv("CERT_FILE")
	KeyFile  = os.Getenv("KEY_FILE")
	ServAddr = os.Getenv("SERV_ADDR")
)

func main() {
	logger := log.New(os.Stdout, "vodacomgo: ", log.LstdFlags|log.Lshortfile)
	r := mux.NewRouter()
	ipgCli := client.NewIPGClient("thirdpartyc2bw", "thirdpartyc2bw", true)
	h := server.NewHandlers(logger, &ipgCli)
	r.HandleFunc("/login", h.LoggingMw(h.Login)).Methods("POST", "GET")
	r.HandleFunc("/c2b", h.LoggingMw(h.C2B)).Methods("POST", "GET")
	r.HandleFunc("/b2c", h.LoggingMw(h.B2C)).Methods("POST")
	r.HandleFunc("/c2b_callback", h.LoggingMw(h.C2BCallback)).Methods("POST")
	r.HandleFunc("/b2c_callback", h.LoggingMw(h.B2CCallback)).Methods("POST")

	// handlers.SetupRoutes(mux)
	logger.Println("Server starting ...")
	// srv := server.New(r, ServAddr)
	err := http.ListenAndServe(ServAddr, r)
	if err != nil {
		logger.Fatalf("Server failed to start: %v", err)
	}
}
