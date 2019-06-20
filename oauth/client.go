// Package oauth provides the OAuth APIs
package oauth

import (
  "bytes"
  "net/url"
  stripe "github.com/stripe/stripe-go"
)

// Client is used to invoke /oauth and related APIs.
type Client struct {
  B   stripe.Backend
  Key string
}

func AuthorizeURL(params *stripe.AuthorizeURLParams) string {
  return getC().AuthorizeURL(params)
}

func (c Client) AuthorizeURL(params *stripe.AuthorizeURLParams) string {
  var buf bytes.Buffer
  buf.WriteString(stripe.CONNECTURL)

  if stripe.BoolValue(params.Express) {
    buf.WriteString("/express")
  } else {
    buf.WriteString("/oauth/authorize")
  }

  v := url.Values{
    "client_id":     {stripe.ClientID},
  }

  if stripe.StringValue(params.ResponseType) == "" {
    v.Set("response_type", "code")
  } else {
    v.Set("response_type", *params.ResponseType)
  }

  if stripe.StringValue(params.Scope) != "" {
    v.Set("scope", *params.Scope)
  }
  if stripe.StringValue(params.State) != "" {
    v.Set("state", *params.State)
  }
  if stripe.StringValue(params.RedirectURI) != "" {
    v.Set("redirect_uri", *params.RedirectURI)
  }
  if stripe.StringValue(params.StripeLanding) != "" {
    v.Set("stripe_landing", *params.StripeLanding)
  }
  if stripe.BoolValue(params.AlwaysPrompt) {
    v.Set("always_prompt", "true")
  }

  buf.WriteByte('?')
  buf.WriteString(v.Encode())
  return buf.String()
}

func getC() Client {
  return Client{stripe.GetBackend(stripe.ConnectBackend), stripe.Key}
}
