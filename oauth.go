package stripe

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
