package authentication

import (
	"context"

	"github.com/auth0/go-auth0/authentication/database"
)

// Database manager.
type Database manager

// Signup given a user's credentials and a connection, will create a new user using active authentication.
//
// This endpoint only works for database connections.
// See: https://auth0.com/docs/api/authentication?http#signup
func (d *Database) Signup(ctx context.Context, params database.SignupRequest) (r *database.SignupResponse, err error) {
	if params.ClientID == "" {
		params.ClientID = d.authentication.clientID
	}

	err = d.authentication.Request(ctx, "POST", d.authentication.URI("dbconnections", "signup"), params, &r)
	return
}
