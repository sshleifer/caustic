package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sshleifer/caustic/oauth"
)

func main() {
	r := mux.NewRouter()
	// Redirect to coursera
	r.Methods("GET").Path("/coursera-login").HandlerFunc(oauth.RedirectToCoursera)
	// URL that coursera redirects to
	r.Methods("GET").Path("/back-at-you").HandlerFunc(oauth.RedeemCode)
	log.Panicf("Server crashed, %s", http.ListenAndServe(":5001", r))
}
