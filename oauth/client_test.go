package oauth

import (
  "testing"
  assert "github.com/stretchr/testify/require"
  stripe "github.com/stripe/stripe-go"
  _ "github.com/stripe/stripe-go/testing"
)

func TestAuthorizeURL(t *testing.T) {
  stripe.ClientID = "ca_123"
  url := AuthorizeURL(&stripe.AuthorizeURLParams{
  })

  assert.Contains(t, url, "https://connect.stripe.com/oauth/authorize?")
  assert.Contains(t, url, "client_id=ca_123")
  assert.Contains(t, url, "response_type=code")
}

func TestAuthorizeURLWithOptionalArgs(t *testing.T) {
  stripe.ClientID = "ca_123"
  url := AuthorizeURL(&stripe.AuthorizeURLParams{
    State:          stripe.String("test-state"),
    Scope:          stripe.String("read_only"),
    RedirectURI:    stripe.String("https://t.example.com"),
    ResponseType:   stripe.String("test-code"),
    StripeLanding:  stripe.String("register"),
    AlwaysPrompt:   stripe.Bool(true),
    Express:        stripe.Bool(true),
  })

  assert.Contains(t, url, "https://connect.stripe.com/express?")
  assert.Contains(t, url, "client_id=ca_123")
  assert.Contains(t, url, "redirect_uri=https%3A%2F%2Ft.example.com")
  assert.Contains(t, url, "response_type=test-code")
  assert.Contains(t, url, "scope=read_only")
  assert.Contains(t, url, "state=test-state")
  assert.Contains(t, url, "stripe_landing=register")
  assert.Contains(t, url, "always_prompt=true")
}
