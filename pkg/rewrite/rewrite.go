package rewrite

import (
	"net/http"
)

// Rule is the simplest redirection rule: for a given URI redirect to a destination URL
type Rule struct {
	HttpStatus int    `json:"http_status"`
	Source     string `json:"source"`
	Target     string `json:"target"`
}

// NewRule creates a new rule from the provided parameters
func NewRule(sourceURI, targetURL string, status int) *Rule {
	return &Rule{
		HttpStatus: status,
		Source:     sourceURI,
		Target:     targetURL,
	}
}

// BaseHandler creates a rewriter whose purpose is to iterate on a slice of rules
// find a matching URI and redirect to the destination URL.
// Will return 404 if there is no match
func BaseHandler(rules []*Rule) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, rule := range rules {
			if r.RequestURI == rule.Source {
				http.Redirect(w, r, rule.Target, rule.HttpStatus)
				return
			}
		}
		http.NotFound(w, r)
		return
	})
}
