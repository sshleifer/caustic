package main

import (
	"github.com/Christopher-Bradshaw/caustic/oauth"
)

func main() {
	oauth.RedirectToCoursera()
	oauth.GetToken("blah")
}
