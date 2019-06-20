// Package oauth provides the OAuth APIs
package oauth

import (
  "bytes"
  "net/url"
  // "net/http"
  //
  stripe "github.com/stripe/stripe-go"
  // "github.com/stripe/stripe-go/form"
)

// Client is used to invoke /oauth and related APIs.
type Client struct {
  B   stripe.Backend
  Key string
}

// ResponseTypeType is the list of allowed values for response_type.
type ResponseTypeType string

// List of values that ResponseTypeType can take.
const (
  ResponseTypeTypeCode  ResponseTypeType = "code"
)

// ScopeType is the list of OAuth scopes supported.
type ScopeType string

// List of values that ScopeType can take.
const (
  ScopeTypeReadOnly   ScopeType = "read_only"
  ScopeTypeReadWrite  ScopeType = "read_write"
)

// StripeLandingType is the list of allowed values for stripe_landing.
type StripeLandingType string

// List of allowed values for StripeLandingType's.
const (
  StripeLandingTypeLogin    StripeLandingType = "login"
  StripeLandingTypeRegister StripeLandingType = "register"
)

// Params for creating OAuth AuthorizeURL's.
type AuthorizeURLParams struct {
  State           string
  ResponseType    ResponseTypeType
  Scope           ScopeType
  RedirectURI     string
  StripeLanding   StripeLandingType
  AlwaysPrompt    bool

  // TODO?
  // Express         bool
}

func AuthorizeURL(params AuthorizeURLParams) string {
  return getC().AuthorizeURL(params)
}

func (c Client) AuthorizeURL(params AuthorizeURLParams) string {
  var buf bytes.Buffer
  buf.WriteString(stripe.CONNECTURL)

  // TODO: If express, we want to use /express take in options
  buf.WriteString("/oauth/authorize")

  v := url.Values{
    "client_id":     {stripe.ClientID},
  }
  if params.ResponseType != "" {
    v.Set("response_type", string(params.ResponseType))
  } else {
    v.Set("response_type", "code")
  }
  if params.Scope != "" {
    v.Set("scope", string(params.Scope))
  }
  if params.State != "" {
    v.Set("state", params.State)
  }
  if params.RedirectURI != "" {
    v.Set("redirect_uri", params.RedirectURI)
  }
  if params.StripeLanding != "" {
    v.Set("stripe_landing", string(params.StripeLanding))
  }
  // TODO How to handle AlwaysPrompt?

  buf.WriteByte('?')
  buf.WriteString(v.Encode())
  return buf.String()
}

func getC() Client {
  return Client{stripe.GetBackend(stripe.ConnectBackend), stripe.Key}
}
