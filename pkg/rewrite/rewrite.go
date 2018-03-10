package rewrite

import (
	"net/http"
)

// Rule is the simplest redirection rule: for a given URI redirect to a destination URL
type Rule struct {
	httpStatus int    `json:"http_status"`
	source     string `json:"source"`
	target     string `json:"target"`
}

// NewRule creates a new rule from the provided parameters
func NewRule(sourceURI, targetURL string, status int) *Rule {
	return &Rule{
		httpStatus: status,
		source:     sourceURI,
		target:     targetURL,
	}
}

// RewriteHandler creates a rewriter whose purpose is to iterate on a slice of rules
// find a matching URI and redirect to the destination
// Will return 404 if there is no match
func RewriteHandler(rules []*Rule) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, rule := range rules {
			if r.RequestURI == rule.source {
				http.Redirect(w, r, rule.target, rule.httpStatus)
				return
			}
		}
		http.NotFound(w, r)
		return
	})
}
