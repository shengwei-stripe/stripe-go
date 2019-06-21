// Package oauth provides the OAuth APIs
package oauth

import (
  "fmt"
  "net/http"
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
  response_type := "code"
  if params.ResponseType != nil {
    response_type = stripe.StringValue(params.ResponseType)
  }
  always_prompt := "false"
  if stripe.BoolValue(params.AlwaysPrompt) {
    always_prompt = "true"
  }

  path := stripe.FormatURLPath(
    "client_id=%s&response_type=%s&scope=%s&state=%s&redirect_uri=%s&stripe_landing=%s&always_prompt=%s",
    stripe.ClientID,
    response_type,
    stripe.StringValue(params.Scope),
    stripe.StringValue(params.State),
    stripe.StringValue(params.RedirectURI),
    stripe.StringValue(params.StripeLanding),
    always_prompt,
  )
  return fmt.Sprintf("https://connect.stripe.com%s/oauth/authorize?%s", express, path)
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

func getC() Client {
  return Client{stripe.GetBackend(stripe.ConnectBackend), stripe.Key}
}
