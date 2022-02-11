package management

import (
	"net/http"
	"testing"
	"time"

	"github.com/auth0/go-auth0"
)

func TestCustomDomain(t *testing.T) {
	c := &CustomDomain{
		Domain:               auth0.Stringf("%d.auth.uat.alexkappa.com", time.Now().UTC().Unix()),
		Type:                 auth0.String("auth0_managed_certs"),
		VerificationMethod:   auth0.String("txt"),
		TLSPolicy:            auth0.String("recommended"),
		CustomClientIPHeader: auth0.String("cf-connecting-ip"),
	}

	var err error

	t.Run("Create", func(t *testing.T) {
		err = m.CustomDomain.Create(c)
		if err != nil {
			if err, ok := err.(Error); ok && err.Status() == http.StatusForbidden {
				t.Skip(err)
			} else {
				t.Error(err)
			}
		}
		t.Logf("%v\n", c)
	})

	t.Run("Update", func(t *testing.T) {
		err = m.CustomDomain.Update(c.GetID(), &CustomDomain{TLSPolicy: auth0.String("recommended")})
		if err != nil {
			if err, ok := err.(Error); ok && err.Status() == http.StatusForbidden {
				t.Skip(err)
			} else {
				t.Error(err)
			}
		}
		t.Logf("%v\n", c)
	})

	t.Run("Read", func(t *testing.T) {
		c, err = m.CustomDomain.Read(c.GetID())
		if err != nil {
			if err, ok := err.(Error); ok && err.Status() == http.StatusNotFound {
				t.Skip(err)
			} else {
				t.Error(err)
			}
		}
		t.Logf("%v\n", c)
	})

	t.Run("Verify", func(t *testing.T) {
		c, err := m.CustomDomain.Verify(c.GetID())
		if err != nil {
			if err, ok := err.(Error); ok && err.Status() == http.StatusNotFound {
				t.Skip(err)
			} else {
				t.Error(err)
			}
		}
		t.Logf("%v\n", c)
	})

	t.Run("Delete", func(t *testing.T) {
		err = m.CustomDomain.Delete(c.GetID())
		if err != nil {
			if err, ok := err.(Error); ok && err.Status() == http.StatusNotFound {
				t.Skip(err)
			} else {
				t.Error(err)
			}
		}
	})
}
