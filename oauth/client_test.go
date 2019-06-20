package oauth

import (
  "testing"
  assert "github.com/stretchr/testify/require"
  stripe "github.com/stripe/stripe-go"
  _ "github.com/stripe/stripe-go/testing"
)

func TestAuthorizeURL(t *testing.T) {
  stripe.ClientID = "ca_123"
  url := AuthorizeURL(AuthorizeURLParams{
    Scope:       "read_write",
    RedirectURI: "https://t.example.com",
  })

  assert.Contains(t, url, "https://connect.stripe.com/oauth/authorize?")
  assert.Contains(t, url, "client_id=ca_123")
  assert.Contains(t, url, "redirect_uri=https%3A%2F%2Ft.example.com")
  assert.Contains(t, url, "response_type=code")
  assert.Contains(t, url, "scope=read_write")
}

func TestAuthorizeURLWithOptionalArgs(t *testing.T) {
  stripe.ClientID = "ca_123"
  url := AuthorizeURL(AuthorizeURLParams{
    State:          "test-state",
    Scope:          "read_only",
    RedirectURI:    "https://t.example.com",
    ResponseType:   "test-code",
    StripeLanding:  "register",
  })

  assert.Contains(t, url, "https://connect.stripe.com/oauth/authorize?")
  assert.Contains(t, url, "client_id=ca_123")
  assert.Contains(t, url, "redirect_uri=https%3A%2F%2Ft.example.com")
  assert.Contains(t, url, "response_type=test-code")
  assert.Contains(t, url, "scope=read_only")
  assert.Contains(t, url, "state=test-state")
  assert.Contains(t, url, "stripe_landing=register")
}
