// Package oauth provides the OAuth APIs
package oauth

import (
  "fmt"
  "net/http"
  "github.com/stripe/stripe-go/form"
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
  express := ""
  if stripe.BoolValue(params.Express) {
    express = "/express"
  }
  if params.ClientID == nil {
    params.ClientID = stripe.String(stripe.ClientID)
  }
  if params.ResponseType == nil {
    params.ResponseType = stripe.String("code")
  }
  qs := &form.Values{}
  form.AppendTo(qs, params)
  return fmt.Sprintf(
    "https://connect.stripe.com%s/oauth/authorize?%s",
    express,
    qs.EncodeValues(),
  )
}

func New(params *stripe.OAuthTokenParams) (*stripe.OAuthToken, error) {
  return getC().New(params)
}

func (c Client) New(params *stripe.OAuthTokenParams) (*stripe.OAuthToken, error) {
  // client_secret is sent in the post body for this endpoint.
  if stripe.StringValue(params.ClientSecret) == "" {
    params.ClientSecret = stripe.String(stripe.Key)
  }
  if stripe.StringValue(params.GrantType) == "" {
    params.GrantType = stripe.String("authorization_code")
  }

  oauth_token := &stripe.OAuthToken{}
  err := c.B.Call(http.MethodPost, "/oauth/token", c.Key, params, oauth_token)

  return oauth_token, err
}

func Del(params *stripe.DeauthorizeParams) (*stripe.Deauthorization, error) {
  return getC().Del(params)
}

func (c Client) Del(params *stripe.DeauthorizeParams) (*stripe.Deauthorization, error) {
  deauthorization := &stripe.Deauthorization{}
  err := c.B.Call(
    http.MethodPost,
    "/oauth/deauthorize",
    c.Key,
    params,
    deauthorization,
  )
  return deauthorization, err
}

func getC() Client {
  return Client{stripe.GetBackend(stripe.ConnectBackend), stripe.Key}
}
