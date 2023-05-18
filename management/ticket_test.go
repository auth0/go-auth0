package management

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/auth0/go-auth0"
)

func TestTicketManager_VerifyEmail(t *testing.T) {
	configureHTTPTestRecordings(t)

	user := givenAUser(t)
	ticket := &Ticket{
		ResultURL: auth0.String("https://example.com/verify-email"),
		UserID:    user.ID,
		TTLSec:    auth0.Int(3600),
	}

	err := api.Ticket.VerifyEmail(context.Background(), ticket)
	assert.NoError(t, err)
}

func TestTicketManager_ChangePassword(t *testing.T) {
	configureHTTPTestRecordings(t)

	user := givenAUser(t)
	ticket := &Ticket{
		ResultURL:              auth0.String("https://example.com/change-password"),
		UserID:                 user.ID,
		TTLSec:                 auth0.Int(3600),
		MarkEmailAsVerified:    auth0.Bool(true),
		IncludeEmailInRedirect: auth0.Bool(true),
	}

	err := api.Ticket.ChangePassword(context.Background(), ticket)
	assert.NoError(t, err)
}
