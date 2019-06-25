package stripe

// Type of OAuth scope.
type OAuthScopeType string

// List of possible values for OAuth scopes.
const (
  OAuthScopeTypeReadOnly  OAuthScopeType = "read_only"
  OAuthScopeTypeReadWrite OAuthScopeType = "read_write"
)

// Type of token. This will always be "bearer."
type OAuthTokenType string

// List of possible OAuthTokenType values.
const (
  OAuthTokenTypeBearer  OAuthTokenType = "bearer"
)

// Type of the business for the Stripe oauth user.
type OAuthStripeUserBusinessType string

// List of supported values for business type.
const (
  OAuthStripeUserBusinessTypeCorporation OAuthStripeUserBusinessType = "corporation"
  OAuthStripeUserBusinessTypeLLC         OAuthStripeUserBusinessType = "llc"
  OAuthStripeUserBusinessTypeNonProfit   OAuthStripeUserBusinessType = "non_profit"
  OAuthStripeUserBusinessTypePartnership OAuthStripeUserBusinessType = "partnership"
  OAuthStripeUserBusinessTypeSoleProp    OAuthStripeUserBusinessType = "sole_prop"
)

// The gender of the person who  will be filling out a Stripe application.
// (International regulations require either male or female.)
type OAuthStripeUserGender string

// The gender of the person who  will be filling out a Stripe application.
// (International regulations require either male or female.)
const (
  OAuthStripeUserGenderFemale OAuthStripeUserGender = "female"
  OAuthStripeUserGenderMale   OAuthStripeUserGender = "male"
)

// Type of Errors raised when failing authorization.
type OAuthError string

// List of supported OAuthError values.
const (
  OAuthErrorInvalidGrant            OAuthError = "invalid_grant"
  OAuthErrorInvalidRequest          OAuthError = "invalid_request"
  OAuthErrorInvalidScope            OAuthError = "invalid_scope"
  OAuthErrorUnsupportedGrantType    OAuthError = "unsupported_grant_type"
  OAuthErrorUnsupportedResponseType OAuthError = "unsupported_response_type"
)


// Type of Errors raised when failing authorization.
type DeauthorizationError string

// List of supported DeauthorizationError values.
const (
  DeauthorizationErrorInvalidClient   DeauthorizationError = "invalid_client"
  DeauthorizationErrorInvalidRequest  DeauthorizationError = "invalid_request"
)

// Params for the stripe_user OAuth Authorize params.
type OAuthStripeUserParams struct {
  BlockKana          *string                       `form:"block_kana"`
  BlockKanji         *string                       `form:"block_kanji"`
  BuildingKana       *string                       `form:"building_kana"`
  BuildingKanji      *string                       `form:"building_kanji"`
  BusinessName       *string                       `form:"business_name"`
  BusinessType       OAuthStripeUserBusinessType   `form:"business_type"`
  City               *string                       `form:"city"`
  Country            *string                       `form:"country"`
  Currency           *string                       `form:"currency"`
  DobDay             uint64                        `form:"dob_day"`
  DobMonth           uint64                        `form:"dob_month"`
  DobYear            uint64                        `form:"dob_year"`
  Email              *string                       `form:"email"`
  FirstName          *string                       `form:"first_name"`
  FirstNameKana      *string                       `form:"first_name_kana"`
  FirstNameKanji     *string                       `form:"first_name_kanji"`
  Gender             OAuthStripeUserGender         `form:"gender"`
  LastName           *string                       `form:"last_name"`
  LastNameKana       *string                       `form:"last_name_kana"`
  LastNameKanji      *string                       `form:"last_name_kanji"`
  PhoneNumber        *string                       `form:"phone_number"`
  PhysicalProduct    *bool                         `form:"physical_product"`
  ProductDescription *string                       `form:"product_description"`
  State              *string                       `form:"state"`
  StreetAddress      *string                       `form:"street_address"`
  Url                *string                       `form:"url"`
  Zip                *string                       `form:"zip"`
}

// Params for creating OAuth AuthorizeURL's.
type AuthorizeURLParams struct {
  Params          `form:"*"`
  AlwaysPrompt    *bool                   `form:"always_prompt"`
  ClientID        *string                 `form:"client_id"`
  Express         *bool                   `form:"-"`
  RedirectURI     *string                 `form:"redirect_uri"`
  ResponseType    *string                 `form:"response_type"`
  Scope           *string                 `form:"scope"`
  State           *string                 `form:"state"`
  StripeLanding   *string                 `form:"stripe_landing"`
  StripeUser      *OAuthStripeUserParams  `form:"stripe_user"`
}

// Params for deauthorizing an account.
type DeauthorizeParams struct {
  Params           `form:"*"`
  ClientID         *string    `form:"client_id"`
  StripeUserID     *string    `form:"stripe_user_id"`
}

// OAuthTokenParams is the set of paramaters that can be used to request
// OAuthTokens.
type OAuthTokenParams struct {
  Params         `form:"*"`
  ClientSecret   *string  `form:"client_secret"`
  Code           *string  `form:"code"`
  GrantType      *string  `form:"grant_type"`
  RefreshToken   *string  `form:"refresh_token"`
  Scope          *string  `form:"scope"`
}

// OAuthToken is the value of the OAuthToken from OAuth flow.
// https://stripe.com/docs/connect/oauth-reference#post-token
type OAuthToken struct {
  AccessToken            string          `json:"access_token"`
  Error                  OAuthError      `json:"error"`
  ErrorDescription       string          `json:"error_description"`
  Livemode               bool            `json:"livemode"`
  RefreshToken           string          `json:"refresh_token"`
  Scope                  OAuthScopeType  `json:"scope"`
  StripePublishableKey   string          `json:"stripe_publishable_key"`
  StripeUserID           string          `json:"stripe_user_id"`
  TokenType              OAuthTokenType  `json:"token_type"`
}

// Deauthorization is the value of the return from deauthorizing.
// https://stripe.com/docs/connect/oauth-reference#post-deauthorize
type Deauthorization struct {
  Error                  DeauthorizationError `json:"error"`
  ErrorDescription       string               `json:"error_description"`
  StripeUserID           string               `json:"stripe_user_id"`
}
