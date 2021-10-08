package gogemredirect

import (
	"strings"

	h "github.com/Jackd4w/goGEM/pkg/Handler"
)

/*
var AWARDURLS = []string{
	"Attributions",
	"Description",
	"Contribution",
	"Engineering",
	"Collaborations",
	"Human_Practices",
	"Implementation",
	"Model",
	"Proof_Of_Concept",
	"Partnership",
	"Communication",
	"Education",
	"Hardware",
	"Inclusivity",
	"Measurement",
	"Model",
	"Plant",
	"Software",
	"Entrepreneurship",
	"Sustainable",
	"Safety",
}
*/
/*
* Creates redirects from the uppercase addresses defined by iGEM to the "normal" lowercase URLs
 */
func CreateRedirects(urls map[string]string, h *h.Handler) {

	h.Redirect("", "/") // Redirects from https...igem.org/Team:teamname to https...igem.org/Team:teamname/

	for _, url := range urls {
		h.Redirect(url, strings.ToLower(url))
	}
}