package management

import (
	"testing"

	"github.com/auth0/go-auth0"
	"github.com/auth0/go-auth0/internal/testing/expect"
)

func TestEmail(t *testing.T) {
	e := &Email{
		Name:               auth0.String("smtp"),
		Enabled:            auth0.Bool(true),
		DefaultFromAddress: auth0.String("no-reply@example.com"),
		Credentials: &EmailCredentials{
			SMTPHost: auth0.String("smtp.example.com"),
			SMTPPort: auth0.Int(587),
			SMTPUser: auth0.String("user"),
			SMTPPass: auth0.String("pass"),
		},
	}

	var err error

	t.Run("Create", func(t *testing.T) {
		err = m.Email.Create(e)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%v\n", e)
	})

	t.Run("Read", func(t *testing.T) {
		e, err = m.Email.Read()
		if err != nil {
			t.Error(err)
		}
		expect.Expect(t, e.GetName(), "smtp")
		expect.Expect(t, e.GetEnabled(), true)
		expect.Expect(t, e.GetDefaultFromAddress(), "no-reply@example.com")
		expect.Expect(t, e.GetCredentials().GetSMTPUser(), "user")
		expect.Expect(t, e.GetCredentials().GetSMTPPass(), "") // passwords are not returned from Auth0
		t.Logf("%v\n", e)
	})

	t.Run("Update", func(t *testing.T) {
		e.Enabled = auth0.Bool(false)
		e.DefaultFromAddress = auth0.String("info@example.com")

		err = m.Email.Update(e)
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", e)
	})

	t.Run("Delete", func(t *testing.T) {
		err = m.Email.Delete()
		if err != nil {
			t.Error(err)
		}
	})
}
