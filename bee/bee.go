package main

import (
	"fmt"
	"net/http"
	"net/url"
	//"time"
)

const (
	authToken = "Aj68CNeg_CpJ5dUYniGp"
	userName  = "sshleifer"
)

// post to beeminder.com/api/v1/users/sshleifer/goals/coursera/datapoints.json when someone does something good
// TODO: can also leave a comment, and a timestamp
func postAccomplishment(value string, user string, goal string) {
	extraStuff := fmt.Sprintf("%v/goals/%v/datapoints.json", user, goal)
	baseUrl := fmt.Sprintf("https://www.beeminder.com/api/v1/users/%v", extraStuff)
	v := url.Values{}
	v.Set("auth_token", authToken)
	//v.Set("timestamp", time.Now())
	v.Set("comment", "from GO") // also have sendmail option
	v.Set("value", value)
	// fmt.Println(time.Now().Format(time.RFC850))
	fullURL := baseUrl + "?" + v.Encode()
	fmt.Println(fullURL)
	resp, err := http.Post(fullURL, "blah", nil)
	if err != nil {
		panic(fmt.Sprintf("err in parsing code response: %s", err.Error()))
	}
	fmt.Printf("%#v\n", resp)
}

func main() {
	postAccomplishment("2", "sshleifer", "coursera")
}
