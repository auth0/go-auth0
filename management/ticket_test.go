package management

import (
	"testing"

	"github.com/auth0/go-auth0"
)

func TestTicket(t *testing.T) {
	var err error

	u := &User{
		Connection: auth0.String("Username-Password-Authentication"),
		Email:      auth0.String("chuck@chucknorris.com"),
		Username:   auth0.String("chuck"),
		Password:   auth0.String("I have a password and its a secret"),
	}
	if err = m.User.Create(u); err != nil {
		t.Fatal(err)
	}

	userID := u.GetID()

	t.Cleanup(func() { m.User.Delete(userID) })

	t.Run("VerifyEmail", func(t *testing.T) {
		v := &Ticket{
			ResultURL: auth0.String("https://example.com/verify-email"),
			UserID:    auth0.String(userID),
			TTLSec:    auth0.Int(3600),
		}

		err = m.Ticket.VerifyEmail(v)
		if err != nil {
			t.Error(err)
		}

		t.Logf("%v\n", v)
	})

	t.Run("ChangePassword", func(t *testing.T) {
		v := &Ticket{
			ResultURL:              auth0.String("https://example.com/change-password"),
			UserID:                 auth0.String(userID),
			TTLSec:                 auth0.Int(3600),
			MarkEmailAsVerified:    auth0.Bool(true),
			IncludeEmailInRedirect: auth0.Bool(true),
		}

		err = m.Ticket.ChangePassword(v)
		if err != nil {
			t.Error(err)
		}

		t.Logf("%v\n", v)
	})
}
