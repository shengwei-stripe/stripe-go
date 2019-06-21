package stripe

// Type of OAuth scope.
type ScopeType string

// List of possible values for OAuth scopes.
const (
  ScopeTypeReadWrite ScopeType = "read_write"
  ScopeTypeReadOnly  ScopeType = "read_only"
)

// Type of token. This will always be "bearer."
type OAuthTokenType string

// List of possible OAuthTokenType values.
const (
  OAuthTokenTypeBearer  OAuthTokenType = "bearer"
)

// Params for creating OAuth AuthorizeURL's.
type AuthorizeURLParams struct {
  State           *string
  ResponseType    *string
  Scope           *string
  RedirectURI     *string
  StripeLanding   *string
  AlwaysPrompt    *bool
  Express         *bool
}

// OAuthTokenParams is the set of paramaters that can be used to request
// OAuthTokens.
type OAuthTokenParams struct {
  Params         `form:"*"`
  GrantType      *string  `form:"grant_type"`
  Code           *string  `form:"code"`
  RefreshToken   *string  `form:"refresh_token"`
  Scope          *string  `form:"scope"`
  ClientSecret   *string  `form:"client_secret"`
}

// OAuthToken is the value of the OAuthToken from OAuth flow.
// https://stripe.com/docs/connect/oauth-reference#post-token
type OAuthToken struct {
  AccessToken            string          `json:"access_token"`
  Scope                  ScopeType       `json:"scope"`
  Livemode               bool            `json:"livemode"`
  OAuthTokenType         OAuthTokenType  `json:"token_type"`
  RefreshToken           string          `json:"refresh_token"`
  StripeUserID           string          `json:"stripe_user_id"`
  StripePublishableKey   string          `json:"stripe_publishable_key"`
}
