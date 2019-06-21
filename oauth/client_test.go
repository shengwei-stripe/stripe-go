package oauth

import (
  "bytes"
	"io/ioutil"
  // "os"
  "net/http"
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

  assert.Contains(t, url, "https://connect.stripe.com/express/oauth/authorize?")
  assert.Contains(t, url, "client_id=ca_123")
  assert.Contains(t, url, "redirect_uri=https%3A%2F%2Ft.example.com")
  assert.Contains(t, url, "response_type=test-code")
  assert.Contains(t, url, "scope=read_only")
  assert.Contains(t, url, "state=test-state")
  assert.Contains(t, url, "stripe_landing=register")
  assert.Contains(t, url, "always_prompt=true")
}


// RoundTripFunc .
type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip .
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
  return f(req), nil
}

//NewTestClient returns *http.Client with Transport replaced to avoid making real calls
func NewTestClient(fn RoundTripFunc) *http.Client {
  return &http.Client{
    Transport: RoundTripFunc(fn),
  }
}

func StubConnectBackend(httpClient *http.Client) {
  mockBackend := stripe.GetBackendWithConfig(
    stripe.ConnectBackend,
    &stripe.BackendConfig{
      URL:        "https://localhost:12113",
      HTTPClient: httpClient,
    },
  )
  stripe.SetBackend(stripe.ConnectBackend, mockBackend)
}

func TestNewOAuthToken(t *testing.T) {
  stripe.Key = "sk_123"
  // stripe-mock doesn't support connect URL's so this stubs out the server.
  httpClient := NewTestClient(func(req *http.Request) *http.Response {
    buf := new(bytes.Buffer)
    buf.ReadFrom(req.Body)
    reqBody := buf.String()

    assert.Contains(t, req.URL.String(), "https://localhost:12113/oauth/token")
    assert.Contains(t, reqBody, "client_secret=sk_123")
    assert.Contains(t, reqBody, "grant_type=authorization_code")
    assert.Contains(t, reqBody, "code=code")

    responseBody := `{
      "access_token":"sk_123",
      "livemode":false,
      "refresh_token":"rt_123",
      "token_type":"bearer",
      "stripe_publishable_key":"pk_123",
      "stripe_user_id":"acct_123",
      "scope":"read_write"
    }`
    return &http.Response{
      StatusCode: 200,
      // Send response to be tested
      Body:       ioutil.NopCloser(bytes.NewBufferString(responseBody)),
      Header:     make(http.Header),
    }
  })
  StubConnectBackend(httpClient)

  token, err := New(&stripe.OAuthTokenParams{
    GrantType:   stripe.String("authorization_code"),
    Code:        stripe.String("code"),
  })
  assert.Nil(t, err)
  assert.NotNil(t, token)
  assert.Equal(t, token.AccessToken, "sk_123")
  assert.Equal(t, token.Livemode, false)
  assert.Equal(t, token.RefreshToken, "rt_123")
  assert.Equal(t, token.TokenType, stripe.OAuthTokenTypeBearer)
  assert.Equal(t, token.StripePublishableKey, "pk_123")
  assert.Equal(t, token.StripeUserID, "acct_123")
  assert.Equal(t, token.Scope, stripe.ScopeTypeReadWrite)
}

func TestNewOAuthTokenWithCustomKey(t *testing.T) {
  stripe.Key = "sk_123"
  // stripe-mock doesn't support connect URL's so this stubs out the server.
  httpClient := NewTestClient(func(req *http.Request) *http.Response {
    buf := new(bytes.Buffer)
    buf.ReadFrom(req.Body)
    reqBody := buf.String()
    assert.Contains(t, reqBody, "client_secret=sk_999")

    return &http.Response{
      StatusCode: 200,
      // Send response to be tested
      Body:       ioutil.NopCloser(bytes.NewBufferString(`{}`)),
      Header:     make(http.Header),
    }
  })
  StubConnectBackend(httpClient)

  token, err := New(&stripe.OAuthTokenParams{
    ClientSecret: stripe.String("sk_999"),
  })
  assert.Nil(t, err)
  assert.NotNil(t, token)
}
