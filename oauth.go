package stripe

// Type of OAuth scope.
type ScopeType string

// List of possible values for OAuth scopes.
const (
  ScopeTypeReadOnly  ScopeType = "read_only"
  ScopeTypeReadWrite ScopeType = "read_write"
)

// Type of token. This will always be "bearer."
type OAuthTokenType string

// List of possible OAuthTokenType values.
const (
  OAuthTokenTypeBearer  OAuthTokenType = "bearer"
)

// Params for creating OAuth AuthorizeURL's.
type AuthorizeURLParams struct {
  AlwaysPrompt    *bool
  Express         *bool
  RedirectURI     *string
  ResponseType    *string
  Scope           *string
  State           *string
  StripeLanding   *string
}

// OAuthTokenParams is the set of paramaters that can be used to request
// OAuthTokens.
type OAuthTokenParams struct {
  ClientSecret   *string  `form:"client_secret"`
  Code           *string  `form:"code"`
  GrantType      *string  `form:"grant_type"`
  Params         `form:"*"`
  RefreshToken   *string  `form:"refresh_token"`
  Scope          *string  `form:"scope"`
}

// OAuthToken is the value of the OAuthToken from OAuth flow.
// https://stripe.com/docs/connect/oauth-reference#post-token
type OAuthToken struct {
  AccessToken            string          `json:"access_token"`
  Livemode               bool            `json:"livemode"`
  RefreshToken           string          `json:"refresh_token"`
  Scope                  ScopeType       `json:"scope"`
  StripePublishableKey   string          `json:"stripe_publishable_key"`
  StripeUserID           string          `json:"stripe_user_id"`
  TokenType              OAuthTokenType  `json:"token_type"`
}
