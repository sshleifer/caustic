package main

import (
	"github.com/sshleifer/caustic/oauth"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	// Redirect to coursera
	r.Methods("GET").Path("/coursera-login").HandlerFunc(oauth.RedirectToCoursera)
	// URL that coursera redirects to
	r.Methods("GET").Path("/back-at-you").HandlerFunc(oauth.RedeemCode)
}
