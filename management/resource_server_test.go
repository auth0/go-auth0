package management

import (
	"testing"
	"time"

	"github.com/auth0/go-auth0"
)

func TestResourceServer(t *testing.T) {
	s := &ResourceServer{
		Name:             auth0.Stringf("Test Resource Server (%s)", time.Now().Format(time.StampMilli)),
		Identifier:       auth0.String("https://api.example.com/"),
		SigningAlgorithm: auth0.String("HS256"),

		TokenLifetime:       auth0.Int(7200),
		TokenLifetimeForWeb: auth0.Int(3600),

		Scopes: []*ResourceServerScope{
			{
				Value:       auth0.String("create:resource"),
				Description: auth0.String("Create Resource"),
			},
		},
	}

	var err error

	t.Run("Create", func(t *testing.T) {
		err = m.ResourceServer.Create(s)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%v\n", s)
	})

	t.Run("Read", func(t *testing.T) {
		s, err = m.ResourceServer.Read(auth0.StringValue(s.ID))
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", s)
	})

	t.Run("Update", func(t *testing.T) {
		id := auth0.StringValue(s.ID)

		s.ID = nil         // read-only
		s.Identifier = nil // read-only

		s.AllowOfflineAccess = auth0.Bool(true)
		s.SigningAlgorithm = auth0.String("RS256")
		s.SkipConsentForVerifiableFirstPartyClients = auth0.Bool(true)

		s.TokenLifetime = auth0.Int(7200)
		s.TokenLifetimeForWeb = auth0.Int(5400)

		s.Scopes = append(s.Scopes, &ResourceServerScope{
			Value:       auth0.String("update:resource"),
			Description: auth0.String("Update Resource"),
		})

		err = m.ResourceServer.Update(id, s)
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", s)
	})

	t.Run("Delete", func(t *testing.T) {
		err = m.ResourceServer.Delete(auth0.StringValue(s.ID))
		if err != nil {
			t.Error(err)
		}
	})
}
