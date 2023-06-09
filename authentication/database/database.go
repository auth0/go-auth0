package database

import (
	"encoding/json"

	"go.devnw.com/structs"
)

// SignUpRequest is a sign up request.
type SignUpRequest struct {
	// The client_id of your client.
	ClientID string `json:"client_id,omitempty"`
	// The user's email address.
	Email string `json:"email,omitempty"`
	// The user's desired password.
	Password string `json:"password,omitempty"`
	// The name of the database configured to your client.
	Connection string `json:"connection,omitempty"`
	// The user's username. Only valid if the connection requires a username.
	Username string `json:"username,omitempty"`
	// The user's given name(s).
	GivenName string `json:"given_name,omitempty"`
	// The user's family name(s).
	FamilyName string `json:"family_name,omitempty"`
	// The user's full name.
	Name string `json:"name,omitempty"`
	// The user's nickname.
	Nickname string `json:"nickname,omitempty"`
	// A URI pointing to the user's picture.
	Picture string `json:"picture,omitempty"`
	// The user metadata to be associated with the user. If set, the field must be an object containing no more than ten properties.
	// Property names can have a maximum of 100 characters, and property values must be strings of no more than 500 characters.
	UserMetadata *map[string]interface{} `json:"user_metadata,omitempty"`
	// Extra parameters to be merged into the request body. Values set here will override any existing values
	ExtraParameters map[string]string `json:"-"`
}

// SignUpResponse is a sign up response.
type SignUpResponse struct {
	// The user's email address.
	Email string `json:"email,omitempty"`
	// Indicates whether a user has verified their email address.
	EmailVerified bool `json:"email_verified,omitempty"`
	// The users ID
	ID string `json:"_id,omitempty"`
	// The user's username. Only valid if the connection requires a username.
	Username string `json:"username,omitempty"`
	// The user's given name(s).
	GivenName string `json:"given_name,omitempty"`
	// The user's family name(s).
	FamilyName string `json:"family_name,omitempty"`
	// The user's full name.
	Name string `json:"name,omitempty"`
	// The user's nickname.
	Nickname string `json:"nickname,omitempty"`
	// A URI pointing to the user's picture.
	Picture string `json:"picture,omitempty"`
	// The user metadata to be associated with the user. If set, the field must be an object containing no more than ten properties.
	// Property names can have a maximum of 100 characters, and property values must be strings of no more than 500 characters.
	UserMetadata *map[string]interface{} `json:"user_metadata,omitempty"`
}

// MarshalJSON implements the json.Unmarshaler interface.
//
// It is required to support adding parameters from the `ExtraParameters`
// field onto the request object.
func (s *SignUpRequest) MarshalJSON() ([]byte, error) {
	n := structs.New(s)
	n.TagName = "json"

	m := n.Map()
	for k, v := range s.ExtraParameters {
		m[k] = v
	}

	return json.Marshal(m)
}
