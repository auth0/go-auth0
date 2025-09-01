package authentication

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/auth0/go-auth0/v2/authentication/database"
)

// Database manager.
type Database manager

// Signup given a user's credentials and a connection, will create a new user using active authentication.
//
// This endpoint only works for database connections.
// See: https://auth0.com/docs/api/authentication?http#signup
func (d *Database) Signup(ctx context.Context, params database.SignupRequest, opts ...RequestOption) (r *database.SignupResponse, err error) {
	if params.ClientID == "" {
		params.ClientID = d.authentication.clientID
	}

	err = d.authentication.Request(ctx, "POST", d.authentication.URI("dbconnections", "signup"), params, &r, opts...)

	return
}

// ChangePassword given a user's email address and a connection, Auth0 will send a change password email.
//
// This endpoint only works for database connections.
// See: https://auth0.com/docs/api/authentication?http#change-password
func (d *Database) ChangePassword(ctx context.Context, params database.ChangePasswordRequest, opts ...RequestOption) (string, error) {
	if params.ClientID == "" {
		params.ClientID = d.authentication.clientID
	}

	request, err := d.authentication.NewRequest(ctx, "POST", d.authentication.URI("dbconnections", "change_password"), params, opts...)

	if err != nil {
		return "", fmt.Errorf("failed to create a new request: %w", err)
	}

	response, err := d.authentication.Do(request)
	if err != nil {
		return "", fmt.Errorf("failed to send the request: %w", err)
	}
	defer response.Body.Close()

	// If the response contains a client or a server error then return the error.
	if response.StatusCode >= http.StatusBadRequest {
		return "", newError(response)
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read the response body: %w", err)
	}

	return string(responseBody), nil
}
