package oauth

// Oauth's you into coursera

import (
	"fmt"
	"net/url"
)

const (
	clientID     = "cBiYd-vLLx-7_BV8AZ1jGQ"
	clientSecret = "A5HoPutMFYTXcySPk4c6sw"
	redirectURI  = "http://christopherbradshaw.me:5001"
	csrfState    = "yoyoyo"
)

// This is the url that we need to redirect incoming users to
func RedirectToCoursera() string {
	baseUrl := "https://accounts.coursera.org/oauth2/v1/auth"

	v := url.Values{}
	v.Set("response_type", "code")
	v.Set("client_id", clientID)
	v.Set("redirect_uri", redirectURI)
	v.Set("scope", "view_profile")
	v.Set("state", csrfState)

	fmt.Println(baseUrl + "?" + v.Encode())
	return baseUrl + "?" + v.Encode()
}

func parseResponse(inUrl string) {
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
}

// Once users are back with us we can take the code they gave us to go to coursera and get a token
func GetToken(code string) string {
	baseUrl := "https://accounts.coursera.org/oauth2/v1/token"

	v := url.Values{}
	v.Set("code", code)
	v.Set("client_id", clientID)
	v.Set("client_secret", clientSecret)
	v.Set("redirect_uri", redirectURI)
	v.Set("grant_type", "authorization_code")

	fmt.Println(baseUrl + "?" + v.Encode())
	return baseUrl + "?" + v.Encode()
	// Post to this url to get the token
}
