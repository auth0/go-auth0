package management

import (
	"testing"
)

func TestSigningKey(t *testing.T) {
	var kid string

	// Our last test revokes the key used to sign the token we're currently
	// using, so we need to re-authenticate so that subsequent tests still work.
	t.Cleanup(func() {
		initTestManagement()
	})

	t.Run("List", func(t *testing.T) {
		ks, err := m.SigningKey.List()
		if err != nil {
			t.Fatal(err)
		}
		if len(ks) == 0 {
			t.Error("expected at least one key to be returned")
		}

		// save kid for later use
		kid = ks[0].GetKID()

		t.Logf("%v\n", ks)
	})

	t.Run("Read", func(t *testing.T) {
		k, err := m.SigningKey.Read(kid)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%v\n", k)
	})

	t.Run("Rotate", func(t *testing.T) {
		k, err := m.SigningKey.Rotate()
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%v\n", k)
	})

	t.Run("Revoke", func(t *testing.T) {
		ks, err := m.SigningKey.List()
		if err != nil {
			t.Fatal(err)
		}

		var pk *SigningKey
		for _, k := range ks {
			if k.GetPrevious() {
				pk = k
				break
			}
		}

		if pk == nil {
			t.Fatal("previous key not found, nothing to revoke")
		}

		r, err := m.SigningKey.Revoke(pk.GetKID())
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%v\n", r)
	})
}
