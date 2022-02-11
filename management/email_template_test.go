package management

import (
	"net/http"
	"testing"

	"github.com/auth0/go-auth0"
)

func TestEmailTemplate(t *testing.T) {
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

	err := m.Email.Create(e)
	if err != nil {
		t.Fatal(err)
	}
	defer m.Email.Delete()

	et := &EmailTemplate{
		Template:  auth0.String("verify_email"),
		Body:      auth0.String("<html><body><h1>Verify your email</h1></body></html>"),
		From:      auth0.String("me@example.com"),
		ResultURL: auth0.String("https://www.example.com/verify-email"),
		Subject:   auth0.String("Verify your email"),
		Syntax:    auth0.String("liquid"),
		Enabled:   auth0.Bool(true),
	}

	t.Run("Create", func(t *testing.T) {
		err = m.EmailTemplate.Create(et)
		if err != nil {
			if err, ok := err.(Error); ok && err.Status() != http.StatusConflict {
				t.Fatal(err)
			}
		}
		t.Logf("%v\n", et)
	})

	t.Run("Read", func(t *testing.T) {
		et, err = m.EmailTemplate.Read(auth0.StringValue(et.Template))
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", et)
	})

	t.Run("Update", func(t *testing.T) {
		err = m.EmailTemplate.Update(auth0.StringValue(et.Template), &EmailTemplate{
			Body: auth0.String("<html><body><h1>Let's get you verified!</h1></body></html>"),
		})
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", et)
	})

	t.Run("Replace", func(t *testing.T) {
		et.Subject = auth0.String("Let's get you verified!")
		et.Body = auth0.String("<html><body><h1>Let's get you verified!</h1></body></html>")
		et.From = auth0.String("someone@example.com")

		err = m.EmailTemplate.Replace(auth0.StringValue(et.Template), et)
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", et)
	})

	t.Run("Delete", func(t *testing.T) {
		et.Enabled = auth0.Bool(false)
		err = m.EmailTemplate.Update(auth0.StringValue(et.Template), et)
		if err != nil {
			t.Error(err)
		}
	})
}
