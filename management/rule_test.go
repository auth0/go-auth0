package management

import (
	"testing"

	"github.com/auth0/go-auth0"
)

func TestRule(t *testing.T) {
	r := &Rule{
		Name:    auth0.String("test-rule"),
		Script:  auth0.String("function (user, context, callback) { callback(null, user, context); }"),
		Enabled: auth0.Bool(false),
	}

	var err error

	t.Run("Create", func(t *testing.T) {
		err = m.Rule.Create(r)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%v\n", r)
	})

	t.Run("Read", func(t *testing.T) {
		r, err = m.Rule.Read(auth0.StringValue(r.ID))
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", r)
	})

	t.Run("Update", func(t *testing.T) {
		id := auth0.StringValue(r.ID)

		r.ID = nil // read-only
		r.Order = auth0.Int(5)
		r.Enabled = auth0.Bool(true)

		err = m.Rule.Update(id, r)
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", r)
	})

	t.Run("List", func(t *testing.T) {
		r, err := m.Rule.List()
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", r)
	})

	t.Run("Delete", func(t *testing.T) {
		err = m.Rule.Delete(auth0.StringValue(r.ID))
		if err != nil {
			t.Error(err)
		}
	})
}
