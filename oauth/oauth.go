package oauth

// Oauth's you into coursera

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	clientID     = "cBiYd-vLLx-7_BV8AZ1jGQ"
	clientSecret = "A5HoPutMFYTXcySPk4c6sw"
	redirectURI  = "http://christopherbradshaw.me:5001/back-at-you"
	csrfState    = "yoyoyo"
)

// This is the url that we need to redirect incoming users to
func RedirectToCoursera(w http.ResponseWriter, r *http.Request) {
	baseUrl := "https://accounts.coursera.org/oauth2/v1/auth"

	v := url.Values{}
	v.Set("response_type", "code")
	v.Set("client_id", clientID)
	v.Set("redirect_uri", redirectURI)
	v.Set("scope", "view_profile")
	v.Set("state", csrfState)

	fullURL := baseUrl + "?" + v.Encode()
	fmt.Println(fullURL)
	http.Redirect(w, r, fullURL, 301)
}

// Once users are back with us we can take the code they gave us to go to coursera and get a token
func RedeemCode(w http.ResponseWriter, r *http.Request) {
	baseUrl := "https://accounts.coursera.org/oauth2/v1/token"

	data := url.Values{}
	data.Set("code", getCodeFromUrl(r.URL.String()))
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)
	data.Set("redirect_uri", redirectURI)
	data.Set("grant_type", "authorization_code")

	resp, err := http.Post(baseUrl, "application/x-www-form-urlencoded", bytes.NewBufferString(data.Encode()))
	if err != nil {
		panic(fmt.Sprintf("err in parsing code response: %s", err.Error()))
	}
	fmt.Printf("%#v\n", resp.Body)

	body := map[string]interface{}{}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		panic(fmt.Sprintf("Could not unmarshal body: %s", err.Error()))
	}
	fmt.Printf("%#v\n", body)
	http.Redirect(w, r, fmt.Sprintf("http://christopherbradshaw.me:5001/with-token?token=%s", body["access_token"]), 301)
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Hello\n")
	w.Write([]byte("Hello!"))
}

func getCodeFromUrl(inUrl string) string {
	parsedUrl, err := url.Parse(inUrl)

	if err != nil {
		panic(fmt.Sprintf("err in parsing response: %s", err.Error()))
	}

	v, err := url.ParseQuery(parsedUrl.RawQuery)

	if err != nil {
		panic(fmt.Sprintf("err in parsing query: %s", err.Error()))
	}
	if v["state"][0] != csrfState {
		panic(fmt.Sprintf("csrf state yo!: %s", v["state"]))
	}

	if len(v["code"]) != 1 {
		panic(fmt.Sprintf("no code yo!: %s", v["code"]))
	}
	return v["code"][0]
}
