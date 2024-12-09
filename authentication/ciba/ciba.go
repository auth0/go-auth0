package ciba

// Request defines the request body for calling the bc-authorize endpoint
type Request struct {
	// The client_id of your client.
	ClientID string `json:"client_id,omitempty"`
	// The client_secret of your client.
	ClientSecret string `json:"client_secret,omitempty"`
	// This is a required field containing format, iss and sub
	LoginHint map[string]string `json:"login_hint,omitempty"`
	// The scope for the flow
	Scope          string `json:"scope,omitempty"`
	Audience       string `json:"audience,omitempty"`
	BindingMessage string `json:"binding_message,omitempty"`
}

// Response defines the response of the CIBA request
type Response struct {
	AuthReqID string `json:"auth_req_id,omitempty"`
	ExpiresIn int64  `json:"expires_in,omitempty"`
	Interval  int64  `json:"interval,omitempty"`
}
