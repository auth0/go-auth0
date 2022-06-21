package management

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/auth0/go-auth0"
)

func TestTicketManager_VerifyEmail(t *testing.T) {
	setupHTTPRecordings(t)

	user := givenAUser(t)
	ticket := &Ticket{
		ResultURL: auth0.String("https://example.com/verify-email"),
		UserID:    user.ID,
		TTLSec:    auth0.Int(3600),
	}

	err := m.Ticket.VerifyEmail(ticket)
	assert.NoError(t, err)
}

func TestTicketManager_ChangePassword(t *testing.T) {
	setupHTTPRecordings(t)

	user := givenAUser(t)
	ticket := &Ticket{
		ResultURL:              auth0.String("https://example.com/change-password"),
		UserID:                 user.ID,
		TTLSec:                 auth0.Int(3600),
		MarkEmailAsVerified:    auth0.Bool(true),
		IncludeEmailInRedirect: auth0.Bool(true),
	}

	err := m.Ticket.ChangePassword(ticket)
	assert.NoError(t, err)
}
