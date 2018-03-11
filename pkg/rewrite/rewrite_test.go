package rewrite_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "git.bravofly.com/fgualazzi/rewrite-engine.git/pkg/rewrite"
	"github.com/stretchr/testify/assert"
)

const testStatusOK = 288

var fakeBackend = func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(testStatusOK)
}

func TestRewriteRuleSuccessfulRedirection(t *testing.T) {
	// create a fake beackend
	fakeServer := httptest.NewServer(http.HandlerFunc(fakeBackend))
	defer fakeServer.Close()
	// set a rule to test
	rule := NewRule("/test", fakeServer.URL, http.StatusMovedPermanently)
	rewriter := BaseHandler([]*Rule{rule})
	server := httptest.NewServer(rewriter)
	defer server.Close()
	// perform the test request and assert
	response, err := http.Get(server.URL + "/test")
	assert.Nil(t, err)
	assert.Equal(t, testStatusOK, response.StatusCode)
}

func TestRewriteRuleNotFound(t *testing.T) {
	rule := NewRule("/test", "/never-reached", http.StatusMovedPermanently)
	rewriter := BaseHandler([]*Rule{rule})
	server := httptest.NewServer(rewriter)
	defer server.Close()

	response, err := http.Get(server.URL + "/error-not-found")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNotFound, response.StatusCode)
}
